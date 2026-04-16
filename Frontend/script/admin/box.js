function GetBox() {
  const container = document.getElementById("box-container");
  const statContainer = document.getElementById("box-stats");
  let libreCount = 0;
  let occupeCount = 0;
  let maintenanceCount = 0;
  if (!container) return;

  container.innerHTML = "";

  fetch("http://localhost:8081/admin/boxs", {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((response) => response.json())
    .then((boxs) => {
      boxs.forEach((box) => {
        let statusClass = "";
        let ledColor = "";

        // .toUpperCase() permet d'être sûr que ça marche même si Go envoie "Libre" ou "libre"
        const etat = box.etat ? box.etat.toUpperCase() : "";

        if (etat === "LIBRE") {
          statusClass = "ok";
          ledColor = "g";
          libreCount++;
        } else if (etat === "OCCUPE") {
          statusClass = "used";
          ledColor = "b";
          occupeCount++;
        } else if (etat === "MAINTENANCE") {
          statusClass = "err";
          ledColor = "r";
          maintenanceCount++;
        } else {
          statusClass = "wait"; // Pour l'état "Vérif" orange si besoin
          ledColor = "o";
        }

        // On injecte directement dans le container pour garder la grille CSS intacte
        container.innerHTML += `
                    <div class="box ${statusClass}" onclick="openBoxDetail('${box.id}', '${box.etat}', '${box.localisation}')">
                        <div class="box-id">BOX-${box.id.toString().padStart(3, "0")}</div>
                        <div class="box-led led-${ledColor}"></div>
                        <div class="box-label">${box.etat}</div>
                    </div>
                `;

        statContainer.innerHTML = `<div class="log-stat">
              <div
                class="log-stat-ico"
                style="background: rgba(46, 204, 113, 0.1)"
              >
                ✅
              </div>
              <div>
                <div class="log-stat-val">${libreCount} / ${boxs.length}</div>
                <div class="log-stat-lbl">Box disponibles</div>
              </div>
            </div>

            <div class="log-stat">
              <div
                class="log-stat-ico"
                style="background: rgba(74, 144, 240, 0.1)"
              >
                📦
              </div>
              <div>
                <div class="log-stat-val">${occupeCount}</div>
                <div class="log-stat-lbl">Dépôts actifs</div>
              </div>
            </div>

            <div
              class="log-stat"
              style="
                border-color: rgba(240, 80, 80, 0.2);
                background: rgba(240, 80, 80, 0.03);
              "
            >
              <div
                class="log-stat-ico"
                style="background: rgba(240, 80, 80, 0.1)"
              >
                ⚠️
              </div>
              <div>
                <div class="log-stat-val" style="color: var(--red)">${maintenanceCount}</div>
                <div class="log-stat-lbl">Erreurs détectées</div>
              </div>
            </div><div style="margin-top:4px;display:flex;flex-direction:column;gap:7px">
          <button class="btn btn-g btn-sm btn-full" onclick="openNewBox()">＋ Ajouter une box</button>
          <button class="btn btn-o btn-sm btn-full" onclick="alert('Rapport logistique PDF généré.')">📄 Rapport logistique</button>
        </div>`;
      });
    })
    .catch((error) =>
      console.error("Erreur lors de la récupération des boxs:", error),
    );
}

function openBoxDetail(id, status, localisation) {
  document.getElementById("boxModalTitle").textContent = id + " — " + status;
  document.getElementById("boxModalStatus").textContent = status;
  document.getElementById("boxModalLocation").textContent = localisation;
  document.getElementById("boxModal").classList.add("open");
}
function openNewBox() {
  document.getElementById("newBoxModal").classList.add("open");
}

function CreateBox() {
  const locationInput = document.getElementById("boxLocation").value;
  const typeInput = document.getElementById("boxType").value;
  const capacityInput = parseInt(document.getElementById("boxCapacity").value);

  const newBoxData = {
    location: locationInput,
    type: typeInput,
    etat: "Libre",
    capacite: capacityInput,
  };

  fetch("http://localhost:8081/admin/box/create", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
    body: JSON.stringify(newBoxData),
  })
    .then((response) => {
      if (response.ok) {
        alert("Box créé avec succès !");
        GetBox(); // Rafraîchir la liste des boxs
        document.getElementById("newBoxModal").classList.remove("open");
      } else {
        alert("Erreur lors de la création du box.");
      }
    })
    .catch((error) => {
      console.error("Erreur lors de la création du box:", error);
      alert("Erreur lors de la création du box.");
    });
}
document.addEventListener("DOMContentLoaded", GetBox);
