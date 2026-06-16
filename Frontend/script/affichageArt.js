const monToken = localStorage.getItem("token");
let articlesData = [];

document.addEventListener("DOMContentLoaded", () => {
  chargerArticlesClient();

  const searchInput = document.getElementById("searchInput");
  if (searchInput) {
    searchInput.addEventListener("keypress", function (event) {
      if (event.key === "Enter") {
        event.preventDefault();
        LancerRecherche();
      }
    });
  }
});

function LancerRecherche() {
  const motCle = document.getElementById("searchInput").value.trim();
  chargerArticlesClient(motCle);
}

function chargerArticlesClient(motCle = "") {
  const container = document.getElementById("liste-articles");
  container.innerHTML =
    "<p style='color: var(--txt-m); text-align: center; grid-column: 1 / -1;' data-i18n=\"article.loading\">Chargement des articles...</p>";
  if (typeof appliquerTraductions === "function") appliquerTraductions();

  let url = `http://localhost:8081/admin/articles`;

  if (motCle !== "") {
    url += `?search=${encodeURIComponent(motCle)}`;
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
    .then((articles) => {
      container.innerHTML = "";

      if (!articles || articles.length === 0) {
        if (motCle !== "") {
          container.innerHTML = `<p style="color: var(--txt-m); text-align: center; grid-column: 1 / -1;">Aucun résultat trouvé pour "<b>${motCle}</b>".</p>`;
        } else {
          container.innerHTML =
            "<p style='color: var(--txt-m); text-align: center; grid-column: 1 / -1;' data-i18n=\"article.empty\">Aucun article publié pour le moment.</p>";
        }
        if (typeof appliquerTraductions === "function") appliquerTraductions();
        return;
      }

      articlesData = articles;
      let htmlContent = "";
      let articlesAffiches = 0;

      articles.forEach((art) => {
        if (art.statut === "valide") {
          articlesAffiches++;

          const idArt = art.id;

          // 🟢 CORRECTION ICI : Ajout de l'adresse du serveur pour les images des cartes
          let imageCover =
            "https://images.unsplash.com/photo-1556910103-1c02745aae4d?w=500"; // Image par défaut
          if (art.image_url && art.image_url.trim() !== "") {
            imageCover = "http://localhost:8081/" + art.image_url;
          }

          // Un petit résumé de 100 caractères
          const resume =
            art.contenu.length > 100
              ? art.contenu.substring(0, 100) + "..."
              : art.contenu;

          htmlContent += `
                <div class="article-card" onclick="OuvrirArticle(${idArt})">
                    <img src="${imageCover}" alt="Image article" class="card-img">
                    <div style="padding: 15px; flex-grow: 1; display: flex; flex-direction: column; align-items: center; text-align: center;">
                        <h3 style="margin-top: 0; color: #ffffff; font-size: 18px;">${art.titre}</h3>
                        <p style="color: var(--txt-m); font-size: 12px; margin-bottom: 8px;"><span class="material-symbols-outlined" style="font-size: 14px; vertical-align: middle;">calendar_today</span> ${art.created_at}</p>
                        <p style="color: var(--txt-m); font-size: 14px; flex-grow: 1;">${resume}</p>
                        <span style="color: var(--blue); font-weight: 600; font-size: 14px; margin-top: 10px;" data-i18n="article.read_more">Lire l'article ➔</span>
                    </div>
                </div>
            `;
        }
      });

      if (articlesAffiches === 0) {
        if (motCle !== "") {
          container.innerHTML = `<p style="color: var(--txt-m); text-align: center; grid-column: 1 / -1;">Aucun article valide trouvé pour "<b>${motCle}</b>".</p>`;
        } else {
          container.innerHTML = `<p style="color: var(--txt-m); text-align: center; grid-column: 1 / -1;" data-i18n="article.empty_valid">Aucun article valide pour le moment.</p>`;
        }
      } else {
        container.innerHTML = htmlContent;
      }

      // On traduit les cartes / messages qu'on vient d'injecter
      if (typeof appliquerTraductions === "function") appliquerTraductions();
    })
    .catch((err) => {
      console.error("Erreur chargement articles :", err);
      container.innerHTML =
        "<p style='color: var(--red); text-align: center; grid-column: 1 / -1;' data-i18n=\"article.error\">Erreur de connexion au serveur.</p>";
      if (typeof appliquerTraductions === "function") appliquerTraductions();
    });
}

function OuvrirArticle(id) {
  const art = articlesData.find((a) => a.id === id);
  if (!art) return;

  // 🟢 CORRECTION ICI AUSSI : Ajout de l'adresse du serveur pour la fenêtre modale
  let imageCover =
    "https://images.unsplash.com/photo-1556910103-1c02745aae4d?w=500";
  if (art.image_url && art.image_url.trim() !== "") {
    imageCover = "http://localhost:8081/" + art.image_url;
  }

  document.getElementById("modalImage").src = imageCover;
  document.getElementById("modalTitre").textContent = art.titre;

  const auteur = `${art.prenom_auteur || ""} ${art.nom_auteur || ""}`.trim();

  const tFn = typeof t === "function" ? t : (k) => k;
  document.getElementById("modalMeta").textContent =
    `${tFn("article.published_on")} ${art.created_at} ${tFn("article.by")} ${auteur}`;
  document.getElementById("modalContenu").innerHTML = art.contenu;

  document.getElementById("articleModal").style.display = "flex";
}

function FermerArticle() {
  document.getElementById("articleModal").style.display = "none";
}

window.onclick = function (event) {
  const modal = document.getElementById("articleModal");
  if (event.target === modal) {
    modal.style.display = "none";
  }
};

const role = localStorage.getItem("role") || "client";
const linkCSS = document.createElement("link");
linkCSS.rel = "stylesheet";

if (role === "Pro") {
  linkCSS.href = "../Frontend/style/pro.css";
} else {
  linkCSS.href = "../Frontend/style/client.css";
}

document.head.appendChild(linkCSS);

function goToProfile() {
  const userId = localStorage.getItem("userId");
  if (!userId) {
    window.location.href = "login.html";
    return;
  }
  window.location.href = `profil.html?id=${userId}`;
}
