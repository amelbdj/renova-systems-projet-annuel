let monToken = localStorage.getItem("token");

function GetEvent() {
  const container = document.getElementById("result");
  if (!container) return;

  fetch("http://localhost:8081/admin/evenements", {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => res.json())
    .then((events) => {
      const currentTabId = document.querySelector(".vtab.on").id;
      if (currentTabId !== "tout") container.innerHTML = "";
      let lieu = "";
      let htmlContent = "";

      events.forEach((event) => {
        console.log("Événement:", event); 

        if (
          event.statut_validation &&
          event.statut_validation.toLowerCase() === "en attente"
        ) {
          if (event.lieu != "0") {
            lieu = ` <span data-i18n="backoffice.events.location">Lieu :</span> ${event.lieu}`;
          } else {
            lieu = "";
          }

<<<<<<< HEAD
          
=======
>>>>>>> origin/faty
          let dateStr = event.date_debut ? event.date_debut : "Date inconnue";

          htmlContent += `
                    <div class="val-list fu fu1">
                        <div class="val-item ann">
                            <div class="val-body">
                                <div class="val-title">${event.titre}</div>
                                <div class="val-meta">
                                    <span class="material-symbols-outlined">calendar_today</span> 
                                    <span data-i18n="backoffice.events.event">${event.type}</span> · ${event.prenomSalarie} ${event.nomSalarie} · 
                                    <span data-i18n="backoffice.events.seats">Places :</span> ${event.nb_places} · 
                                    ${lieu}
                                </div>
                                <div class="val-desc">${event.description}</div>
                                
                                <span data-i18n="backoffice.events.on_date">Le</span> ${dateStr}
                                
                                <div class="val-actions">
                                    <button class="va-btn va-ok" onclick="ValidateEvent(${event.id})" data-i18n="backoffice.btn.approve">✓ Approuver</button>
                                    <button class="va-btn va-no" onclick="RefuseEvent(${event.id})" data-i18n="backoffice.btn.refuse">✕ Refuser</button>
                                </div>
                            </div>
                            <span class="tag t-pu" style="flex-shrink: 0; font-size: 10px" data-i18n="backoffice.events.event_tag">${event.type}</span>
                        </div>
                    </div>`;
        }
      });

      if (htmlContent) {
        container.innerHTML += htmlContent;
      } else if (!container.innerHTML) {
        container.innerHTML = `<div style="padding:20px" data-i18n="backoffice.events.no_events">Aucun événement en attente.</div>`;
      }

<<<<<<< HEAD
      
=======
>>>>>>> origin/faty
      if (typeof appliquerTraductions === "function") {
        appliquerTraductions();
      }
    })
    .catch((err) => console.error(err));
}

function ValidateEvent(id) {
  fetch(`http://localhost:8081/admin/evenements/validate/${id}`, {
    method: "PUT",
    headers: {
      Authorization: "Bearer " + monToken,
    },
  }).then(() => {
    GetEvent();
    UpdateValidationCount();
  });
}

function RefuseEvent(id) {
  fetch(`http://localhost:8081/admin/evenements/refuse/${id}`, {
    method: "PUT",
    headers: {
      Authorization: "Bearer " + monToken,
    },
  }).then(() => {
    GetEvent();
    UpdateValidationCount();
  });
}

document.addEventListener("DOMContentLoaded", () => {
  GetEvent();
});
