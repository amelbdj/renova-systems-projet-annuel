function setVtab(element, type) {
  document
    .querySelectorAll(".vtab")
    .forEach((btn) => btn.classList.remove("on"));
  element.classList.add("on");

  const container = document.getElementById("result");
  if (!container) return;

  container.innerHTML = "<p style='padding:20px'>Chargement</p>";

  if (type === "ann") {
    GetAnnonce();
  } else if (type === "evt") {
    GetEvent();
  } else if (type === "all") {
    container.innerHTML = "";
    GetAnnonce();
    GetEvent();
  } else {
    container.innerHTML = "<p style='padding:20px'>En cours de dev</p>";
  }
}

function openNewUser() {
  const modal = document.getElementById("NewuserModal");
  if (modal) modal.style.display = "flex";
}

function closeModal(modalId) {
  const modal = document.getElementById(modalId);
  if (modal) modal.style.display = "none";
}

function OpenEditModalAPI(id, nom, prenom, email, role) {
  const modal = document.getElementById("editUserModal");
  if (!modal) return;
  modal.style.display = "flex";

  const inputs = modal.querySelectorAll("input");
  if (inputs[0]) inputs[0].value = prenom;
  if (inputs[1]) inputs[1].value = nom;
  if (inputs[2]) inputs[2].value = email;

  const select = modal.querySelector("select");
  if (select) select.value = role;

  modal.dataset.userId = id;
}

function AfficherTableau(users) {
  const container = document.querySelector(".u-table");
  const totalStat = document.getElementById("totalUser");
  const badge = document.querySelector("#users .sec-label .tag.t-or");

  if (!users) {
    users = [];
  }

  if (badge) badge.textContent = `${users.length} comptes`;
  if (totalStat) totalStat.innerHTML = users.length;

  const headerHTML = `
      <div class="u-thead">
          <div></div>
          <div>Utilisateur</div>
          <div>Rôle</div>
          <div>Score</div>
          <div>ID</div>
          <div>Actions</div>
      </div>`;

  if (users.length === 0) {
    //tab vide
    container.innerHTML =
      headerHTML +
      `<div style="padding:20px; text-align:center; color:var(--txt-m);">Aucun utilisateur trouvé.</div>`;
    return;
  }

  let rowsHTML = headerHTML;
  users.forEach((user) => {
    const init =
      ((user.prenom?.[0] || "") + (user.nom?.[0] || "")).toUpperCase() || "?";

    let tagClass = "t-blue";
    if (user.role === "Admin") tagClass = "t-or";
    else if (user.role === "Salarié" || user.role === "Salarie")
      tagClass = "t-pu";
    else if (user.role === "Professionnel" || user.role === "Pro")
      tagClass = "t-or";

    rowsHTML += `
      <div class="u-row">
          <input type="checkbox" class="u-chk">
          <div class="u-info">
              <div class="u-ava">${init}</div>
              <div>
                  <div class="u-name">${user.prenom} ${user.nom}</div>
                  <div class="u-email">${user.email}</div>
              </div>
          </div>
          <div><span class="tag ${tagClass}">${user.role}</span></div>
          <div><span class="tag t-grn">${user.score || 0}</span></div>
          <div style="font-size:11.5px; color:var(--txt-m)">ID : ${user.id}</div>
          <div class="u-actions">
              <button class="btn btn-xs btn-g" onclick="OpenEditModalAPI(${user.id}, '${user.nom}', '${user.prenom}', '${user.email}', '${user.role}')">✏️</button>
              <button class="btn btn-xs btn-red" onclick="DeleteUser(${user.id})">🚫</button>
          </div>
      </div>`;
  });

  container.innerHTML = rowsHTML;
}

function GetUsers() {
  fetch("http://localhost:8081/admin/users")
    .then((res) => res.json())
    .then(AfficherTableau)
    .catch((err) => console.error("Erreur GET Users:", err));
}

function GetUserByRole(role) {
  fetch(`http://localhost:8081/admin/users/role/${role}`)
    .then((res) => res.json())
    .then(AfficherTableau)
    .catch((err) => console.error("Erreur GET Role:", err));
}

function Search(query, role) {
  url = `http://localhost:8081/admin/users/search?name=${encodeURIComponent(query)}`;
  if (role && role !== "Tous les rôles" && role !== "") {
    url += `&role=${encodeURIComponent(role)}`;
  }

  fetch(url)
    .then((res) => {
      if (!res.ok) throw new Error("Erreur de recherche");
      return res.json();
    })
    .then(AfficherTableau)
    .catch((err) => {
      console.error("Erreur GET Search:", err);
      const container = document.querySelector(".u-table");
      if (container)
        container.innerHTML = `<div style="padding:20px; color:red;">Erreur lors de la recherche.</div>`;
    });
}

function DeleteUser(userId) {
  if (confirm("Supprimer cet utilisateur ?")) {
    fetch(`http://localhost:8081/admin/users/delete/${userId}`, {
      method: "DELETE",
    }).then(() => GetUsers());
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
      closeModal("NewuserModal");
      GetUsers();
    })
    .catch((error) => console.error("Erreur API :", error));
}

function UpdateUser() {
  const modal = document.getElementById("editUserModal");
  const userId = modal.dataset.userId;

  const data = {
    prenom: document.getElementById("edit-prenom").value.trim(),
    nom: document.getElementById("edit-nom").value.trim(),
    email: document.getElementById("edit-email").value.trim(),
    role: document.getElementById("edit-role").value,
  };

  fetch(`http://localhost:8081/admin/users/modify/${userId}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  })
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur lors de l'UPDATE");
      modal.style.display = "none";
      GetUsers();
      alert("✅ Utilisateur mis à jour avec succès.");
    })
    .catch((error) => console.error("Erreur API :", error));
}

function UpdateValidationCount() {
  Promise.all([
    fetch("http://localhost:8081/admin/annonces").then((res) => res.json()),
    fetch("http://localhost:8081/admin/evenements").then((res) => res.json()),
  ])
    .then(([annonces, events]) => {
      let total = 0;

      annonces.forEach((annonce) => {
        if (
          annonce.statut_validation &&
          annonce.statut_validation.toLowerCase() === "en attente"
        )
          total++;
      });

      events.forEach((event) => {
        if (
          event.statut_validation &&
          event.statut_validation.toLowerCase() === "en attente"
        )
          total++;
      });

      const badgeTotal = document.querySelector(
        "#validations .sec-label .tag.t-red",
      );
      if (badgeTotal) badgeTotal.textContent = `${total} en attente`;

      const kpiTotal = document.getElementById("kpi-validations");
      const totalStat = document.getElementById("totalValidation");
      if (kpiTotal) kpiTotal.textContent = total;
      if (totalStat) totalStat.textContent = total;
    })
    .catch((err) => console.error("Erreur comptage validations :", err));
}

document.addEventListener("DOMContentLoaded", () => {
  GetUsers();
  UpdateValidationCount();

  const searchInput = document.getElementById("search-user-input");
  const roleFilter = document.getElementById("role-filter");

  if (searchInput) {
    searchInput.addEventListener("keydown", function (event) {
      if (event.key === "Enter") {
        event.preventDefault();

        const valeurRecherche = searchInput.value.trim();
        const role = roleFilter ? roleFilter.value : "";

        if (
          valeurRecherche === "" &&
          (role === "Tous les rôles" || role === "")
        ) {
          GetUsers();
        } else {
          Search(valeurRecherche, role);
        }
      }
    });
  }
});
