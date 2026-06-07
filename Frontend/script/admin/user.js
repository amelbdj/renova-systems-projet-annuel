monToken = localStorage.getItem("token");
function setVtab(element, type) {
  document
    .querySelectorAll(".vtab")
    .forEach((btn) => btn.classList.remove("on"));
  element.classList.add("on");

  const container = document.getElementById("result");
  if (!container) return;

  container.innerHTML = `<p style='padding:20px'>${t("backoffice.common.loading")}</p>`;

  if (type === "ann") {
    GetAnnonce();
  } else if (type === "evt") {
    GetEvent();
  } else if (type === "all") {
    container.innerHTML = "";
    GetAnnonce();
    GetEvent();
  } else if (type === "con") {
    GetArticle();
  } else {
    container.innerHTML = `<p style='padding:20px'>${t("backoffice.common.in_dev")}</p>`;
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

  // 🛡️ LE FAMEUX BOUCLIER : Si le tableau n'est pas sur la page, on arrête tout !
  if (!container) return;

  const totalStat = document.getElementById("totalUser");

  if (!users) users = [];
  if (totalStat) totalStat.innerHTML = users.length;

  const headerHTML = `
        <div class="u-thead">
            <div></div>
            <div data-i18n="backoffice.users.th_user">${t("backoffice.users.th_user")}</div>
            <div data-i18n="backoffice.form.role">${t("backoffice.form.role")}</div>
            <div data-i18n="backoffice.users.th_document">${t("backoffice.users.th_document")}</div>
            <div data-i18n="backoffice.users.th_status">${t("backoffice.users.th_status")}</div>
            <div data-i18n="backoffice.users.th_score">${t("backoffice.users.th_score")}</div>
            <div>ID</div>
            <div data-i18n="backoffice.users.th_actions">${t("backoffice.users.th_actions")}</div>
        </div>`;

  if (users.length === 0) {
    container.innerHTML =
      headerHTML +
      `<div style="padding:20px; text-align:center; color:var(--txt-m);" data-i18n="backoffice.users.no_users_found">${t("backoffice.users.no_users_found")}</div>`;
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

    let docContent = "";
    if (user.chemin_fichier) {
      const fileUrl = `http://localhost:8081/view-uploads/${user.chemin_fichier.replace(/\\/g, "/").replace("uploads/", "")}`;
      docContent = `
                <button class="btn btn-xs btn-g" onclick="window.open('${fileUrl}', '_blank')">
                    <span class="material-symbols-outlined" style="font-size:16px;">description</span>
                    <span data-i18n="backoffice.users.doc_view">${t("backoffice.users.doc_view")}</span>
                </button>`;
    } else {
      docContent = `<span style="color:var(--txt-m); font-size:12px;" data-i18n="backoffice.users.doc_none">${t("backoffice.users.doc_none")}</span>`;
    }

    let validationContent = "";
    if (user.validation === "En attente") {
      validationContent = `
                <div style="display:flex; gap:6px;">
                    <button class="btn btn-xs btn-grn" onclick="ValidateUser(${user.id})">
                        <span class="material-symbols-outlined" style="font-size: 16px;">check</span>
                    </button>
                    <button class="btn btn-xs btn-red" onclick="RefuseUser(${user.id})">
                        <span class="material-symbols-outlined" style="font-size: 16px;">close</span>
                    </button>
                </div>`;
    } else if (user.validation === "Validé") {
      validationContent = `
                <span style="display:flex; align-items:center; gap:4px; color:#2ecc71; font-size:12px; font-weight:600;">
                    <span class="material-symbols-outlined" style="font-size:16px;">verified</span>
                    <span data-i18n="backoffice.users.status_approved">${t("backoffice.users.status_approved")}</span>
                </span>`;
    } else {
      validationContent = `
                <span style="display:flex; align-items:center; gap:4px; color:#f05050; font-size:12px; font-weight:600;">
                    <span class="material-symbols-outlined" style="font-size:16px;">block</span>
                    <span data-i18n="backoffice.users.status_rejected">${t("backoffice.users.status_rejected")}</span>
                </span>`;
    }

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
                <div>${docContent}</div>
                <div>${validationContent}</div>
                <div><span class="tag t-grn">${user.score || 0}</span></div>
                <div style="font-size:11.5px; color:var(--txt-m)">ID : ${user.id}</div>
                <div class="u-actions" style="display: flex; gap: 6px; align-items: center;">
                    <button class="btn btn-xs btn-g" onclick="OpenEditModalAPI(${user.id}, '${user.nom}', '${user.prenom}', '${user.email}', '${user.role}')">
                        <span class="material-symbols-outlined" style="font-size: 16px;">person_edit</span>
                    </button>
                    <button class="btn btn-xs btn-red" onclick="DeleteUser(${user.id})">
                        <span class="material-symbols-outlined" style="font-size: 16px;">delete</span>
                    </button>
                </div>
            </div>`;
  });

  container.innerHTML = rowsHTML;

  if (typeof appliquerTraductions === "function") {
    appliquerTraductions();
  }
}

function GetUsers() {
  fetch("http://localhost:8081/admin/users", {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => res.json())
    .then(AfficherTableau)
    .catch((err) => console.error("Erreur GET Users:", err));
}

let currentUserIdToRefuse = null; // Variable temporaire pour stocker l'ID

function RefuseUser(userId) {
  currentUserIdToRefuse = userId; // On mémorise quel user on veut refuser
  document.getElementById("modalRefus").style.display = "flex";
  document.getElementById("motifTexte").value = ""; // On vide le texte
}

function ValidateUser(userId) {
  fetch(`http://localhost:8081/admin/users/validate/${userId}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
  })
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
  let url = `http://localhost:8081/admin/users/search?name=${encodeURIComponent(query)}`;

  const roleAllText =
    t("backoffice.role.all") !== "backoffice.role.all"
      ? t("backoffice.role.all")
      : "Tous les rôles";

  if (
    role &&
    role !== roleAllText &&
    role !== "Tous les rôles" &&
    role !== ""
  ) {
    url += `&role=${encodeURIComponent(role)}`;
  }

  fetch(url, {
    method: "GET",
    headers: {
      Authorization: "Bearer " + localStorage.getItem("token"), // 👈 C'est ce passe-partout qui manquait !
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur de recherche");
      return res.json();
    })
    .then(AfficherTableau)
    .catch((err) => {
      console.error("Erreur GET Search:", err);
      const container = document.querySelector(".u-table");
      if (container)
        container.innerHTML = `<div style="padding:20px; color:red;">${t("backoffice.users.search_error")}</div>`;
    });
}

function DeleteUser(userId) {
  if (confirm(t("backoffice.users.confirm_delete"))) {
    fetch(`http://localhost:8081/admin/users/delete/${userId}`, {
      method: "DELETE",
      headers: {
        Authorization: "Bearer " + monToken,
      },
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
    alert(t("backoffice.users.alert_empty"));
    return;
  }

  fetch("http://localhost:8081/admin/users/add", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
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
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
    body: JSON.stringify(data),
  })
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur lors de l'UPDATE");
      modal.style.display = "none";
      GetUsers();
      alert(t("backoffice.users.success_update"));
    })
    .catch((error) => console.error("Erreur API :", error));
}

function UpdateValidationCount() {
  // 🟢 1. On récupère le token

  // 🟢 2. On prépare les options avec l'en-tête d'autorisation
  const fetchOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${monToken}`,
    },
  };

  // 🟢 3. On passe fetchOptions en deuxième paramètre de tes fetch
  Promise.all([
    fetch("http://localhost:8081/admin/annonces", fetchOptions).then((res) =>
      res.json(),
    ),
    fetch("http://localhost:8081/admin/evenements", fetchOptions).then((res) =>
      res.json(),
    ),
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
      if (badgeTotal)
        badgeTotal.textContent = `${total} ${t("backoffice.kpi.pending_badge")}`;

      const kpiTotal = document.getElementById("kpi-validations");
      const totalStat = document.getElementById("totalValidation");
      if (kpiTotal) kpiTotal.textContent = total;
      if (totalStat) totalStat.textContent = total;
    })
    .catch((err) => console.error("Erreur comptage validations :", err));
}

function FermerModaleRefus() {
  document.getElementById("modalRefus").style.display = "none";
}

const btnConfirmerRefus = document.getElementById("btnConfirmerRefus");

// 🛡️ LE BOUCLIER : On ne met le onclick que si le bouton existe sur la page !
if (btnConfirmerRefus) {
  btnConfirmerRefus.onclick = function () {
    const raison = document.getElementById("motifTexte").value;

    if (!raison) {
      alert("Merci de saisir un motif pour l'utilisateur.");
      return;
    }

    fetch(`http://localhost:8081/admin/users/refuse/${currentUserIdToRefuse}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + monToken,
      },
      body: JSON.stringify({ motif: raison }),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data.message);
        FermerModaleRefus();
        GetUsers(); // On rafraîchit le tableau instantanément
      })
      .catch((err) => console.error("Erreur refus:", err));
  };
}

function logout() {
  localStorage.clear();
  window.location.href = "login.html";
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

        const roleAllText =
          t("backoffice.role.all") !== "backoffice.role.all"
            ? t("backoffice.role.all")
            : "Tous les rôles";

        if (
          valeurRecherche === "" &&
          (role === roleAllText || role === "Tous les rôles" || role === "")
        ) {
          GetUsers();
        } else {
          Search(valeurRecherche, role);
        }
      }
    });
  }
  // const role = localStorage.getItem("role");

  // if (role != "Administrateur") {
  //   window.location.href = "403.html";
  // }
});
