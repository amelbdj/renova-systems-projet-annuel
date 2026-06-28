function openNewCategory() {
  const modal = document.getElementById("NewCategoryModal");
  if (modal) modal.style.display = "flex";
}

function AfficherCategories() {
  const container = document.getElementById("resultC");
  if (!container) return;
  container.innerHTML = "";

  fetch(`${API_BASE_URL}/admin/categories`, {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((response) => response.json())
    .then((categories) => {
      categories.forEach((category) => {
        const categoryElement = document.createElement("li");
        categoryElement.className = "category";
        categoryElement.innerHTML = `<h5>${category.libelle}</h5><button class="btn btn-xs btn-red" onclick="DeleteCategory(${category.id})"><span class="material-symbols-outlined">delete</span></button>`;
        container.appendChild(categoryElement);
      });
    });
}

function CreateCategory() {
  const libelle = document.getElementById("add-libelle").value.trim();

  if (libelle === "") {
    alert(t("backoffice.categories.alert_empty"));
    return;
  }

  fetch(`${API_BASE_URL}/admin/categories/add`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + monToken,
    },
    body: JSON.stringify({ libelle }),
  })
    .then((response) => {
      if (!response.ok) throw new Error("Erreur serveur");
      closeModal("NewCategoryModal");
      AfficherCategories();
    })
    .catch((error) => {
      console.error("Erreur API :", error);
      const container = document.getElementById("resultC");
      if (container)
        container.innerHTML = `<div style="padding:20px; color:red;" data-i18n="backoffice.categories.create_error">${t("backoffice.categories.create_error")}</div>`;
    });
}

function DeleteCategory(categoryId) {
  if (confirm(t("backoffice.categories.confirm_delete"))) {
    fetch(`${API_BASE_URL}/admin/categories/delete/${categoryId}`, {
      method: "DELETE",
      headers: {
        Authorization: "Bearer " + monToken,
      },
    }).then(() => AfficherCategories());
  }
}

document.addEventListener("DOMContentLoaded", () => {
  AfficherCategories();
});
