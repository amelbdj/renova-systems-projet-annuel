let sujetActifId = null;
let sondageForum = null; 

function initForum() {
  chargerForum();
}

function chargerForum() {
  const token = localStorage.getItem("token");
  const container = document.getElementById("liste-sujets-forum");
  if (!container) return;

  container.innerHTML = "<p>Chargement du forum...</p>";

  fetch(API_BASE_URL + "/user/forums", {
    headers: { Authorization: "Bearer " + token },
  })
    .then(function (response) {
      if (!response.ok) throw new Error("Erreur réseau");
      return response.json();
    })
    .then(function (topics) {
      container.innerHTML = "";

      if (!topics || topics.length === 0) {
        container.innerHTML =
          '<p data-i18n="forum.empty">Aucun sujet pour le moment. Soyez le premier !</p>';
        if (typeof appliquerTraductions === "function") appliquerTraductions();
        return;
      }

      topics.forEach(function (sujet) {
        const safeTitre = sujet.titre.replace(/'/g, "\\'");

        container.innerHTML += `
        <div class="sujet-card" onclick="ouvrirSujet(${sujet.id_topic}, '${safeTitre}')">
            <div class="sujet-header">
                <div class="sujet-titre">${sujet.titre}</div>
                <div class="sujet-badge"><span class="material-symbols-outlined">chat_bubble</span> ${sujet.nb_reponses}</div>
            </div>
            <div class="sujet-meta">
                <span data-i18n="forum.by">Par</span> <strong>${sujet.auteur}</strong> <span data-i18n="forum.on">le</span> ${sujet.date_creation}
            </div>
        </div>
    `;
      });

      
      if (typeof appliquerTraductions === "function") appliquerTraductions();
    })
    .catch(function (error) {
      console.error("Erreur chargement forum :", error);
      container.innerHTML =
        '<p data-i18n="forum.load_error">Impossible de charger le forum.</p>';
      if (typeof appliquerTraductions === "function") appliquerTraductions();
    });
}

window.ouvrirSujet = function (idTopic, titre) {
  sujetActifId = idTopic;

  document.getElementById("vue-liste-forums").style.display = "none";
  document.getElementById("vue-sujet-actif").style.display = "flex";
  const titreEl = document.getElementById("titre-sujet-actif");
  
  
  titreEl.removeAttribute("data-i18n");
  titreEl.textContent = titre;

  chargerMessagesSujet(idTopic);

  
  
  
  if (sondageForum) clearInterval(sondageForum);
  sondageForum = setInterval(function () {
    if (sujetActifId) chargerMessagesSujet(sujetActifId, true);
  }, 4000);
};

window.retourListeForums = function () {
  sujetActifId = null;

  
  if (sondageForum) {
    clearInterval(sondageForum);
    sondageForum = null;
  }

  document.getElementById("vue-sujet-actif").style.display = "none";
  document.getElementById("vue-liste-forums").style.display = "block";

  chargerForum();
};

function chargerMessagesSujet(idTopic, silencieux) {
  const token = localStorage.getItem("token");
  const zone = document.getElementById("zone-messages");
  
  
  if (!silencieux) zone.innerHTML = "<p>Chargement...</p>";

  fetch(API_BASE_URL + "/user/forums/messages?topic_id=" + idTopic, {
    headers: { Authorization: "Bearer " + token },
  })
    .then(function (res) {
      if (!res.ok) throw new Error("Erreur serveur lors de la récupération");
      return res.json();
    })
    .then(function (messages) {
      console.log(
        "🔍 Payload reçu du serveur pour le sujet " + idTopic + " :",
        messages,
      );
      zone.innerHTML = "";

      
      if (messages === null || !messages) {
        messages = [];
      }

      if (messages.length === 0) {
        zone.innerHTML =
          "<p style='text-align:center; padding:20px; color:gray;' data-i18n=\"forum.no_replies\">Aucune réponse visible pour le moment.</p>";
        if (typeof appliquerTraductions === "function") appliquerTraductions();
        return;
      }

      messages.forEach(function (m) {
        const prenom = m.prenom_auteur || m.PrenomAuteur || "Utilisateur";
        const nom = m.nom_auteur || m.NomAuteur || "Anonyme";
        const date = m.date_creation || m.DateCreation || "";
        const contenu = m.contenu || m.Contenu || "";

        zone.innerHTML += `
        <div class="msg-card">
            <div class="msg-meta">
                <span class="msg-auteur">${prenom} ${nom}</span> 
                <span>• ${date}</span>
            </div>
            <div class="msg-contenu">
                ${contenu}
            </div>
        </div>
    `;
      });
      zone.scrollTop = zone.scrollHeight;
    })
    .catch(function (err) {
      console.error("Erreur messages:", err);
      zone.innerHTML =
        "<p style='color:red; text-align:center; padding:20px;' data-i18n=\"forum.messages_error\">Erreur lors du chargement des messages.</p>";
      if (typeof appliquerTraductions === "function") appliquerTraductions();
    });
}

window.envoyerMessage = function () {
  const inputElement = document.getElementById("input-nouveau-message");
  const texte = inputElement.value.trim();
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("token");

  if (!texte || !sujetActifId || !userId) return;

  const payload = {
    id_topic: parseInt(sujetActifId),
    id_user: parseInt(userId),
    contenu: texte,
  };

  fetch(API_BASE_URL + "/user/forums/messages", {
    method: "POST",
    headers: {
      Authorization: "Bearer " + token,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  })
    .then(function (res) {
      if (!res.ok) throw new Error("Erreur d'envoi");
      inputElement.value = "";
      chargerMessagesSujet(sujetActifId);
    })
    .catch(function (err) {
      alert("Erreur lors de l'envoi du message");
      console.error(err);
    });
};

window.ouvrirNouveauSujet = function () {
  document.getElementById("modal-nouveau-sujet").style.display = "flex";
};

window.fermerModalSujet = function () {
  document.getElementById("modal-nouveau-sujet").style.display = "none";
  document.getElementById("nouveau-sujet-titre").value = "";
  document.getElementById("nouveau-sujet-message").value = "";
};

window.validerNouveauSujet = function () {
  const titre = document.getElementById("nouveau-sujet-titre").value.trim();
  const message = document.getElementById("nouveau-sujet-message").value.trim();
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("token");

  if (!titre || !message) {
    alert("Veuillez remplir le titre et le message.");
    return;
  }

  const payload = {
    id_user: parseInt(userId),
    titre: titre,
    message: message,
  };

  fetch(API_BASE_URL + "/user/forums", {
    method: "POST",
    headers: {
      Authorization: "Bearer " + token,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  })
    .then(function (res) {
      if (!res.ok) throw new Error("Erreur de création");
      fermerModalSujet();
      chargerForum();
    })
    .catch(function (err) {
      console.error("Erreur:", err);
      alert("Impossible de créer le sujet.");
    });
};

function goToProfile() {
  const userId = localStorage.getItem("userId");
  if (!userId) {
    window.location.href = "/login";
    return;
  }
  window.location.href = `/profil?id=${userId}`;
}

document.addEventListener("DOMContentLoaded", initForum);
