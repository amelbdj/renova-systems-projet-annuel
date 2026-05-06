monToken = localStorage.getItem("token");
function GetConteneurs() {
  const container = document.getElementById("box-container");
  const statContainer = document.getElementById("box-stats");
  if (!container) return;

  container.innerHTML =
    "<div style='color:var(--txt-m)'>Chargement des conteneurs...</div>";

  fetch("http://localhost:8081/api/admin/conteneurs", {
    headers: { Authorization: "Bearer " + monToken },
  })
    .then((response) => response.json())
    .then((conteneurs) => {
      container.innerHTML = "";
      let totalBoxes = 0;

      if (!conteneurs || conteneurs.length === 0) {
        container.innerHTML =
          "<div style='color:var(--txt-m)'>Aucun conteneur déployé.</div>";
      } else {
        conteneurs.forEach((c) => {
          totalBoxes += c.total_boxes;

          container.innerHTML += `
              <div class="box ok" onclick="GetBoxesForConteneur(${c.id}, '${c.nom}', '${c.adresse}')" style="cursor: pointer; width: 100%; grid-column: span 1;">
                <div class="box-id" style="font-size: 14px; text-transform: uppercase;">${c.nom}</div>
                <div class="box-label" style="margin-top: 15px; color: #fff; font-size: 11px;"><span class="material-symbols-outlined">location_on</span> ${c.adresse}</div>
                <div class="box-label" style="margin-top: 8px; color: var(--txt-m); font-size: 11px;">
                    <span class="material-symbols-outlined">inventory_2</span> ${c.total_boxes} portes au total
                </div>
              </div>
            `;
        });
      }

      statContainer.innerHTML = `
        <div class="log-stat">
          <div class="log-stat-ico" style="background: rgba(46, 204, 113, 0.1)">🗺️</div>
          <div>
            <div class="log-stat-val">${conteneurs ? conteneurs.length : 0}</div>
            <div class="log-stat-lbl">Conteneurs déployés</div>
          </div>
        </div>
        <div class="log-stat">
          <div class="log-stat-ico" style="background: rgba(74, 144, 240, 0.1)">📦</div>
          <div>
            <div class="log-stat-val">${totalBoxes}</div>
            <div class="log-stat-lbl">Casiers au total</div>
          </div>
        </div>
        <div style="margin-top:8px;display:flex;flex-direction:column;gap:7px">
          <button class="btn btn-g btn-sm btn-full" onclick="openNewBox()">＋ Ajouter un conteneur</button>
          <button class="btn btn-o btn-sm btn-full" onclick="alert('Rapport logistique PDF généré.')">📄 Rapport logistique</button>
        </div>`;
    })
    .catch((error) => console.error("Erreur de récupération :", error));
}

function GetBoxesForConteneur(conteneurId, nomConteneur, adresseConteneur) {
  const container = document.getElementById("box-container");
  const statContainer = document.getElementById("box-stats");

  container.innerHTML =
    "<div style='color:var(--txt-m)'>Chargement des casiers</div>";

  fetch(`http://localhost:8081/api/admin/conteneur/${conteneurId}/boxes`, {
    headers: { Authorization: "Bearer " + monToken },
  })
    .then((response) => response.json())
    .then((boxs) => {
      container.innerHTML = `
        <div style="grid-column: 1 / -1; display: flex; justify-content: space-between; align-items: center; margin-bottom: 10px; padding-bottom: 10px; border-bottom: 1px solid var(--b1);">
            <div style="color: white; font-weight: bold;">Casiers du : ${nomConteneur}</div>
            <button class="btn btn-sm btn-o" onclick="GetConteneurs()">← Retour</button>
        </div>
      `;

      let libreCount = 0;
      let occupeCount = 0;
      let maintenanceCount = 0;

      if (!boxs || boxs.length === 0) {
        container.innerHTML += `<div style="grid-column: 1/-1; color: var(--txt-m);">Ce conteneur est vide.</div>`;
      } else {
        boxs.forEach((box) => {
          let statusClass = "";
          let ledColor = "";
          const etat = box.statut ? box.statut.toUpperCase() : "INCONNU";

          if (etat === "LIBRE") {
            statusClass = "ok";
            ledColor = "g";
            libreCount++;
          } else if (etat === "OCCUPE" || etat === "RESERVEE") {
            statusClass = "used";
            ledColor = "b";
            occupeCount++;
          } else {
            statusClass = "err";
            ledColor = "r";
            maintenanceCount++;
          }

          container.innerHTML += `
<div class="box ${statusClass}" onclick="openBoxDetail('${box.id}', '${box.numero}', '${box.statut}', '${nomConteneur}', ${conteneurId})">
                  <div class="box-id">PORTE-${box.numero.toString().padStart(2, "0")}</div>
                  <div class="box-led led-${ledColor}"></div>
                  <div class="box-label">${box.statut}</div>
              </div>
            `;
        });
      }

      statContainer.innerHTML = `
        <div class="log-stat">
          <div class="log-stat-ico" style="background: rgba(46, 204, 113, 0.1)">✅</div>
          <div>
            <div class="log-stat-val">${libreCount} / ${boxs ? boxs.length : 0}</div>
            <div class="log-stat-lbl">Box disponibles ici</div>
          </div>
        </div>
        <div class="log-stat">
          <div class="log-stat-ico" style="background: rgba(74, 144, 240, 0.1)">📦</div>
          <div>
            <div class="log-stat-val">${occupeCount}</div>
            <div class="log-stat-lbl">Dépôts actifs ici</div>
          </div>
        </div>
        <div class="log-stat" style="border-color: rgba(240, 80, 80, 0.2); background: rgba(240, 80, 80, 0.03);">
          <div class="log-stat-ico" style="background: rgba(240, 80, 80, 0.1)">⚠️</div>
          <div>
            <div class="log-stat-val" style="color: var(--red)">${maintenanceCount}</div>
            <div class="log-stat-lbl">Erreurs détectées</div>
          </div>
        </div>
        
        <!-- C'EST ICI QU'ON RAJOUTE LE BOUTON 👇 -->
        <div style="margin-top:8px; display:flex; flex-direction:column; gap:8px;">
          <button class="btn btn-o btn-sm btn-full" onclick="AddSingleBox(${conteneurId}, '${nomConteneur}')">
            ＋ Ajouter une porte (Box)
          </button>
          
          <button class="btn btn-red btn-sm btn-full" onclick="alert('Maintenance demandée pour le meuble.')">
            🛠️ Envoyer maintenance
          </button>
        </div>`;
    })
    .catch((error) => console.error("Erreur des casiers:", error));
}

let currentBoxId = null;
let currentConteneurId = null;
let currentConteneurNom = "";

function openBoxDetail(id, numero, status, localisation, conteneurId) {
  currentBoxId = id;
  currentConteneurId = conteneurId;
  currentConteneurNom = localisation;

  document.getElementById("boxModalTitle").textContent =
    "PORTE n°" + numero + " — " + status.toUpperCase();
  document.getElementById("boxModalStatus").textContent = status.toUpperCase();
  document.getElementById("boxModalLocation").textContent = localisation;

  document.getElementById("boxModal").classList.add("open");
}

function openNewBox() {
  document.getElementById("newBoxModal").classList.add("open");
}

function CreateBox() {
  const adresseInput = document.getElementById("boxLocation").value;
  const nomInput = document.getElementById("boxType").value;
  const capacityInput = parseInt(document.getElementById("boxCapacity").value);

  if (!adresseInput || !nomInput || isNaN(capacityInput)) {
    alert("Veuillez remplir tous les champs !");
    return;
  }

  const newConteneurData = {
    nom: nomInput,
    adresse: adresseInput,
    nombre_de_boxs: capacityInput,
  };

  fetch("http://localhost:8081/api/admin/conteneur/create", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
    body: JSON.stringify(newConteneurData),
  })
    .then((response) => {
      if (response.ok) {
        alert("Nouveau Conteneur déployé avec succès !");
        GetConteneurs();
        document.getElementById("newBoxModal").classList.remove("open");
      } else {
        alert("Erreur lors du déploiement.");
      }
    })
    .catch((error) => console.error("Erreur de création:", error));
}

function AddSingleBox(conteneurId, nomConteneur) {
  const taille = prompt(
    "Quelle taille pour cette nouvelle box ? (Tapez S, M ou L)",
    "M",
  );

  if (!taille) return;

  fetch("http://localhost:8081/api/admin/box/add", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
    body: JSON.stringify({
      id_conteneur: parseInt(conteneurId),
      taille: taille.toUpperCase(),
    }),
  })
    .then((response) => {
      if (response.ok) {
        alert("Nouvelle porte ajoutée !");
        GetBoxesForConteneur(conteneurId, nomConteneur, "");
      } else {
        alert("Erreur lors de l'ajout de la box.");
      }
    })
    .catch((error) => console.error("Erreur:", error));
}

function UpdateBoxStatusAPI() {
  const selectStatut = document.getElementById("selectBoxStatus");
  if (!selectStatut) return;

  const nouveauStatut = selectStatut.value;

  fetch("http://localhost:8081/api/admin/box/update", {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
    body: JSON.stringify({
      box_id: parseInt(currentBoxId),
      statut: nouveauStatut,
    }),
  })
    .then((response) => {
      if (response.ok) {
        alert("Statut mis à jour avec succès !");
        closeModal("boxModal");
        GetBoxesForConteneur(currentConteneurId, currentConteneurNom, "");
      } else {
        alert("Erreur lors de la mise à jour.");
      }
    })
    .catch((error) => console.error("Erreur:", error));
}

document.addEventListener("DOMContentLoaded", GetConteneurs);
