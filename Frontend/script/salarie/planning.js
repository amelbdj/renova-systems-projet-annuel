let tousMesEvenements = [];
let dateAffichee = new Date(2026, 2, 1); // Mars 2026 (les mois JS commencent à 0)
let dateSelectionnee = "";

function parseDateSql(dateStr) {
  if (!dateStr) return { isoDate: "", time: "00:00" };

  let datePart = dateStr;
  let timePart = "00:00";

  if (dateStr.includes(" a ")) {
    const parts = dateStr.split(" a ");
    datePart = parts[0]; // "20/03/2026"
    timePart = parts[1]; // "14:00"
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
      dateAffichee.setMonth(dateAffichee.getMonth() - 1);
      genererJoursAvecEvenements();
    });
  }

  if (btnNext) {
    btnNext.addEventListener("click", () => {
      dateAffichee.setMonth(dateAffichee.getMonth() + 1);
      genererJoursAvecEvenements();
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

      genererJoursAvecEvenements();
    })
    .catch((err) => console.error("Erreur chargement:", err));
}

function genererJoursAvecEvenements() {
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
  const joursNoms = ["Dim", "Lun", "Mar", "Mer", "Jeu", "Ven", "Sam"]; // Dimanche est 0 en JS

  const annee = dateAffichee.getFullYear();
  const moisIndex = dateAffichee.getMonth();
  const moisStr = String(moisIndex + 1).padStart(2, "0");

  const titreMois = document.querySelector(".week-month");
  if (titreMois) titreMois.textContent = `${moisNoms[moisIndex]} ${annee}`;

  let datesDuMois = new Set();

  tousMesEvenements.forEach((ev) => {
    const dateEv = parseDateSql(ev.date_debut).isoDate;

    if (dateEv.startsWith(`${annee}-${moisStr}`)) {
      datesDuMois.add(dateEv);
    }
  });

  const datesTriees = Array.from(datesDuMois).sort();

  const conteneurJours = document.querySelector(".week-days");
  if (!conteneurJours) return;
  conteneurJours.innerHTML = "";

  if (datesTriees.length === 0) {
    conteneurJours.innerHTML = `<div style="color:var(--txt-d); padding:10px 20px; font-size:14px;">Aucun événement prévu en ${moisNoms[moisIndex]}.</div>`;

    const timelineContainer = document.querySelector(".timeline");
    if (timelineContainer) timelineContainer.innerHTML = "";

    const titreTimeline = document.querySelector(".card-ht");
    if (titreTimeline)
      titreTimeline.innerHTML = `<div class="cht-i" style="background: rgba(240, 106, 170, 0.1)">🗓️</div> Mois vide`;

    return;
  }

  if (!datesTriees.includes(dateSelectionnee)) {
    dateSelectionnee = datesTriees[0];
  }

  datesTriees.forEach((dateStr) => {
    const anneeJour = parseInt(dateStr.substring(0, 4));
    const moisJour = parseInt(dateStr.substring(5, 7)) - 1;
    const jourDuMois = parseInt(dateStr.substring(8, 10));

    const dateObj = new Date(anneeJour, moisJour, jourDuMois);
    const nomJour = joursNoms[dateObj.getDay()];

    const jourDiv = document.createElement("div");
    jourDiv.className = "wday";

    let dotHtml = `<div class="wdot" style="background: var(--vi)"></div>`;

    if (dateStr === dateSelectionnee) {
      jourDiv.classList.add("today");
    }

    jourDiv.innerHTML = `
        <div class="wday-num">${jourDuMois}</div>
        <div class="wday-name">${nomJour}</div>
        <div class="wday-dots">${dotHtml}</div>
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
            ${nomJour} ${jourDuMois} ${moisNoms[moisIndex]} ${annee}
        `;
      }

      afficherTimeline(dateSelectionnee);
    });

    conteneurJours.appendChild(jourDiv);
  });

  const premierJourCree = document.querySelector(".wday.today");
  if (premierJourCree) {
    premierJourCree.click();
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

  timelineContainer.innerHTML += `
    <div class="tl-hour">
        <div class="tl-time">17:00</div>
        <div class="tl-col">
            <div style="font-size: 12px; color: var(--txt-d); padding: 8px 0;">Fin de journée</div>
        </div>
    </div>
  `;
}

document.addEventListener("DOMContentLoaded", initPlanning);
