
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
        alert("Payment success! vous etes maintenant Premium. Profitez de votre abonnement!");
        window.history.replaceState(null, "", window.location.pathname); 
      }
    } catch (err) {
      console.error("Error upgrading account:", err);
    }
  } else if (urlParams.get("abo") === "cancel") {
    alert("Payment cancelled. You can upgrade anytime!");
    window.history.replaceState(null, "", window.location.pathname);
  }

  await checkPremiumStatus(token, userId);
});


async function subscribeToPremium() {
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
      popup.style.display = "flex"; 
      return; 
    }
    const stripeResponse = await fetch(`http://localhost:8081/api/pro/subscribe?id=${userId}`, {
      method: "POST",
      headers: { Authorization: `Bearer ${token}` }
    });

    if (!stripeResponse.ok) throw new Error("Erreur lors de la creation de la session Stripe");

    const stripeData = await stripeResponse.json();
    
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
    const premiumDashboard = document.getElementById("premiumDashboard"); // 👈 GRAB THE DASHBOARD

    if (user.est_premium === 1 || user.est_premium === "1") {
      if (btnPasserPremium) btnPasserPremium.style.display = "none";
      if (btnPremiumActif) btnPremiumActif.style.display = "block"; 
      if (planSection) planSection.style.display = "none"; 
      if (premiumDashboard) premiumDashboard.style.display = "block"; // 👈 SHOW DASHBOARD
      
    } else {
      if (btnPasserPremium) btnPasserPremium.style.display = "block"; 
      if (btnPremiumActif) btnPremiumActif.style.display = "none";
      if (planSection) planSection.style.display = "block"; 
      if (premiumDashboard) premiumDashboard.style.display = "none"; // 👈 HIDE DASHBOARD
    }
    
  } catch (err) {
    console.error("Erreur lors de la verif du profil :", err);
  }
}


async function cancelPremium() {
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("token");

  try {
    const response = await fetch(`http://localhost:8081/api/pro/portal?id=${userId}`, {
      method: "POST",
      headers: { Authorization: `Bearer ${token}` }
    });

    if (!response.ok) throw new Error("Could not load portal");

    const data = await response.json();
    
    if (data.url) {
      window.location.href = data.url; 
    }
  } catch (err) {
    console.error("Error:", err);
    alert("Peut pas redirect vers la page abbonement");
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

    if (!projets || projets.length === 0) return;

    projets.forEach(p => {
        const card = document.createElement("div");
        card.className = "proj-card";
        card.innerHTML = `
            <div class="proj-banner" style="background:linear-gradient(135deg,#100e03,#1e1a06)">
                ${p.photo ? `<img src="http://localhost:8081${p.photo}" style="width:100%;height:100%;object-fit:cover;border-radius:inherit">` : '🔨'}
                <div class="proj-impact">🌱 Impact projet</div>
            </div>
            <div class="proj-body">
                <div class="proj-name">${p.titre}</div>
                <div class="proj-desc">${p.description || ''}</div>
                <div class="before-after">
                    <div class="ba-col before">
                        <div class="bai">🗑️</div>
                        <div class="bal">Avant</div>
                        ${p.avant_desc || ''}
                    </div>
                    <div class="ba-col after">
                        <div class="bai">✨</div>
                        <div class="bal">Après</div>
                        ${p.apres_desc || ''}
                    </div>
                </div>
                <div class="proj-foot">
                    <span style="font-size:12px;color:var(--txt-m)">Projet</span>
                    <div style="display:flex;gap:6px">
                        <button class="btn btn-g btn-xs" onclick="openEtapesModal(${p.id}, '${p.titre.replace(/'/g, "\\'")}')">Voir →</button>
                        <button class="btn btn-g btn-xs" onclick="openEditProjet(${p.id}, '${p.titre}', '${(p.description||'').replace(/'/g,"\\'")}', '${(p.avant_desc||'').replace(/'/g,"\\'")}', '${(p.apres_desc||'').replace(/'/g,"\\'")}', '${p.photo||''}')">✏️ Modifier</button>
                        <button class="btn btn-xs" style="background:rgba(255,90,90,.1);color:#ff5a5a;border:1px solid rgba(255,90,90,.2)" onclick="deleteProjet(${p.id})">🗑️ Supprimer</button>
                    </div>
                </div>
            </div>
        `;
        grid.insertBefore(card, addBtn);
    });
}

function openEditProjet(id, titre, description, avantDesc, apresDesc, photo) {
    document.getElementById("projModal").classList.add("open");
    const inputs = document.querySelectorAll("#projModal input[type='text']");
    const textarea = document.querySelector("#projModal textarea");
    inputs[0].value = titre;
    inputs[1].value = avantDesc;
    inputs[2].value = apresDesc;
    textarea.value = description;
    document.getElementById("photoLabel").textContent = photo ? "Photo actuelle conservée" : "Ajouter une photo";
    // Stocke l'id et l'ancienne photo pour la sauvegarde
    document.getElementById("projModal").dataset.editId = id;
    document.getElementById("projModal").dataset.oldPhoto = photo;
}

async function saveProjet() {
    const token = localStorage.getItem("token");
    const userId = localStorage.getItem("userId");
    const modal = document.getElementById("projModal");

    const inputs = document.querySelectorAll("#projModal input[type='text']");
    const textarea = document.querySelector("#projModal textarea");
    const photoInput = document.getElementById("photoInput");

    const titre = inputs[0].value.trim();
    const avantDesc = inputs[1].value.trim();
    const apresDesc = inputs[2].value.trim();
    const description = textarea.value.trim();

    if (!titre) { alert("Le nom du projet est obligatoire."); return; }

    const formData = new FormData();
    formData.append("titre", titre);
    formData.append("description", description);
    formData.append("avant_desc", avantDesc);
    formData.append("apres_desc", apresDesc);
    if (photoInput.files[0]) formData.append("photo", photoInput.files[0]);

    const editId = modal.dataset.editId;

    if (editId) {
        // Mode modification
        formData.append("id", editId);
        formData.append("old_photo", modal.dataset.oldPhoto || "");
        await fetch("http://localhost:8081/api/pro/projets/update", {
            method: "PUT",
            headers: { "Authorization": `Bearer ${token}` },
            body: formData
        });
        delete modal.dataset.editId;
        delete modal.dataset.oldPhoto;
    } else {
        // Mode création
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
}