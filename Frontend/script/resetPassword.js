function togglePwd(id, btn) {
  const inp = document.getElementById(id);
  if (!inp) return;
  const show = inp.type === "password";
  inp.type = show ? "text" : "password";
  btn.textContent = show ? "🔒" : "👁";
}

document.addEventListener("DOMContentLoaded", () => {
  const params = new URLSearchParams(window.location.search);
  const email = params.get("email");

  if (email) {
    document.getElementById("resetEmail").value = email;
  }
});

async function submitReset() {
  const params = new URLSearchParams(window.location.search);
  const token = params.get("token");
  const email = document.getElementById("resetEmail").value.trim();
  const pwd = document.getElementById("resetPwd").value;
  const confirmation = document.getElementById("resetConfirm").value;

  const errEl = document.getElementById("resetErr");
  const okEl = document.getElementById("resetOk");
  errEl.style.display = "none";
  okEl.style.display = "none";

  if (!token) {
    errEl.textContent = "Lien de reinitialisation invalide.";
    errEl.style.display = "block";
    return;
  }

  if (!email || !pwd || !confirmation) {
    errEl.textContent = "Veuillez remplir tous les champs.";
    errEl.style.display = "block";
    return;
  }

  if (pwd.length < 6) {
    errEl.textContent = "Le mot de passe doit faire au moins 6 caracteres.";
    errEl.style.display = "block";
    return;
  }

  if (pwd !== confirmation) {
    errEl.textContent = "Les mots de passe ne correspondent pas.";
    errEl.style.display = "block";
    return;
  }

  try {
    const reponse = await fetch(API_BASE_URL + "/auth/reset-password", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: email,
        token: token,
        new_password: pwd,
      }),
    });

    const data = await reponse.json();

    if (reponse.ok) {
      okEl.textContent = "Mot de passe reinitialise. Vous pouvez vous connecter.";
      okEl.style.display = "block";
      setTimeout(() => {
        window.location.href = "login.html";
      }, 2000);
    } else {
      errEl.textContent = data.error || "Impossible de reinitialiser le mot de passe.";
      errEl.style.display = "block";
    }
  } catch (error) {
    console.error("Erreur reset:", error);
    errEl.textContent = "Impossible de joindre le serveur.";
    errEl.style.display = "block";
  }
}
