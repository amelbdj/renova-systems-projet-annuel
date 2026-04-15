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

  const eventData = {
    idSalarie: parseInt(userId),
    titre: titre,
    type: type,
    description: desc,
    date_debut: datetimeDebut,
    date_fin: datetimeFin,
    lieu: lieu,
    capacite: parseInt(capacite) || 0,
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
    })
    .catch((err) => {
      console.error(err);
      alert("Une erreur est survenue.");
    });
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

function GetEvenements() {
  const container = document.getElementById("event-grid");
  if (!container) return;

  fetch(`http://localhost:8081/admin/evenements/salarie/${userId}`, {
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

      if (!evenements || evenements.length === 0) {
        container.innerHTML = `<div style="color:var(--txt-m); padding:20px;">Aucun événement trouvé.</div>
        <div class="evt-add" onclick="openNewEvt()"><div class="plus">＋</div><span>Créer un événement</span></div>`;
        return;
      }

      evenements.forEach((evt) => {
        let statusBadge = "";
        let actionButtons = "";

        const statut = evt.statut_validation
          ? evt.statut_validation.toLowerCase()
          : "en attente";

        if (statut === "validé" || statut === "en ligne") {
          statusBadge = `<div class="evt-status t-green">✓ En ligne</div>`;
          actionButtons = `
            <span class="tag t-green">Publiée</span>
            <div style="display:flex;gap:5px">
              <button class="btn btn-g btn-xs">Modifier</button>
              <button class="btn btn-danger btn-xs" onclick="DeleteEvenement(${evt.id})">Annuler</button>
            </div>`;
        } else {
          statusBadge = `<div class="evt-status t-amber">⏳ En attente</div>`;
          actionButtons = `
            <span class="tag t-amber">Validation en cours</span>
            <button class="btn btn-g btn-xs">Modifier</button>`;
        }

        let dateFormatee = "Date inconnue";
        let heureFormatee = "";
        if (evt.date_debut) {
          const d = new Date(evt.date_debut);
          dateFormatee = d.toLocaleDateString("fr-FR", {
            day: "numeric",
            month: "short",
          }); // "15 mars"
          heureFormatee = d
            .toLocaleTimeString("fr-FR", { hour: "2-digit", minute: "2-digit" })
            .replace(":", "h"); // "10h00"
        }
        // mettre une img specail event et formation par def
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
              <span class="tag t-vi">${dateFormatee}</span>
              <span class="tag t-blue">${heureFormatee}</span>
              <span class="tag" style="color:var(--txt-m);background:var(--bg3);border:1px solid var(--b0)">👥 Max ${evt.nb_places || evt.capacite || 0}</span>
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
    })
    .catch((err) => console.error(err));
}

document.addEventListener("DOMContentLoaded", () => {
  GetEvenements();
});
