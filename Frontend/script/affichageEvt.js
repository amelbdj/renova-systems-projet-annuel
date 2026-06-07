// Variable globale pour empêcher le spam des boutons de paiement/inscription
window.isProcessingPayment = false;

function LancerRecherche() {
  const motCle = document.getElementById("searchInput").value.trim();
  chargerEvenementsClient(motCle);
}

function chargerEvenementsClient(motCle = "") {
  const container = document.getElementById("liste-evenements");
  if (!container) return; // Sécurité : on arrête si on n'est pas sur la bonne page

  container.innerHTML =
    "<p style='color: var(--txt-m); text-align: center; grid-column: 1 / -1;'>Chargement des événements...</p>";

  let url = `http://localhost:8081/admin/evenements`;
  if (motCle !== "") {
    url += `?search=${encodeURIComponent(motCle)}`;
  }

  fetch(url, {
    method: "GET",
    headers: {
      Authorization: "Bearer " + localStorage.getItem("token"),
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur réseau");
      return res.json();
    })
    .then((evenements) => {
      container.innerHTML = "";

      // 🟢 AJOUT : On sauvegarde les données globalement pour la sécurité et la modale
      window.evenementData = evenements;

      if (!evenements || evenements.length === 0) {
        container.innerHTML = `<p style="color: var(--txt-m); text-align: center; grid-column: 1 / -1;">Aucun événement publié pour le moment.</p>`;
        return;
      }

      let htmlContent = "";
      let evenementsAffiches = 0;
      const maintenant = new Date();

      evenements.forEach((evt) => {
        // --- CORRECTION DU FORMAT DE DATE ---
        let parts = evt.date_debut.split(" a ");
        let dateParts = parts[0].split("/");
        let timeParts = parts[1].split(":");
        const dateEvenement = new Date(
          dateParts[2],
          dateParts[1] - 1,
          dateParts[0],
          timeParts[0],
          timeParts[1],
        );

        if (evt.statut_validation === "valide" && dateEvenement >= maintenant) {
          evenementsAffiches++;

          const idEvt = evt.id;
          const imageCover =
            evt.image_url ||
            "https://images.unsplash.com/photo-1542601906990-b4d3fb778b09?w=500";
          const resume =
            (evt.description || "Pas de description.").substring(0, 100) +
            "...";

          // 🟢 MODIFICATION : Gestion dynamique Inscription (Bleu) / Désinscription (Rouge)
          let boutonAction = evt.deja_inscrit
            ? `<span style="background-color: #ef4444; color: white; padding: 8px 15px; border-radius: 6px; font-weight: 600; font-size: 14px; margin-top: 10px; display: inline-block; cursor: pointer;" onclick="SeDesinscrire(${idEvt}); event.stopPropagation();">Se désinscrire ➔</span>`
            : `<span style="background-color: var(--blue); color: white; padding: 8px 15px; border-radius: 6px; font-weight: 600; font-size: 14px; margin-top: 10px; display: inline-block; cursor: pointer;" onclick="sinscrireEvenement(${idEvt}, ${evt.prix}); event.stopPropagation();">S'inscrire ➔</span>`;

          htmlContent += `
                <div class="article-card" onclick="OuvrirEvenement(${idEvt})">
                    <img src="${imageCover}" alt="Image" class="card-img">
                    <div style="padding: 15px; flex-grow: 1; display: flex; flex-direction: column; align-items: center; text-align: center;">
                        <h3 style="margin-top: 0; color: #ffffff; font-size: 18px;">${evt.titre}</h3>
                        <p style="color: var(--blue-l); font-size: 13px; font-weight: bold;">${evt.date_debut}</p>
                        <p style="color: var(--txt-m); font-size: 14px;">${resume}</p>
                        ${boutonAction}
                    </div>
                </div>`;
        }
      });
      container.innerHTML =
        evenementsAffiches === 0
          ? "<p>Aucun événement à venir.</p>"
          : htmlContent;
    })
    .catch((err) => {
      console.error("Erreur :", err);
      container.innerHTML = "<p>Erreur de connexion.</p>";
    });
}

// --- FONCTION D'INSCRIPTION ULTRA-SÉCURISÉE ---
function sinscrireEvenement(idEvent, prixEvent) {
  const idUser = localStorage.getItem("userId");
  const monToken = localStorage.getItem("token");

  if (!idUser || idUser === "null") {
    alert("Erreur : Vous devez être connecté pour vous inscrire.");
    return;
  }

  // 🛡️ SÉCURITÉ 1 : Blocage si déjà inscrit dans les données locales
  if (window.evenementData) {
    const currentEvt = window.evenementData.find((e) => e.id === idEvent);
    if (currentEvt && currentEvt.deja_inscrit) {
      alert("⚠️ Vous êtes déjà inscrit à cet événement !");
      return;
    }
  }

  // 🛡️ SÉCURITÉ 2 : Verrou anti-spam au clic
  if (window.isProcessingPayment) return;
  window.isProcessingPayment = true;

  // SCÉNARIO 1 : L'ÉVÉNEMENT EST PAYANT (Prix > 0)
  if (prixEvent > 0) {
    fetch("http://localhost:8081/api/web/checkout/evenement", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id_user: parseInt(idUser),
        id_event: parseInt(idEvent),
      }),
    })
      .then(function (response) {
        return response.json();
      })
      .then(function (data) {
        if (data.checkout_url) {
          window.location.href = data.checkout_url;
        } else {
          alert("Erreur lors de l'initialisation du paiement.");
          window.isProcessingPayment = false;
        }
      })
      .catch(function (error) {
        console.error("Erreur Stripe :", error);
        alert("Impossible de contacter le serveur de paiement.");
        window.isProcessingPayment = false;
      });
  }

  // SCÉNARIO 2 : L'ÉVÉNEMENT EST GRATUIT
  else {
    fetch("http://localhost:8081/admin/evenements/inscription", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + monToken,
      },
      body: JSON.stringify({
        id_user: parseInt(idUser),
        id_event: parseInt(idEvent),
      }),
    })
      .then(function (response) {
        return response.json().then(function (data) {
          if (!response.ok) {
            throw data.erreur || "Erreur lors de l'inscription";
          }
          return data;
        });
      })
      .then(function (data) {
        alert("Succès : Inscription validée à l'événement gratuit !");
        window.isProcessingPayment = false;
        FermerEvenement();
        LancerRecherche();
      })
      .catch(function (error) {
        console.error("Erreur d'inscription directe :", error);
        alert("Erreur : " + error);
        window.isProcessingPayment = false;
      });
  }
}

function OuvrirEvenement(id) {
  if (!window.evenementData) return;
  const evt = window.evenementData.find((a) => a.id === id);
  if (!evt) return;

  const imageCover =
    evt.image_url ||
    "https://images.unsplash.com/photo-1542601906990-b4d3fb778b09?w=500";

  document.getElementById("modalImage").src = imageCover;
  document.getElementById("modalTitre").textContent = evt.titre;

  const auteur =
    `${evt.nomSalarie || ""} ${evt.prenomSalarie || ""}`.trim() ||
    "UpcycleConnect";

  document.getElementById("modalMeta").textContent =
    `Animé par ${auteur} • Le ${evt.date_debut}`;

  document.getElementById("modalContenu").innerHTML =
    evt.description || "Pas de description disponible pour cet événement.";

  document.getElementById("articleModal").style.display = "flex";
}

function FermerEvenement() {
  const modal = document.getElementById("articleModal");
  // 🟢 SÉCURITÉ : On vérifie si la modale existe avant de toucher à son style
  if (modal) {
    modal.style.display = "none";
  }
}
function SeDesinscrire(idEvent) {
  if (
    !confirm("Voulez-vous vraiment annuler votre inscription à cet événement ?")
  ) {
    return;
  }

  const idUser = localStorage.getItem("userId");
  const monToken = localStorage.getItem("token");

  if (!idUser || idUser === "null") {
    alert("Erreur : Vous devez être connecté pour faire cette action.");
    return;
  }

  fetch("http://localhost:8081/admin/evenements/desinscription", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
    body: JSON.stringify({
      id_user: parseInt(idUser),
      id_event: idEvent,
    }),
  })
    .then(function (res) {
      return res.json().then(function (data) {
        if (!res.ok) {
          throw data.erreur || "Erreur lors de la désinscription";
        }
        return data;
      });
    })
    .then(function (data) {
      alert("Succès : " + (data.message || "Désinscription validée"));
      FermerEvenement();
      // On recharge la liste : le bouton redeviendra bleu automatiquement !
      chargerEvenementsClient();
    })
    .catch(function (errorMessage) {
      alert("Attention : " + errorMessage);
    });
}

document.addEventListener("DOMContentLoaded", () => {
  chargerEvenementsClient();

  const searchInput = document.getElementById("searchInput");
  if (searchInput) {
    searchInput.addEventListener("keypress", function (event) {
      if (event.key === "Enter") {
        event.preventDefault();
        LancerRecherche();
      }
    });
  }
});
