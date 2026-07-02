function checkSession(requiredRole) {
    const token = localStorage.getItem("token");
    let userRole = localStorage.getItem("userRole");

    if (userRole === "SalariÃ©") {
        userRole = "Salarié";
    }

    if (!token) {
        if (window.location.pathname.includes("/salarie/")) {
            window.location.href = "../login.html";
        } else {
            window.location.href = "login.html";
        }
        return;
    }

    let roleOk = userRole === requiredRole;

    if (requiredRole && requiredRole.indexOf("Salari") === 0 && userRole && userRole.indexOf("Salari") === 0) {
        roleOk = true;
    }

    if (requiredRole && !roleOk) {
        if (window.location.pathname.includes("/salarie/")) {
            window.location.href = "../403.html";
        } else {
            window.location.href = "403.html";
        }
        return;
    }
}

// Surligne dans la sidebar le lien correspondant a la page ouverte
// (evite le surlignage fige sur "Vue d'ensemble").
function surlignerLienActif() {
    var pageCourante = window.location.pathname.split("/").pop();
    var liens = document.querySelectorAll("nav .nav-link");
    liens.forEach(function (lien) {
        var href = lien.getAttribute("href");
        if (href && href.split("/").pop() === pageCourante) {
            lien.classList.add("active");
        } else {
            lien.classList.remove("active");
        }
    });
}

// Met a jour le petit badge rouge (nb de validations en attente) dans la sidebar
function majBadgeValidations() {
    var token = localStorage.getItem("token");
    if (!token || typeof API_BASE_URL === "undefined") return;

    var badge = document.querySelector('nav a[href$="admin_validations.html"] .tag');
    if (!badge) return;

    var opts = { headers: { Authorization: "Bearer " + token } };

    Promise.all([
        fetch(API_BASE_URL + "/admin/annonces", opts).then(function (r) { return r.json(); }).catch(function () { return []; }),
        fetch(API_BASE_URL + "/admin/evenements", opts).then(function (r) { return r.json(); }).catch(function () { return []; }),
        fetch(API_BASE_URL + "/admin/articles", opts).then(function (r) { return r.json(); }).catch(function () { return []; }),
    ]).then(function (res) {
        var annonces = res[0] || [];
        var events = res[1] || [];
        var articles = res[2] || [];
        var total = 0;

        annonces.forEach(function (a) {
            if (a.statut_validation && a.statut_validation.toLowerCase() === "en attente") total++;
        });
        events.forEach(function (e) {
            if (e.statut_validation && e.statut_validation.toLowerCase() === "en attente") total++;
        });
        articles.forEach(function (a) {
            if (a.statut && a.statut.toLowerCase() === "en attente") total++;
        });

        badge.textContent = total;
        badge.style.display = total > 0 ? "" : "none"; // on cache le badge si rien en attente
    });
}

document.addEventListener("DOMContentLoaded", function () {
    surlignerLienActif();
    majBadgeValidations();
});
