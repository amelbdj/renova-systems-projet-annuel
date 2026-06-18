if (!monToken || !userId) {
  window.location.href = "../login.html";
}

let mesEvenements = [];
let dateAffichee = new Date();

function initPlanningSalarie() {
  fetchEvenementsSalarie();
}

function fetchEvenementsSalarie() {
  fetch("http://localhost:8081/admin/evenements", {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then(function (response) {
      if (!response.ok) {
        throw new Error("Erreur HTTP : " + response.status);
      }
      return response.json();
    })
    .then(function (data) {
      if (!data) data = [];

      mesEvenements = [];
      for (let i = 0; i < data.length; i++) {
        let ev = data[i];
        let statut = ev.statut_validation
          ? ev.statut_validation.toLowerCase()
          : "";
        let aMoi = ev.idSalarie == userId || ev.id_salarie == userId;
        if ((statut === "valide" || statut === "en ligne") && aMoi) {
          mesEvenements.push(ev);
        }
      }

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

function getJourEvenement(ev) {
  let datePart = ev.date_debut || "";
  if (datePart.indexOf(" a ") !== -1) {
    datePart = datePart.split(" a ")[0];
  }
  return datePart;
}

function getHeureEvenement(ev) {
  let s = ev.date_debut || "";
  if (s.indexOf(" a ") !== -1) {
    return s.split(" a ")[1];
  }
  return "";
}

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
    const dateKey = jourStr + "/" + moisStr + "/" + annee;

    const evenementsDuJour = mesEvenements.filter(function (ev) {
      return getJourEvenement(ev) === dateKey;
    });

    if (evenementsDuJour.length > 0) {
      divJour.classList.add("has");

      evenementsDuJour.forEach(function (ev) {
        let cssColor = "bl";
        let typeTxt = (ev.type || "").toLowerCase();
        if (typeTxt === "formation") cssColor = "bl";
        else if (typeTxt === "atelier") cssColor = "gr";
        else if (typeTxt === "reunion") cssColor = "am";

        const heure = getHeureEvenement(ev);
        const lieuTxt = ev.lieu ? ev.lieu : "Lieu non précisé";
        const dObj = new Date(annee, moisIndex, j);
        const nomJour = joursShort[dObj.getDay()];

        list.innerHTML += `
                    <div class="ev ${cssColor}" onclick="ouvrirDetailEvent(${ev.id})" style="cursor: pointer;">
                        <div class="ev-time">${nomJour} ${j}<br>${heure}</div>
                        <div>
                            <div class="ev-title">${ev.titre}</div>
                            <div class="ev-meta">📍 ${lieuTxt} · <strong>${ev.type || "Événement"}</strong></div>
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

window.ouvrirDetailEvent = function (id) {
  let evt = null;
  for (let i = 0; i < mesEvenements.length; i++) {
    if (mesEvenements[i].id == id) {
      evt = mesEvenements[i];
    }
  }
  if (!evt) return;

  document.getElementById("detail-titre").textContent = evt.titre;

  let html = "";
  html += "<p><strong>Type :</strong> " + (evt.type || "Événement") + "</p>";
  html +=
    "<p><strong>Statut :</strong> " + (evt.statut_validation || "") + "</p>";
  html += "<p><strong>Début :</strong> " + (evt.date_debut || "") + "</p>";
  html += "<p><strong>Fin :</strong> " + (evt.date_fin || "") + "</p>";
  html +=
    "<p><strong>Lieu :</strong> " +
    (evt.lieu ? evt.lieu : "Non précisé") +
    "</p>";
  html += "<p><strong>Places :</strong> " + (evt.nb_places || 0) + "</p>";
  html +=
    "<p><strong>Tarif :</strong> " +
    (evt.prix ? evt.prix + " €" : "Gratuit") +
    "</p>";
  html +=
    "<p style='margin-top:12px;'>" +
    (evt.description || "Pas de description.") +
    "</p>";

  if (evt.pdf_url && evt.pdf_url !== "") {
    html +=
      "<p style='margin-top:12px;'><a href='http://localhost:8081/" +
      evt.pdf_url +
      "' target='_blank' style='color:#fff; background:var(--vi); padding:8px 14px; border-radius:6px; text-decoration:none;'>📄 Télécharger le support PDF</a></p>";
  }

  html +=
    "<div style='margin-top:18px; border-top:1px solid var(--b0); padding-top:14px;'>" +
    "<strong>Personnes inscrites</strong>" +
    "<div id='detail-inscrits' style='margin-top:8px;'>Chargement...</div>" +
    "</div>";

  document.getElementById("detail-body").innerHTML = html;
  document.getElementById("detailModal").style.display = "flex";

  chargerInscrits(evt.id);
};

function chargerInscrits(id) {
  fetch("http://localhost:8081/admin/evenements/inscrits/" + id, {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then(function (response) {
      return response.json();
    })
    .then(function (inscrits) {
      let zone = document.getElementById("detail-inscrits");
      if (!zone) return;

      if (!inscrits || inscrits.length === 0) {
        zone.innerHTML =
          "<p style='color:var(--txt-d);'>Aucune personne inscrite pour le moment.</p>";
        return;
      }

      let liste = "<ul style='margin:0; padding-left:18px;'>";
      for (let i = 0; i < inscrits.length; i++) {
        liste +=
          "<li>" +
          inscrits[i].prenom +
          " " +
          inscrits[i].nom +
          " — " +
          inscrits[i].email +
          "</li>";
      }
      liste += "</ul>";
      zone.innerHTML = liste;
    })
    .catch(function (error) {
      let zone = document.getElementById("detail-inscrits");
      if (zone) {
        zone.innerHTML =
          "<p style='color:var(--txt-d);'>Erreur lors du chargement des inscrits.</p>";
      }
      console.error(error);
    });
}

window.fermerDetailEvent = function () {
  document.getElementById("detailModal").style.display = "none";
};

document.addEventListener("DOMContentLoaded", initPlanningSalarie);
