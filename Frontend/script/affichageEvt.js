
window.isProcessingPayment = false;

function LancerRecherche() {
  const motCle = document.getElementById("searchInput").value.trim();
  chargerEvenementsClient(motCle);
}

function chargerEvenementsClient(motCle = "") {
  const container = document.getElementById("liste-evenements");
  if (!container) return; 

  container.innerHTML =
    "<p style='color: var(--txt-m); text-align: center; grid-column: 1 / -1;' data-i18n=\"evenement.loading\">Chargement des événements...</p>";
  if (typeof appliquerTraductions === "function") appliquerTraductions();

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

      
      window.evenementData = evenements;

      if (!evenements || evenements.length === 0) {
        container.innerHTML = `<p style="color: var(--txt-m); text-align: center; grid-column: 1 / -1;" data-i18n="evenement.empty">Aucun événement publié pour le moment.</p>`;
        if (typeof appliquerTraductions === "function") appliquerTraductions();
        return;
      }

      let htmlContent = "";
      let evenementsAffiches = 0;
      const maintenant = new Date();

      evenements.forEach((evt) => {
        
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
            evt.image_url && evt.image_url.trim() !== ""
              ? `http://localhost:8081/${evt.image_url}`
              : "https://images.unsplash.com/photo-1542601906990-b4d3fb778b09?w=500";

          const resume =
            (evt.description || "Pas de description.").substring(0, 100) +
            "...";

          
          let boutonAction = evt.deja_inscrit
            ? `<span style="background-color: #ef4444; color: white; padding: 8px 15px; border-radius: 6px; font-weight: 600; font-size: 14px; margin-top: 10px; display: inline-block; cursor: pointer;" onclick="SeDesinscrire(${idEvt}); event.stopPropagation();" data-i18n="evenement.unsubscribe">Se désinscrire ➔</span>`
            : `<span style="background-color: var(--blue); color: white; padding: 8px 15px; border-radius: 6px; font-weight: 600; font-size: 14px; margin-top: 10px; display: inline-block; cursor: pointer;" onclick="sinscrireEvenement(${idEvt}, ${evt.prix}); event.stopPropagation();" data-i18n="evenement.subscribe">S'inscrire ➔</span>`;

          htmlContent += `
                <div class="article-card" onclick="OuvrirEvenement(${idEvt})">
                    <img src="${imageCover}" alt="Image" class="card-img" style="width: 100%; height: 200px; object-fit: cover;">
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
          ? '<p data-i18n="evenement.empty_upcoming">Aucun événement à venir.</p>'
          : htmlContent;

      
      if (typeof appliquerTraductions === "function") appliquerTraductions();
    })
    .catch((err) => {
      console.error("Erreur :", err);
      container.innerHTML = '<p data-i18n="evenement.error">Erreur de connexion.</p>';
      if (typeof appliquerTraductions === "function") appliquerTraductions();
    });
}


function sinscrireEvenement(idEvent, prixEvent) {
  const idUser = localStorage.getItem("userId");
  const monToken = localStorage.getItem("token");

  if (!idUser || idUser === "null") {
    alert(t("evenement.login_required_sub"));
    return;
  }

  
  if (window.evenementData) {
    const currentEvt = window.evenementData.find((e) => e.id === idEvent);
    if (currentEvt && currentEvt.deja_inscrit) {
      alert(t("evenement.already_subscribed"));
      return;
    }
  }

  
  if (window.isProcessingPayment) return;
  window.isProcessingPayment = true;

  
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
          alert(t("evenement.payment_init_error"));
          window.isProcessingPayment = false;
        }
      })
      .catch(function (error) {
        console.error("Erreur Stripe :", error);
        alert(t("evenement.payment_server_error"));
        window.isProcessingPayment = false;
      });
  }

  
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
        alert(t("evenement.sub_success"));
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
    evt.image_url && evt.image_url.trim() !== ""
      ? `http://localhost:8081/${evt.image_url}`
      : "https://images.unsplash.com/photo-1542601906990-b4d3fb778b09?w=500";

  document.getElementById("modalImage").src = imageCover;
  document.getElementById("modalTitre").textContent = evt.titre;

  const auteur =
    `${evt.nomSalarie || ""} ${evt.prenomSalarie || ""}`.trim() ||
    "UpcycleConnect";

  const tFn = typeof t === "function" ? t : (k) => k;
  document.getElementById("modalMeta").textContent =
    `${tFn("evenement.hosted_by")} ${auteur} • ${tFn("evenement.on")} ${evt.date_debut}`;

  let contenuHtml =
    evt.description || "Pas de description disponible pour cet événement.";

  if (evt.pdf_url && evt.pdf_url.trim() !== "") {
    contenuHtml +=
      `<div style="margin-top: 18px;">` +
      `<a href="http://localhost:8081/${evt.pdf_url}" target="_blank" ` +
      `style="display:inline-block; background-color: var(--blue); color:#fff; padding:10px 16px; border-radius:6px; font-weight:600; text-decoration:none;">` +
      `📄 Télécharger le support (PDF)</a></div>`;
  }

  document.getElementById("modalContenu").innerHTML = contenuHtml;

  document.getElementById("articleModal").style.display = "flex";
}

function FermerEvenement() {
  const modal = document.getElementById("articleModal");
  
  if (modal) {
    modal.style.display = "none";
  }
}
function SeDesinscrire(idEvent) {
  if (!confirm(t("evenement.unsub_confirm"))) {
    return;
  }

  const idUser = localStorage.getItem("userId");
  const monToken = localStorage.getItem("token");

  if (!idUser || idUser === "null") {
    alert(t("evenement.login_required_action"));
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
