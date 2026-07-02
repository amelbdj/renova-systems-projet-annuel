let userId = localStorage.getItem("userId");

function GetArticle() {
  const container = document.getElementById("result");
  if (!container) return;

  fetch(`${API_BASE_URL}/admin/articles`, {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur serveur articles");
      return res.json();
    })
    .then((articles) => {
      if (!articles || !Array.isArray(articles)) {
        articles = [];
      }

      const currentTab = document.querySelector(".vtab.on");
      const currentTabId = currentTab ? currentTab.id : "";

      if (currentTabId !== "tout") container.innerHTML = "";

      let htmlContent = "";

      articles.forEach((article) => {
        const statut =
          article.statut || article.Statut || article.statut_validation;

        if (statut && statut.toLowerCase() === "en attente") {
          const articleId = article.id;

          htmlContent += `
   <div class="val-item con" data-type="con">
        <div class="val-ico" style="background:rgba(0,212,232,.1)">✍️</div>
        <div class="val-body">
          <div class="val-title">${article.titre || article.Titre}</div>
          <div class="val-meta"> Article · ${article.prenom_auteur || ""} ${article.nom_auteur || ""} · ${article.type || ""} </div>
          <div class="val-actions">
            <button class="va-btn va-ok"   onclick="ValidateArticle(${articleId})">✓ Publier</button>
            <button class="va-btn va-no"   onclick="RefuseArticle(${articleId})">✕ Refuser</button>
            <button class="va-btn va-view" onclick="openArticleModal(${articleId})">👁 Lire</button>
          </div>
        </div>
        <span class="tag t-cyan" style="flex-shrink:0;font-size:10px">Contenu</span>
      </div>`;
        }
      });

      if (htmlContent) {
        container.innerHTML += htmlContent;
      } else if (!container.innerHTML || container.innerHTML.trim() === "") {
        container.innerHTML = `<div style="padding:20px" data-i18n="backoffice.ads.no_ads">Aucun article en attente.</div>`;
      }

      if (typeof appliquerTraductions === "function") {
        appliquerTraductions();
      }
    })
    .catch((err) => console.error("Erreur dans GetArticle :", err));
}

let articleActuelId = null;

function openArticleModal(id) {
  if (!id) {
    console.error("Erreur : Aucun ID valide fourni.");
    return;
  }

  document.getElementById("articleModModal").style.display = "flex";

  fetch(`${API_BASE_URL}/admin/articles/${id}`, {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Impossible de charger l'article");
      return res.json();
    })
    .then((article) => {
      document.getElementById("modal-art-title").textContent =
        article.titre || article.Titre;

      document.getElementById("modal-art-meta").innerHTML = `
        <span class="tag t-vi">${article.type || article.Type || "Article"}</span>
        • Rédigé par <b>${article.prenom_auteur || ""} ${article.nom_auteur || ""}</b>
      `;

      const imgElement = document.getElementById("modal-art-image");
      const imageUrl = article.image_url || article.ImageUrl;

      if (imgElement) {
        if (imageUrl && imageUrl.trim() !== "") {
          imgElement.src = `${API_BASE_URL}/` + imageUrl;
          imgElement.style.display = "block";
        } else {
          imgElement.src = "";
          imgElement.style.display = "none";
        }
      }

      document.getElementById("modal-art-content").textContent =
        article.contenu || article.Contenu;

      document.getElementById("btn-modal-valider").onclick = function () {
        ValidateArticle(id);
        closeArticleModal();
      };

      document.getElementById("btn-modal-refuser").onclick = function () {
        RefuseArticle(id);
        closeArticleModal();
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
  articleActuelId = null;
}

function ValidateArticle(id) {
  fetch(`${API_BASE_URL}/admin/articles/validate/${id}`, {
    method: "PUT",
    headers: {
      Authorization: "Bearer " + monToken,
    },
  }).then(() => {
    alert("Article validé et publié !");
    rechargerValidations();
  });
}

function RefuseArticle(id) {
  fetch(`${API_BASE_URL}/admin/articles/refuse/${id}`, {
    method: "PUT",
    headers: {
      Authorization: "Bearer " + monToken,
    },
  }).then(() => {
    alert("Article refusé.");
    rechargerValidations();
  });
}

document.addEventListener("DOMContentLoaded", () => {
  GetArticle();
});
