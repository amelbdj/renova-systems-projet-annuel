let stripePayoutStatus = "none";

document.addEventListener("DOMContentLoaded", async () => {
  await loadUserProfile();

  const urlParams = new URLSearchParams(window.location.search);
  if (urlParams.get("stripe") === "success") {
    alert("✅ Stripe configuré !");
    setStripeState("active");
    window.history.replaceState({}, document.title, window.location.pathname);
  }
});

function setStripeState(state) {
  stripePayoutStatus = state;

  const button = document.getElementById("stripe-btn");
  const status = document.getElementById("stripeStatus");

  if (!button || !status) {
    console.error("Stripe elements not found");
    return;
  }

  if (state === "active") {
    button.style.display = "none";
    status.textContent = "✅ Actif";
    status.className = "tag t-grn";
  } else {
    button.style.display = "inline-block";
    status.textContent = "Non configuré";
    status.className = "tag t-red";
  }
}
async function loadUserProfile() {
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("token");
  const avatar = document.querySelector(".profile-ava");

  if (!userId) return;

  try {
    const res = await fetch(`http://localhost:8081/user/profile?id=${userId}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      console.error("Server returned error:", res.status);
      return;
    }

    const user = await res.json();
    console.log("Full User Data:", user);

    // Check for 1 (integer) or true (boolean)
    if (
      user.stripe_verif_completed === true ||
      user.stripe_verif_completed === 1
    ) {
      setStripeState("active");
    } else {
      setStripeState("none");
    }

    document.getElementById("navName").textContent =
      `${user.prenom} ${user.nom}`;
    avatar.textContent = user.prenom.charAt(0);
  } catch (err) {
    console.error("Fetch Error:", err);
  }
}

async function stripeConnect() {
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("token");

  try {
    const res = await fetch(
      `http://localhost:8081/admin/connect-stripe?id=${userId}`,
      {
        method: "POST",
        headers: { Authorization: `Bearer ${token}` },
      },
    );

    const data = await res.json();

    if (data.url) {
      window.location.href = data.url;
    }
  } catch (err) {
    alert("Erreur Stripe");
  }
}

function showTab(name) {
  document
    .querySelectorAll(".tab-panel")
    .forEach((p) => p.classList.remove("active"));
  document
    .querySelectorAll(".ptab")
    .forEach((t) => t.classList.remove("active"));

  document.getElementById("tab-" + name).classList.add("active");

  document.querySelectorAll(".ptab").forEach((t) => {
    if (t.getAttribute("onclick") === `showTab('${name}')`) {
      t.classList.add("active");
    }
  });
}
function logout() {
  localStorage.clear();
  window.location.href = "login.html";
}
