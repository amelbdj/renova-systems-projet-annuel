function GetForumMessages(filtreType) {
  const container = document.getElementById("forum-thread-list");
  if (!container) return;

  container.innerHTML =
    "<p style='padding:14px; color:var(--txt-m);'>Chargement des messages...</p>";

  let url = "http://localhost:8081/admin/forum/messages";
  if (filtreType === "flag") {
    url += "?filter=signales";
  }

  fetch(url, {
    method: "GET",
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur réseau");
      return res.json();
    })
    .then((messages) => {
      container.innerHTML = ""; // On vide le conteneur

      if (!messages || messages.length === 0) {
        container.innerHTML =
          "<p style='padding:14px; color:var(--txt-m);'>Aucun message trouvé.</p>";
        return;
      }

      let htmlContent = "";
      messages.forEach((msg) => {
        const isFlagged = msg.est_signale ? "flagged" : "";
        const flagBadge = msg.est_signale
          ? `<span class="tag t-red" style="font-size: 10px">🚩 Signalé</span>`
          : "";
        const flagIcon = msg.est_signale
          ? `<div class="flag-icon">🚩</div>`
          : "";

        const authorName = `${msg.prenom_auteur} ${msg.nom_auteur}`;

        const timeAgo = formatTimeAgo(msg.date_creation);

        htmlContent += `
            <div class="thread-item ${isFlagged}" data-flag="${msg.est_signale}">
                <div class="thread-ava">👤</div>
                <div class="thread-body">
                    <div class="thread-top">
                        <span class="thread-user">${authorName}</span>
                        <span class="thread-time">${timeAgo}</span>
                        ${flagBadge}
                    </div>
                    <div class="thread-text">
                        ${msg.contenu}
                    </div>
                    <div class="thread-actions">
                        <button class="mod-btn mod-approve" onclick="ModerateMessage(${msg.id_message}, 'approuver')">✓ Approuver</button>
                        <button class="mod-btn mod-hide" onclick="ModerateMessage(${msg.id_message}, 'masquer')">⊘ Masquer</button>
                        <button class="mod-btn mod-ban" onclick="alert('Fonction de bannissement à lier au module User')">🚫 Bannir</button>
                    </div>
                </div>
                ${flagIcon}
            </div>`;
      });

      container.innerHTML = htmlContent;
    })
    .catch((err) => {
      console.error("Erreur chargement messages forum :", err);
      container.innerHTML =
        "<p style='padding:14px; color:red;'>Erreur lors du chargement.</p>";
    });
}

function ModerateMessage(idMessage, action) {
  fetch(`http://localhost:8081/admin/forum/messages/moderate/${idMessage}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
    body: JSON.stringify({ action: action }),
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur modération");
      const currentFilter = document.querySelector(".btn-g.btn-xs")
        ? "all"
        : "flag";
      GetForumMessages(currentFilter);
      GetForumStats();
    })
    .catch((err) => console.error("Erreur :", err));
}

function GetForumStats() {
  fetch("http://localhost:8081/admin/forum/stats", {
    method: "GET",
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => res.json())
    .then((stats) => {
      const statMsg = document.getElementById("stat-messages");
      const statMembres = document.getElementById("stat-membres");
      const statSign = document.getElementById("stat-signalements");

      if (statMsg) statMsg.textContent = stats.messages_semaine || 0;
      if (statMembres) statMembres.textContent = stats.membres_actifs || 0;
      if (statSign) statSign.textContent = stats.signalements_en_attente || 0;
    })
    .catch((err) => console.error("Erreur stats forum :", err));
}

window.setFilter = function (filterType, btnElement) {
  const buttons = btnElement.parentElement.querySelectorAll("button");
  buttons.forEach((btn) => {
    btn.classList.remove("btn-g");
    btn.style.background = "";
    btn.style.color = "";
    btn.style.border = "";
  });

  if (filterType === "all") {
    btnElement.classList.add("btn-g");
  } else {
    btnElement.style.background = "rgba(240, 95, 95, 0.1)";
    btnElement.style.color = "var(--red)";
    btnElement.style.border = "1px solid rgba(240, 95, 95, 0.2)";
  }

  GetForumMessages(filterType);
};

function formatTimeAgo(dateString) {
  if (!dateString) return "";
  const date = new Date(dateString);
  const now = new Date();
  const diffInSeconds = Math.floor((now - date) / 1000);

  if (diffInSeconds < 60) return "Il y a l'instant";
  const diffInMinutes = Math.floor(diffInSeconds / 60);
  if (diffInMinutes < 60) return `Il y a ${diffInMinutes} min`;
  const diffInHours = Math.floor(diffInMinutes / 60);
  if (diffInHours < 24) return `Il y a ${diffInHours} h`;
  const diffInDays = Math.floor(diffInHours / 24);
  return `Il y a ${diffInDays} j`;
}

document.addEventListener("DOMContentLoaded", () => {
  GetForumMessages("all");
  GetForumStats();
});
