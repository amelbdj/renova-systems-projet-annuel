function GetAnnonce() {
  const container = document.getElementById("result");
  if (!container) return;

  fetch("http://localhost:8081/admin/annonces")
    .then((res) => {
      if (!res.ok) throw new Error("Erreur serveur annonces");
      return res.json();
    })
    .then((annonces) => {
      const currentTab = document.querySelector(".vtab.on").textContent;
      if (!currentTab.includes("Tout")) container.innerHTML = "";

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
                                    📦 Annonce · ${annonce.prenom} ${annonce.nom} · 
                                    Publiée le ${new Date(annonce.date_publication).toLocaleDateString()}
                                </div>
                                <div class="val-desc">${annonce.description}</div>
                                <div class="val-actions">
                                    <button class="va-btn va-ok" onclick="ValidateAnnonce(${annonce.id})">✓ Approuver</button>
                                    <button class="va-btn va-no" onclick="RefuseAnnonce(${annonce.id})">✕ Refuser</button>
                                </div>
                            </div>
                            <span class="tag t-or" style="flex-shrink: 0 font-size: 10px">Annonce</span>
                        </div>
                    </div>`;
        }
      });

      if (htmlContent) {
        container.innerHTML += htmlContent;
      } else if (!container.innerHTML) {
        container.innerHTML = `<div style="padding:20px">Aucune annonce en attente.</div>`;
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
