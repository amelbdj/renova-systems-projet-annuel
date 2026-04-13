/* ── FETCH & AFFICHAGE DES ARTICLES ──────────────────────────── */

function chargerArticles() {
  // Appel à ton API Go
  fetch("http://localhost:8081/admin/articles/salarie/1")
    .then((res) => {
      if (!res.ok) {
        throw new Error("Erreur réseau");
      }
      // On convertit la réponse en JSON
      return res.json();
    })
    .then((articles) => {
      // Cette partie s'exécute quand on a reçu les articles
      const conteneurPublies = document.getElementById("liste-publies");
      const conteneurBrouillons = document.getElementById("liste-brouillons");

      // On vide les conteneurs
      conteneurPublies.innerHTML = "";
      conteneurBrouillons.innerHTML = "";

      let compteurPublies = 0;

      articles.forEach((art) => {
        // 1. Choix du style de badge selon le statut
        let badgeStatut = "";
        if (art.statut === "brouillon")
          badgeStatut = `<span class="tag t-amber">Brouillon</span>`;
        else if (art.statut === "en attente")
          badgeStatut = `<span class="tag t-blue">En attente</span>`;
        else if (art.statut === "refuse")
          badgeStatut = `<span class="tag t-red">Refusé</span>`;
        else {
          badgeStatut = `<span class="tag t-green">Publié</span>`;
          compteurPublies++;
        }

        // 2. Choix de l'icône selon le type
        let icone = "📝";
        if (art.type.toLowerCase().includes("conseil")) icone = "💡";
        if (art.type.toLowerCase().includes("news")) icone = "📰";
        if (art.type.toLowerCase().includes("tuto")) icone = "🛠️";

        // 3. Création de la carte HTML en Template Literal
        const card = `
          <div class="post-item" onclick="editerArticle(${art.id_article})">
            <div class="post-ico" style="background:rgba(48,212,192,.09)">${icone}</div>
            <div class="post-body">
              <div class="post-title">${art.titre}</div>
              <div class="post-excerpt">${art.contenu.substring(0, 65)}...</div>
              <div class="post-meta">
                  <span class="tag t-vi">${art.type}</span>
                  ${badgeStatut}
                  <button class="btn btn-v btn-sm" style="flex-shrink: 0" onclick="editerArticle(${art.id_article})">
            ＋ Modifier
              </div>
          </button>
            </div>
          </div>
        `;

        // 4. Tri : Les brouillons et refusés à droite, le reste à gauche
        if (art.statut === "brouillon" || art.statut === "refuse") {
          conteneurBrouillons.innerHTML += card;
        } else {
          conteneurPublies.innerHTML += card;
        }
      });

      // Mettre à jour le petit compteur en haut de section
      document.getElementById("compteur-articles").textContent =
        `${compteurPublies} en ligne`;
    })
    .catch((error) => {
      // Cette partie s'exécute s'il y a une erreur (serveur éteint, etc.)
      console.error("Impossible de récupérer les articles", error);
      document.getElementById("liste-publies").innerHTML =
        `<p style="color:var(--red); font-size:13px;">Serveur indisponible.</p>`;
    });
}

// Fonction appelée quand on clique sur un article
function editerArticle(id) {
  // Ici, on va appeler une route pour récupérer l'article par son ID,
  // remplir le 'postModal' avec les valeurs, et l'ouvrir !
  console.log("Édition de l'article :", id);
  // openNewPost(); // On ouvrira ton modal existant
}

function editerArticle(id) {
  // 1. Appel à ton API Go pour récupérer l'article par son ID
  fetch(`http://localhost:8081/admin/articles/${id}`)
    .then((res) => {
      if (!res.ok) throw new Error("Impossible de récupérer l'article");
      return res.json();
    })
    .then((art) => {
      document.getElementById("edit-article-id").value = art.id_article;
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
  // 1. Récupération des données du formulaire
  const id = document.getElementById("edit-article-id").value;
  const titre = document.getElementById("post-title").value;
  const contenu = document.getElementById("post-content").value;
  const type = document.getElementById("post-type").value;

  if (!titre || !contenu) {
    alert("Le titre et le contenu sont obligatoires !");
    return;
  }

  const articleData = {
    id_salarie: 1, // À remplacer par l'ID réel du salarié connecté apre sync avec faty
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
    },
    body: JSON.stringify(articleData),
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur lors de la sauvegarde");
      return res.json();
    })
    .then((data) => {
      // Succès !
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

// Lancer la fonction au chargement de la page
document.addEventListener("DOMContentLoaded", chargerArticles);
