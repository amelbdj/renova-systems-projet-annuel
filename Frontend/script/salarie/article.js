let monToken = localStorage.getItem("token");
let userId = localStorage.getItem("userId");

if (!monToken || !userId) {
  alert("Vous devez être connecté pour accéder à cette page.");
  window.location.href = "login.html";
}

function chargerArticles() {
  fetch(`http://localhost:8081/admin/articles/salarie/${userId}`, {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) {
        throw new Error("Erreur réseau");
      }
      return res.json();
    })
    .then((articles) => {
      const conteneurPublies = document.getElementById("liste-publies");
      const conteneurBrouillons = document.getElementById("liste-brouillons");
      const statArticle = document.getElementById("stat-article"); // Utilise getElementById      let compteurPublies = 0;
      let compteurPublies = 0;
      // On vide les conteneurs
      conteneurPublies.innerHTML = "";
      conteneurBrouillons.innerHTML = "";

      articles.forEach((art) => {
        // 1. Choix du style de badge selon le statut
        let badgeStatut = "";
        if (art.statut === "brouillon") {
          badgeStatut = `<span class="tag t-amber">Brouillon</span>`;
        } else if (art.statut === "en attente")
          badgeStatut = `<span class="tag t-blue">En attente</span>`;
        else if (art.statut === "refuse")
          badgeStatut = `<span class="tag t-red">Refusé</span>`;
        else {
          badgeStatut = `<span class="tag t-green">Publié</span>`;
          compteurPublies++;
        }

        let icone = "📝";
        if (art.type.toLowerCase().includes("conseil")) icone = "💡";
        if (art.type.toLowerCase().includes("news")) icone = "📰";
        if (art.type.toLowerCase().includes("tuto")) icone = "🛠️";
        const card = `
  <div class="post-item">
    <div class="post-ico" style="background:rgba(48,212,192,.09)">${icone}</div>
    <div class="post-body">
      <div class="post-title">${art.titre}</div>
      <div class="post-excerpt">${art.contenu.substring(0, 65)}...</div>
      <div class="post-meta">
          <span class="tag t-vi">${art.type}</span>
          ${badgeStatut}
          <button class="btn btn-v btn-sm" onclick="editerArticle(${art.id})">＋ Modifier</button>
          <button class="mod-btn mod-ban" onclick="DeleteArticle(${art.id})">Supprimer</button>
      </div>
    </div>
  </div>
`;

        if (art.statut === "brouillon" || art.statut === "refuse") {
          conteneurBrouillons.innerHTML += card;
        } else {
          conteneurPublies.innerHTML += card;
        }
      });

      statArticle.textContent = compteurPublies;
    })
    .catch((error) => {
      console.error("Impossible de récupérer les articles", error);
      document.getElementById("liste-publies").innerHTML =
        `<p style="color:var(--red); font-size:13px;">Serveur indisponible.</p>`;
    });
}

function editerArticle(id) {
  fetch(`http://localhost:8081/admin/articles/${id}`, {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Impossible de récupérer l'article");
      return res.json();
    })
    .then((art) => {
      document.getElementById("edit-article-id").value = art.id;
      document.getElementById("post-title").value = art.titre;
      document.getElementById("post-content").value = art.contenu;
      document.getElementById("post-type").value = art.type;

      const modalTitle = document.querySelector("#postModal .sec-title-text");
      if (modalTitle) modalTitle.textContent = "Modifier l'article #" + id;

      openNewPost();
    })
    .catch((err) => {
      console.error(err);
      alert("Erreur lors de la récupération des données de l'article.");
    });
}

function saveArticle(action) {
  const id = document.getElementById("edit-article-id").value;
  const titre = document.getElementById("post-title").value;
  const contenu = document.getElementById("post-content").value;
  const type = document.getElementById("post-type").value;

  if (!titre || !contenu) {
    alert("Le titre et le contenu sont obligatoires !");
    return;
  }

  const articleData = {
    id_salarie: parseInt(userId), // À remplacer par l'ID réel du salarié connecté apre sync avec faty
    titre: titre,
    contenu: contenu,
    type: type,
  };

  let url = "";
  if (id !== "") {
    // Mode MODIFICATION
    url = `http://localhost:8081/admin/articles/modify/${id}/${action}`;
    method = "PUT";
  } else {
    // Mode CRÉATION
    url = `http://localhost:8081/admin/articles/add/${action}`;
    method = "POST";
  }

  fetch(url, {
    method: method,
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
    body: JSON.stringify(articleData),
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur lors de la sauvegarde");
      return res.json();
    })
    .then((data) => {
      const message =
        action === "publier" ? "Article publié !" : "Brouillon enregistré.";
      alert(message);

      closePost();
      chargerArticles();
    })
    .catch((err) => {
      console.error(err);
      alert("Erreur serveur : " + err.message);
    });
}

function DeleteArticle(id) {
  if (!confirm("Êtes-vous sûr de vouloir supprimer cet article ?")) return;
  fetch(`http://localhost:8081/admin/articles/delete/${id}`, {
    method: "DELETE",
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur lors de la suppression");
      return res.json();
    })
    .then((data) => {
      alert("Article supprimé !");
      chargerArticles();
    })
    .catch((err) => {
      console.error(err);
      alert("Erreur serveur : " + err.message);
    });

  chargerArticles();
}

function openNewPost() {
  // Vérifie bien que l'ID est postModal et pas autre chose !
  const modal = document.getElementById("postModal");
  if (modal) {
    modal.classList.add("open");
    modal.style.display = "flex";
  }
}

function closePost() {
  const modal = document.getElementById("postModal");
  if (modal) {
    modal.classList.remove("open");
    modal.style.display = "none";
  }
}

document.addEventListener("DOMContentLoaded", chargerArticles);
