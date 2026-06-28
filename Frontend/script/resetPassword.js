/* ─── SÉCURITÉ : il faut être connecté pour accéder à cette page ───── */
document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");
  if (!token) {
    alert("Vous devez être connecté pour réinitialiser votre mot de passe.");
    window.location.href = "login.html";
  }
});

/* ─── AFFICHER / CACHER LE MOT DE PASSE ─────────────── */
function togglePwd(id, btn) {
  const inp = document.getElementById(id);
  if (!inp) return;
  const show = inp.type === "password";
  inp.type = show ? "text" : "password";
  btn.textContent = show ? "🔒" : "👁";
}

/* ─── RÉINITIALISATION DU MOT DE PASSE (compte connecté) ──────────── */
async function submitReset() {
  const token = localStorage.getItem("token");
  const email = document.getElementById("resetEmail").value.trim();
  const pwd = document.getElementById("resetPwd").value;
  const confirmation = document.getElementById("resetConfirm").value;

  const errEl = document.getElementById("resetErr");
  const okEl = document.getElementById("resetOk");
  errEl.style.display = "none";
  okEl.style.display = "none";

  if (!token) {
    window.location.href = "login.html";
    return;
  }

  if (!email || !pwd || !confirmation) {
    errEl.textContent = "Veuillez remplir tous les champs.";
    errEl.style.display = "block";
    return;
  }
  if (pwd.length < 6) {
    errEl.textContent = "Le mot de passe doit faire au moins 6 caractères.";
    errEl.style.display = "block";
    return;
  }
  if (pwd !== confirmation) {
    errEl.textContent = "Les mots de passe ne correspondent pas.";
    errEl.style.display = "block";
    return;
  }

  try {
    const reponse = await fetch("http://localhost:8081/auth/reset-password", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ email: email, new_password: pwd }),
    });
    const data = await reponse.json();

    if (reponse.ok) {
      okEl.textContent =
        "✅ Mot de passe réinitialisé ! Redirection vers votre profil…";
      okEl.style.display = "block";
      setTimeout(() => {
        window.location.href = "profil.html";
      }, 2000);
    } else {
      errEl.textContent =
        data.error || "Impossible de réinitialiser le mot de passe.";
      errEl.style.display = "block";
    }
  } catch (error) {
    console.error("Erreur reset:", error);
    errEl.textContent = "Impossible de joindre le serveur.";
    errEl.style.display = "block";
  }
}
