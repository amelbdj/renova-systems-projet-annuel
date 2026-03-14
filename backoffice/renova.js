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
    alert("Veuillez remplir le champ libellé.");
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
      if (container)
        container.innerHTML = `<div style="padding:20px; color:red;">Erreur lors de la création.</div>`;
    });
}

function DeleteCategory(categoryId) {
  if (confirm("Supprimer cette catégorie ?")) {
    fetch(`http://localhost:8081/admin/categories/delete/${categoryId}`, {
      method: "DELETE",
    }).then(() => AfficherCategories());
  }
}

document.addEventListener("DOMContentLoaded", () => {
  AfficherCategories();
});
