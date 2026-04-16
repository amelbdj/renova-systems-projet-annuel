/* ─── TOGGLE PASSWORD ─────────────────────── */
function togglePwd(id, btn) {
  const inp = document.getElementById(id);
  if (!inp) return;
  const show = inp.type === "password";
  inp.type = show ? "text" : "password";
  btn.textContent = show ? "🔒" : "👁";
}

/* ─── VALIDATION HELPERS ──────────────────── */
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

/* ─── SHOW SPACE SELECT ───────────────────── */
function showSpaceSelect() {
  document.getElementById("loginSuccess").style.display = "none";
  document.getElementById("spaceSelect").style.display = "block";
}
function goSpace(name) {
  alert("🚀 Redirection vers l'Espace " + name + "…");
}

/* ─── FORGOT PASSWORD ─────────────────────── */
function showForgot() {
  const email = document.getElementById("loginEmail").value.trim();
  if (email) {
    alert("📧 Un e-mail de réinitialisation a été envoyé à : " + email);
  } else {
    alert(
      'Veuillez entrer votre adresse e-mail, puis cliquer sur "Mot de passe oublié".',
    );
    document.getElementById("loginEmail").focus();
  }
}

/* ─── SOCIAL LOGIN ────────────────────────── */
function socialLogin(p) {
  alert("🔐 Authentification " + p + " — à connecter au back-end OAuth2.");
}

/* ─── SUBMIT LOGIN (API) ────────────────────────── */
/* ─── SUBMIT LOGIN (API) ────────────────────────── */
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
    let reponse = await fetch("http://localhost:8081/admin/login", {
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

      if (donneesServeur.validation === "En attente") {
        window.location.href = "attente.html";
        return;
      }

      if (donneesServeur.role === "Utilisateur") {
        window.location.href = "espace_particulier.html";
      } else if (donneesServeur.role === "Prestataire") {
        window.location.href = "espace_pro.html";
      } else if (donneesServeur.role === "Salarié") {
        window.location.href = "espace_salarie.html";
      } else if (donneesServeur.role === "Administrateur") {
        window.location.href = "espace_admin.html";
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
