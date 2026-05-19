// On récupère les identifiants globaux au démarrage
var monToken = localStorage.getItem("token");
var userId = localStorage.getItem("userId");

function chargerProfil() {
  if (!monToken || !userId) {
    window.location.href = "../login.html";
    return;
  }

  // 1. CHARGEMENT DES INFOS DU PROFIL
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

  // 2. COMPTEUR DES ARTICLES (MES PUBLICATIONS)
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

  // 3. COMPTEURS DES ÉVÉNEMENTS (CRÉÉS & EN LIGNE)
  const statEvent = document.getElementById("stat-event");
  const statValide = document.getElementById("stat-valide");
  if (statEvent || statValide) {
    fetch(`http://localhost:8081/admin/evenements`, {
      headers: { Authorization: "Bearer " + monToken },
    })
      .then((res) => res.json())
      .then((evenements) => {
        if (!evenements) evenements = [];

        // 🛡️ SÉCURITÉ ULTRA-LARGE : On attrape toutes les casquettes possibles de l'ID auteur
        const mesEvts = evenements.filter((e) => {
          const idAuteur =
            e.id_salarie ||
            e.IdSalarie ||
            e.idsalarie ||
            e.idSalarie ||
            e.user_id ||
            e.IdUser;
          // Sécurité de type : On transforme les deux côtés en Chaîne de texte pour éviter les conflits int/string
          return String(idAuteur) === String(userId);
        });

        // Mise à jour du total des événements créés par le salarié
        if (statEvent) statEvent.textContent = mesEvts.length;

        // Mise à jour du total des événements validés / en ligne
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

  // 4. COMPTEUR DU FORUM (SUJETS MODÉRÉS)
  const statForum = document.getElementById("stat-forum");
  if (statForum) {
    fetch(`http://localhost:8081/admin/forum/messages`, {
      headers: { Authorization: "Bearer " + monToken },
    })
      .then((res) => res.json())
      .then((messages) => {
        if (!messages) messages = [];
        // On affiche le volume d'activité sur le forum (ou les messages totaux du flux)
        statForum.textContent = messages.length;
      })
      .catch((err) => console.error("Erreur stat forum:", err));
  }
}

function logout() {
  localStorage.clear();
  window.location.href = "../login.html";
}

// Lancement au chargement du DOM
document.addEventListener("DOMContentLoaded", chargerProfil);
