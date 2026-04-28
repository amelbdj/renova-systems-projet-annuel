function chargerProfil() {
  if (!monToken) {
    window.location.href = "login.html";
    return;
  }

  const hasSeenTutorial = localStorage.getItem("tutorielVu");
  const tutorialOverlay = document.getElementById("tut");

  if (hasSeenTutorial === "true") {
    if (tutorialOverlay) tutorialOverlay.style.display = "none";
  } else {
    console.log("showing the tutorial");
  }

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

function logout() {
  localStorage.clear();
  window.location.href = "login.html";
}

document.addEventListener("DOMContentLoaded", chargerProfil);

function goToProfile() {
  const userId = localStorage.getItem("userId");
  const role = localStorage.getItem("role");

  if (!userId) {
    window.location.href = "login.html";
    return;
  }

  window.location.href = `profil.html?id=${userId}`;
}
