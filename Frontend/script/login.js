function togglePwd(id, btn) {
  const inp = document.getElementById(id);
  if (!inp) return;
  const show = inp.type === "password";
  inp.type = show ? "text" : "password";
  btn.textContent = show ? "🔒" : "👁";
}

function validateEmail(inp) {
  const ok = /^[^@]+@[^@]+\.[^@]+$/.test(inp.value.trim());
  inp.classList.toggle("err", inp.value && !ok);
  inp.classList.toggle("ok", ok);
  const errId = inp.id + "Err";
  const errEl = document.getElementById(errId);
  if (errEl) errEl.style.display = inp.value && !ok ? "block" : "none";
}

function validatePwdLogin(inp) {
  const ok = inp.value.length >= 6;
  const errEl = document.getElementById("loginPwdErr");
  if (errEl) errEl.style.display = !ok && inp.value ? "block" : "none";
}

function showSpaceSelect() {
  document.getElementById("loginSuccess").style.display = "none";
  document.getElementById("spaceSelect").style.display = "block";
}
function goSpace(name) {
  alert("🚀 Redirection vers l'Espace " + name + "…");
}

function showForgot() {
  const box = document.getElementById("forgotBox");
  const loginEmail = document.getElementById("loginEmail").value.trim();
  const forgotEmail = document.getElementById("forgotEmail");

  if (box.style.display === "none") {
    box.style.display = "block";
  } else {
    box.style.display = "none";
  }

  if (loginEmail !== "") {
    forgotEmail.value = loginEmail;
  }
}

async function submitForgot() {
  const email = document.getElementById("forgotEmail").value.trim();
  const msg = document.getElementById("forgotMsg");

  msg.textContent = "";
  msg.style.color = "#00c97a";

  if (email === "") {
    msg.style.color = "#ff5a5a";
    msg.textContent = "Veuillez entrer votre adresse e-mail.";
    return;
  }

  try {
    const reponse = await fetch(API_BASE_URL + "/auth/forgot-password", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email: email }),
    });

    const data = await reponse.json();

    if (reponse.ok) {
      msg.textContent = data.message || "Un e-mail a ete envoye si le compte existe.";
    } else {
      msg.style.color = "#ff5a5a";
      msg.textContent = data.error || "Impossible d'envoyer le mail.";
    }
  } catch (error) {
    console.error("Erreur forgot password:", error);
    msg.style.color = "#ff5a5a";
    msg.textContent = "Impossible de joindre le serveur.";
  }
}

function socialLogin(p) {
  alert("🔐 Authentification " + p + " — à connecter au back-end OAuth2.");
}

async function submitLogin() {
  let emailInfo = document.getElementById("loginEmail").value;
  let motDePasseInfo = document.getElementById("loginPwd").value;

  if (emailInfo === "" || motDePasseInfo === "") {
    alert("Il faut remplir l'email et le mot de passe");
    return;
  }

  let infosAEnvoyer = {
    email: emailInfo,
    mot_de_passe: motDePasseInfo,
  };

  try {
    let reponse = await fetch(API_BASE_URL + "/admin/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(infosAEnvoyer),
    });

    if (reponse.ok === true) {
      let donneesServeur = await reponse.json();

      localStorage.setItem("token", donneesServeur.token);
      localStorage.setItem("userRole", donneesServeur.role);
      localStorage.setItem("userId", donneesServeur.id);
      localStorage.setItem("userName", donneesServeur.prenom);
      localStorage.setItem("userScore", donneesServeur.score || 0);
      localStorage.setItem("tutorielVu", donneesServeur.tutorielVu);

      if (donneesServeur.validation === "En attente") {
        window.location.href = "attente.html";
        return;
      }

      if (donneesServeur.validation === "Rejeté") {
        window.location.href = "403.html";
        return;
      }

      if (donneesServeur.role === "Utilisateur") {
        window.location.href = "/client";
      } else if (donneesServeur.role === "Pro") {
        window.location.href = "/pro";
      } else if (donneesServeur.role === "Salarié") {
        window.location.href = "/salarie";
      } else if (donneesServeur.role === "Administrateur") {
        window.location.href = "/admin";
      } else {
        alert("Rôle inconnu. Contactez l'administrateur.");
      }
    } else {
      alert("Email ou mot de passe incorrect.");
    }
  } catch (error) {
    console.error("Erreur login:", error);
    alert("Impossible de joindre le serveur.");
  }
}
