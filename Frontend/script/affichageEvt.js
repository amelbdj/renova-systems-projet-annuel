function LancerRecherche() {
  const motCle = document.getElementById("searchInput").value.trim();
  chargerEvenementsClient(motCle);
}

function chargerEvenementsClient(motCle = "") {
  const container = document.getElementById("liste-evenements");
  container.innerHTML =
    "<p style='color: var(--txt-m); text-align: center; grid-column: 1 / -1;'>Chargement des événements...</p>";

  let url = `http://localhost:8081/admin/evenements`;

  if (motCle !== "") {
    url += `?search=${encodeURIComponent(motCle)}`;
  }

  fetch(url, {
    method: "GET",
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur réseau");
      return res.json();
    })
    .then((evenements) => {
      container.innerHTML = "";

      if (!evenements || evenements.length === 0) {
        if (motCle !== "") {
          container.innerHTML = `<p style="color: var(--txt-m); text-align: center; grid-column: 1 / -1;">Aucun résultat trouvé pour "<b>${motCle}</b>".</p>`;
        } else {
          container.innerHTML =
            "<p style='color: var(--txt-m); text-align: center; grid-column: 1 / -1;'>Aucun événement publié pour le moment.</p>";
        }
        return;
      }

      evenementData = evenements;
      let htmlContent = "";
      let evenementsAffiches = 0;

      evenements.forEach((evt) => {
        if (evt.statut_validation === "valide") {
          evenementsAffiches++;

          const idEvt = evt.id;
          const imageCover =
            evt.image_url ||
            "https://images.unsplash.com/photo-1542601906990-b4d3fb778b09?w=500"; // Image d'event par défaut

          const textContent =
            evt.description || evt.contenu || "Pas de description.";
          const resume =
            textContent.length > 100
              ? textContent.substring(0, 100) + "..."
              : textContent;

          htmlContent += `
                <div class="article-card" onclick="OuvrirEvenement(${idEvt})">
                    <img src="${imageCover}" alt="Image événement" class="card-img">
                    <div style="padding: 15px; flex-grow: 1; display: flex; flex-direction: column; align-items: center; text-align: center;">
                        <h3 style="margin-top: 0; color: #ffffff; font-size: 18px;">${evt.titre}</h3>
                        <p style="color: var(--blue-l); font-size: 13px; font-weight: bold; margin-bottom: 8px;">
                            <span class="material-symbols-outlined" style="font-size: 14px; vertical-align: middle;">calendar_today</span>  Le ${evt.date_debut}
                        </p>
                        <p style="color: var(--txt-m); font-size: 14px; flex-grow: 1;">${resume}</p>
                        
                        <span style="background-color: var(--blue); color: white; padding: 8px 15px; border-radius: 6px; font-weight: 600; font-size: 14px; margin-top: 10px; display: inline-block;" onclick="Sinscrire(${idEvt}); event.stopPropagation();">
                            S'inscrire ➔
                        </span>
                    </div>
                </div>
            `;
        }
      });

      if (evenementsAffiches === 0) {
        if (motCle !== "") {
          container.innerHTML = `<p style="color: var(--txt-m); text-align: center; grid-column: 1 / -1;">Aucun événement valide trouvé pour "<b>${motCle}</b>".</p>`;
        } else {
          container.innerHTML = `<p style="color: var(--txt-m); text-align: center; grid-column: 1 / -1;">Aucun événement valide pour le moment.</p>`;
        }
      } else {
        container.innerHTML = htmlContent;
      }
    })
    .catch((err) => {
      console.error("Erreur chargement événements :", err);
      container.innerHTML =
        "<p style='color: var(--red); text-align: center; grid-column: 1 / -1;'>Erreur de connexion au serveur.</p>";
    });
}

function OuvrirEvenement(id) {
  const evt = evenementData.find((a) => a.id === id);
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
  document.getElementById("articleModal").style.display = "none";
}

function Sinscrire(idEvent) {
  const btn = window.event ? window.event.target : null;
  const originalText = btn ? btn.innerHTML : "S'inscrire";

  if (btn) {
    btn.disabled = true;
    btn.innerHTML = "Vérification...";
  }

  const idUser = localStorage.getItem("userId");
  const monToken = localStorage.getItem("token");

  if (!idUser || idUser === "null") {
    alert("Attention : Vous devez être connecté pour vous inscrire.");
    if (btn) {
      btn.disabled = false;
      btn.innerHTML = originalText;
    }
    return;
  }

  fetch("http://localhost:8081/admin/evenements/inscription", {
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
          throw data.erreur || "Erreur d'inscription";
        }
        return data;
      });
    })
    .then(function (data) {
      alert("Succès : " + (data.message || "Inscription validée"));
      FermerEvenement();
      if (typeof chargerEvenementsClient === "function") {
        chargerEvenementsClient();
      }
    })
    .catch(function (errorMessage) {
      alert("Attention : " + errorMessage);
    })
    .finally(function () {
      if (btn) {
        btn.disabled = false;
        btn.innerHTML = originalText;
      }
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
