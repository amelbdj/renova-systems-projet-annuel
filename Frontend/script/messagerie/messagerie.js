let chatSocket = null;
let currentChatAnnonceId = null;
let currentChatDestinataireId = null;

function openChat(annonceId, destinataireId) {
  currentChatAnnonceId = parseInt(annonceId);
  currentChatDestinataireId = parseInt(destinataireId);

  const monUserId = parseInt(localStorage.getItem("userId"));
  const chatContainer = document.getElementById("chat-messages");

  chatContainer.innerHTML =
    "<div style='text-align:center; color:var(--txt-m); padding-top:20px;'>Chargement de la conversation...</div>";
  document.getElementById("chatModal").style.display = "flex";

  fetch(
    `${API_BASE_URL}/api/chat/history?annonce_id=${currentChatAnnonceId}&user1=${monUserId}&user2=${currentChatDestinataireId}`,
    {
      headers: { Authorization: "Bearer " + localStorage.getItem("token") },
    },
  )
    .then((res) => {
      if (!res.ok)
        throw new Error("Erreur serveur lors de la récupération des messages.");
      return res.json();
    })
    .then((messages) => {
      chatContainer.innerHTML = "";

      if (messages && messages.length > 0) {
        messages.forEach((msg) => {
          const isMe = msg.expediteur_id === monUserId;
          appendMessageToUI(msg.contenu, isMe);
        });
      } else {
        chatContainer.innerHTML =
          "<div id='empty-chat' style='text-align:center; color:var(--txt-m); padding-top:20px;'>Aucun message. Lancez la discussion !</div>";
      }

      connectWebSocket(monUserId);
    })
    .catch((err) => {
      console.error("Erreur historique:", err);
      chatContainer.innerHTML =
        "<div style='text-align:center; color:red; padding-top:20px;'>Erreur lors du chargement.</div>";
    });
}

function connectWebSocket(monUserId) {
  if (chatSocket) chatSocket.close();

  chatSocket = new WebSocket(`${WS_BASE_URL}/ws/chat?userId=${monUserId}`);

  chatSocket.onmessage = function (event) {
    const msg = JSON.parse(event.data);
    if (msg.annonce_id === currentChatAnnonceId) {
      appendMessageToUI(msg.contenu, false);
    }
  };
}

function sendChatMessage() {
  const input = document.getElementById("chat-input");
  const texte = input.value.trim();
  if (!texte || !chatSocket) return;

  const msgData = {
    annonce_id: currentChatAnnonceId,
    expediteur_id: parseInt(localStorage.getItem("userId")),
    destinataire_id: currentChatDestinataireId,
    contenu: texte,
  };

  chatSocket.send(JSON.stringify(msgData));
  appendMessageToUI(texte, true);
  input.value = "";
}

function closeChat() {
  document.getElementById("chatModal").style.display = "none";
  if (chatSocket) chatSocket.close();
}

function appendMessageToUI(texte, isMe) {
  const container = document.getElementById("chat-messages");

  const emptyMsg = document.getElementById("empty-chat");
  if (emptyMsg) emptyMsg.remove();

  const alignement = isMe ? "flex-end" : "flex-start";
  const bgColor = isMe ? "var(--vi)" : "var(--bg3)";
  const textColor = isMe ? "#fff" : "var(--txt)";
  const borderRadius = isMe ? "12px 12px 0 12px" : "12px 12px 12px 0";

  const bulle = `
        <div style="display: flex; flex-direction: column; align-items: ${alignement};">
            <div style="background: ${bgColor}; color: ${textColor}; padding: 10px 14px; border-radius: ${borderRadius}; max-width: 80%; font-size: 14px; word-wrap: break-word;">
                ${texte}
            </div>
        </div>
    `;

  container.innerHTML += bulle;
  container.scrollTop = container.scrollHeight;
}
const chatInput = document.getElementById("chat-input");

document.addEventListener("DOMContentLoaded", () => {
  const chatInput = document.getElementById("chat-input");

  if (chatInput) {
    chatInput.addEventListener("keypress", function (e) {
      if (e.key === "Enter") {
        sendChatMessage();
      }
    });
  }
});

function loadMyMessages() {
  const userId = localStorage.getItem("userId");

  const container = document.getElementById("conversations-list");
  const noMsg = document.getElementById("no-messages");

  if (!container) return;

  fetch(`${API_BASE_URL}/api/chat/conversations?userId=${userId}`, {
    headers: { Authorization: "Bearer " + localStorage.getItem("token") },
  })
    .then((res) => res.json())
    .then((data) => {
      container.innerHTML = "";

      if (!data || data.length === 0) {
        if (noMsg) noMsg.style.display = "block";
        return;
      }

      if (noMsg) noMsg.style.display = "none";

      data.forEach((conv) => {
        const div = document.createElement("div");
        div.className = "conversation-item";
        div.style =
          "padding: 15px; background: var(--bg2); margin-bottom: 10px; border-radius: 8px; cursor: pointer; border: 1px solid var(--b0);";

        div.innerHTML = `
                <div>
                    <strong style="color: var(--txt)">${conv.prenom} ${conv.nom}</strong><br>
                    <small style="color: var(--txt-m)">Annonce : ${conv.annonce_titre}</small>
                </div>
            `;

        div.onclick = () => openChat(conv.annonce_id, conv.contact_id);
        container.appendChild(div);
      });
    })
    .catch((err) => {
      console.error("Erreur:", err);
      container.innerHTML = "<p style='color:red'>Erreur de chargement.</p>";
    });
}
function showSection(sectionId) {
  const sections = document.querySelectorAll(
    ".hero, .ann-section, .cont-section, .cat-section, .bottom-row, .dashboard-section",
  );
  sections.forEach((s) => (s.style.display = "none"));

  const target = document.getElementById(sectionId);
  if (target) {
    target.style.display = "block";

    if (sectionId === "messages-section") {
      loadMyMessages();
    }
  }
}
function appendMessageToUI(texte, isMe) {
  const container = document.getElementById("chat-messages");

  const emptyMsg = document.getElementById("empty-chat");
  if (emptyMsg) emptyMsg.remove();

  if (container.innerText.includes("Chargement")) {
    container.innerHTML = "";
  }

  const messageRow = document.createElement("div");
  messageRow.className = `message-row ${isMe ? "me" : "them"}`;

  const bubble = document.createElement("div");
  bubble.className = "message-bubble";
  bubble.textContent = texte;

  messageRow.appendChild(bubble);
  container.appendChild(messageRow);

  container.scrollTo({
    top: container.scrollHeight,
    behavior: "smooth",
  });
}
