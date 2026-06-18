
var monToken = localStorage.getItem("token");
var userId = localStorage.getItem("userId");

function chargerProfil() {
  if (!monToken || !userId) {
    window.location.href = "../login.html";
    return;
  }

  
  fetch(`http://localhost:8081/admin/users/${userId}`, {
    headers: { Authorization: "Bearer " + monToken },
  })
    .then((res) => {
      if (!res.ok) throw new Error("Erreur lors du chargement du profil");
      return res.json();
    })
    .then((data) => {
      const profilName = document.querySelector(".profile-name");
      const heroTitleEm = document.getElementById("user-name-display");
      const avatar = document.querySelector(".profile-ava");
      const profilRole = document.querySelector(".profile-role");

      if (profilName) profilName.textContent = data.nom + " " + data.prenom;
      if (heroTitleEm) heroTitleEm.textContent = data.prenom;
      if (avatar) avatar.textContent = data.prenom.charAt(0);
      if (profilRole) profilRole.textContent = data.role || "Salarié";
    })
    .catch((err) => console.error("Erreur Profil:", err));

  
  const statArticle = document.getElementById("stat-article");
  if (statArticle) {
    fetch(`http://localhost:8081/admin/articles/salarie/${userId}`, {
      headers: { Authorization: "Bearer " + monToken },
    })
      .then((res) => res.json())
      .then((articles) => {
        if (!articles) articles = [];
        statArticle.textContent = articles.length;
      })
      .catch((err) => console.error("Erreur stat articles:", err));
  }

  
  const statEvent = document.getElementById("stat-event");
  const statValide = document.getElementById("stat-valide");
  if (statEvent || statValide) {
    fetch(`http://localhost:8081/admin/evenements`, {
      headers: { Authorization: "Bearer " + monToken },
    })
      .then((res) => res.json())
      .then((evenements) => {
        if (!evenements) evenements = [];

        
        const mesEvts = evenements.filter((e) => {
          const idAuteur =
            e.id_salarie ||
            e.IdSalarie ||
            e.idsalarie ||
            e.idSalarie ||
            e.user_id ||
            e.IdUser;
          
          return String(idAuteur) === String(userId);
        });

        
        if (statEvent) statEvent.textContent = mesEvts.length;

        
        if (statValide) {
          const enLigne = mesEvts.filter((evt) => {
            const statut = (
              evt.statut_validation ||
              evt.statut ||
              evt.Statut ||
              "en attente"
            ).toLowerCase();
            return (
              statut === "valide" ||
              statut === "en ligne" ||
              statut === "publié" ||
              statut === "publie"
            );
          });
          statValide.textContent = enLigne.length;
        }
      })
      .catch((err) => console.error("Erreur stat événements:", err));
  }

  
  const statForum = document.getElementById("stat-forum");
  if (statForum) {
    fetch(`http://localhost:8081/admin/forum/messages`, {
      headers: { Authorization: "Bearer " + monToken },
    })
      .then((res) => res.json())
      .then((messages) => {
        if (!messages) messages = [];
        
        statForum.textContent = messages.length;
      })
      .catch((err) => console.error("Erreur stat forum:", err));
  }
}

function logout() {
  localStorage.clear();
  window.location.href = "../login.html";
}
function goToProfile() {
  const userId = localStorage.getItem("userId");
  const role = localStorage.getItem("role");

  if (!userId) {
    window.location.href = "../login.html";
    return;
  }

  window.location.href = `../profil.html?id=${userId}`;
}

document.addEventListener("DOMContentLoaded", chargerProfil);
