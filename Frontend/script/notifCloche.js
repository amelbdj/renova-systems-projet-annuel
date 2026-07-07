




(function () {
  function creerClocheFlottante() {
    var btn = document.createElement("div");
    btn.style.cssText =
      "position:fixed; top:14px; right:64px; z-index:4000; cursor:pointer;" +
      "width:40px; height:40px; border-radius:50%; background:rgba(0,0,0,.4);" +
      "display:flex; align-items:center; justify-content:center;";
    btn.innerHTML =
      '<span class="material-symbols-outlined" style="color:#fff; font-size:22px;">notifications</span>' +
      '<span class="notif-dot" style="display:none; position:absolute; top:9px; right:9px;' +
      'width:9px; height:9px; border-radius:50%; background:#f05050; border:1px solid #000;"></span>';
    document.body.appendChild(btn);
    return btn.querySelector(".notif-dot");
  }

  function initNotifs() {
    var token = localStorage.getItem("token");
    var userId = localStorage.getItem("userId");
    if (!token || !userId) return;
    if (window.__notifClocheInit) return; 
    window.__notifClocheInit = true;

    var pastille = document.querySelector(".notif-dot") || creerClocheFlottante();
    var cloche = pastille.parentElement;

    fetch(API_BASE_URL + "/admin/notifications/user/" + userId, {
      headers: { Authorization: "Bearer " + token },
    })
      .then(function (res) {
        return res.json();
      })
      .then(function (notifs) {
        if (!notifs) notifs = [];

        var nonLues = 0;
        for (var i = 0; i < notifs.length; i++) {
          if (notifs[i].est_lu == 0) nonLues++;
        }
        pastille.style.display = nonLues > 0 ? "block" : "none";

        var panneau = document.getElementById("notif-panel");
        if (!panneau) {
          panneau = document.createElement("div");
          panneau.id = "notif-panel";
          panneau.style.cssText =
            "display:none; position:fixed; top:64px; right:20px; width:320px;" +
            "max-height:420px; overflow-y:auto; background:var(--bg2,#111);" +
            "border:1px solid var(--b0,#333); border-radius:12px;" +
            "box-shadow:0 10px 30px rgba(0,0,0,0.45); z-index:5000; padding:8px;";
          document.body.appendChild(panneau);
        }

        var html =
          "<div style='font-family:Syne,sans-serif; font-weight:700; padding:10px; color:var(--txt,#eee);'>Notifications</div>";
        if (notifs.length === 0) {
          html +=
            "<div style='padding:14px; color:var(--txt-m,#999); font-size:13px;'>Aucune notification.</div>";
        } else {
          for (var j = 0; j < notifs.length; j++) {
            var fond = notifs[j].est_lu == 0 ? "rgba(120,120,120,0.14)" : "transparent";
            html +=
              "<div style='padding:10px 12px; border-bottom:1px solid var(--b0,#333);" +
              "font-size:13px; color:var(--txt-m,#bbb); background:" + fond + ";'>" +
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
              fetch(API_BASE_URL + "/admin/notifications/user/" + userId + "/read", {
                method: "POST",
                headers: { Authorization: "Bearer " + token },
              }).then(function () {
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

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", initNotifs);
  } else {
    initNotifs();
  }
})();
