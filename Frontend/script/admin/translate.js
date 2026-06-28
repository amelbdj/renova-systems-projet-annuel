let currentTranslations = {};
let monTokenTraduction = localStorage.getItem("token");

function changerLangue(langue) {
  localStorage.setItem("langue", langue);

  fetch(`${API_BASE_URL}/api/translations?lang=${langue}`, {
    cache: "no-store",
    headers: {
      Authorization: "Bearer " + monTokenTraduction,
    },
  })
    .then((res) => res.json())
    .then((data) => {
      currentTranslations = data;
      appliquerTraductions();
    })
    .catch((err) => {
      console.error("Erreur de chargement des traductions:", err);
    });
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

  if (typeof texteTraduit === "string") {
    return texteTraduit;
  }

  return cle;
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

  // Si la page possède le tutoriel pro (en-tête géré par JS), on le rafraîchit
  // pour qu'il s'affiche dans la bonne langue.
  if (typeof updTut === "function") {
    try {
      updTut();
    } catch (e) {}
  }
}

function ImporterLangue() {
  const code = document.getElementById("input_lang_code").value.trim().toLowerCase();
  const nom = document.getElementById("input_lang_name").value.trim();
  const fichierInput = document.getElementById("input_fichier_json");

  if (code === "" || nom === "") {
    alert("Merci de remplir le code et le nom de la langue.");
    return;
  }

  if (!fichierInput || fichierInput.files.length === 0) {
    alert("Merci de choisir un fichier JSON.");
    return;
  }

  const fichier = fichierInput.files[0];
  const lecteur = new FileReader();

  lecteur.onload = function () {
    let contenuJson;

    try {
      contenuJson = JSON.parse(lecteur.result);
    } catch (e) {
      alert("Le fichier n'est pas un JSON valide.");
      return;
    }

    const dataToSend = {
      lang_code: code,
      lang_name: nom,
      data: contenuJson,
    };

    fetch(`${API_BASE_URL}/admin/translations/add`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + monTokenTraduction,
      },
      body: JSON.stringify(dataToSend),
    })
      .then((res) => {
        if (!res.ok) {
          throw new Error("Erreur serveur");
        }

        alert("La langue a ete importee avec succes !");
        document.getElementById("input_lang_code").value = "";
        document.getElementById("input_lang_name").value = "";
        fichierInput.value = "";
        GetLanguages();
      })
      .catch((err) => {
        console.error(err);
        alert("Erreur lors de l'import.");
      });
  };

  lecteur.readAsText(fichier);
}

function ExporterLangue(code) {
  fetch(`${API_BASE_URL}/api/translations?lang=${code}`, {
    headers: {
      Authorization: "Bearer " + monTokenTraduction,
    },
  })
    .then((res) => res.json())
    .then((data) => {
      const texte = JSON.stringify(data, null, 2);
      const blob = new Blob([texte], { type: "application/json" });
      const url = URL.createObjectURL(blob);

      const lien = document.createElement("a");
      lien.href = url;
      lien.download = code + ".json";
      lien.click();

      URL.revokeObjectURL(url);
    })
    .catch((err) => {
      console.error("Erreur export:", err);
    });
}

function GetLanguages() {
  fetch(`${API_BASE_URL}/api/languages`, {
    headers: {
      Authorization: "Bearer " + monTokenTraduction,
    },
  })
    .then((res) => {
      if (!res.ok) {
        throw new Error("Erreur reseau");
      }
      return res.json();
    })
    .then((languages) => {
      const container = document.getElementById("wrapper");

      if (!container) {
        return;
      }

      container.innerHTML = "";

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

document.addEventListener("DOMContentLoaded", () => {
  GetLanguages();

  const langueSauvegardee = localStorage.getItem("langue") || "fr";
  changerLangue(langueSauvegardee);

  const btnToggleForm = document.getElementById("btn-toggle-form");
  if (btnToggleForm) {
    btnToggleForm.addEventListener("click", function () {
      const formContainer = document.getElementById("form-container");

      if (!formContainer) {
        return;
      }

      if (formContainer.style.display === "none") {
        formContainer.style.display = "block";
        this.innerHTML = "Masquer le formulaire";
      } else {
        formContainer.style.display = "none";
        this.innerHTML = "Ajouter une nouvelle langue";
      }
    });
  }
});
