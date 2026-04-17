// On s'assure de récupérer les données de session au cas où

function chargerProfil() {
  // Sécurité : on ne lance pas le fetch si les infos sont manquantes
  if (!userId || !monToken) return;

  fetch(`http://localhost:8081/admin/users/${userId}`, {
    headers: {
      Authorization: "Bearer " + monToken,
    },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur lors du chargement du profil");
      return res.json();
    })
    .then((data) => {
      const profilName = document.querySelector(".profile-name");
      const profilRole = document.querySelector(".profile-role");
      const heroTitleEm = document.querySelector(".hero-title em");
      const avatar = document.querySelector(".profile-ava");

      profilName.textContent = data.nom + " " + data.prenom;
      profilRole.textContent = data.role;
      heroTitleEm.textContent = data.prenom;
      avatar.textContent = data.prenom.charAt(0);
    })
    .catch((err) => {
      console.error("Erreur Profil:", err);
    });
}

document.addEventListener("DOMContentLoaded", chargerProfil);
