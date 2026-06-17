let currentTranslations = {};
monToken = localStorage.getItem("token");
function changerLangue(langue) {
  localStorage.setItem("langue", langue); // On mémorise le choix pour toutes les pages
  fetch(`http://localhost:8081/api/translations?lang=${langue}`, {
    cache: "no-store", // on veut toujours les traductions à jour (sinon le navigateur garde l'ancienne version)
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
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

// Import d'une nouvelle langue à partir d'un fichier JSON
function ImporterLangue() {
  const code = document
    .getElementById("input_lang_code")
    .value.trim()
    .toLowerCase();
  const nom = document.getElementById("input_lang_name").value.trim();
  const fichierInput = document.getElementById("input_fichier_json");

  // Petites vérifications avant d'envoyer
  if (code === "" || nom === "") {
    alert("Merci de remplir le code ET le nom de la langue.");
    return;
  }
  if (fichierInput.files.length === 0) {
    alert("Merci de choisir un fichier JSON.");
    return;
  }

  const fichier = fichierInput.files[0];
  const lecteur = new FileReader();

  // Cette fonction se lance UNE FOIS que le fichier est lu
  lecteur.onload = function () {
    let contenuJson;
    try {
      // On transforme le texte du fichier en objet JavaScript
      contenuJson = JSON.parse(lecteur.result);
    } catch (e) {
      alert("❌ Le fichier n'est pas un JSON valide.");
      return;
    }

    const dataToSend = {
      lang_code: code,
      lang_name: nom,
      data: contenuJson,
    };

    fetch("http://localhost:8081/admin/translations/add", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + monToken,
      },
      body: JSON.stringify(dataToSend),
    })
      .then((res) => {
        if (!res.ok) throw new Error("Erreur serveur");

        alert("✅ La langue a été importée avec succès !");
        document.getElementById("input_lang_code").value = "";
        document.getElementById("input_lang_name").value = "";
        fichierInput.value = "";

        // 🌟 On met à jour les boutons de langue tout de suite !
        GetLanguages();
      })
      .catch((err) => {
        console.error(err);
        alert("❌ Erreur lors de l'import.");
      });
  };

  // On lance la lecture du fichier (en texte)
  lecteur.readAsText(fichier);
}

// Export d'une langue en fichier JSON (sert de modèle à traduire)
function ExporterLangue(code) {
  fetch(`http://localhost:8081/api/translations?lang=${code}`, {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => res.json())
    .then((data) => {
      // On transforme l'objet en texte JSON bien indenté
      const texte = JSON.stringify(data, null, 2);

      // On crée un fichier en mémoire et on déclenche le téléchargement
      const blob = new Blob([texte], { type: "application/json" });
      const url = URL.createObjectURL(blob);

      const lien = document.createElement("a");
      lien.href = url;
      lien.download = code + ".json";
      lien.click();

      URL.revokeObjectURL(url); // On nettoie
    })
    .catch((err) => console.error("Erreur export:", err));
}

const btnToggleForm = document.getElementById("btn-toggle-form");

// 🛡️ SÉCURITÉ ICI : On vérifie si l'élément btn-toggle-form existe
if (btnToggleForm) {
  btnToggleForm.addEventListener("click", function () {
    const formContainer = document.getElementById("form-container");

    if (formContainer.style.display === "none") {
      formContainer.style.display = "block";
      this.innerHTML = "➖ Masquer le formulaire";
    } else {
      formContainer.style.display = "none";
      this.innerHTML = "➕ Ajouter une nouvelle langue";
    }
  });
}

// Affichage dynamique des boutons de langue
function GetLanguages() {
  fetch("http://localhost:8081/api/languages", {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur réseau");
      return res.json();
    })
    .then((languages) => {
      const container = document.getElementById("wrapper");

      // 🛡️ SÉCURITÉ ICI : On vérifie si le conteneur des boutons de langue existe
      if (!container) return;

      container.innerHTML = ""; // 🌟 CORRECTION : On vide le conteneur pour éviter de dupliquer les boutons

      languages.forEach((lang) => {
        container.innerHTML += `
          <button class="btn btn-sm btn-o" onclick="changerLangue('${lang.code}')" style="margin-right: 5px;">
            ${lang.name}
          </button>
        `;
      });
    })
    .catch((err) => {
      console.error("Erreur de chargement des boutons :", err);
    });
}

// Initialisation au chargement de la page
document.addEventListener("DOMContentLoaded", () => {
  GetLanguages();
  // On reprend la langue choisie précédemment, sinon français par défaut
  const langueSauvegardee = localStorage.getItem("langue") || "fr";
  changerLangue(langueSauvegardee);
});
