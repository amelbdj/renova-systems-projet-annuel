function GetAnnonce() {
  const container = document.getElementById("validations");

  fetch("http://localhost:8081/admin/annonces")
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur");
      return response.json();
    })
    .then((annonces) => {
      let htmlContent = "";

      annonces.forEach((annonce) => {
        if (
          annonce.statut_validation &&
          annonce.statut_validation.toLowerCase() === "en attente"
        ) {
          htmlContent += `
          <div class="val-list fu fu1">
            <div class="val-item ann" data-type="ann">
              <div class="val-body">
                <div class="val-title">${annonce.titre}</div>
                <div class="val-meta">
                  📦 Annonce · ${annonce.prenom} ${annonce.nom} · 
                  Publiée le ${new Date(annonce.date_publication).toLocaleDateString()}
                </div>
                <div class="val-desc">${annonce.description}</div>
                
                <div class="val-actions">
                  <button class="va-btn va-ok" onclick="ValidateAnnonce(${annonce.id})">
                    ✓ Approuver
                  </button>
                  <button class="va-btn va-no" onclick="RefuseAnnonce(${annonce.id})">
                    ✕ Refuser
                  </button>
                  <button class="va-btn va-view" onclick="showPreview(${annonce.id})">
                    👁 Aperçu
                  </button>
                </div>
              </div>
              <span class="tag t-or" style="flex-shrink: 0; font-size: 10px">Annonce</span>
            </div>
          </div>`;
        }
      });

      container.innerHTML +=
        htmlContent ||
        `<div style="padding:20px;">Aucune annonce en attente.</div>`;
    })
    .catch((error) => {
      console.error("Erreur API :", error);
      container.innerHTML = `<div style="padding: 20px; color: red;">Impossible de charger les données.</div>`;
    });
}

function ValidateAnnonce(id) {
  fetch(`http://localhost:8081/admin/annonces/validate/${id}`, {
    method: "PUT",
  })
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur");
      GetAnnonce();
    })
    .catch((error) => {
      console.error("Erreur API :", error);
    });
}

function RefuseAnnonce(id) {
  fetch(`http://localhost:8081/admin/annonces/refuse/${id}`, {
    method: "PUT",
  })
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur");
      GetAnnonce();
    })
    .catch((error) => {
      console.error("Erreur API :", error);
    });
}
function setVtab(element, type) {
  // 1. Gère l'apparence visuelle des onglets
  document
    .querySelectorAll(".vtab")
    .forEach((btn) => btn.classList.remove("on"));
  element.classList.add("on");

  // 2. Récupère le container et vide-le
  const container = document.getElementById("result");
  container.innerHTML = "<p style='padding:20px;'>Chargement...</p>";

  // 3. Appelle la fonction correspondante au type
  if (type === "ann") {
    GetAnnonce(); // Ta fonction pour les annonces
  } else if (type === "evt") {
    GetEvent(); // Ta fonction pour les événements
  } else if (type === "all") {
    // Tu peux choisir d'afficher les deux ou juste un message
    container.innerHTML = "";
    GetAnnonce();
    // Optionnel : GetEvent(); (attention au container.innerHTML = "" dans les fonctions)
  }
}
document.addEventListener("DOMContentLoaded", () => {
  GetAnnonce();
});
