let mesActivites = [];
let dateAffichee = new Date(); // Par défaut : aujourd'hui

function initPlanningClient() {
  fetchActivitesClient();
}

function fetchActivitesClient() {
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("token");

  if (!userId) {
    console.error("Erreur : ID utilisateur introuvable dans le localStorage");
    return;
  }

  fetch("http://localhost:8081/user/planning?id=" + userId, {
    headers: {
      Authorization: "Bearer " + token,
    },
  })
    .then(function (response) {
      if (!response.ok) {
        throw new Error("Erreur HTTP : " + response.status);
      }
      return response.json();
    })
    .then(function (data) {
      mesActivites = data || [];
      console.log("Données reçues du serveur :", mesActivites);

      genererGrilleMois();
    })
    .catch(function (error) {
      console.error("Erreur lors du fetch :", error);
    });
}

window.changerMois = function (direction) {
  dateAffichee.setMonth(dateAffichee.getMonth() + direction);
  genererGrilleMois();
};

function genererGrilleMois() {
  const grid = document.querySelector(".cal-grid");
  const list = document.querySelector(".ev-list");
  if (!grid || !list) return;

  const moisNoms = [
    "Janvier",
    "Février",
    "Mars",
    "Avril",
    "Mai",
    "Juin",
    "Juillet",
    "Août",
    "Septembre",
    "Octobre",
    "Novembre",
    "Décembre",
  ];
  const joursShort = ["Dim", "Lun", "Mar", "Mer", "Jeu", "Ven", "Sam"];

  const moisIndex = dateAffichee.getMonth();
  const annee = dateAffichee.getFullYear();

  const labelMois = document.getElementById("labelCurrentMonth");
  if (labelMois) {
    labelMois.textContent = moisNoms[moisIndex] + " " + annee;
  }

  while (grid.children.length > 7) {
    grid.removeChild(grid.lastChild);
  }
  list.innerHTML = "";

  let premierJour = new Date(annee, moisIndex, 1).getDay();
  premierJour = premierJour === 0 ? 6 : premierJour - 1;
  const nbJoursMois = new Date(annee, moisIndex + 1, 0).getDate();

  for (let i = 0; i < premierJour; i++) {
    const divEmpty = document.createElement("div");
    divEmpty.className = "cal-d empty";
    grid.appendChild(divEmpty);
  }

  for (let j = 1; j <= nbJoursMois; j++) {
    const divJour = document.createElement("div");
    divJour.className = "cal-d";
    divJour.textContent = j;

    const jourStr = String(j).padStart(2, "0");
    const moisStr = String(moisIndex + 1).padStart(2, "0");
    const dateKey = jourStr + "-" + moisStr + "-" + annee;

    const activitesDuJour = mesActivites.filter(function (item) {
      return item.date === dateKey;
    });

    if (activitesDuJour.length > 0) {
      divJour.classList.add("has"); //point sur le calendrier

      activitesDuJour.forEach(function (ev) {
        let cssColor = "bl";
        let displayTitle = ev.titre;
        let metaTxt = ev.meta || "";

        if (ev.type === "evenement") {
          cssColor = "bl";
          metaTxt = "Inscrit";
        } else if (ev.type === "box_depot") {
          cssColor = "am"; // Orange
          displayTitle = "📦 À Déposer : " + ev.titre;
          metaTxt = ev.meta ? "Code : " + ev.meta : "Code à venir";
        } else if (ev.type === "box_retrait") {
          cssColor = "gr"; // Vert
          displayTitle = "✅ À Récupérer : " + ev.titre;
          metaTxt = ev.meta ? "Code : " + ev.meta : "Code à venir";
        }

        const dObj = new Date(annee, moisIndex, j);
        const nomJour = joursShort[dObj.getDay()];

        list.innerHTML += `
                    <div class="ev ${cssColor}">
                        <div class="ev-time">${nomJour} ${j}</div>
                        <div>
                            <div class="ev-title">${displayTitle}</div>
                            <div class="ev-meta">📍 ${ev.lieu} · <strong>${metaTxt}</strong></div>
                        </div>
                    </div>
                `;
      });
    }

    grid.appendChild(divJour);
  }

  if (list.innerHTML === "") {
    list.innerHTML =
      '<div style="color:var(--txt-m); text-align:center; padding: 20px;">Aucune activité ce mois-ci.</div>';
  }
}

document.addEventListener("DOMContentLoaded", initPlanningClient);
