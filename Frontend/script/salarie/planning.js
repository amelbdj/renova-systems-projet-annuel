let tousMesEvenements = [];
let dateAffichee = new Date(2026, 3, 1);
let dateSelectionnee = "";

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
      tousMesEvenements = evenements.filter((ev) => {
        const statutEv = ev.statut_validation
          ? ev.statut_validation.toLowerCase()
          : "";
        return statutEv === "valide" && ev.idSalarie == userId;
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
  const joursNoms = ["Lun", "Mar", "Mer", "Jeu", "Ven", "Sam", "Dim"];

  const annee = dateAffichee.getFullYear();
  const moisIndex = dateAffichee.getMonth();
  const moisStr = String(moisIndex + 1).padStart(2, "0");

  const titreMois = document.querySelector(".week-month");
  if (titreMois) titreMois.textContent = `${moisNoms[moisIndex]} ${annee}`;

  let datesDuMois = new Set();

  tousMesEvenements.forEach((ev) => {
    const dateEv = ev.date_debut ? ev.date_debut.substring(0, 10) : "";

    if (dateEv.startsWith(`${annee}-${moisStr}`)) {
      datesDuMois.add(dateEv);
    }
  });

  const datesTriees = Array.from(datesDuMois).sort();

  const conteneurJours = document.querySelector(".week-days");
  if (!conteneurJours) return;
  conteneurJours.innerHTML = "";

  // Cas où il n'y a aucun événement validé ce mois-ci
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
    const dateEv = ev.date_debut ? ev.date_debut.substring(0, 10) : "";
    return dateEv === dateCible;
  });

  eventsDuJour.sort((a, b) => {
    const hA = a.date_debut || "";
    const hB = b.date_debut || "";
    return hA.localeCompare(hB);
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

    const heureDebut = ev.date_debut
      ? ev.date_debut.substring(11, 16)
      : "00:00";
    const heureFin = ev.date_fin ? ev.date_fin.substring(11, 16) : "00:00";

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
