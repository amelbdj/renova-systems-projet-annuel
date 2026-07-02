var monToken = localStorage.getItem("token");
var userId = localStorage.getItem("userId");

function chargerProfil() {
  if (!monToken || !userId) {
    window.location.href = "/login";
    return;
  }

  fetch(`${API_BASE_URL}/admin/users/${userId}`, {
    headers: { Authorization: "Bearer " + monToken },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur lors du chargement du profil");
      return res.json();
    })
    .then((data) => {
      const profilName = document.querySelector(".profile-name");
      const heroTitleEm = document.getElementById("user-name-display");
      const avatar = document.querySelector(".profile-ava");
      const profilRole = document.querySelector(".profile-role");

      if (profilName) profilName.textContent = data.nom + " " + data.prenom;
      if (heroTitleEm) heroTitleEm.textContent = data.prenom;
      if (avatar) avatar.textContent = data.prenom.charAt(0);
      if (profilRole) profilRole.textContent = data.role || "Salarié";
    })
    .catch((err) => console.error("Erreur Profil:", err));

  const statArticle = document.getElementById("stat-article");
  if (statArticle) {
    fetch(`${API_BASE_URL}/admin/articles/salarie/${userId}`, {
      headers: { Authorization: "Bearer " + monToken },
    })
      .then((res) => res.json())
      .then((articles) => {
        if (!articles) articles = [];
        statArticle.textContent = articles.length;
      })
      .catch((err) => console.error("Erreur stat articles:", err));
  }

  const statEvent = document.getElementById("stat-event");
  const statValide = document.getElementById("stat-valide");
  if (statEvent || statValide) {
    fetch(`${API_BASE_URL}/admin/evenements`, {
      headers: { Authorization: "Bearer " + monToken },
    })
      .then((res) => res.json())
      .then((evenements) => {
        if (!evenements) evenements = [];

        const mesEvts = evenements.filter((e) => {
          const idAuteur =
            e.id_salarie ||
            e.IdSalarie ||
            e.idsalarie ||
            e.idSalarie ||
            e.user_id ||
            e.IdUser;

          return String(idAuteur) === String(userId);
        });

        if (statEvent) statEvent.textContent = mesEvts.length;

        if (statValide) {
          const enLigne = mesEvts.filter((evt) => {
            const statut = (
              evt.statut_validation ||
              evt.statut ||
              evt.Statut ||
              "en attente"
            ).toLowerCase();
            return (
              statut === "valide" ||
              statut === "en ligne" ||
              statut === "publié" ||
              statut === "publie"
            );
          });
          statValide.textContent = enLigne.length;
        }
      })
      .catch((err) => console.error("Erreur stat événements:", err));
  }

  const statForum = document.getElementById("stat-forum");
  if (statForum) {
    fetch(`${API_BASE_URL}/admin/forum/messages`, {
      headers: { Authorization: "Bearer " + monToken },
    })
      .then((res) => res.json())
      .then((messages) => {
        if (!messages) messages = [];

        let nbSignales = 0;
        for (let i = 0; i < messages.length; i++) {
          if (messages[i].est_signale) {
            nbSignales++;
          }
        }
        statForum.textContent = nbSignales;
        statForum.textContent = messages.length;
      })
      .catch((err) => console.error("Erreur stat forum:", err));
  }

  chargerNotifications();
}

function chargerNotifications() {
  var pastille = document.querySelector(".notif-dot");
  if (!pastille) return;
  var cloche = pastille.parentElement;

  fetch(`${API_BASE_URL}/admin/notifications/user/` + userId, {
    headers: { Authorization: "Bearer " + monToken },
  })
    .then(function (res) {
      return res.json();
    })
    .then(function (notifs) {
      if (!notifs) notifs = [];

      var nonLues = 0;
      for (var i = 0; i < notifs.length; i++) {
        if (notifs[i].est_lu == 0) {
          nonLues++;
        }
      }

      pastille.style.display = nonLues > 0 ? "block" : "none";

      var panneau = document.getElementById("notif-panel");
      if (!panneau) {
        panneau = document.createElement("div");
        panneau.id = "notif-panel";
        panneau.style.cssText =
          "display:none; position:fixed; top:64px; right:20px; width:320px; max-height:420px; overflow-y:auto; background:var(--bg2); border:1px solid var(--b0); border-radius:12px; box-shadow:0 10px 30px rgba(0,0,0,0.45); z-index:5000; padding:8px;";
        document.body.appendChild(panneau);
      }

      var html =
        "<div style='font-family:Syne,sans-serif; font-weight:700; padding:10px; color:var(--txt);'>Notifications</div>";
      if (notifs.length === 0) {
        html +=
          "<div style='padding:14px; color:var(--txt-m); font-size:13px;'>Aucune notification.</div>";
      } else {
        for (var j = 0; j < notifs.length; j++) {
          var fond =
            notifs[j].est_lu == 0 ? "rgba(138,100,255,0.10)" : "transparent";
          html +=
            "<div style='padding:10px 12px; border-bottom:1px solid var(--b0); font-size:13px; color:var(--txt-m); background:" +
            fond +
            ";'>" +
            notifs[j].contenu +
            "</div>";
        }
      }
      panneau.innerHTML = html;

      cloche.style.cursor = "pointer";
      cloche.onclick = function () {
        if (panneau.style.display === "none") {
          panneau.style.display = "block";
          if (nonLues > 0) {
            fetch(
              `${API_BASE_URL}/admin/notifications/user/` + userId + "/read",
              {
                method: "POST",
                headers: { Authorization: "Bearer " + monToken },
              },
            ).then(function () {
              pastille.style.display = "none";
              nonLues = 0;
            });
          }
        } else {
          panneau.style.display = "none";
        }
      };
    })
    .catch(function (err) {
      console.error("Erreur notifications:", err);
    });
}

function logout() {
  localStorage.clear();
  window.location.href = "/login";
}
function goToProfile() {
  const userId = localStorage.getItem("userId");
  const role = localStorage.getItem("role");

  if (!userId) {
    window.location.href = "/login";
    return;
  }

  window.location.href = `/profil?id=${userId}`;
}

document.addEventListener("DOMContentLoaded", chargerProfil);
