let allAnnonces = [];
let isListView = false;

// Si l'utilisateur est un Pro, on colore la page en teal (comme le dashboard pro)
if (localStorage.getItem("userRole") === "Pro") {
  document.documentElement.classList.add("theme-pro");
}

if (!localStorage.getItem("token") || !localStorage.getItem("userId")) {
  window.location.replace("/login");
}

async function loadAllAnnonces() {
  const grid = document.getElementById("listingsGrid");
  if (!grid) return;

  const userId = localStorage.getItem("userId") || 0;

  try {
    const response = await fetch(
      `${API_BASE_URL}/api/annonces/all?id=${userId}`,
    );
    const data = await response.json();

    allAnnonces = data.filter(
      (ann) =>
        ann.statut_vente !== "VENDU" &&
        ann.statut_vente !== "Vendu" &&
        ann.statut_vente !== "EN ATTENTE DEPOT" &&
        ann.statut_vente !== "EN BOX",
    );

    displayAnnonces(allAnnonces);
    buildCategoryFilter();
  } catch (err) {
    console.error("Erreur chargement marketplace:", err);
    grid.innerHTML =
      '<p style="color:var(--red); text-align:center; padding:50px;" data-i18n="annonce.error.connection">Erreur de connexion au serveur.</p>';
    if (typeof appliquerTraductions === "function") appliquerTraductions();
  }
}

function displayAnnonces(items) {
  const grid = document.getElementById("listingsGrid");
  grid.innerHTML = "";

  if (!items || items.length === 0) {
    grid.innerHTML =
      '<div class="empty"><div class="empty-ico">🔍</div><div class="empty-title" data-i18n="annonce.empty.title">Aucune annonce validée</div></div>';
    document.getElementById("resultCount").textContent = "0";
    if (typeof appliquerTraductions === "function") appliquerTraductions();
    return;
  }

  document.getElementById("resultCount").textContent = items.length;

  items.forEach((ann) => {
    const imgSrc = ann.image ? `${API_BASE_URL}${ann.image}` : null;

    const card = document.createElement("a");
    card.className = "listing-card fu";

    card.href = `/annonce?id=${ann.id}`;

    const isFree = ann.prix <= 0 || ann.type === "don";

    const isSponsored = ann.is_sponsored === true || ann.is_sponsored === 1 || ann.is_sponsored === "1";
    let promoBadge = "";
    if (isSponsored) {
      promoBadge = '<span class="badge" data-i18n="annonce.badge.sponsored" style="background:rgba(166,124,255,.92);color:#fff">⭐ Sponsorisé</span>';
    } else if (ann.plan_abo === "plus" || ann.plan_abo === "pro") {
      promoBadge = '<span class="badge" data-i18n="annonce.badge.priority" style="background:rgba(58,142,255,.92);color:#fff">⚡ Prioritaire</span>';
    }

    card.innerHTML = `
            <div class="card-thumb" style="${imgSrc ? `background: url('${imgSrc}') center/cover no-repeat;` : `background: var(--bg4);`}">
                ${imgSrc ? "" : '<div style="font-size:3rem"><i class="fas fa-box"></i></div>'}
                <div class="card-badges">
                    ${isFree ? '<span class="badge b-don" data-i18n="annonce.type.donation">Don gratuit</span>' : '<span class="badge b-ven" data-i18n="annonce.type.sale">Vente</span>'}
                    ${promoBadge}
                </div>
            </div>
            <div class="card-body">
                <div class="card-cat">${ann.categorie || '<span data-i18n="annonce.card.object">Objet</span>'}</div>
                <div class="card-title">${ann.titre}</div>
                <div class="card-desc">${ann.description}</div>
                <div class="card-meta">
                    <div class="card-seller"><div class="seller-ava"><i class="fas fa-user"></i></div>${ann.prenom} ${ann.nom}</div>
                    <span style="color:var(--txt-d)">·</span>
                    <span class="card-loc"><i class="fas fa-map-marker-alt"></i> ${ann.ville}</span>
                </div>
                <div class="card-foot">
                    <div>
                        <div class="card-price ${isFree ? "free" : ""}">${isFree ? '<span data-i18n="annonce.card.free">Gratuit</span>' : ann.prix + " €"}</div>
                        <div class="card-price-sub">${ann.etat || '<span data-i18n="annonce.card.good_condition">Bon état</span>'}</div>
                    </div>
                    <div class="card-arrow" data-i18n="annonce.card.see">Voir l'annonce →</div>
                </div>
            </div>
        `;
    grid.appendChild(card);
  });

  
  if (typeof appliquerTraductions === "function") appliquerTraductions();
}

// Construit la liste des catégories du filtre à partir de celles en base
// (table categorie via /admin/categories). Le nombre affiché est calculé
// sur les annonces déjà chargées.
async function buildCategoryFilter() {
  const container = document.getElementById("catFilterOpts");
  if (!container) return;

  const token = localStorage.getItem("token");

  try {
    const res = await fetch(`${API_BASE_URL}/admin/categories`, {
      headers: { Authorization: "Bearer " + token },
    });
    const categories = await res.json();

    // Compteur du bouton "Toutes"
    const countAll = document.getElementById("catCountAll");
    if (countAll) countAll.textContent = allAnnonces.length;

    categories.forEach((cat) => {
      const nb = allAnnonces.filter(
        (a) => (a.categorie || "").toLowerCase() === cat.libelle.toLowerCase(),
      ).length;

      const opt = document.createElement("div");
      opt.className = "f-opt";
      opt.setAttribute("onclick", `toggleFilter(this,'cat','${cat.libelle}')`);
      opt.innerHTML = `<div class="f-opt-left"><div class="check-box"></div><span><i class="fa-solid fa-tag"></i></span><span>${cat.libelle}</span></div><div class="f-count">${nb}</div>`;
      container.appendChild(opt);
    });
  } catch (err) {
    console.error("Erreur chargement catégories:", err);
  }
}

function filterListings() {
  const q = document.getElementById("searchInput").value.toLowerCase();
  const filtered = allAnnonces.filter((item) => {
    return (
      item.titre.toLowerCase().includes(q) ||
      item.description.toLowerCase().includes(q) ||
      item.ville.toLowerCase().includes(q)
    );
  });
  displayAnnonces(filtered);
}

function setView(v) {
  isListView = v === "list";
  document
    .getElementById("listingsGrid")
    .classList.toggle("list-view", isListView);
  document.getElementById("vGrid").classList.toggle("on", !isListView);
  document.getElementById("vList").classList.toggle("on", isListView);
}

let activeCat = "all";
let activeType = "all";

function toggleFilter(el, group, val) {
  el.closest(".filter-opts")
    .querySelectorAll(".f-opt")
    .forEach((o) => o.classList.remove("on"));
  el.classList.add("on");

  if (group === "cat") activeCat = val;
  if (group === "type") activeType = val;

  const filtered = allAnnonces.filter((item) => {
    const matchCat =
      activeCat === "all" ||
      item.categorie?.toLowerCase() === activeCat.toLowerCase();
    const matchType =
      activeType === "all" ||
      item.type?.toLowerCase() === activeType.toLowerCase();
    return matchCat && matchType;
  });

  displayAnnonces(filtered);
}

function sortListings(val) {
  const arr = [...allAnnonces];
  if (val === "price-asc") arr.sort((a, b) => a.prix - b.prix);
  if (val === "price-desc") arr.sort((a, b) => b.prix - a.prix);
  if (val === "alpha") arr.sort((a, b) => a.titre.localeCompare(b.titre));
  displayAnnonces(arr);
}

document.addEventListener("DOMContentLoaded", loadAllAnnonces);
