function checkSession(requiredRole) {
    const token = localStorage.getItem('token');
    const userRole = localStorage.getItem('userRole');

    if (!token) {
        window.location.href = "login.html";
        return;
    }

    if (requiredRole && userRole !== requiredRole) {
        window.location.href = "/Frontend/403.html";
        return;
    }
}