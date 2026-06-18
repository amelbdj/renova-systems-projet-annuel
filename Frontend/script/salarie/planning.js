if (!monToken || !userId) {
  window.location.href = "../login.html";
}

let tousMesEvenements = [];
let lundiAffiche = getLundi(new Date());
let dateSelectionnee = "";

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
const joursNoms = ["Dim", "Lun", "Mar", "Mer", "Jeu", "Ven", "Sam"];
const moisCourts = [
  "Jan",
  "Fév",
  "Mars",
  "Avr",
  "Mai",
  "Juin",
  "Juil",
  "Août",
  "Sept",
  "Oct",
  "Nov",
  "Déc",
];

function getLundi(d) {
  const date = new Date(d.getFullYear(), d.getMonth(), d.getDate());
  const decalage = (date.getDay() + 6) % 7;
  date.setDate(date.getDate() - decalage);
  return date;
}

function formatDateIso(d) {
  const annee = d.getFullYear();
  const mois = String(d.getMonth() + 1).padStart(2, "0");
  const jour = String(d.getDate()).padStart(2, "0");
  return `${annee}-${mois}-${jour}`;
}

function parseDateSql(dateStr) {
  if (!dateStr) return { isoDate: "", time: "00:00" };

  let datePart = dateStr;
  let timePart = "00:00";

  if (dateStr.includes(" a ")) {
    const parts = dateStr.split(" a ");
    datePart = parts[0];
    timePart = parts[1];
  }

  if (datePart.includes("/")) {
    const d = datePart.split("/");
    if (d.length === 3) {
      return { isoDate: `${d[2]}-${d[1]}-${d[0]}`, time: timePart };
    }
  }

  return { isoDate: datePart.substring(0, 10), time: timePart };
}

function initPlanning() {
  const btnPrev = document.getElementById("btn-prev-month");
  const btnNext = document.getElementById("btn-next-month");

  if (btnPrev) {
    btnPrev.addEventListener("click", () => {
      lundiAffiche.setDate(lundiAffiche.getDate() - 7);
      genererSemaine();
    });
  }

  if (btnNext) {
    btnNext.addEventListener("click", () => {
      lundiAffiche.setDate(lundiAffiche.getDate() + 7);
      genererSemaine();
    });
  }

  fetchEvenements();
}

function fetchEvenements() {
  fetch(`http://localhost:8081/admin/evenements`, {
    headers: { Authorization: "Bearer " + monToken },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur réseau API");
      return res.json();
    })
    .then((evenements) => {
      if (!evenements) evenements = [];

      tousMesEvenements = evenements.filter((ev) => {
        const statutEv = ev.statut_validation
          ? ev.statut_validation.toLowerCase()
          : "";
        return (
          (statutEv === "valide" || statutEv === "en ligne") &&
          (ev.idSalarie == userId || ev.id_salarie == userId)
        );
      });

      genererSemaine();
    })
    .catch((err) => console.error("Erreur chargement:", err));
}

function genererSemaine() {
  const conteneurJours = document.querySelector(".week-days");
  if (!conteneurJours) return;
  conteneurJours.innerHTML = "";

  const jours = [];
  for (let i = 0; i < 7; i++) {
    const d = new Date(lundiAffiche);
    d.setDate(lundiAffiche.getDate() + i);
    jours.push(d);
  }

  const dernierJour = jours[6];
  const titreMois = document.querySelector(".week-month");
  if (titreMois) {
    if (lundiAffiche.getMonth() === dernierJour.getMonth()) {
      titreMois.textContent = `${lundiAffiche.getDate()} – ${dernierJour.getDate()} ${moisCourts[dernierJour.getMonth()]}`;
    } else {
      titreMois.textContent = `${lundiAffiche.getDate()} ${moisCourts[lundiAffiche.getMonth()]} – ${dernierJour.getDate()} ${moisCourts[dernierJour.getMonth()]}`;
    }
  }

  const isoSemaine = jours.map((d) => formatDateIso(d));
  const todayIso = formatDateIso(new Date());

  if (!isoSemaine.includes(dateSelectionnee)) {
    if (isoSemaine.includes(todayIso)) {
      dateSelectionnee = todayIso;
    } else {
      const premierAvecEvt = isoSemaine.find((iso) =>
        tousMesEvenements.some(
          (ev) => parseDateSql(ev.date_debut).isoDate === iso,
        ),
      );
      dateSelectionnee = premierAvecEvt || isoSemaine[0];
    }
  }

  jours.forEach((dateObj) => {
    const dateStr = formatDateIso(dateObj);
    const nomJour = joursNoms[dateObj.getDay()];
    const jourDuMois = dateObj.getDate();

    const nbEvents = tousMesEvenements.filter(
      (ev) => parseDateSql(ev.date_debut).isoDate === dateStr,
    ).length;

    const jourDiv = document.createElement("div");
    jourDiv.className = "wday";
    if (dateStr === dateSelectionnee) jourDiv.classList.add("today");

    let dotsHtml = "";
    for (let k = 0; k < Math.min(nbEvents, 3); k++) {
      dotsHtml += `<div class="wdot" style="background: var(--vi)"></div>`;
    }

    jourDiv.innerHTML = `
        <div class="wday-num">${jourDuMois}</div>
        <div class="wday-name">${nomJour}</div>
        <div class="wday-dots">${dotsHtml}</div>
    `;

    jourDiv.addEventListener("click", () => {
      document
        .querySelectorAll(".wday")
        .forEach((j) => j.classList.remove("today"));
      jourDiv.classList.add("today");

      dateSelectionnee = dateStr;

      const titreTimeline = document.querySelector(".card-ht");
      if (titreTimeline) {
        titreTimeline.innerHTML = `
            <div class="cht-i" style="background: rgba(240, 106, 170, 0.1)">🗓️</div>
            ${nomJour} ${jourDuMois} ${moisNoms[dateObj.getMonth()]} ${dateObj.getFullYear()}
        `;
      }

      afficherTimeline(dateSelectionnee);
    });

    conteneurJours.appendChild(jourDiv);
  });

  const jourSelectionne = document.querySelector(".wday.today");
  if (jourSelectionne) {
    jourSelectionne.click();
  }
}

function afficherTimeline(dateCible) {
  const eventsDuJour = tousMesEvenements.filter((ev) => {
    const dateEv = parseDateSql(ev.date_debut).isoDate;
    return dateEv === dateCible;
  });

  eventsDuJour.sort((a, b) => {
    const timeA = parseDateSql(a.date_debut).time;
    const timeB = parseDateSql(b.date_debut).time;
    return timeA.localeCompare(timeB);
  });

  const timelineContainer = document.querySelector(".timeline");
  if (!timelineContainer) return;
  timelineContainer.innerHTML = "";

  if (eventsDuJour.length === 0) {
    timelineContainer.innerHTML = `<div style="color:var(--txt-d); padding:10px 0; font-size:14px;">Aucun événement ce jour-là.</div>`;
    return;
  }

  eventsDuJour.forEach((ev) => {
    const typeStr = ((ev.type || "") + " " + (ev.titre || "")).toLowerCase();
    let colorClass = "vi";

    if (typeStr.includes("formation")) colorClass = "vi";
    else if (typeStr.includes("réunion") || typeStr.includes("reunion"))
      colorClass = "pink";
    else if (typeStr.includes("atelier")) colorClass = "teal";
    else if (typeStr.includes("admin")) colorClass = "amber";

    const lieuTxt = ev.lieu ? ` · ${ev.lieu}` : "";

    const heureDebut = parseDateSql(ev.date_debut).time;
    const heureFin = parseDateSql(ev.date_fin).time;

    const blocHTML = `
        <div class="tl-hour">
            <div class="tl-time">${heureDebut}</div>
            <div class="tl-col">
                <div class="tl-event tl-ev ${colorClass}">
                    <div class="tl-ev-title">${ev.titre}</div>
                    <div class="tl-ev-meta">
                        ${heureDebut} – ${heureFin}${lieuTxt}
                    </div>
                </div>
            </div>
        </div>
    `;
    timelineContainer.innerHTML += blocHTML;
  });
}

document.addEventListener("DOMContentLoaded", initPlanning);
