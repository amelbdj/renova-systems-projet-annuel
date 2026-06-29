document.addEventListener("DOMContentLoaded", () => {
  chargerDocuments();
});

function chargerDocuments() {
  const token = localStorage.getItem("token");
  const tbody = document.getElementById("table-documents");
  if (!tbody) {
    return;
  }

  fetch(`${API_BASE_URL}/admin/documents`, {
    headers: { Authorization: `Bearer ${token}` },
  })
    .then((res) => {
      if (!res.ok) {
        throw new Error("Erreur réseau");
      }
      return res.json();
    })
    .then((docs) => {
      if (docs && docs.length > 0) {
        tbody.innerHTML = "";

        docs.forEach((doc) => {
          const dateObj = new Date(doc.date);
          const dateStr =
            dateObj.toLocaleDateString("fr-FR") +
            " " +
            dateObj.toLocaleTimeString("fr-FR", {
              hour: "2-digit",
              minute: "2-digit",
            });

          let typeLabel = doc.type === "facture" ? "Facture" : doc.type;

          const tr = document.createElement("tr");
          tr.innerHTML = `
            <td style="font-weight: 700; color: white;">${typeLabel}</td>
            <td>${doc.utilisateur || "—"}</td>
            <td style="color: var(--txt-m);">${dateStr}</td>
            <td><a href="${API_BASE_URL}${doc.url_pdf}" target="_blank" class="btn btn-g btn-xs">Ouvrir le PDF</a></td>
          `;
          tbody.appendChild(tr);
        });
      } else {
        tbody.innerHTML =
          '<tr><td colspan="4" style="text-align: center; color: var(--txt-m);">Aucun document généré pour le moment.</td></tr>';
      }
    })
    .catch((err) => {
      console.error("Erreur documents:", err);
      tbody.innerHTML =
        '<tr><td colspan="4" style="text-align: center; color: var(--red);">Erreur lors du chargement des documents.</td></tr>';
    });
}
