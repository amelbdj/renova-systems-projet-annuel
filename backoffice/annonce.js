function GetAnnonce() {
  const container = document.getElementById("result");
  if (!container) return;

  fetch("http://localhost:8081/admin/annonces")
    .then((res) => {
      if (!res.ok) throw new Error("Erreur serveur annonces");
      return res.json();
    })
    .then((annonces) => {
      // On vérifie l'ID et non le texte (qui va changer avec la trad)
      const currentTabId = document.querySelector(".vtab.on").id;
      if (currentTabId !== "tout") container.innerHTML = "";

      let htmlContent = "";
      annonces.forEach((annonce) => {
        if (
          annonce.statut_validation &&
          annonce.statut_validation.toLowerCase() === "en attente"
        ) {
          htmlContent += `
                    <div class="val-list fu fu1">
                        <div class="val-item ann">
                            <div class="val-body">
                                <div class="val-title">${annonce.titre}</div>
                                <div class="val-meta">
                                    📦 <span data-i18n="backoffice.ads.ad">Annonce</span> · ${annonce.prenom} ${annonce.nom} · 
                                    <span data-i18n="backoffice.ads.published_on">Publiée le</span> ${new Date(annonce.date_publication).toLocaleDateString()}
                                </div>
                                <div class="val-desc">${annonce.description}</div>
                                <div class="val-actions">
                                    <button class="va-btn va-ok" onclick="ValidateAnnonce(${annonce.id})" data-i18n="backoffice.btn.approve">✓ Approuver</button>
                                    <button class="va-btn va-no" onclick="RefuseAnnonce(${annonce.id})" data-i18n="backoffice.btn.refuse">✕ Refuser</button>
                                </div>
                            </div>
                            <span class="tag t-or" style="flex-shrink: 0; font-size: 10px" data-i18n="backoffice.ads.ad_tag">Annonce</span>
                        </div>
                    </div>`;
        }
      });

      if (htmlContent) {
        container.innerHTML += htmlContent;
      } else if (!container.innerHTML) {
        container.innerHTML = `<div style="padding:20px" data-i18n="backoffice.ads.no_ads">Aucune annonce en attente.</div>`;
      }

      // On demande au script de traduire les nouveaux éléments fraîchement injectés
      if (typeof appliquerTraductions === "function") {
        appliquerTraductions();
      }
    })
    .catch((err) => console.error(err));
}

function ValidateAnnonce(id) {
  fetch(`http://localhost:8081/admin/annonces/validate/${id}`, {
    method: "PUT",
  }).then(() => {
    GetAnnonce();
    UpdateValidationCount(); // Met à jour le total rouge en haut
  });
}

function RefuseAnnonce(id) {
  fetch(`http://localhost:8081/admin/annonces/refuse/${id}`, {
    method: "PUT",
  }).then(() => {
    GetAnnonce();
    UpdateValidationCount();
  });
}

document.addEventListener("DOMContentLoaded", () => {
  GetAnnonce();
});
