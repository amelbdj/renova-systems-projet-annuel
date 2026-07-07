



(function () {
  var role = localStorage.getItem("userRole") || "";

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
