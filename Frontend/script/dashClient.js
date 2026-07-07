let currentEditingImagePath = "";
document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");
  if (!token) {
    window.location.href = "/login";
    return;
  }

  const firstName = localStorage.getItem("userName");
  const hasSeenTutorial = localStorage.getItem("tutorielVu");
  const tutorialOverlay = document.getElementById("tut");

  const navNameDisplay = document.getElementById("navName");
  const heroNameDisplay = document.getElementById("heroName");

  if (firstName) {
    if (navNameDisplay) navNameDisplay.textContent = firstName;
    if (heroNameDisplay) heroNameDisplay.textContent = firstName;
  }

  if (hasSeenTutorial === "true") {
    if (tutorialOverlay) tutorialOverlay.style.display = "none";
  } else {
    console.log("showing the tutorial");
  }

  loadMyAnnonces();
  loadAllUserBoxes();
  loadEcoScore();
});

function logout() {
  localStorage.clear();
  window.location.href = "/login";
}

async function loadMyAnnonces() {
  const userId = localStorage.getItem("userId");
  const annGrid = document.querySelector(".ann-grid");
  if (!annGrid) return;

  try {
    const response = await fetch(
      `${API_BASE_URL}/mes-annonces?id=${userId}`,
    );
    const toutesAnnonces = await response.json();

    annGrid.innerHTML = "";

    let annoncesActives = [];
    if (toutesAnnonces && toutesAnnonces.length > 0) {
      annoncesActives = toutesAnnonces.filter(
        (ann) =>
          ann.statut_vente !== "VENDU" &&
          ann.statut_vente !== "EN ATTENTE DEPOT" &&
          ann.statut_vente !== "EN BOX",
      );
    }

    
    const statAnnonces = document.getElementById("statAnnonces");
    if (statAnnonces) statAnnonces.textContent = annoncesActives.length;

    if (!annoncesActives || annoncesActives.length === 0) {
      annGrid.innerHTML =
        "<p style=\"color:var(--txt-m); padding:20px;\">Vous n'avez pas encore d'annonces.</p>";
    } else {
      annoncesActives.forEach((ann) => {
        console.log("Données de l'annonce:", ann);
        const imgSrc = ann.image ? `${API_BASE_URL}${ann.image}` : null;
        const card = document.createElement("div");
        card.className = "ann-card fu";
        const statusClass =
          ann.statut_validation === "Validée" ? "t-green" : "t-amber";

        card.innerHTML = `
                    <div class="ann-thumb" style="${imgSrc ? `background: url('${imgSrc}') center/cover no-repeat;` : `background: linear-gradient(135deg,#0a0f1e,#101828);`}">
            ${imgSrc ? "" : '<i class="fa-solid fa-box-archive"></i>'}
        </div>
                    <div class="ann-body">
                        <div class="ann-name">${ann.titre}</div>
                        <div class="ann-meta">${ann.prix > 0 ? ann.prix + " €" : "Don gratuit"}</div>
                        <div class="ann-foot">
                            <span class="tag ${statusClass}">${ann.statut_validation}</span>
                            <button class="btn btn-xs" 
                                    style="background:rgba(255,92,92,.08); color:var(--red); border:1px solid rgba(255,92,92,.2)" 
                                    onclick="deleteAnnonce(${ann.id})">
                                Supprimer
                            </button>
                            <button class="btn btn-g btn-xs" onclick='openEditForm(${JSON.stringify(ann)})'>Modifier</button>
                        </div>
                    </div>
                `;
        annGrid.appendChild(card);
      });
    }

    const addBox = document.createElement("div");
    addBox.className = "ann-add";
    addBox.onclick = toggleAnnForm;
    addBox.innerHTML =
      `<div class="plus">＋</div><span>${t("client.add_ad")}</span>`;
    annGrid.appendChild(addBox);
  } catch (err) {
    console.error("Erreur chargement annonces:", err);
  }
}

async function deleteAnnonce(id) {
  if (!confirm("Voulez-vous supprimez l'annonce ?")) return;

  const token = localStorage.getItem("token");
  try {
    const response = await fetch(
      `${API_BASE_URL}/admin/annonces/delete/${id}`,
      {
        method: "DELETE",
        headers: { Authorization: `Bearer ${token}` },
      },
    );

    if (response.ok) {
      loadMyAnnonces();
    }
  } catch (err) {
    console.error("Delete error:", err);
  }
}

function openEditForm(ann) {
  document.getElementById("editAnnId").value = ann.id;
  document.getElementById("formTitle").textContent = "Modifier l'annonce";

  document.querySelector('#annForm input[type="text"]').value = ann.titre;
  document.querySelector("#annForm select").value = ann.type;
  document.getElementById("annCategorie").value = ann.id_categorie;
  document.querySelector('#annForm input[type="number"]').value = ann.prix;
  document.querySelector("#annForm textarea").value = ann.description;

  if (ann.image) {
    currentEditingImagePath = ann.image;
  } else {
    currentEditingImagePath = "";
  }

  toggleAnnForm();
}

function goToProfile() {
  const userId = localStorage.getItem("userId");
  const role = localStorage.getItem("role");

  if (!userId) {
    window.location.href = "/login";
    return;
  }

  window.location.href = `/profil?id=${userId}`;
}

async function loadAllUserBoxes() {
  const userId = localStorage.getItem("userId");
  const grid = document.getElementById("systeme-conteneurs"); 

  if (!grid) return;

  try {
    
    const [resDeposits, resPickups] = await Promise.all([
      fetch(`${API_BASE_URL}/api/user/boxes?user_id=${userId}`),
      fetch(`${API_BASE_URL}/api/user/pickups/${userId}`),
    ]);

    
    const deposits = (await resDeposits.json()) || [];
    const pickups = (await resPickups.json()) || [];

    
    const statDepots = document.getElementById("statDepots");
    if (statDepots) statDepots.textContent = deposits.length;

    
    const allItems = [
      ...deposits.map((item) => ({ ...item, typeAction: "depot" })),
      ...pickups.map((item) => ({ ...item, typeAction: "recuperation" })),
    ];

    
    if (allItems.length === 0) {
      grid.innerHTML = `
                <div class="cont-card avail">
                    <div class="cont-body">Vous n'avez aucun objet en box pour le moment.</div>
                    <span class="tag t-green">Prêt pour de nouveaux achats ou ventes</span>
                </div>`;
      return;
    }

    
    grid.innerHTML = allItems
      .map((box) => {
        
        const isDepot = box.typeAction === "depot";
        const pastilleText = isDepot ? t("client.box.deposit_tag") : t("client.box.pickup_tag");
        const pastilleColor = isDepot
          ? "background-color: #f39c12;"
          : "background-color: #27ae60;"; 
        const statutAffichage = isDepot ? box.etat : box.statut_vente;
        const dateAffichage = isDepot
          ? `${t("client.box.reserved_on")} ${new Date(box.date).toLocaleDateString()}`
          : `${t("client.box.deposited_on")} ${new Date(box.date_depot).toLocaleDateString()}`;
        const btnText = isDepot
          ? t("client.box.print_deposit")
          : t("client.box.print_pickup");

        return `
            <div class="cont-card active" style="position: relative;">
                
                <div style="position: absolute; top: -10px; right: -10px; ${pastilleColor} color: white; padding: 4px 10px; border-radius: 12px; font-size: 11px; font-weight: bold; box-shadow: 0 2px 4px rgba(0,0,0,0.2); z-index: 10;">
                    ${pastilleText}
                </div>

                <div class="cont-top">
                    <div>
                        <div class="cont-id">${box.numero_box}</div> 
                        <div style="font-size: 11px; color: var(--blue-l); margin-top: 2px">Statut : ${statutAffichage}</div>
                    </div>
                    <div class="cont-led ${isDepot ? "led-b" : ""}" style="${!isDepot ? "background-color: #27ae60; box-shadow: 0 0 8px #27ae60;" : ""}"></div>
                </div>
                
                <div class="cont-body">
                    <strong>Objet : ${box.objet}</strong><br>
                    <i class="fas fa-map-marker-alt"></i> ${box.lieu}<br>
                    <small>${dateAffichage}</small>
                </div>
                
                <div class="cont-codes">
                    ${
                      isDepot
                        ? `<span class="ccode blue" style="font-size: 1.1em; font-weight: bold;">${t("client.box.deposit_pin")} : ${box.code_pin}</span>`
                        : `<span class="ccode purple" style="font-size: 1.1em; font-weight: bold;">${t("client.box.pickup_barcode")} : ${box.barcode}</span>`
                    }
                </div>
                
                ${
                  !isDepot
                    ? `<div style="background:white; padding:8px; border-radius:4px; margin-top:12px; text-align:center;">
                    <svg class="barcode-img"
                         jsbarcode-value="${box.barcode}"
                         jsbarcode-width="1.2"
                         jsbarcode-height="30"
                         jsbarcode-fontsize="10">
                    </svg>
                </div>`
                    : ""
                }

                <button class="btn btn-g btn-sm" style="margin-top: 12px; width:100%" onclick="window.print()">
                    <i class="fas fa-download"></i> ${btnText}
                </button>
            </div>
        `;
      })
      .join("");

    if (window.JsBarcode) {
      JsBarcode(".barcode-img").init();
    }
  } catch (error) {
    console.error("Erreur lors du chargement des boxes:", error);
    grid.innerHTML = `<p>${t("client.containers_error")}</p>`;
  }
}


window.addEventListener("DOMContentLoaded", loadAllUserBoxes);
function togglePriceField() {
  const typeSelect = document.querySelector("#annForm select").value;
  const priceContainer = document.getElementById("priceContainer");
  const priceInput = document.querySelector('#annForm input[type="number"]');

  if (typeSelect === "Don gratuit" || typeSelect === "don") {
    priceContainer.style.display = "none";
    priceInput.value = "0";
  } else {
    priceContainer.style.display = "block";
    if (priceInput.value === "0") {
      priceInput.value = "";
    }
  }
}

async function loadEcoScore() {
  const userId = localStorage.getItem("userId");

  try {
    const response = await fetch(
      API_BASE_URL + "/api/user/stats?user_id=" + userId,
    );
    const stats = await response.json();

    const monScore = stats.score;

    const elementsScore = document.querySelectorAll(".user-score");

    elementsScore.forEach((element) => {
      element.textContent = monScore;
    });

    const pointsTexte = document.getElementById("dynamic-pts");
    if (pointsTexte) {
      pointsTexte.textContent = monScore + " pts";
    }

    const barreProgression = document.querySelector(".score-lfill");
    if (barreProgression) {
      let pourcentage = (monScore / 1000) * 100;

      if (pourcentage > 100) {
        pourcentage = 100;
      }

      barreProgression.style.setProperty(
        "width",
        pourcentage + "%",
        "important",
      );
    }

    const casesStats = document.querySelectorAll(".sstat-v");
    if (casesStats.length >= 4) {
      casesStats[0].textContent = stats.objets_donnes || 0;
      casesStats[1].textContent = (stats.dechets_evites || 0) + " kg";
      casesStats[2].textContent = "-";
      casesStats[3].textContent = "-";
    }
  } catch (error) {
    console.error("Erreur lors du chargement du score :", error);
  }
}
document.addEventListener("DOMContentLoaded", () => {
  loadMyAnnonces();
  loadAllUserBoxes(); 
  loadEcoScore();
});
