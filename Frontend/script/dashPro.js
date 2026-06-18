
document.addEventListener("DOMContentLoaded", async () => {
  const token = localStorage.getItem("token");
  if (!token) {
    window.location.href = "login.html";
    return;
  }

  const userId = localStorage.getItem("userId");
  const firstName = localStorage.getItem("userName");

  const navNameDisplay = document.getElementById("navName");
  const heroNameDisplay = document.getElementById("heroName");

  if (firstName) {
    if (navNameDisplay) navNameDisplay.textContent = firstName;
    if (heroNameDisplay) heroNameDisplay.textContent = firstName;
  }

  const urlParams = new URLSearchParams(window.location.search);
  
  if (urlParams.get("abo") === "success") {
    const sessionId = urlParams.get("session_id");
    try {
      const upgradeRes = await fetch(`http://localhost:8081/api/pro/upgrade?id=${userId}&session_id=${sessionId}`, {
        method: "POST",
        headers: { "Authorization": `Bearer ${token}` }
      });

      if (upgradeRes.ok) {
        alert("Payment success! vous etes maintenant Premium. Profitez de votre abonnement!");
        window.history.replaceState(null, "", window.location.pathname); 
      }
    } catch (err) {
      console.error("Error upgrading account:", err);
    }
  } else if (urlParams.get("abo") === "cancel") {
    alert("Payment cancelled. You can upgrade anytime!");
    window.history.replaceState(null, "", window.location.pathname);
  }

  await checkPremiumStatus(token, userId);
});


async function subscribeToPremium() {
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("token");

  if (!userId || !token) {
    window.location.href = "login.html";
    return;
  }

  try {
    const response = await fetch(`http://localhost:8081/user/profile?id=${userId}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (!response.ok) throw new Error("Erreur lors de la récupération du profil");

    const data = await response.json();
    const user = Array.isArray(data) ? data[0] : data;

    if (!user.stripe_account_id || user.stripe_account_id.trim() === "") {
      const popup = document.getElementById("stripeWarningPopup");
      popup.style.display = "flex"; 
      return; 
    }
    const stripeResponse = await fetch(`http://localhost:8081/api/pro/subscribe?id=${userId}`, {
      method: "POST",
      headers: { Authorization: `Bearer ${token}` }
    });

    if (!stripeResponse.ok) throw new Error("Erreur lors de la creation de la session Stripe");

    const stripeData = await stripeResponse.json();
    
    if (stripeData.url) {
      window.location.href = stripeData.url;
    } else {
      alert("Erreur réseau avec Stripe.");
    }

  } catch (error) {
    console.error("Erreur:", error);
    alert("Impossible de vérifier l'état de votre compte.");
  }
}

function goToProfile() {
  const userId = localStorage.getItem("userId");

  if (!userId) {
    window.location.href = "login.html";
    return;
  }

  window.location.href = `profil.html?id=${userId}`;
}


async function checkPremiumStatus(token, userId) {
  try {
    await fetch(`http://localhost:8081/api/pro/sync?id=${userId}`, {
      method: "GET",
      headers: { Authorization: `Bearer ${token}` }
    });
  } catch (syncErr) {
    console.log("sync failed. log page anyways");
  }

try {
    const response = await fetch(`http://localhost:8081/user/profile?id=${userId}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    const data = await response.json();
    const user = Array.isArray(data) ? data[0] : data;

    const btnPasserPremium = document.getElementById("btnPasserPremium");
    const btnPremiumActif = document.getElementById("btnPremiumActif");
    const planSection = document.getElementById("plan"); 
    const premiumDashboard = document.getElementById("premiumDashboard"); 

    if (user.est_premium === 1 || user.est_premium === "1") {
      if (btnPasserPremium) btnPasserPremium.style.display = "none";
      if (btnPremiumActif) btnPremiumActif.style.display = "block"; 
      if (planSection) planSection.style.display = "none"; 
      if (premiumDashboard) premiumDashboard.style.display = "block"; 
      
    } else {
      if (btnPasserPremium) btnPasserPremium.style.display = "block"; 
      if (btnPremiumActif) btnPremiumActif.style.display = "none";
      if (planSection) planSection.style.display = "block"; 
      if (premiumDashboard) premiumDashboard.style.display = "none"; 
    }
    
  } catch (err) {
    console.error("Erreur lors de la verif du profil :", err);
  }
}


async function cancelPremium() {
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("token");

  try {
    const response = await fetch(`http://localhost:8081/api/pro/portal?id=${userId}`, {
      method: "POST",
      headers: { Authorization: `Bearer ${token}` }
    });

    if (!response.ok) throw new Error("Could not load portal");

    const data = await response.json();
    
    if (data.url) {
      window.location.href = data.url; 
    }
  } catch (err) {
    console.error("Error:", err);
    alert("Peut pas redirect vers la page abbonement");
  }
}