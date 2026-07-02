let tokenAdmin = localStorage.getItem("token");

function LoadDashboardData() {
  fetch(`${API_BASE_URL}/admin/users`, {
    headers: { Authorization: "Bearer " + tokenAdmin },
  })
    .then((res) => res.json())
    .then((users) => {
      const el = document.getElementById("dash-users");
      if (el && users) el.textContent = users.length;
    })
    .catch((err) => console.error("Erreur KPI Users:", err));

  fetch(`${API_BASE_URL}/admin/annonces`, {
    headers: { Authorization: "Bearer " + tokenAdmin },
  })
    .then((res) => res.json())
    .then((annonces) => {
      let totalPoidsKg = 0;
      let aModerer = 0;

      if (annonces) {
        annonces.forEach((a) => {
          if (
            a.statut_validation &&
            a.statut_validation.toLowerCase() === "valide"
          ) {
            totalPoidsKg += parseFloat(a.poids_kg) || 0;

            totalPoidsKg += parseFloat(a.poids_kg) || 0;

            totalPoidsKg += parseFloat(a.poids_kg) || 0;
          } else if (
            a.statut_validation &&
            a.statut_validation.toLowerCase() === "en attente"
          ) {
            aModerer++;
          }
        });
      }

      const elWaste = document.getElementById("dash-waste");
      if (elWaste) {
        if (totalPoidsKg >= 1000)
          elWaste.textContent = (totalPoidsKg / 1000).toFixed(2) + " T";
        else elWaste.textContent = totalPoidsKg + " kg";
      }

      const elAlertAnn = document.getElementById("alert-annonces");
      if (elAlertAnn) elAlertAnn.textContent = `${aModerer} ${t("backoffice.kpi.ads_to_moderate")}`;

      AfficherActiviteRecente(annonces);
    })
    .catch((err) => console.error("Erreur KPI Annonces:", err));

  fetch(`${API_BASE_URL}/admin/evenements`, {
    headers: { Authorization: "Bearer " + tokenAdmin },
  })
    .then((res) => res.json())
    .then((events) => {
      let aValider = 0;
      if (events) {
        events.forEach((e) => {
          if (
            e.statut_validation &&
            e.statut_validation.toLowerCase() === "en attente"
          ) {
            aValider++;
          }
        });
      }
      const elAlertEvt = document.getElementById("alert-events");
      if (elAlertEvt)
        elAlertEvt.textContent = `${aValider} ${t("backoffice.kpi.events_pending")}`;
    })
    .catch((err) => console.error("Erreur KPI Evènements:", err));

  fetch(`${API_BASE_URL}/api/admin/conteneurs`, {
    headers: { Authorization: "Bearer " + tokenAdmin },
  })
    .then((res) => res.json())
    .then((conteneurs) => {
      const elBoxes = document.getElementById("dash-boxes");
      const elAlertMaint = document.getElementById("alert-maintenance");

      if (conteneurs) {
        let totalCasiers = 0;
        conteneurs.forEach((c) => (totalCasiers += c.total_boxes));

        if (elBoxes)
          elBoxes.textContent = `${conteneurs.length} / ${totalCasiers}`;

        let fetchPromises = conteneurs.map((c) =>
          fetch(`${API_BASE_URL}/api/admin/conteneur/${c.id}/boxes`, {
            headers: { Authorization: "Bearer " + tokenAdmin },
          })
            .then((res) => res.json())
            .catch(() => []),
        );

        Promise.all(fetchPromises).then((results) => {
          let maintenanceCount = 0;
          results.forEach((boxArray) => {
            if (boxArray && boxArray.length) {
              boxArray.forEach((box) => {
                const etat = box.statut ? box.statut.toUpperCase() : "";
                if (
                  etat !== "LIBRE" &&
                  !etat.startsWith("OCCUP") &&
                  !etat.startsWith("RESERV")
                ) {
                  maintenanceCount++;
                }
              });
            }
          });

          if (elAlertMaint) {
            if (maintenanceCount === 0) {
              elAlertMaint.textContent = `0 ${t("backoffice.kpi.lockers_maintenance")}`;
              elAlertMaint.style.color = "var(--teal)";
            } else {
              elAlertMaint.textContent = `${maintenanceCount} ${t("backoffice.kpi.lockers_down")}`;
            }
          }
        });
      }
    })
    .catch((err) => console.error("Erreur KPI Conteneurs:", err));

  fetch(`${API_BASE_URL}/admin/articles`, {
    headers: { Authorization: "Bearer " + tokenAdmin },
  })
    .then((res) => res.json())
    .then((articles) => {
      let artAValider = 0;
      if (articles) {
        articles.forEach((art) => {
          if (art.statut && art.statut.toLowerCase() === "en attente") {
            artAValider++;
          }
        });
      }

      const elAlertArt = document.getElementById("alert-articles");
      if (elAlertArt) {
        if (artAValider === 0) {
          elAlertArt.textContent = "0 Article en attente";
          elAlertArt.style.color = "var(--txt-m)";
        } else {
          elAlertArt.textContent = `${artAValider} ${t("backoffice.kpi.articles_to_moderate")}`;
          elAlertArt.style.color = "#2ecc71";
        }
      }
    })
    .catch((err) => console.error("Erreur KPI Articles:", err));

  // Revenus du mois (commission plateforme) via l'endpoint finances
  fetch(`${API_BASE_URL}/admin/finance/overview`, {
    headers: { Authorization: "Bearer " + tokenAdmin },
  })
    .then((res) => res.json())
    .then((data) => {
      const elRev = document.getElementById("rev-mois-montant");
      if (elRev) {
        elRev.textContent =
          new Intl.NumberFormat("fr-FR").format(data.revenuMois || 0) + " €";
      }
      // Pas de comparaison mois precedent disponible : on cache le "+...%"
      const elPct = document.getElementById("rev-mois-pourcentage");
      if (elPct) elPct.style.display = "none";
    })
    .catch((err) => console.error("Erreur KPI Revenus:", err));
}

// Activite recente : on affiche les 5 dernieres annonces publiees
function AfficherActiviteRecente(annonces) {
  const zone = document.getElementById("recent-activity");
  if (!zone) return;

  if (!annonces || annonces.length === 0) {
    zone.innerHTML = `<div style="padding:15px; color:var(--txt-m); font-size:13px;">Aucune activité récente.</div>`;
    return;
  }

  // Les plus recentes d'abord (id le plus grand = plus recent)
  const dernieres = annonces.slice().sort((a, b) => b.id - a.id).slice(0, 5);

  let html = "";
  dernieres.forEach((a) => {
    const gratuit = a.type === "Don" || a.prix <= 0;
    const montant = gratuit ? "Don" : a.prix + " €";
    html += `
      <div class="tx-row">
        <div class="tx-ico" style="background: rgba(74, 144, 240, 0.1)">
          <span class="material-symbols-outlined"> sell </span>
        </div>
        <div class="tx-info">
          <div class="tx-name">${a.titre}</div>
          <div class="tx-date">${a.prenom} ${a.nom} · ${a.created_at || ""}</div>
        </div>
        <div class="tx-amt ${gratuit ? "" : "in"}">${montant}</div>
      </div>`;
  });

  zone.innerHTML = html;
}

document.addEventListener("DOMContentLoaded", LoadDashboardData);
