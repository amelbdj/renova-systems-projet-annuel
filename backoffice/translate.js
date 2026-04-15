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

function LoadFormulaireTraduction() {
  fetch("http://localhost:8081/admin/translations/keys")
    .then((res) => {
      if (!res.ok) throw new Error("Erreur réseau");
      return res.json();
    })
    .then((keys) => {
      const container = document.getElementById("dynamic-fields-container");
      container.innerHTML = ""; // On vide avant de remplir

      keys.forEach((key) => {
        container.innerHTML += `
                <div style="margin-bottom: 10px;">
                    <label>Traduire : <strong>${key}</strong></label><br>
                    <input type="text" data-key="${key}" class="input-traduction" required style="width: 100%; padding: 5px;">
                </div>
            `;
      });
    })
    .catch((err) => {
      console.error("Erreur de chargement des clés :", err);
    });
}

// On lance le dessin du formulaire tout de suite
LoadFormulaireTraduction();

// Gestion de la soumission du formulaire d'ajout de langue
document
  .getElementById("form-add-language")
  .addEventListener("submit", function (event) {
    event.preventDefault();

    const codeLangue = document
      .getElementById("input_lang_code")
      .value.toLowerCase();
    const dataToSend = { lang_code: codeLangue, translations: [] };
    const inputs = document.querySelectorAll(".input-traduction");

    inputs.forEach((input) => {
      dataToSend.translations.push({
        msg_key: input.getAttribute("data-key"),
        msg_value: input.value,
      });
    });

    fetch("http://localhost:8081/admin/translations/add", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(dataToSend),
    })
      .then((res) => {
        if (!res.ok) throw new Error("Erreur serveur");

        alert("✅ La langue a été ajoutée avec succès !");
        document.getElementById("form-add-language").reset();

        // 🌟 L'ASTUCE DE PRO : On met à jour les boutons en haut instantanément !
        GetLanguages();
      })
      .catch((err) => {
        console.error(err);
        alert("❌ Erreur lors de l'enregistrement.");
      });
  });

document
  .getElementById("btn-toggle-form")
  .addEventListener("click", function () {
    const formContainer = document.getElementById("form-container");

    if (formContainer.style.display === "none") {
      formContainer.style.display = "block";
      this.innerHTML = "➖ Masquer le formulaire";
    } else {
      formContainer.style.display = "none";
      this.innerHTML = "➕ Ajouter une nouvelle langue";
    }
  });

// Affichage dynamique des boutons de langue
function GetLanguages() {
  fetch("http://localhost:8081/api/languages")
    .then((res) => {
      if (!res.ok) throw new Error("Erreur réseau");
      return res.json();
    })
    .then((languages) => {
      const container = document.getElementById("wrapper");

      // 🌟 CORRECTION : On vide le conteneur pour éviter de dupliquer les boutons

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
  changerLangue("fr"); // On charge le français par défaut
});
