
var monToken = localStorage.getItem("token");
var userId = localStorage.getItem("userId");

// Redirection si non connecté
if (!monToken || !userId) {
  window.location.href = "../login.html";
}

function openNewEvt() {
  document.getElementById("evtModal").style.display = "flex";
}

function closeEvt() {
  document.getElementById("evtModal").style.display = "none";
}

function resetEvtForm() {
  document.getElementById("evt-titre").value = "";
  document.getElementById("evt-desc").value = "";
  document.getElementById("evt-date").value = "";
  document.getElementById("evt-lieu").value = "";
  document.getElementById("evt-heure-debut").value = "10:00";
  document.getElementById("evt-heure-fin").value = "13:00";
  document.getElementById("evt-capacite").value = "";
  document.getElementById("evt-tarif").value = "";

  var imageInput = document.getElementById("evt-image");
  if (imageInput) {
    imageInput.value = "";
  }
}

function CreateEvent() {
  var titre = document.getElementById("evt-titre").value.trim();
  var type = document.getElementById("evt-type").value;
  var desc = document.getElementById("evt-desc").value.trim();
  var date = document.getElementById("evt-date").value;
  var lieu = document.getElementById("evt-lieu").value.trim();
  var heureDebut = document.getElementById("evt-heure-debut").value;
  var heureFin = document.getElementById("evt-heure-fin").value;
  var capacite = document.getElementById("evt-capacite").value;
  var tarif = document.getElementById("evt-tarif").value;

  // L'image
  var imageInput = document.getElementById("evt-image");
  var imageFile = null;
  if (imageInput && imageInput.files.length > 0) {
    imageFile = imageInput.files[0];
  }

  if (!titre || !desc || !date) {
    alert(
      "Veuillez remplir les champs obligatoires : Titre, Description et Date.",
    );
    return;
  }

  var datetimeDebut = date + " " + (heureDebut || "00:00") + ":00";
  var datetimeFin = date + " " + (heureFin || "00:00") + ":00";

  var formData = new FormData();
  formData.append("idSalarie", userId);
  formData.append("titre", titre);
  formData.append("type", type);
  formData.append("description", desc);
  formData.append("date_debut", datetimeDebut);
  formData.append("date_fin", datetimeFin);
  formData.append("lieu", lieu);

  if (capacite) {
    formData.append("capacite", capacite);
  } else {
    formData.append("capacite", 0);
  }
  if (tarif) {
    formData.append("tarif", tarif);
  } else {
    formData.append("tarif", 0);
  }

  if (imageFile) {
    formData.append("image", imageFile);
  }

  fetch("http://localhost:8081/admin/evenements/add", {
    method: "POST",
    headers: {
      Authorization: "Bearer " + monToken,
    },
    body: formData, // Le navigateur gère le format tout seul !
  })
    .then(function (reponse) {
      if (!reponse.ok) {
        throw new Error("Erreur lors de la création");
      }
      return reponse.text();
    })
    .then(function () {
      alert("Événement soumis avec succès. Il est en attente de validation.");
      closeEvt();
      resetEvtForm();
      GetEvenements(); // On met à jour la liste
    })
    .catch(function (erreur) {
      console.log(erreur);
      alert("Erreur lors de l'enregistrement de l'événement.");
    });
}

function GetEvenements() {
  var conteneur = document.getElementById("event-grid");
  if (!conteneur) return;

  var counterEvt = 0;
  var counterValide = 0;

  fetch("http://localhost:8081/admin/evenements", {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then(function (reponse) {
      return reponse.json();
    })
    .then(function (evenements) {
      if (!evenements) {
        evenements = [];
      }

      var htmlContent = "";
      var statEvent = document.getElementById("stat-event");
      var statAttente = document.getElementById("stat-valide");

      var mesEvenements = [];
      for (var i = 0; i < evenements.length; i++) {
        if (
          evenements[i].id_salarie == userId ||
          evenements[i].IdSalarie == userId ||
          evenements[i].idSalarie == userId
        ) {
          mesEvenements.push(evenements[i]);
        }
      }

      if (mesEvenements.length === 0) {
        conteneur.innerHTML =
          "<p style='color:var(--txt-m); padding:20px;'>Vous n'avez créé aucun événement pour le moment.</p>";
        if (statEvent) statEvent.textContent = 0;
        if (statAttente) statAttente.textContent = 0;
        return;
      }

      for (var j = 0; j < mesEvenements.length; j++) {
        var evt = mesEvenements[j];
        counterEvt++;

        var statut = evt.statut_validation
          ? evt.statut_validation.toLowerCase()
          : "en attente";
        var statusBadge = "";
        var actionButtons = "";

        if (statut === "valide" || statut === "en ligne") {
          counterValide++;
          statusBadge = "<div class='evt-status t-green'>✓ En ligne</div>";
          actionButtons =
            "<button class='btn btn-danger btn-xs' onclick='DeleteEvenement(" +
            evt.id +
            ")'>Annuler</button>";
        } else {
          statusBadge = "<div class='evt-status t-amber'>⏳ En attente</div>";
          actionButtons =
            "<button class='btn btn-danger btn-xs' onclick='DeleteEvenement(" +
            evt.id +
            ")'>Annuler</button>";
        }

        var dateFormatee = evt.date_debut;
        var typeAffichage = evt.type || evt.format || "Événement";

        var topSectionHtml = "";

        if (evt.image_url && evt.image_url !== "") {
          topSectionHtml =
            `
        <div style="position: relative;">
          <img src="http://localhost:8081/` +
            evt.image_url +
            `" style="width: 100%; height: 160px; object-fit: cover; border-radius: 12px 12px 0 0; display: block;" />
          <div style="position: absolute; top: 12px; right: 12px; box-shadow: 0 4px 6px rgba(0,0,0,0.3); border-radius: 20px;">
            ` +
            statusBadge +
            `
          </div>
          <div style="position: absolute; top: 12px; left: 12px; box-shadow: 0 4px 6px rgba(0,0,0,0.3);" class="evt-type-badge etb-formation">
            ` +
            typeAffichage +
            `
          </div>
        </div>`;
        } else {
          topSectionHtml =
            `
        <div class="evt-banner" style="background:linear-gradient(135deg,#100820,#1c1040); margin: 0; border-radius: 12px 12px 0 0;">
          📅 
          <div class="evt-type-badge etb-formation">` +
            typeAffichage +
            `</div>
          ` +
            statusBadge +
            `
        </div>`;
        }

        htmlContent +=
          `
      <div class="evt-card" style="padding: 0; border: 1px solid var(--b0); border-radius: 12px; background: var(--bg2); margin-bottom: 20px;">
        
        ` +
          topSectionHtml +
          `
        
        <div class="evt-body" style="padding: 20px;">
          <div class="evt-name">` +
          evt.titre +
          `</div>
          <div class="evt-desc" style="margin-top: 10px;">` +
          evt.description +
          `</div>
          <div class="evt-meta" style="margin-top: 15px;">
            <span class="tag t-vi">📅 ` +
          dateFormatee +
          `</span>
          </div>
          <div class="evt-foot" style="margin-top: 15px;">
            ` +
          actionButtons +
          `
          </div>
        </div>
      </div>`;
      }

      conteneur.innerHTML = htmlContent;

      if (statEvent) statEvent.textContent = counterEvt;
      if (statAttente) statAttente.textContent = counterValide;
    })
    .catch(function (erreur) {
      console.log(erreur);
    });
}

function DeleteEvenement(id) {
  if (!confirm("Êtes-vous sûr de vouloir annuler cet événement ?")) {
    return;
  }

  fetch("http://localhost:8081/admin/evenements/delete/" + id, {
    method: "DELETE",
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then(function (reponse) {
      if (!reponse.ok) throw new Error("Erreur lors de la suppression");
      alert("Événement annulé avec succès.");
      GetEvenements();
    })
    .catch(function (erreur) {
      console.log(erreur);
      alert("Une erreur est survenue lors de l'annulation.");
    });
}

document.addEventListener("DOMContentLoaded", function () {
  GetEvenements();

  var dateInput = document.getElementById("evt-date");
  if (dateInput) {
    var today = new Date().toISOString().split("T")[0];
    dateInput.setAttribute("min", today);
  }
});
