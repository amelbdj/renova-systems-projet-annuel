// On enlève le mot "async" devant la fonction
function GetUsers() {
  const container = document.getElementById("users-container");
  container.innerHTML = "";

  fetch("http://localhost:8081/admin/users")
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur");
      return response.json();
    })
    .then((users) => {
      const badge = document.querySelector(".cnt-badge");
      if (badge) badge.textContent = `${users.length} comptes`;

      users.forEach((user) => {
        const prenom = user.prenom;
        const nom = user.nom;
        const initiales = (prenom[0] + nom[0]).toUpperCase();

        let tagClass = "t-ind";
        let avatarClass = "uav-blue";
        let roleAffiche = user.role || "Utilisateur";

        if (roleAffiche === "Admin") {
          tagClass = "t-actif";
          avatarClass = "uav-green";
        } else if (roleAffiche === "Prestataire") {
          tagClass = "t-pro";
          avatarClass = "uav-gold";
        }

        container.innerHTML += `
                <div class="urow">
                  
                  <div class="uav ${avatarClass}">${initiales}</div>
                  
                  <div class="uinfo">
                    <div class="uname">${prenom} ${nom}</div>
                    <div class="uemail">${user.email}</div>
                  </div>
                  
                  <div class="utags">
                    <span class="t ${tagClass}">${roleAffiche}</span>
                  </div>
                  
                  <div class="umeta">ID : ${user.id}</div>
                  
                  <div class="uicons">
                    <div class="ic" title="Modifier" onclick="OpenEditModal(${user.id}, '${user.nom}', '${user.prenom}', '${user.email}', '${user.role}')">✏</div>
                    <div class="ic del" title="Supprimer" onclick="DeleteUser(${user.id})">🗑</div>
                  </div>
                </div>`;
      });
    })
    .catch((error) => {
      console.error("Erreur API :", error);
      container.innerHTML = `<div style="padding: 20px; color: var(--red);">Impossible de charger les données.</div>`;
    });
}

function DeleteUser(userId) {
  if (confirm("Êtes-vous sûr de vouloir supprimer cet utilisateur ?")) {
    fetch(`http://localhost:8081/admin/users/delete/${userId}`, {
      method: "DELETE", // Utiliser OPTIONS pour contourner les problèmes de CORS
    })
      .then((response) => {
        if (!response.ok) throw new Error("Erreur serveur");
        GetUsers(); // Recharger la liste des utilisateurs
      })
      .catch((error) => {
        console.error("Erreur API :", error);
      });
  }
}

function CreateUser() {
  const nom = document.getElementById("add-nom").value.trim();
  const prenom = document.getElementById("add-prenom").value.trim();
  const email = document.getElementById("add-email").value.trim();
  const role = document.getElementById("add-role").value;
  const mdp = document.getElementById("add-mdp").value;

  if (nom === "" || prenom === "" || email === "" || mdp === "") {
    alert("Veuillez remplir tous les champs.");
    return;
  }

  fetch("http://localhost:8081/admin/users/add", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ nom, prenom, email, role, mot_de_passe: mdp }),
  })
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur");
      GetUsers(); // Recharger la liste des utilisateurs
    })
    .catch((error) => {
      console.error("Erreur API :", error);
    });
}

function OpenEditModal(id, nom, prenom, email, role) {
  const modal = document.getElementById("edit-user-modal");
  modal.style.display = "flex";

  // 2. On remplit les champs avec les données reçues
  document.getElementById("edit-id").value = id;
  document.getElementById("edit-nom").value = nom;
  document.getElementById("edit-prenom").value = prenom;
  document.getElementById("edit-email").value = email;
  document.getElementById("edit-role").value = role;
}

function CloseEditModal() {
  document.getElementById("edit-user-modal").style.display = "none";
}

function UpdateUser(userId) {
  const nom = document.getElementById("edit-nom").value.trim();
  const prenom = document.getElementById("edit-prenom").value.trim();
  const email = document.getElementById("edit-email").value.trim();
  const role = document.getElementById("edit-role").value;

  if (nom === "" || prenom === "" || email === "") {
    alert("Veuillez remplir tous les champs.");
    return;
  }

  fetch(`http://localhost:8081/admin/users/modify/${userId}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ nom, prenom, email, role }),
  })
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur");
      CloseEditModal();
      GetUsers();
    }) // <--- On ferme proprement le .then ici
    .catch((error) => {
      console.error("Erreur API :", error);
    });
}

document.addEventListener("DOMContentLoaded", () => {
  GetUsers();
});
