
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

function validateName(inp) {
  const ok = inp.value.trim().length >= 2;
  inp.classList.toggle("ok", ok);
  const errEl = document.getElementById(inp.id + "Err");
  if (errEl) errEl.style.display = !ok && inp.value ? "block" : "none";
}

function validateConfirm(inp) {
  const pwd = document.getElementById("rPwd").value;
  const ok = inp.value === pwd && inp.value.length > 0;
  document.getElementById("rPwd2Err").style.display =
    inp.value && !ok ? "block" : "none";
  document.getElementById("rPwd2Ok").style.display = ok ? "block" : "none";
  inp.classList.toggle("err", inp.value && !ok);
  inp.classList.toggle("ok", ok);
}


function checkStrength(inp) {
  const v = inp.value;
  let score = 0;
  if (v.length >= 8) score++;
  if (/[A-Z]/.test(v)) score++;
  if (/[0-9]/.test(v)) score++;
  if (/[^A-Za-z0-9]/.test(v)) score++;
  const bar = document.getElementById("sBar");
  const lbl = document.getElementById("sLbl");
  const data = [
    { w: "0%", c: var_("--red"), t: "Trop court" },
    { w: "25%", c: var_("--red"), t: "Faible" },
    { w: "50%", c: var_("--amber"), t: "Moyen" },
    { w: "75%", c: var_("--amber"), t: "Bon" },
    { w: "100%", c: var_("--green"), t: "Excellent 💪" },
  ];
  const d = data[v.length === 0 ? 0 : score];
  bar.style.width = d.w;
  bar.style.background = d.c;
  lbl.textContent = v.length === 0 ? "Force du mot de passe" : d.t;
  lbl.style.color = v.length === 0 ? var_("--txt-m") : d.c;
}
function var_(n) {
  return getComputedStyle(document.documentElement).getPropertyValue(n).trim();
}


let curStep = 1;
function goStep(n) {
  if (n > curStep) {
    if (curStep === 1) {
      const fn = document.getElementById("rFirstName").value.trim();
      const ln = document.getElementById("rLastName").value.trim();
      const em = document.getElementById("rEmail").value.trim();
      const pw = document.getElementById("rPwd").value;
      const pw2 = document.getElementById("rPwd2").value;
      if (!fn || !ln) {
        alert("Veuillez renseigner votre prénom et nom.");
        return;
      }
      if (!/^[^@]+@[^@]+\.[^@]+$/.test(em)) {
        alert("Adresse e-mail invalide.");
        return;
      }
      if (pw.length < 8) {
        alert("Mot de passe trop court (8 caractères min.).");
        return;
      }
      if (pw !== pw2) {
        alert("Les mots de passe ne correspondent pas.");
        return;
      }
    }
    if (curStep === 2) {
      const fn = document.getElementById("rFirstName").value.trim();
      const ln = document.getElementById("rLastName").value.trim();
      const em = document.getElementById("rEmail").value.trim();
      document.getElementById("sumName").textContent = fn + " " + ln;
      document.getElementById("sumEmail").textContent = em;
    }
  }
  document.getElementById("sp" + curStep).classList.remove("active");
  document.getElementById("s" + curStep).classList.remove("active");
  document.getElementById("s" + curStep).classList.add("done");
  curStep = n;
  document.getElementById("sp" + curStep).classList.add("active");
  document.getElementById("s" + curStep).classList.add("active");
  if (n < 3) {
    for (let i = n + 1; i <= 3; i++) {
      document.getElementById("s" + i).classList.remove("done", "active");
    }
  }
}


let roleChoisi = "";

function selectRole(element, codeRole) {
  document.querySelectorAll(".role-card").forEach((carte) => {
    carte.classList.remove("selected");
  });

  element.classList.add("selected");

  document.getElementById("partFields").style.display = "none";
  document.getElementById("proFields").style.display = "none";

  if (codeRole === "part") {
    document.getElementById("partFields").style.display = "block";
    roleChoisi = "Utilisateur";
  } else if (codeRole === "pro") {
    document.getElementById("proFields").style.display = "block";
    roleChoisi = "Pro";
  } else if (codeRole === "sal") {
    roleChoisi = "Salarié";
  } else if (codeRole === "adm") {
    roleChoisi = "Administrateur";
  }

  const labels = {
    part: "Particulier",
    pro: "Professionnel",
    sal: "Salarié UpcycleConnect",
    adm: "Administrateur",
  };
  document.getElementById("sumRole").textContent = labels[codeRole];
}

function olderThan18(dateTexte) {
  if (!dateTexte) return false;

  let dateNaissance = new Date(dateTexte);
  let aujourdHui = new Date();

  let anneeNaissance = dateNaissance.getFullYear();
  let moisNaissance = dateNaissance.getMonth();
  let jourNaissance = dateNaissance.getDate();

  let anneeActuelle = aujourdHui.getFullYear();
  let moisActuel = aujourdHui.getMonth();
  let jourActuel = aujourdHui.getDate();

  let age = anneeActuelle - anneeNaissance;

  if (moisActuel < moisNaissance) {
    age = age - 1;
  } else if (moisActuel === moisNaissance && jourActuel < jourNaissance) {
    age = age - 1;
  }

  if (age >= 18) {
    return true;
  } else {
    return false;
  }
}

// Verifie un SIRET : 14 chiffres + cle de controle de Luhn.
function siretValide(siret) {
  if (!siret || siret.length !== 14) return false;
  let somme = 0;
  for (let i = 0; i < 14; i++) {
    let c = siret[i];
    if (c < "0" || c > "9") return false;
    let n = parseInt(c, 10);
    if ((14 - i) % 2 === 0) {
      n = n * 2;
      if (n > 9) n = n - 9;
    }
    somme = somme + n;
  }
  return somme % 10 === 0;
}


async function submitRegister() {
  if (!document.getElementById("chkCgu").checked) {
    alert("Veuillez accepter les CGU pour continuer.");
    return;
  }

  let infosAEnvoyer = {
    email: document.getElementById("rEmail").value,
    mot_de_passe: document.getElementById("rPwd").value,
    nom: document.getElementById("rLastName").value,
    prenom: document.getElementById("rFirstName").value,
    role: roleChoisi,
    validation: roleChoisi === "Utilisateur" ? "Validé" : "En attente",
  };

  if (roleChoisi === "Utilisateur") {
    let dateNaissance = document.getElementById("partDob").value;
    if (olderThan18(dateNaissance) === false) {
      alert("Désolé, il faut avoir au moins 18 ans pour s'inscrire !");
      return;
    }
    infosAEnvoyer.date_naissance = dateNaissance;
  } else if (roleChoisi === "Pro") {
    let siret = document.getElementById("proSiret").value.replace(/\s/g, "");
    if (siretValide(siret) === false) {
      alert("SIRET invalide : il doit contenir 14 chiffres avec une clé de contrôle correcte.");
      return;
    }
    infosAEnvoyer.nom_entreprise = document.getElementById("proName").value;
    infosAEnvoyer.siret = siret;
  }

  try {
    let reponse = await fetch(API_BASE_URL + "/auth/inscription", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(infosAEnvoyer),
    });

    if (reponse.ok === true) {
      document.getElementById("panelReg").style.display = "none";
      document.getElementById("regSuccess").style.display = "block";

      setTimeout(function () {
        window.location.href = "/login";
      }, 3000);
    } else {
      alert("Erreur lors de l'inscription. L'email existe peut-être déjà.");
    }
  } catch (error) {
    console.error(error);
    alert("Impossible de joindre le serveur.");
  }
}
