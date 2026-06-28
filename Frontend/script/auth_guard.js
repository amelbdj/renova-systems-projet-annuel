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
