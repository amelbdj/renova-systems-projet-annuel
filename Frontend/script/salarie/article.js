if (!monToken || !userId) {
  window.location.href = "../login.html";
}

function chargerArticles() {
  fetch(`${API_BASE_URL}/admin/articles/salarie/${userId}`, {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur réseau");
      return res.json();
    })
    .then((articles) => {
      const conteneurPublies = document.getElementById("liste-publies");
      const conteneurBrouillons = document.getElementById("liste-brouillons");
      const statArticle = document.getElementById("stat-article");

      let compteurPublies = 0;

      if (conteneurPublies) conteneurPublies.innerHTML = "";
      if (conteneurBrouillons) conteneurBrouillons.innerHTML = "";

      if (!articles) articles = [];

      articles.forEach((art) => {
        let badgeStatut = "";
        let statut = (art.statut || "").toLowerCase();

        if (statut === "brouillon")
          badgeStatut = `<span class="tag t-amber">Brouillon</span>`;
        else if (statut === "en attente")
          badgeStatut = `<span class="tag t-blue">En attente</span>`;
        else if (statut === "refuse")
          badgeStatut = `<span class="tag t-red">Refusé</span>`;
        else {
          badgeStatut = `<span class="tag t-green">Publié</span>`;
          compteurPublies++;
        }

        let typeArt = (art.type || "").toLowerCase();


        let imageHtml = `<div class="post-ico" style="background:rgba(48,212,192,.09)">📝</div>`;
        if (art.image_url && art.image_url !== "") {
          imageHtml = `<img src="${API_BASE_URL}/${art.image_url}" style="width: 50px; height: 50px; border-radius: 8px; object-fit: cover;">`;
        } else if (typeArt.includes("conseil")) {
          imageHtml = `<div class="post-ico" style="background:rgba(48,212,192,.09)"><span class="material-symbols-outlined">lightbulb</span></div>`;
        }

        const card = `
        <div class="post-item" style="display:flex; align-items:center; gap:15px; margin-bottom:15px;">
          ${imageHtml}
          <div class="post-body" style="flex:1;">
            <div class="post-title" style="font-weight:bold;">${art.titre}</div>
            <div class="post-excerpt" style="font-size:13px; color:#666;">${art.contenu ? art.contenu.substring(0, 65) : ""}...</div>
            <div class="post-meta" style="margin-top:8px;">
                <span class="tag t-vi">${art.type || "Article"}</span>
                ${badgeStatut}
                <button class="btn btn-v btn-sm" onclick="editerArticle(${art.id})">＋ Modifier</button>
                <button class="mod-btn mod-ban" style="color:red; background:none; border:none; cursor:pointer;" onclick="DeleteArticle(${art.id})">Supprimer</button>
            </div>
          </div>
        </div>`;

        if (statut === "brouillon" || statut === "refuse") {
          if (conteneurBrouillons) conteneurBrouillons.innerHTML += card;
        } else {
          if (conteneurPublies) conteneurPublies.innerHTML += card;
        }
      });

      if (articles.length === 0) {
        if (conteneurPublies)
          conteneurPublies.innerHTML = `<p style="color:var(--txt-m); font-size:13px; padding: 10px 0;">Aucun article publié.</p>`;
        if (conteneurBrouillons)
          conteneurBrouillons.innerHTML = `<p style="color:var(--txt-m); font-size:13px; padding: 10px 0;">Aucun brouillon en cours.</p>`;
      }

      if (statArticle) statArticle.textContent = compteurPublies;
    })
    .catch((error) =>
      console.error("Impossible de récupérer les articles", error),
    );
}

function editerArticle(id) {
  fetch(`${API_BASE_URL}/admin/articles/${id}`, {
    headers: { Authorization: "Bearer " + monToken },
  })
    .then((res) => res.json())
    .then((art) => {
      document.getElementById("edit-article-id").value = art.id;
      document.getElementById("post-title").value = art.titre;
      document.getElementById("post-content").value = art.contenu;
      document.getElementById("post-type").value = art.type;

      const modalTitle = document.querySelector("#postModal .sec-title-text");
      if (modalTitle) modalTitle.textContent = "Modifier l'article #" + id;

      openNewPost();
    })
    .catch((err) =>
      alert("Erreur lors de la récupération des données de l'article."),
    );
}

function sauvegarderBrouillon() {
  saveArticle("brouillon");
}
function soumettreAValidation() {
  saveArticle("publier");
}

function saveArticle(action) {
  const id = document.getElementById("edit-article-id").value;
  const titre = document.getElementById("post-title").value.trim();
  const contenu = document.getElementById("post-content").value.trim();
  const type = document.getElementById("post-type").value.trim();

  const imageInput = document.getElementById("post-image");
  const imageFile =
    imageInput && imageInput.files.length > 0 ? imageInput.files[0] : null;

  if (!titre || !contenu) {
    alert("Le titre et le contenu sont obligatoires !");
    return;
  }

  const formData = new FormData();
  formData.append("id_salarie", userId);
  formData.append("titre", titre);
  formData.append("contenu", contenu);
  formData.append("type", type);

  if (imageFile) {
    formData.append("image", imageFile);
  }

  let url = id
    ? `${API_BASE_URL}/admin/articles/modify/${id}/${action}`
    : `${API_BASE_URL}/admin/articles/add/${action}`;
  let method = id ? "PUT" : "POST";

  fetch(url, {
    method: method,
    headers: {
      Authorization: "Bearer " + monToken,
    },
    body: formData,
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur lors de la sauvegarde");
      return res.json();
    })
    .then((data) => {
      alert(
        action === "publier"
          ? "Article publié / En attente !"
          : "Brouillon enregistré.",
      );
      closePost();
      chargerArticles();
    })
    .catch((err) => alert("Erreur serveur : " + err.message));
}

function DeleteArticle(id) {
  if (!confirm("Êtes-vous sûr de vouloir supprimer cet article ?")) return;
  fetch(`${API_BASE_URL}/admin/articles/delete/${id}`, {
    method: "DELETE",
    headers: { Authorization: "Bearer " + monToken },
  })
    .then(() => {
      alert("Article supprimé !");
      chargerArticles();
    })
    .catch((err) => alert("Erreur serveur : " + err.message));
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
    document.getElementById("edit-article-id").value = "";
    document.getElementById("post-title").value = "";
    document.getElementById("post-content").value = "";
    document.getElementById("post-type").value = "";

    const imageInput = document.getElementById("post-image");
    if (imageInput) imageInput.value = "";

    const modalTitle = document.querySelector("#postModal .sec-title-text");
    if (modalTitle) modalTitle.textContent = "Rédiger un article";
  }
}

document.addEventListener("DOMContentLoaded", chargerArticles);
