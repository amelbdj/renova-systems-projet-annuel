function openNewCategory() {
  const modal = document.getElementById("NewCategoryModal");
  if (modal) modal.style.display = "flex";
}

function AfficherCategories() {
  const container = document.getElementById("resultC");
  if (!container) return;
  container.innerHTML = "";

  fetch("http://localhost:8081/admin/categories")
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

  fetch("http://localhost:8081/admin/categories/add", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
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
    fetch(`http://localhost:8081/admin/categories/delete/${categoryId}`, {
      method: "DELETE",
    }).then(() => AfficherCategories());
  }
}

document.addEventListener("DOMContentLoaded", () => {
  AfficherCategories();
});

let currentTranslations = {};

function changerLangue(langue) {
  fetch(`http://localhost:8081/api/translations?lang=${langue}`)
    .then((res) => res.json())
    .then((data) => {
      currentTranslations = data;
      appliquerTraductions();
    })
    .catch((err) =>
      console.error("Erreur de chargement des traductions:", err),
    );
}

function t(cle) {
  const keys = cle.split(".");
  let texteTraduit = currentTranslations;

  for (let k of keys) {
    if (texteTraduit && texteTraduit[k]) {
      texteTraduit = texteTraduit[k];
    } else {
      return cle;
    }
  }
  return typeof texteTraduit === "string" ? texteTraduit : cle;
}

function appliquerTraductions() {
  document.querySelectorAll("[data-i18n]").forEach((element) => {
    const cle = element.getAttribute("data-i18n");
    const texteTraduit = t(cle);
    if (texteTraduit !== cle) {
      element.innerHTML = texteTraduit;
    }
  });

  document.querySelectorAll("[data-i18n-placeholder]").forEach((element) => {
    const cle = element.getAttribute("data-i18n-placeholder");
    const texteTraduit = t(cle);
    if (texteTraduit !== cle) {
      element.placeholder = texteTraduit;
    }
  });
}

document.addEventListener("DOMContentLoaded", () => {
  changerLangue("fr");
});
