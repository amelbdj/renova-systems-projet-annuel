function GetEvent() {
  const container = document.getElementById("result");
  if (!container) return;

  fetch("http://localhost:8081/admin/evenements")
    .then((res) => res.json())
    .then((events) => {
      const currentTab = document.querySelector(".vtab.on").textContent;
      if (!currentTab.includes("Tout")) container.innerHTML = "";

      let htmlContent = "";
      events.forEach((event) => {
        if (
          event.statut_validation &&
          event.statut_validation.toLowerCase() === "en attente"
        ) {
          htmlContent += `
                    <div class="val-list fu fu1">
                        <div class="val-item ann">
                            <div class="val-body">
                                <div class="val-title">${event.titre}</div>
                                <div class="val-meta">
                                    📅 Event · ${event.prenom || "Organisateur"} · 
                                    Places : ${event.nb_places} · Le ${new Date(event.date_debut).toLocaleDateString()}
                                </div>
                                <div class="val-desc">${event.description}</div>
                                <div class="val-actions">
                                    <button class="va-btn va-ok" onclick="ValidateEvent(${event.id})">✓ Approuver</button>
                                    <button class="va-btn va-no" onclick="RefuseEvent(${event.id})">✕ Refuser</button>
                                </div>
                            </div>
                            <span class="tag t-pu" style="flex-shrink: 0 font-size: 10px">Événement</span>
                        </div>
                    </div>`;
        }
      });

      if (htmlContent) {
        container.innerHTML += htmlContent;
      } else if (!container.innerHTML) {
        container.innerHTML = `<div style="padding:20px">Aucun événement en attente.</div>`;
      }
    })
    .catch((err) => console.error(err));
}

function ValidateEvent(id) {
  fetch(`http://localhost:8081/admin/evenements/validate/${id}`, {
    method: "PUT",
  }).then(() => {
    GetEvent();
    UpdateValidationCount();
  });
}

function RefuseEvent(id) {
  fetch(`http://localhost:8081/admin/evenements/refuse/${id}`, {
    method: "PUT",
  }).then(() => {
    GetEvent();
    UpdateValidationCount();
  });
}
document.addEventListener("DOMContentLoaded", () => {
  GetEvent();
});
