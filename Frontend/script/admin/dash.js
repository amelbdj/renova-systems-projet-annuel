let tokenAdmin = localStorage.getItem("token");

function LoadDashboardData() {
  fetch("http://localhost:8081/admin/users", {
    headers: { Authorization: "Bearer " + tokenAdmin },
  })
    .then((res) => res.json())
    .then((users) => {
      const el = document.getElementById("dash-users");
      if (el && users) el.textContent = users.length;
    })
    .catch((err) => console.error("Erreur KPI Users:", err));

  fetch("http://localhost:8081/admin/annonces", {
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
            totalPoidsKg += parseFloat(a.poids_kg) || 0; // Addition pour le KPI des déchets
          } else if (
            a.statut_validation &&
            a.statut_validation.toLowerCase() === "en attente"
          ) {
            aModerer++; // On compte pour l'alerte
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
      if (elAlertAnn) elAlertAnn.textContent = `${aModerer} Annonces à modérer`;
    })
    .catch((err) => console.error("Erreur KPI Annonces:", err));

  fetch("http://localhost:8081/admin/evenements", {
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
            aValider++; // On compte
          }
        });
      }
      const elAlertEvt = document.getElementById("alert-events");
      if (elAlertEvt)
        elAlertEvt.textContent = `${aValider} Événements en attente`;
    })
    .catch((err) => console.error("Erreur KPI Evènements:", err));

  fetch("http://localhost:8081/api/admin/conteneurs", {
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

        // Astuce : On interroge l'API pour voir à l'intérieur de TOUS les conteneurs
        let fetchPromises = conteneurs.map((c) =>
          fetch(`http://localhost:8081/api/admin/conteneur/${c.id}/boxes`, {
            headers: { Authorization: "Bearer " + tokenAdmin },
          })
            .then((res) => res.json())
            .catch(() => []),
        );

        // Quand on a scanné tous les casiers de France...
        Promise.all(fetchPromises).then((results) => {
          let maintenanceCount = 0;
          results.forEach((boxArray) => {
            if (boxArray && boxArray.length) {
              boxArray.forEach((box) => {
                const etat = box.statut ? box.statut.toUpperCase() : "";
                if (
                  etat !== "LIBRE" &&
                  etat !== "OCCUPE" &&
                  etat !== "RESERVEE"
                ) {
                  maintenanceCount++;
                }
              });
            }
          });

          // Mise à jour de l'alerte
          if (elAlertMaint) {
            if (maintenanceCount === 0) {
              elAlertMaint.textContent = "0 Casier en maintenance";
              elAlertMaint.style.color = "var(--teal)"; // Vert si tout va bien !
            } else {
              elAlertMaint.textContent = `${maintenanceCount} Casiers en panne !`;
            }
          }
        });
      }
    })
    .catch((err) => console.error("Erreur KPI Conteneurs:", err));

  fetch("http://localhost:8081/admin/articles", {
    headers: { Authorization: "Bearer " + tokenAdmin },
  })
    .then((res) => res.json())
    .then((articles) => {
      let artAValider = 0;
      if (articles) {
        articles.forEach((art) => {
          // Attention, dans ton fichier articles.js, la variable s'appelle souvent "statut" tout court
          if (art.statut && art.statut.toLowerCase() === "en attente") {
            artAValider++;
          }
        });
      }

      const elAlertArt = document.getElementById("alert-articles");
      if (elAlertArt) {
        if (artAValider === 0) {
          elAlertArt.textContent = "0 Article en attente";
          elAlertArt.style.color = "var(--txt-m)"; // Gris si tout est à jour
        } else {
          elAlertArt.textContent = `${artAValider} Articles à modérer`;
          elAlertArt.style.color = "#2ecc71"; // Vert s'il y a de l'action
        }
      }
    })
    .catch((err) => console.error("Erreur KPI Articles:", err));
}

document.addEventListener("DOMContentLoaded", LoadDashboardData);
