
document.addEventListener("DOMContentLoaded", async () => {
  const token = localStorage.getItem("token");
  if (!token) {
    window.location.href = "login.html";
    return;
  }

  const userId = localStorage.getItem("userId");
  const firstName = localStorage.getItem("userName");

  const navNameDisplay = document.getElementById("navName");
  const heroNameDisplay = document.getElementById("heroName");

  if (firstName) {
    if (navNameDisplay) navNameDisplay.textContent = firstName;
    if (heroNameDisplay) heroNameDisplay.textContent = firstName;
  }

  const urlParams = new URLSearchParams(window.location.search);
  
  if (urlParams.get("abo") === "success") {
    const sessionId = urlParams.get("session_id");
    try {
      const upgradeRes = await fetch(`http://localhost:8081/api/pro/upgrade?id=${userId}&session_id=${sessionId}`, {
        method: "POST",
        headers: { "Authorization": `Bearer ${token}` }
      });

      if (upgradeRes.ok) {
<<<<<<< HEAD
        alert("Payment success! vous etes maintenant Premium. Profitez de votre abonnement!");
        window.history.replaceState(null, "", window.location.pathname); 
=======
        alert("Paiement réussi ! Votre abonnement est maintenant actif. 🎉");
        window.history.replaceState(null, "", window.location.pathname);
>>>>>>> origin/faty
      }
    } catch (err) {
      console.error("Error upgrading account:", err);
    }
  } else if (urlParams.get("abo") === "cancel") {
    alert("Payment cancelled. You can upgrade anytime!");
    window.history.replaceState(null, "", window.location.pathname);
  }

  await checkPremiumStatus(token, userId);
<<<<<<< HEAD
});


async function subscribeToPremium() {
=======
  loadMyProAnnonces();
});

let currentEditingImagePath = "";

function toggleAnnForm() {
  const f = document.getElementById("annForm");
  const open = f.style.display === "block";
  f.style.display = open ? "none" : "block";
  if (!open) f.scrollIntoView({ behavior: "smooth", block: "start" });
}

function togglePriceField() {
  const typeSelect = document.querySelector("#annForm select").value;
  const priceContainer = document.getElementById("priceContainer");
  const priceInput = document.querySelector('#annForm input[type="number"]');
  if (typeSelect === "Don gratuit" || typeSelect === "don") {
    priceContainer.style.display = "none";
    priceInput.value = "0";
  } else {
    priceContainer.style.display = "block";
    if (priceInput.value === "0") priceInput.value = "";
  }
}

async function loadMyProAnnonces() {
  const userId = localStorage.getItem("userId");
  const list = document.getElementById("proAnnList");
  if (!list) return;

  try {
    const res = await fetch(`http://localhost:8081/mes-annonces?id=${userId}`);
    const toutes = await res.json();

    const actives = (toutes || []).filter(
      (ann) =>
        ann.statut_vente !== "VENDU" &&
        ann.statut_vente !== "EN ATTENTE DEPOT" &&
        ann.statut_vente !== "EN BOX",
    );

    if (actives.length === 0) {
      list.innerHTML = `<div style="text-align:center;color:var(--txt-m);font-size:13px;padding:16px">Vous n'avez pas encore d'annonces.</div>`;
      return;
    }

    list.innerHTML = "";
    actives.forEach((ann) => {
      const imgSrc = ann.image ? `http://localhost:8081${ann.image}` : null;
      const valide = (ann.statut_validation || "").toLowerCase().startsWith("valid");
      const tagCls = valide ? "t-green" : "t-amber";

      const item = document.createElement("div");
      item.className = "mat-item";
      item.innerHTML = `
        <div class="mat-thumb" style="${imgSrc ? `background:url('${imgSrc}') center/cover no-repeat` : "background:rgba(58,142,255,.08)"}">${imgSrc ? "" : "📦"}</div>
        <div class="mat-info">
          <div class="mat-name">${ann.titre} <span class="tag ${tagCls}" style="margin-left:4px">${ann.statut_validation || ""}</span></div>
          <div class="mat-meta">${ann.prix > 0 ? ann.prix + " €" : "Don gratuit"}</div>
        </div>
        <div class="mat-action" style="display:flex;gap:6px">
          <button class="btn btn-g btn-xs" onclick='openEditForm(${JSON.stringify(ann)})'>Modifier</button>
          <button class="btn btn-xs" style="background:rgba(255,90,90,.1);color:#ff5a5a;border:1px solid rgba(255,90,90,.2)" onclick="deleteAnnoncePro(${ann.id})">Supprimer</button>
        </div>`;
      list.appendChild(item);
    });
  } catch (err) {
    console.error("Erreur chargement de mes annonces :", err);
    list.innerHTML = `<div style="text-align:center;color:var(--red);font-size:13px;padding:16px">Erreur de chargement.</div>`;
  }
}

function openEditForm(ann) {
  document.getElementById("editAnnId").value = ann.id;
  document.getElementById("formTitle").textContent = "Modifier l'annonce";
  document.querySelector('#annForm input[type="text"]').value = ann.titre;
  document.querySelector("#annForm select").value = ann.type;
  document.getElementById("annCategorie").value = ann.id_categorie;
  document.querySelector('#annForm input[type="number"]').value = ann.prix;
  document.querySelector("#annForm textarea").value = ann.description || "";
  currentEditingImagePath = ann.image || "";
  togglePriceField();
  const f = document.getElementById("annForm");
  if (f.style.display !== "block") toggleAnnForm();
}

async function submitAnn() {
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("token");
  const editId = document.getElementById("editAnnId").value;

  const formData = new FormData();
  formData.append("titre", document.querySelector('#annForm input[type="text"]').value);
  formData.append("type", document.querySelector("#annForm select").value);
  formData.append("id_categorie", document.getElementById("annCategorie").value);
  let prixValue = document.querySelector('#annForm input[type="number"]').value;
  if (!prixValue) prixValue = "0";
  formData.append("prix", prixValue);
  formData.append("description", document.querySelector("#annForm textarea").value);
  formData.append("id_user", userId);
  formData.append("code_postal", "75000");
  formData.append("ville", "Paris");
  formData.append("etat", "Bon état");
  formData.append("poids_kg", "1.0");
  formData.append("quantite", "1");

  const photoInput = document.getElementById("annPhoto");
  if (photoInput.files[0]) {
    formData.append("image", photoInput.files[0]);
  } else if (editId) {
    formData.append("old_image_path", currentEditingImagePath);
  }

  const url = editId
    ? `http://localhost:8081/admin/annonces/modify/${editId}`
    : "http://localhost:8081/admin/annonces/add";
  const method = editId ? "PUT" : "POST";

  try {
    const response = await fetch(url, {
      method,
      headers: { Authorization: `Bearer ${token}` },
      body: formData,
    });
    if (response.ok) {
      document.getElementById("editAnnId").value = "";
      document.getElementById("formTitle").textContent = "Créer une annonce";
      currentEditingImagePath = "";
      toggleAnnForm();
      loadMyProAnnonces();
      if (document.getElementById("sponsorSection")?.style.display === "block") {
        loadSponsorAnnonces();
      }
    } else {
      alert("Erreur: " + (await response.text()));
    }
  } catch (err) {
    console.error("Submit error:", err);
  }
}

async function deleteAnnoncePro(id) {
  if (!confirm("Voulez-vous supprimer l'annonce ?")) return;
  const token = localStorage.getItem("token");
  try {
    const response = await fetch(`http://localhost:8081/admin/annonces/delete/${id}`, {
      method: "DELETE",
      headers: { Authorization: `Bearer ${token}` },
    });
    if (response.ok) {
      loadMyProAnnonces();
      if (document.getElementById("sponsorSection")?.style.display === "block") {
        loadSponsorAnnonces();
      }
    }
  } catch (err) {
    console.error("Delete error:", err);
  }
}


function subscribeToPremium() {
  return subscribeToPlan("premium");
}

async function subscribeToPlan(plan = "premium") {
>>>>>>> origin/faty
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("token");

  if (!userId || !token) {
    window.location.href = "login.html";
    return;
  }

  try {
    const response = await fetch(`http://localhost:8081/user/profile?id=${userId}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (!response.ok) throw new Error("Erreur lors de la récupération du profil");

    const data = await response.json();
    const user = Array.isArray(data) ? data[0] : data;

    if (!user.stripe_account_id || user.stripe_account_id.trim() === "") {
      const popup = document.getElementById("stripeWarningPopup");
<<<<<<< HEAD
      popup.style.display = "flex"; 
      return; 
    }
    const stripeResponse = await fetch(`http://localhost:8081/api/pro/subscribe?id=${userId}`, {
=======
      popup.style.display = "flex";
      return;
    }
    const stripeResponse = await fetch(`http://localhost:8081/api/pro/subscribe?id=${userId}&plan=${plan}`, {
>>>>>>> origin/faty
      method: "POST",
      headers: { Authorization: `Bearer ${token}` }
    });

    if (!stripeResponse.ok) throw new Error("Erreur lors de la creation de la session Stripe");

    const stripeData = await stripeResponse.json();
<<<<<<< HEAD
    
=======

>>>>>>> origin/faty
    if (stripeData.url) {
      window.location.href = stripeData.url;
    } else {
      alert("Erreur réseau avec Stripe.");
    }

  } catch (error) {
    console.error("Erreur:", error);
    alert("Impossible de vérifier l'état de votre compte.");
  }
}

function goToProfile() {
  const userId = localStorage.getItem("userId");

  if (!userId) {
    window.location.href = "login.html";
    return;
  }

  window.location.href = `profil.html?id=${userId}`;
}


async function checkPremiumStatus(token, userId) {
  try {
    await fetch(`http://localhost:8081/api/pro/sync?id=${userId}`, {
      method: "GET",
      headers: { Authorization: `Bearer ${token}` }
    });
  } catch (syncErr) {
    console.log("sync failed. log page anyways");
  }

try {
    const response = await fetch(`http://localhost:8081/user/profile?id=${userId}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    const data = await response.json();
    const user = Array.isArray(data) ? data[0] : data;

    const btnPasserPremium = document.getElementById("btnPasserPremium");
    const btnPremiumActif = document.getElementById("btnPremiumActif");
    const planSection = document.getElementById("plan"); 
<<<<<<< HEAD
    const premiumDashboard = document.getElementById("premiumDashboard"); 

    if (user.est_premium === 1 || user.est_premium === "1") {
      if (btnPasserPremium) btnPasserPremium.style.display = "none";
      if (btnPremiumActif) btnPremiumActif.style.display = "block"; 
      if (planSection) planSection.style.display = "none"; 
      if (premiumDashboard) premiumDashboard.style.display = "block"; 
      
    } else {
      if (btnPasserPremium) btnPasserPremium.style.display = "block"; 
      if (btnPremiumActif) btnPremiumActif.style.display = "none";
      if (planSection) planSection.style.display = "block"; 
      if (premiumDashboard) premiumDashboard.style.display = "none"; 
    }
    
=======
    const premiumDashboard = document.getElementById("premiumDashboard");
    const invoiceSection = document.getElementById("inv"); 

    const sponsorSection = document.getElementById("sponsorSection");
    const plan = (user.plan_abo || "premium").toLowerCase();

    if (user.est_premium === 1 || user.est_premium === "1") {
      if (btnPasserPremium) btnPasserPremium.style.display = "none";
      if (btnPremiumActif) btnPremiumActif.style.display = "block";
      if (premiumDashboard) premiumDashboard.style.display = "block";
      if (invoiceSection) invoiceSection.classList.remove("blurred");

      applyPlanLabels(plan);
      showUpgradeCards(plan);

      if (sponsorSection) {
        if (plan === "pro") {
          sponsorSection.style.display = "block";
          loadSponsorAnnonces();
        } else {
          sponsorSection.style.display = "none";
        }
      }

    } else {
      if (btnPasserPremium) btnPasserPremium.style.display = "block";
      if (btnPremiumActif) btnPremiumActif.style.display = "none";
      if (planSection) planSection.style.display = "block";
      showAllPlanCards();
      if (premiumDashboard) premiumDashboard.style.display = "none";
      if (invoiceSection) invoiceSection.classList.add("blurred");
      if (sponsorSection) sponsorSection.style.display = "none";
    }

>>>>>>> origin/faty
  } catch (err) {
    console.error("Erreur lors de la verif du profil :", err);
  }
}

<<<<<<< HEAD
=======
function applyPlanLabels(plan) {
  const meta = {
    premium: { label: "Premium actif", name: "Plan Premium", price: "25 €", icon: "fa-star",  color: "var(--gold-l)" },
    plus:    { label: "Plus actif",    name: "Plan Plus",    price: "45 €", icon: "fa-medal", color: "#7fb2ff" },
    pro:     { label: "Pro actif",     name: "Plan Pro",     price: "99 €", icon: "fa-rocket", color: "#c4a7ff" },
  }[plan] || null;
  if (!meta) return;

  const tag = document.getElementById("planActiveTag");
  if (tag) tag.innerHTML = `<i class="fa-solid ${meta.icon}"></i> ${meta.label}`;

  const name = document.getElementById("contractPlanName");
  if (name) {
    name.innerHTML = `<i class="fa-solid ${meta.icon}"></i> ${meta.name}`;
    name.style.color = meta.color;
  }

  const price = document.getElementById("contractPlanPrice");
  if (price) price.innerHTML = `${meta.price}<span>/mois</span>`;
}

function showUpgradeCards(plan) {
  const order = ["premium", "plus", "pro"];
  const idx = order.indexOf(plan);
  const higher = idx === -1 ? [] : order.slice(idx + 1);

  const free = document.getElementById("card-free");
  if (free) free.style.display = "none";

  order.forEach((p) => {
    const card = document.getElementById("card-" + p);
    if (card) card.style.display = higher.includes(p) ? "" : "none";
  });

  const section = document.getElementById("plan");
  const title = document.getElementById("planSectionTitle");
  if (!section) return;

  if (higher.length === 0) {
    section.style.display = "none";
  } else {
    section.style.display = "block";
    if (title) title.textContent = "Améliorer mon abonnement";
  }
}

function showAllPlanCards() {
  const title = document.getElementById("planSectionTitle");
  if (title) title.textContent = "Votre Abonnement";
  ["free", "premium", "plus", "pro"].forEach((p) => {
    const card = document.getElementById("card-" + p);
    if (card) card.style.display = "";
  });
}

async function loadSponsorAnnonces() {
  const userId = localStorage.getItem("userId");
  const list = document.getElementById("sponsorList");
  if (!list) return;

  try {
    const res = await fetch(`http://localhost:8081/mes-annonces?id=${userId}`);
    const annonces = await res.json();

    const sponsorisables = (annonces || []).filter(
      (ann) =>
        ann.statut_vente === "LIBRE" &&
        (ann.statut_validation || "").toLowerCase().startsWith("valid"),
    );

    if (sponsorisables.length === 0) {
      list.innerHTML = `<div style="text-align:center;color:var(--txt-m);font-size:13px;padding:16px">Vous n'avez aucune annonce disponible à sponsoriser pour l'instant.</div>`;
      return;
    }

    list.innerHTML = "";
    sponsorisables.forEach((ann) => {
      const boosted = ann.is_sponsored === true || ann.is_sponsored === 1 || ann.is_sponsored === "1";
      const imgSrc = ann.image ? `http://localhost:8081${ann.image}` : null;

      const item = document.createElement("div");
      item.className = "mat-item";
      item.innerHTML = `
        <div class="mat-thumb" style="${imgSrc ? `background:url('${imgSrc}') center/cover no-repeat` : "background:rgba(166,124,255,.08)"}">${imgSrc ? "" : "📦"}</div>
        <div class="mat-info">
          <div class="mat-name">${ann.titre} ${boosted ? '<span class="tag t-gold" style="margin-left:4px">⭐ Sponsorisé</span>' : ""}</div>
          <div class="mat-meta">${ann.prix > 0 ? ann.prix + " €" : "Don"} · ${ann.statut_validation || ""}</div>
        </div>
        <div class="mat-action">
          <button class="btn ${boosted ? "btn-danger" : "btn-pro"} btn-xs" onclick="toggleSponsor(${ann.id})">
            ${boosted ? "Retirer le boost" : "🚀 Booster"}
          </button>
        </div>`;
      list.appendChild(item);
    });
  } catch (err) {
    console.error("Erreur chargement annonces à sponsoriser :", err);
    list.innerHTML = `<div style="text-align:center;color:var(--red);font-size:13px;padding:16px">Erreur de chargement.</div>`;
  }
}

async function toggleSponsor(annonceId) {
  const token = localStorage.getItem("token");
  const userId = localStorage.getItem("userId");

  try {
    const res = await fetch(`http://localhost:8081/api/pro/annonces/sponsor?id=${userId}&annonce_id=${annonceId}`, {
      method: "POST",
      headers: { Authorization: `Bearer ${token}` },
    });
    if (!res.ok) {
      const e = await res.json().catch(() => ({}));
      alert(e.error || "Impossible de modifier le boost.");
      return;
    }
    await loadSponsorAnnonces();
  } catch (err) {
    console.error("Erreur toggle sponsor :", err);
    alert("Erreur réseau.");
  }
}

>>>>>>> origin/faty

async function cancelPremium() {
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("token");

<<<<<<< HEAD
  try {
    const response = await fetch(`http://localhost:8081/api/pro/portal?id=${userId}`, {
=======
  if (!confirm("Êtes-vous sûr de vouloir résilier votre abonnement ? L'accès Premium sera retiré immédiatement.")) {
    return;
  }

  try {
    const response = await fetch(`http://localhost:8081/api/pro/cancel?id=${userId}`, {
>>>>>>> origin/faty
      method: "POST",
      headers: { Authorization: `Bearer ${token}` }
    });

<<<<<<< HEAD
    if (!response.ok) throw new Error("Could not load portal");

    const data = await response.json();
    
    if (data.url) {
      window.location.href = data.url; 
    }
  } catch (err) {
    console.error("Error:", err);
    alert("Peut pas redirect vers la page abbonement");
  }
=======
    if (!response.ok) throw new Error("Could not cancel subscription");

    alert("Votre abonnement a bien été résilié.");
    await checkPremiumStatus(token, userId);
  } catch (err) {
    console.error("Error:", err);
    alert("Impossible de résilier l'abonnement. Réessayez.");
  }
}

async function loadProjets() {
    const token = localStorage.getItem("token");
    const userId = localStorage.getItem("userId");

    const response = await fetch(`http://localhost:8081/api/pro/projets?id_user=${userId}`, {
        headers: { "Authorization": `Bearer ${token}` }
    });

    const projets = await response.json();
    const grid = document.querySelector(".proj-grid");
    const addBtn = document.querySelector(".proj-add");
    grid.innerHTML = "";
    grid.appendChild(addBtn);

    let nbFinis = 0;
    let nbEnCours = 0;
    let nbEnPause = 0;
    let totalCo2 = 0;
    
    if (projets && projets.length > 0) {
        projets.forEach(p => {
            if (p.statut === 'termine' || p.statut === 'Terminé') {
                nbFinis++; 
                if (p.co2_evite) {
                    totalCo2 += parseFloat(p.co2_evite); 
                }
            } else if (p.statut === 'en_pause' || p.statut === 'En pause') {
                nbEnPause++; 
            } else {
                nbEnCours++; 
            }
        });
    }

    const compteursTermines = document.querySelectorAll(".count-projets-termines");
    compteursTermines.forEach(c => c.textContent = nbFinis);

    const compteursEnCours = document.querySelectorAll(".count-projets-en-cours");
    compteursEnCours.forEach(c => c.textContent = nbEnCours);

    const compteursEnPause = document.querySelectorAll(".count-projets-en-pause");
    compteursEnPause.forEach(c => c.textContent = nbEnPause);

    const compteursCo2 = document.querySelectorAll(".count-co2-total");
    compteursCo2.forEach(c => c.textContent = totalCo2.toFixed(1) + " kg");

    if (!projets || projets.length === 0) return;

    projets.forEach(p => {
        const card = document.createElement("div");
        card.className = "proj-card";
        
        const isTermine = (p.statut === 'termine' || p.statut === 'Terminé');
        const badgeHTML = isTermine 
            ? '<span style="background: rgba(46, 204, 113, 0.2); color: #2ecc71; padding: 4px 8px; border-radius: 12px; font-size: 11px; letter-spacing: 0.5px;">✅ Terminé</span>' 
            : '<span style="background: rgba(245, 197, 66, 0.2); color: #f5c542; padding: 4px 8px; border-radius: 12px; font-size: 11px; letter-spacing: 0.5px;">🔄 En cours</span>';

        const co2Value = p.co2_evite ? parseFloat(p.co2_evite).toFixed(1) : "0.0";

        const avantImg = p.photo_avant ? `http://localhost:8081${p.photo_avant}` : null;
        const apresImg = p.photo_apres ? `http://localhost:8081${p.photo_apres}` : null;
        const heroImg = apresImg || avantImg;
        const noPhoto = `<div style="height:90px;display:flex;align-items:center;justify-content:center;color:var(--txt-m);font-size:11px;background:var(--bg3);border-radius:8px;margin-top:6px">Pas de photo</div>`;

        card.innerHTML = `
            <div class="proj-banner" style="background:linear-gradient(135deg,#100e03,#1e1a06)">
                ${heroImg ? `<img src="${heroImg}" style="width:100%;height:100%;object-fit:cover;border-radius:inherit">` : '🔨'}
                <div class="proj-impact" style="font-weight: bold; color: #2ecc71; background: rgba(0,0,0,0.6); padding: 4px 8px; border-radius: 6px;">🌱 ${co2Value} kg CO₂ évités</div>
            </div>
            <div class="proj-body">
                <div class="proj-name" style="display: flex; align-items: center; justify-content: space-between;">
                    <span>${p.titre}</span>
                    ${badgeHTML}
                </div>
                <div class="proj-desc">${p.description || ''}</div>
                <div class="before-after">
                    <div class="ba-col before">
                        <div class="bal">🗑️ Avant</div>
                        ${avantImg ? `<img src="${avantImg}" style="width:100%;height:90px;object-fit:cover;border-radius:8px;margin-top:6px">` : noPhoto}
                    </div>
                    <div class="ba-col after">
                        <div class="bal">✨ Après</div>
                        ${apresImg ? `<img src="${apresImg}" style="width:100%;height:90px;object-fit:cover;border-radius:8px;margin-top:6px">` : noPhoto}
                    </div>
                </div>
                <div class="proj-foot">
                    <span style="font-size:12px;color:var(--txt-m)">Projet</span>
                    <div style="display:flex;gap:6px">
                        <button class="btn btn-g btn-xs" onclick="openEtapesModal(${p.id}, '${p.titre.replace(/'/g, "\\'")}')">Voir →</button>

                        <button class="btn btn-g btn-xs" onclick="openEditProjet(${p.id}, '${p.titre.replace(/'/g, "\\'")}', '${(p.description||'').replace(/'/g,"\\'")}', '${p.photo_avant||''}', '${p.photo_apres||''}', '${p.statut||'en_cours'}', '${p.co2_evite||0}')">✏️ Modifier</button>

                        <button class="btn btn-xs" style="background:rgba(255,90,90,.1);color:#ff5a5a;border:1px solid rgba(255,90,90,.2)" onclick="deleteProjet(${p.id})">🗑️ Supprimer</button>
                    </div>
                </div>
            </div>
        `;
        grid.insertBefore(card, addBtn);
    });
}

function openEditProjet(id, titre, description, photoAvant, photoApres, statut, co2) {
    const modal = document.getElementById("projModal");
    modal.classList.add("open");

    const inputs = document.querySelectorAll("#projModal input[type='text']");
    const textarea = document.querySelector("#projModal textarea");

    inputs[0].value = titre;
    textarea.value = description;

    document.getElementById("photoAvantLabel").textContent = photoAvant ? "Photo avant conservée" : "Objet récupéré (avant)";
    document.getElementById("photoApresLabel").textContent = photoApres ? "Photo après conservée" : "Après transformation";
    document.getElementById("photoAvantInput").value = "";
    document.getElementById("photoApresInput").value = "";

    const statutInput = document.getElementById("modalStatut");
    const co2Input = document.getElementById("modalCo2");
    if (statutInput) statutInput.value = statut || "en_cours";
    if (co2Input) co2Input.value = co2 || 0;

    modal.dataset.editId = id;
    modal.dataset.oldPhotoAvant = photoAvant || "";
    modal.dataset.oldPhotoApres = photoApres || "";
}

async function saveProjet() {
    const token = localStorage.getItem("token");
    const userId = localStorage.getItem("userId");
    const modal = document.getElementById("projModal");

    const inputs = document.querySelectorAll("#projModal input[type='text']");
    const textarea = document.querySelector("#projModal textarea");
    const photoAvantInput = document.getElementById("photoAvantInput");
    const photoApresInput = document.getElementById("photoApresInput");

    const titre = inputs[0].value.trim();
    if (!titre) { alert("Le nom du projet est obligatoire."); return; }

    const statutInput = document.getElementById("modalStatut");
    const co2Input = document.getElementById("modalCo2");
    const statutVal = statutInput ? statutInput.value : "en_cours";
    const co2Val = co2Input ? co2Input.value.replace(',', '.') : 0;

    const formData = new FormData();
    formData.append("titre", titre);
    formData.append("description", textarea.value.trim());
    formData.append("statut", statutVal);
    formData.append("co2_evite", co2Val);

    if (photoAvantInput.files[0]) formData.append("photo_avant", photoAvantInput.files[0]);
    if (photoApresInput.files[0]) formData.append("photo_apres", photoApresInput.files[0]);

    const editId = modal.dataset.editId;

    if (editId) {
        formData.append("id", editId);
        formData.append("old_photo_avant", modal.dataset.oldPhotoAvant || "");
        formData.append("old_photo_apres", modal.dataset.oldPhotoApres || "");
        await fetch("http://localhost:8081/api/pro/projets/update", {
            method: "PUT",
            headers: { "Authorization": `Bearer ${token}` },
            body: formData
        });
        delete modal.dataset.editId;
        delete modal.dataset.oldPhotoAvant;
        delete modal.dataset.oldPhotoApres;
    } else {
        formData.append("id_user", userId);
        await fetch("http://localhost:8081/api/pro/projets/create", {
            method: "POST",
            headers: { "Authorization": `Bearer ${token}` },
            body: formData
        });
    }

    modal.classList.remove("open");
    await loadProjets();
}

async function deleteProjet(id) {
    if (!confirm("Supprimer ce projet ?")) return;
    const token = localStorage.getItem("token");
    const response = await fetch(`http://localhost:8081/api/pro/projets/delete?id=${id}`, {
        method: "DELETE",
        headers: { "Authorization": `Bearer ${token}` }
    });
    if (response.ok) {
        loadProjets();
    } else {
        alert("Erreur lors de la suppression.");
    }
}

// --- ETAPES ---

let currentProjetId = null;

function openEtapesModal(idProjet, titreProjet) {
    currentProjetId = idProjet;
    document.getElementById("etapesModalTitre").textContent = "🔨 " + titreProjet;
    document.getElementById("etapesModal").classList.add("open");
    loadEtapes(idProjet);
}

function closeEtapesModal() {
    document.getElementById("etapesModal").classList.remove("open");
    currentProjetId = null;
}

async function loadEtapes(idProjet) {
    const token = localStorage.getItem("token");
    const response = await fetch(`http://localhost:8081/api/pro/etapes?id_projet=${idProjet}`, {
        headers: { "Authorization": `Bearer ${token}` }
    });

    const etapes = await response.json();
    const list = document.getElementById("etapesList");
    list.innerHTML = "";

    if (!etapes || etapes.length === 0) {
        list.innerHTML = `<div style="text-align:center;color:var(--txt-m);font-size:13px;padding:16px">Aucune étape pour l'instant.</div>`;
        document.getElementById("etapesCount").textContent = "0 / 0 terminées";
        document.getElementById("etapesProgressBar").style.width = "0%";
        return;
    }

    const terminees = etapes.filter(e => e.statut === "termine").length;
    document.getElementById("etapesCount").textContent = `${terminees} / ${etapes.length} terminées`;
    const pct = Math.round((terminees / etapes.length) * 100);
    document.getElementById("etapesProgressBar").style.width = pct + "%";

    etapes.forEach(e => {
        const statutLabel = { a_faire: "📋 À faire", en_cours: "🔄 En cours", termine: "✅ Terminé" }[e.statut] || e.statut;
        const statutColor = { a_faire: "var(--txt-m)", en_cours: "var(--blue)", termine: "var(--teal-l)" }[e.statut];

        const item = document.createElement("div");
        item.style.cssText = "background:var(--bg3);border:1px solid var(--b1);border-radius:10px;padding:12px 14px;display:flex;align-items:flex-start;gap:10px";
        item.innerHTML = `
            <div style="flex:1">
                <div style="font-size:14px;font-weight:600;color:var(--txt);margin-bottom:4px">${e.titre}</div>
                ${e.description ? `<div style="font-size:12.5px;color:var(--txt-m);margin-bottom:6px">${e.description}</div>` : ''}
                <select onchange="updateStatutEtape(${e.id}, this.value)" style="font-size:12px;padding:4px 8px;border-radius:6px;background:var(--bg);border:1px solid var(--b1);color:${statutColor}">
                    <option value="a_faire" ${e.statut === 'a_faire' ? 'selected' : ''}>📋 À faire</option>
                    <option value="en_cours" ${e.statut === 'en_cours' ? 'selected' : ''}>🔄 En cours</option>
                    <option value="termine" ${e.statut === 'termine' ? 'selected' : ''}>✅ Terminé</option>
                </select>
            </div>
            <button onclick="deleteEtape(${e.id})" style="background:rgba(255,90,90,.1);color:#ff5a5a;border:1px solid rgba(255,90,90,.2);border-radius:6px;padding:4px 8px;font-size:12px;cursor:pointer;flex-shrink:0">🗑️</button>
        `;
        list.appendChild(item);
    });
}

async function addEtape() {
    const token = localStorage.getItem("token");
    const titre = document.getElementById("etapeTitre").value.trim();
    if (!titre) { alert("Le titre est obligatoire."); return; }

    const body = {
        id_projet: currentProjetId,
        titre: titre,
        description: document.getElementById("etapeDescription").value.trim(),
        statut: document.getElementById("etapeStatut").value
    };

    await fetch("http://localhost:8081/api/pro/etapes/create", {
        method: "POST",
        headers: { "Authorization": `Bearer ${token}`, "Content-Type": "application/json" },
        body: JSON.stringify(body)
    });

    document.getElementById("etapeTitre").value = "";
    document.getElementById("etapeDescription").value = "";
    document.getElementById("etapeStatut").value = "a_faire";
    await loadEtapes(currentProjetId);
}

async function deleteEtape(id) {
    if (!confirm("Supprimer cette étape ?")) return;
    const token = localStorage.getItem("token");
    await fetch(`http://localhost:8081/api/pro/etapes/delete?id=${id}`, {
        method: "DELETE",
        headers: { "Authorization": `Bearer ${token}` }
    });
    await loadEtapes(currentProjetId);
}

async function updateStatutEtape(id, statut) {
    const token = localStorage.getItem("token");
    await fetch("http://localhost:8081/api/pro/etapes/statut", {
        method: "PUT",
        headers: { "Authorization": `Bearer ${token}`, "Content-Type": "application/json" },
        body: JSON.stringify({ id, statut })
    });
    await loadEtapes(currentProjetId);
>>>>>>> origin/faty
}