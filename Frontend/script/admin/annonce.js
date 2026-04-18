const Auth = {
  getToken: () => localStorage.getItem("token"),
  getUserId: () => localStorage.getItem("userId"),
  getUserRole: () => localStorage.getItem("userRole"),
  isConnected: () => !!localStorage.getItem("token"),

  // Pour se déconnecter proprement
  logout: () => {
    localStorage.clear();
    window.location.href = "login.html";
  },
};

function GetAnnonce() {
  const container = document.getElementById("result");
  if (!container) return;

  fetch("http://localhost:8081/admin/annonces", {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
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
                <div style="display:flex; justify-content:space-between; align-items:start;">
                    <div class="val-title">${annonce.titre}</div>
                    <span class="tag ${annonce.type === "Don" ? "t-grn" : "t-blue"}">
                        ${annonce.type === "Don" ? t("backoffice.ads.free") : annonce.prix + " €"}
                    </span>
                </div>

                <div class="val-meta">
                    <span class="material-symbols-outlined" style="font-size:14px; vertical-align:middle;">category</span> 
                    <b>${annonce.categorie}</b> · 
                    ${annonce.prenom} ${annonce.nom} · 
                    <span data-i18n="backoffice.ads.published_on">${t("backoffice.ads.published_on")}</span> ${new Date(annonce.date_publication).toLocaleDateString()}
                </div>

                <div class="val-desc">${annonce.description}</div>

                <div style="margin-top:10px; font-size:11px; color:var(--txt-m); display:flex; gap:15px;">
                    <span><span class="material-symbols-outlined" style="font-size:12px;">scale</span> ${annonce.poids_kg} kg</span>
                    <span><span class="material-symbols-outlined" style="font-size:12px;">location_on</span> ${annonce.ville} (${annonce.code_postal})</span>
                    <span><span class="material-symbols-outlined" style="font-size:12px;">inventory_2</span> ${t("backoffice.ads.condition")}: ${annonce.etat}</span>
                </div>

                <div class="val-actions" style="margin-top:15px;">
                    <button class="va-btn va-ok" onclick="ValidateAnnonce(${annonce.id})">
                        <span class="material-symbols-outlined" style="font-size:18px;">check_circle</span>
                        <span data-i18n="backoffice.btn.approve">${t("backoffice.btn.approve")}</span>
                    </button>
                    <button class="va-btn va-no" onclick="RefuseAnnonce(${annonce.id})">
                        <span class="material-symbols-outlined" style="font-size:18px;">cancel</span>
                        <span data-i18n="backoffice.btn.refuse">${t("backoffice.btn.refuse")}</span>
                    </button>
                </div>
            </div>
            <span class="tag t-or" style="flex-shrink: 0; font-size: 10px" data-i18n="backoffice.ads.ad_tag">${t("backoffice.ads.ad_tag")}</span>
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
    headers: {
      Authorization: "Bearer " + monToken,
    },
  }).then(() => {
    GetAnnonce();
    UpdateValidationCount(); // Met à jour le total rouge en haut
  });
}

function RefuseAnnonce(id) {
  fetch(`http://localhost:8081/admin/annonces/refuse/${id}`, {
    method: "PUT",
    headers: {
      Authorization: "Bearer " + monToken,
    },
  }).then(() => {
    GetAnnonce();
    UpdateValidationCount();
  });
}

document.addEventListener("DOMContentLoaded", () => {
  GetAnnonce();
});
