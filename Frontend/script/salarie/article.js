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
      const statArticle = document.getElementById("stat-article");

      let compteurPublies = 0;

      // On vide les conteneurs
      if (conteneurPublies) conteneurPublies.innerHTML = "";
      if (conteneurBrouillons) conteneurBrouillons.innerHTML = "";

      // 🛡️ LE FAMEUX BOUCLIER ANTI-NULL (Spécial Golang)
      if (!articles) {
        articles = [];
      }

      articles.forEach((art) => {
        // Choix du style de badge selon le statut
        let badgeStatut = "";
        let statut = (art.statut || "").toLowerCase(); // Sécurité pour les majuscules

        if (statut === "brouillon") {
          badgeStatut = `<span class="tag t-amber">Brouillon</span>`;
        } else if (statut === "en attente") {
          badgeStatut = `<span class="tag t-blue">En attente</span>`;
        } else if (statut === "refuse") {
          badgeStatut = `<span class="tag t-red">Refusé</span>`;
        } else {
          badgeStatut = `<span class="tag t-green">Publié</span>`;
          compteurPublies++;
        }

        let icone = "📝";
        let typeArt = (art.type || "").toLowerCase();

        if (typeArt.includes("conseil"))
          icone = `<span class="material-symbols-outlined">lightbulb</span>`;
        if (typeArt.includes("news"))
          icone = `<span class="material-symbols-outlined">newspaper</span>`;
        if (typeArt.includes("tuto"))
          icone = `<span class="material-symbols-outlined">build</span>`;

        const card = `
  <div class="post-item">
    <div class="post-ico" style="background:rgba(48,212,192,.09)">${icone}</div>
    <div class="post-body">
      <div class="post-title">${art.titre}</div>
      <div class="post-excerpt">${art.contenu ? art.contenu.substring(0, 65) : ""}...</div>
      <div class="post-meta">
          <span class="tag t-vi">${art.type || "Article"}</span>
          ${badgeStatut}
          <button class="btn btn-v btn-sm" onclick="editerArticle(${art.id})">＋ Modifier</button>
          <button class="mod-btn mod-ban" onclick="DeleteArticle(${art.id})">Supprimer</button>
      </div>
    </div>
  </div>
`;

        if (statut === "brouillon" || statut === "refuse") {
          if (conteneurBrouillons) conteneurBrouillons.innerHTML += card;
        } else {
          if (conteneurPublies) conteneurPublies.innerHTML += card;
        }
      });

      // Si le tableau est vide, on affiche un petit message sympa
      if (articles.length === 0) {
        if (conteneurPublies)
          conteneurPublies.innerHTML = `<p style="color:var(--txt-m); font-size:13px; padding: 10px 0;">Aucun article publié.</p>`;
        if (conteneurBrouillons)
          conteneurBrouillons.innerHTML = `<p style="color:var(--txt-m); font-size:13px; padding: 10px 0;">Aucun brouillon en cours.</p>`;
      }

      if (statArticle) statArticle.textContent = compteurPublies;
    })
    .catch((error) => {
      console.error("Impossible de récupérer les articles", error);
      const conteneurPublies = document.getElementById("liste-publies");
      if (conteneurPublies)
        conteneurPublies.innerHTML = `<p style="color:var(--red); font-size:13px;">Serveur indisponible.</p>`;
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
    id_salarie: parseInt(userId),
    titre: titre,
    contenu: contenu,
    type: type,
  };

  let url = "";
  let method = "";
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
        action === "publier"
          ? "Article publié / En attente !"
          : "Brouillon enregistré.";
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
      chargerArticles(); // Rafraîchit les articles APRES la suppression
    })
    .catch((err) => {
      console.error(err);
      alert("Erreur serveur : " + err.message);
    });
}

function openNewPost() {
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

    // On vide les champs du formulaire à la fermeture pour que la prochaine création soit propre
    document.getElementById("edit-article-id").value = "";
    document.getElementById("post-title").value = "";
    document.getElementById("post-content").value = "";
    const modalTitle = document.querySelector("#postModal .sec-title-text");
    if (modalTitle) modalTitle.textContent = "Rédiger un article";
  }
}

document.addEventListener("DOMContentLoaded", chargerArticles);
