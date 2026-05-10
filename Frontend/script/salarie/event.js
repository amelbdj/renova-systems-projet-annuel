if (!monToken || !userId) {
  alert("Vous devez être connecté pour accéder à cette page.");
  window.location.href = "login.html";
}

// ==========================================
// 1. CRÉATION D'UN ÉVÉNEMENT
// ==========================================
function CreateEvent() {
  const titre = document.getElementById("evt-titre").value.trim();
  const type = document.getElementById("evt-type").value;
  const desc = document.getElementById("evt-desc").value.trim();
  const date = document.getElementById("evt-date").value;
  const lieu = document.getElementById("evt-lieu").value.trim();
  const heureDebut = document.getElementById("evt-heure-debut").value;
  const heureFin = document.getElementById("evt-heure-fin").value;
  const capacite = document.getElementById("evt-capacite").value;
  const tarif = document.getElementById("evt-tarif").value;

  const datetimeDebut = `${date} ${heureDebut || "00:00"}:00`;
  const datetimeFin = `${date} ${heureFin || "00:00"}:00`;

  if (!titre || !desc || !date) {
    alert(
      "Veuillez remplir les champs obligatoires : Titre, Description et Date.",
    );
    return;
  }

  // Vérification que la date n'est pas dans le passé
  if (date) {
    const selectedDate = new Date(date);
    const today = new Date();
    today.setHours(0, 0, 0, 0);

    if (selectedDate < today) {
      alert("Erreur : La date de l'événement ne peut pas être dans le passé.");
      return;
    }
  }

  const eventData = {
    idSalarie: parseInt(userId),
    titre: titre,
    type: type,
    description: desc,
    date_debut: datetimeDebut,
    date_fin: datetimeFin,
    lieu: lieu,
    capacite: parseInt(capacite) || 0,
    tarif: parseFloat(tarif) || 0,
  };

  fetch("http://localhost:8081/admin/evenements/add", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
    body: JSON.stringify(eventData),
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur lors de la création");
      return res.text();
    })
    .then(() => {
      alert("Événement soumis avec succès. Il est en attente de validation.");
      closeEvt();
      resetEvtForm();
      GetEvenements();
    })
    .catch((err) => {
      console.error(err);
      alert("Une erreur est survenue.");
    });
}

// ==========================================
// 2. LECTURE DES ÉVÉNEMENTS (GRILLE & KPI)
// ==========================================
function GetEvenements() {
  const container = document.getElementById("event-grid");
  if (!container) return;

  let counterEvt = 0;
  let counterValide = 0;

  fetch(`http://localhost:8081/admin/evenements`, {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur serveur événements");
      return res.json();
    })
    .then((evenements) => {
      let htmlContent = "";

      // Sécurité si la BDD renvoie null
      if (!evenements) evenements = [];

      // On filtre pour ne garder que les événements de ce salarié
      const mesEvenements = evenements.filter(
        (e) =>
          e.id_salarie == userId ||
          e.IdSalarie == userId ||
          e.idSalarie == userId,
      );

      const statEvent = document.getElementById("stat-event");
      const statAttente = document.getElementById("stat-valide");

      if (mesEvenements.length === 0) {
        container.innerHTML = `
          <div style="color:var(--txt-m); padding:20px;">Vous n'avez créé aucun événement pour le moment.</div>
          <div class="evt-add" onclick="openNewEvt()"><div class="plus">＋</div><span>Créer un événement</span></div>`;

        // On met les compteurs à 0
        if (statEvent) statEvent.textContent = 0;
        if (statAttente) statAttente.textContent = 0;
        return;
      }

      mesEvenements.forEach((evt) => {
        let statusBadge = "";
        let actionButtons = "";
        counterEvt++; // +1 événement total

        const statut = evt.statut_validation
          ? evt.statut_validation.toLowerCase()
          : "en attente";

        if (statut === "valide" || statut === "en ligne") {
          counterValide++; // +1 événement validé
          statusBadge = `<div class="evt-status t-green">✓ En ligne</div>`;
          actionButtons = `
            <span class="tag t-green">Publiée</span>
            <div style="display:flex;gap:5px">
              <button class="btn btn-g btn-xs" onclick="editerEvenement(${evt.id})">Modifier</button>
              <button class="btn btn-danger btn-xs" onclick="DeleteEvenement(${evt.id})">Annuler</button>
            </div>`;
        } else {
          statusBadge = `<div class="evt-status t-amber">⏳ En attente</div>`;
          actionButtons = `
            <span class="tag t-amber">Validation en cours</span>
            <button class="btn btn-g btn-xs" onclick="editerEvenement(${evt.id})">Modifier</button>
            <button class="btn btn-danger btn-xs" onclick="DeleteEvenement(${evt.id})">Annuler</button>`;
        }

        // FORMATAGE DE LA DATE (Sépare la date de l'heure du format SQL)
        let dateFormatee = "Date inconnue";
        let heureFormatee = "";

        if (evt.date_debut) {
          if (evt.date_debut.includes(" a ")) {
            const parts = evt.date_debut.split(" a ");
            dateFormatee = parts[0];
            heureFormatee = parts[1].replace(":", "h");
          } else {
            dateFormatee = evt.date_debut;
          }
        }

        // FORMATAGE DU PRIX
        const prixEvt = parseFloat(evt.prix || evt.Prix) || 0;
        let prixAffichage = "";
        let tagPrixClass = "";

        if (prixEvt > 0) {
          prixAffichage = prixEvt.toFixed(2) + " €";
          tagPrixClass = "t-amber";
        } else {
          prixAffichage = "Gratuit";
          tagPrixClass = "t-green";
        }

        htmlContent += `
        <div class="evt-card">
          <div class="evt-banner" style="background:linear-gradient(135deg,#100820,#1c1040)">
            📅 
            <div class="evt-type-badge etb-formation">${evt.type || evt.format || "Événement"}</div>
            ${statusBadge}
          </div>
          <div class="evt-body">
            <div class="evt-name">${evt.titre}</div>
            <div class="evt-desc">${evt.description}</div>
            <div class="evt-meta">
              <span class="tag t-vi">📅 ${dateFormatee}</span>
              ${heureFormatee ? `<span class="tag t-blue">⏰ ${heureFormatee}</span>` : ""}
              <span class="tag" style="color:var(--txt-m);background:var(--bg3);border:1px solid var(--b0)">👥 Max ${evt.nb_places || evt.capacite || 0}</span>
              <span class="tag ${tagPrixClass}">💶 ${prixAffichage}</span>
            </div>
            <div class="evt-foot">
              ${actionButtons}
            </div>
          </div>
        </div>`;
      });

      htmlContent += `
      <div class="evt-add" onclick="openNewEvt()">
        <div class="plus">＋</div>
        <span>Créer un événement ou une formation</span>
      </div>`;

      container.innerHTML = htmlContent;

      // On met à jour les KPI dans le HTML
      if (statEvent) statEvent.textContent = counterEvt;
      if (statAttente) statAttente.textContent = counterValide;
    })
    .catch((err) => console.error(err));
}

// ==========================================
// 3. SUPPRESSION D'UN ÉVÉNEMENT
// ==========================================
function DeleteEvenement(id) {
  if (!confirm("Êtes-vous sûr de vouloir annuler cet événement ?")) return;

  // L'URL corrigée avec /delete/
  fetch(`http://localhost:8081/admin/evenements/delete/${id}`, {
    method: "DELETE",
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur lors de la suppression");
      return res.text();
    })
    .then(() => {
      alert("Événement annulé avec succès.");
      GetEvenements(); // Met à jour la liste
    })
    .catch((err) => {
      console.error(err);
      alert("Une erreur est survenue lors de l'annulation.");
    });
}

// ==========================================
// 4. FONCTIONS UTILITAIRES (MODALES)
// ==========================================
function openNewEvt() {
  document.getElementById("evtModal").classList.add("open");
}

function closeEvt() {
  document.getElementById("evtModal").classList.remove("open");
}

function resetEvtForm() {
  document.getElementById("evt-titre").value = "";
  document.getElementById("evt-desc").value = "";
  document.getElementById("evt-date").value = "";
  document.getElementById("evt-lieu").value = "";
  document.getElementById("evt-capacite").value = "";
  document.getElementById("evt-tarif").value = "";
}

function editerEvenement(id) {
  alert("La modification d'événement sera bientôt disponible !");
  // À implémenter avec une modale pré-remplie
}

// ==========================================
// INITIALISATION AU CHARGEMENT DE LA PAGE
// ==========================================
document.addEventListener("DOMContentLoaded", () => {
  GetEvenements();

  // Bloque la sélection de dates antérieures à aujourd'hui
  const dateInput = document.getElementById("evt-date");
  if (dateInput) {
    const today = new Date().toISOString().split("T")[0];
    dateInput.setAttribute("min", today);
  }
});
