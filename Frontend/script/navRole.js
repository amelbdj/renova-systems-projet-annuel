// Adapte les pages partagées (evenement / forum / article) au rôle connecté :
// - applique le thème "pro" (teal) si l'utilisateur est un Pro ;
// - fait pointer le lien "Accueil" et les liens de dashboard vers l'espace du rôle
//   (un Pro revient sur /pro au lieu de /client).
(function () {
  var role = localStorage.getItem("userRole") || "";

  // Thème teal pour les pros (appliqué tout de suite, avant le rendu)
  if (role === "Pro") {
    document.documentElement.classList.add("theme-pro");
  }

  function dashboardDuRole(r) {
    if (r === "Pro") return "/pro";
    if (r === "Administrateur") return "/admin";
    if (r.indexOf("Salari") === 0) return "/salarie";
    return "/client";
  }

  function corrigerNav() {
    var dash = dashboardDuRole(role);
    var liens = document.querySelectorAll(".nav-links a.nav-link");
    liens.forEach(function (a) {
      var cle = a.getAttribute("data-i18n") || "";
      var href = a.getAttribute("href");
      // "Accueil" (souvent sans href) et tous les liens qui renvoient vers
      // le dashboard particulier -> on les pointe vers le dashboard du rôle.
      if (cle === "landing.nav.home" || href === "/client") {
        a.setAttribute("href", dash);
      }
    });
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", corrigerNav);
  } else {
    corrigerNav();
  }
})();
