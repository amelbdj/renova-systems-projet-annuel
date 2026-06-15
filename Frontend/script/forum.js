let sujetActifId = null;

function initForum() {
  chargerForum();
}

function chargerForum() {
  const token = localStorage.getItem("token");
  const container = document.getElementById("liste-sujets-forum");
  if (!container) return;

  container.innerHTML = "<p>Chargement du forum...</p>";

  fetch("http://localhost:8081/user/forums", {
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
          "<p>Aucun sujet pour le moment. Soyez le premier !</p>";
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
                Par <strong>${sujet.auteur}</strong> le ${sujet.date_creation}
            </div>
        </div>
    `;
      });
    })
    .catch(function (error) {
      console.error("Erreur chargement forum :", error);
      container.innerHTML = "<p>Impossible de charger le forum.</p>";
    });
}

window.ouvrirSujet = function (idTopic, titre) {
  sujetActifId = idTopic;

  document.getElementById("vue-liste-forums").style.display = "none";
  document.getElementById("vue-sujet-actif").style.display = "flex";
  document.getElementById("titre-sujet-actif").textContent = titre;

  chargerMessagesSujet(idTopic);
};

window.retourListeForums = function () {
  sujetActifId = null;

  document.getElementById("vue-sujet-actif").style.display = "none";
  document.getElementById("vue-liste-forums").style.display = "block";

  chargerForum();
};

function chargerMessagesSujet(idTopic) {
  const token = localStorage.getItem("token");
  const zone = document.getElementById("zone-messages");
  zone.innerHTML = "<p>Chargement...</p>";

  fetch("http://localhost:8081/user/forums/messages?topic_id=" + idTopic, {
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

      // 🛡️ BLINDAGE DES DONNÉES LOCALES CONTRE LE RETOUR NULL
      if (messages === null || !messages) {
        messages = [];
      }

      if (messages.length === 0) {
        zone.innerHTML =
          "<p style='text-align:center; padding:20px; color:gray;'>Aucune réponse visible pour le moment.</p>";
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
        "<p style='color:red; text-align:center; padding:20px;'>Erreur lors du chargement des messages.</p>";
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

  fetch("http://localhost:8081/user/forums/messages", {
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

  fetch("http://localhost:8081/user/forums", {
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
    window.location.href = "login.html";
    return;
  }
  window.location.href = `profil.html?id=${userId}`;
}

document.addEventListener("DOMContentLoaded", initForum);
