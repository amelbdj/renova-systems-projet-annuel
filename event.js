function GetEvent() {
  const container = document.getElementById("validations");

  fetch("http://localhost:8081/admin/Events")
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur");
      return response.json();
    })
    .then((events) => {
      let htmlContent = "";

      events.forEach((event) => {
        if (
          event.statut_validation &&
          event.statut_validation.toLowerCase() === "en attente"
        ) {
          htmlContent += `
          <div class="val-list fu fu1">
            <div class="val-item evt" data-type="evt">
              <div class="val-body">
                <div class="val-title">${event.titre}</div>
                <div class="val-meta">
                  📅 Event · ${event.prenom} ${event.nom} ·
                  Date : ${new Date(event.date_debut).toLocaleDateString()}
                  Nombre de place : ${event.nombre_place}
                </div>
                <div class="val-desc">${event.description}</div>

                <div class="val-actions">
                  <button class="va-btn va-ok" onclick="ValidateEvent(${event.id})">
                    ✓ Approuver
                  </button>
                  <button class="va-btn va-no" onclick="RefuseEvent(${event.id})">
                    ✕ Refuser
                  </button>
                  <button class="va-btn va-view" onclick="showPreview(${event.id})">
                    👁 Aperçu
                  </button>
                </div>
              </div>
              <span class="tag t-or" style="flex-shrink: 0; font-size: 10px">Event</span>
            </div>
          </div>`;
        }
      });

      container.innerHTML +=
        htmlContent ||
        `<div style="padding:20px;">Aucun evenement en attente.</div>`;
    })
    .catch((error) => {
      console.error("Erreur API :", error);
      container.innerHTML = `<div style="padding: 20px; color: red;">Impossible de charger les données.</div>`;
    });
}

function ValidateEvent(id) {
  fetch(`http://localhost:8081/admin/evenements/validate/${id}`, {
    method: "PUT",
  })
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur");
      GetEvent();
    })
    .catch((error) => {
      console.error("Erreur API :", error);
    });
}

function RefuseEvent(id) {
  fetch(`http://localhost:8081/admin/evenements/refuse/${id}`, {
    method: "PUT",
  })
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur");
      GetEvent();
    })
    .catch((error) => {
      console.error("Erreur API :", error);
    });
}

const eventbutton = document.getElementById("eventButton");
eventbutton.addEventListener("click", () => {
  GetEvent();
});
document.addEventListener("DOMContentLoaded", () => {
  GetEvent();
});
