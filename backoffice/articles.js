function GetArticle() {
  const container = document.getElementById("result");
  if (!container) return;

  fetch("http://localhost:8081/admin/articles")
    .then((res) => {
      if (!res.ok) throw new Error("Erreur serveur articles");
      return res.json();
    })
    .then((articles) => {
      // On vérifie l'ID et non le texte (qui va changer avec la trad)
      const currentTabId = document.querySelector(".vtab.on").id;
      if (currentTabId !== "tout") container.innerHTML = "";

      let htmlContent = "";
      articles.forEach((article) => {
        if (article.statut && article.statut.toLowerCase() === "en attente") {
          htmlContent += `
   <div class="val-item con" data-type="con">
        <div class="val-ico" style="background:rgba(0,212,232,.1)">✍️</div>
        <div class="val-body">
          <div class="val-title">${article.titre}</div>
          <div class="val-meta"> Article · ${article.prenom_auteur} ${article.nom_auteur} · ${article.type} </div>
          <div class="val-actions">
            <button class="va-btn va-ok"   onclick="ValidateArticle(${article.id})">✓ Publier</button>
            <button class="va-btn va-no"   onclick="RefuseArticle(${article.id})">✕ Refuser</button>
            <button class="va-btn va-view" onclick="openArticleModal(${article.id})">👁 Lire</button>
          </div>
        </div>
        <span class="tag t-cyan" style="flex-shrink:0;font-size:10px">Contenu</span>
      </div>`;
        }
      });

      if (htmlContent) {
        container.innerHTML += htmlContent;
      } else if (!container.innerHTML) {
        container.innerHTML = `<div style="padding:20px" data-i18n="backoffice.ads.no_ads">Aucun article en attente.</div>`;
      }

      // On demande au script de traduire les nouveaux éléments fraîchement injectés
      if (typeof appliquerTraductions === "function") {
        appliquerTraductions();
      }
    })
    .catch((err) => console.error(err));
}

function openArticleModal(id) {
  document.getElementById("articleModModal").style.display = "flex";

  fetch(`http://localhost:8081/admin/articles/${id}`)
    .then((res) => {
      if (!res.ok) throw new Error("Impossible de charger l'article");
      return res.json();
    })
    .then((article) => {
      document.getElementById("modal-art-title").textContent = article.titre;
      document.getElementById("modal-art-meta").innerHTML = `
        <span class="tag t-vi">${article.type}</span> 
        • Rédigé par <b>${article.prenom_auteur} ${article.nom_auteur}</b>
      `;
      document.getElementById("modal-art-content").textContent =
        article.contenu;

      document.getElementById("btn-modal-valider").onclick = function () {
        ValidateArticle(id);
      };
      document.getElementById("btn-modal-refuser").onclick = function () {
        RefuseArticle(id);
      };
    })
    .catch((err) => {
      console.error(err);
      document.getElementById("modal-art-title").textContent = "Erreur";
      document.getElementById("modal-art-content").textContent =
        "Le contenu n'a pas pu être chargé.";
    });
}

function closeArticleModal() {
  document.getElementById("articleModModal").style.display = "none";
}

function closeArticleModal() {
  document.getElementById("articleModModal").style.display = "none";
  articleActuelId = null;
}
function ValidateArticle(id) {
  fetch(`http://localhost:8081/admin/articles/validate/${id}`, {
    method: "PUT",
  }).then(() => {
    alert("Article validé et publié !");
    GetArticle();
  });
}

function RefuseArticle(id) {
  fetch(`http://localhost:8081/admin/articles/refuse/${id}`, {
    method: "PUT",
  }).then(() => {
    alert("Article refusé.");
    GetArticle();
  });
}

document.addEventListener("DOMContentLoaded", () => {
  GetArticle();
});
