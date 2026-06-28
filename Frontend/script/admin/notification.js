function EnvoyerNotification() {
  var cible = document.getElementById("notif-cible").value;
  var message = document.getElementById("notif-message").value.trim();
  var feedback = document.getElementById("notif-feedback");

  if (!message) {
    if (feedback) {
      feedback.style.color = "#e74c3c";
      feedback.textContent = "Veuillez écrire un message.";
    }
    return;
  }

  if (feedback) {
    feedback.style.color = "var(--txt-m)";
    feedback.textContent = "Envoi en cours…";
  }

  var token = localStorage.getItem("token");

  fetch(`${API_BASE_URL}/admin/notifications/send`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + token,
    },
    body: JSON.stringify({ cible: cible, message: message }),
  })
    .then(function (response) {
      return response.json().then(function (data) {
        if (!response.ok) {
          throw new Error(data.erreur || "Erreur lors de l'envoi");
        }
        return data;
      });
    })
    .then(function (data) {
      if (feedback) {
        feedback.style.color = "#2ecc71";
        feedback.textContent =
          "Notification envoyée à " + data.nombre + " personne(s).";
      }
      document.getElementById("notif-message").value = "";
    })
    .catch(function (error) {
      if (feedback) {
        feedback.style.color = "#e74c3c";
        feedback.textContent = error.message;
      }
      console.error(error);
    });
}
