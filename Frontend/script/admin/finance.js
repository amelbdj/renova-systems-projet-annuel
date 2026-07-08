document.addEventListener("DOMContentLoaded", () => {
  chargerKPIsFinance();
  chargerTransactions();
});

function chargerKPIsFinance() {
  const token = localStorage.getItem("token");

  const volumeEl = document.getElementById("total-volume");
  const commissionEl = document.getElementById("total-commission");

  if (!volumeEl || !commissionEl) {
    return;
  }

  fetch(`${API_BASE_URL}/admin/finance/overview`, {
    headers: { Authorization: `Bearer ${token}` },
  })
    .then((res) => {
      if (!res.ok) {
        throw new Error("Erreur réseau");
      }
      return res.json();
    })
    .then((data) => {
      volumeEl.textContent =
        new Intl.NumberFormat("fr-FR").format(data.volumeMois) + " €";
      commissionEl.textContent =
        new Intl.NumberFormat("fr-FR").format(data.revenuMois) + " €";
    })
    .catch((err) => {
      console.error("Erreur lors du chargement des KPIs:", err);
    });
}

function chargerTransactions() {
  const token = localStorage.getItem("token");
  const tbody = document.getElementById("table-transactions");

  if (!tbody) {
    return;
  }

  fetch(`${API_BASE_URL}/admin/finance/transactions`, {
    headers: { Authorization: `Bearer ${token}` },
  })
    .then((res) => {
      if (!res.ok) {
        throw new Error("Erreur réseau");
      }
      return res.json();
    })
    .then((txs) => {
      if (txs && txs.length > 0) {
        tbody.innerHTML = "";

        txs.forEach((tx) => {
          const dateObj = new Date(tx.date);
          const dateStr =
            dateObj.toLocaleDateString("fr-FR") +
            " " +
            dateObj.toLocaleTimeString("fr-FR", {
              hour: "2-digit",
              minute: "2-digit",
            });

          var typeLabel = "Annonce";
          if (tx.type === "evenement") {
            typeLabel = "Formation / Événement";
          }
          if (tx.type === "abonnement") {
            typeLabel = "Abonnement";
          }

          var statut = (tx.statut || "").toLowerCase();
          var statutTexte = tx.statut;
          var statutCouleur = "#4ade80";
          if (statut === "succeeded" || statut === "payé" || statut === "paid" || statut === "actif") {
            statutTexte = "Payé";
            statutCouleur = "#4ade80";
          } else if (statut === "pending") {
            statutTexte = "En attente";
            statutCouleur = "#f5a623";
          } else if (statut === "failed" || statut === "refunded" || statut === "resilie") {
            statutTexte = "Échoué";
            statutCouleur = "#ef4444";
          }

          const tr = document.createElement("tr");
          tr.innerHTML = `
                    <td style="color: var(--txt-m);">#${tx.id}</td>
                    <td>${dateStr}</td>
                    <td style="font-weight: 700; color: white;">${tx.titre}<br><span style="font-size:11px; color: var(--txt-m); font-weight:400;">${typeLabel}</span></td>
                    <td>${new Intl.NumberFormat("fr-FR").format(tx.montant)} €</td>
                    <td style="color: #4ade80; font-weight: bold;">+ ${new Intl.NumberFormat("fr-FR").format(tx.commission)} €</td>
                    <td><span class="badge-stripe" style="color:${statutCouleur};">${statutTexte}</span></td>
                `;
          tbody.appendChild(tr);
        });
      } else {
        tbody.innerHTML =
          '<tr><td colspan="6" style="text-align: center; color: var(--txt-m);">Aucune transaction récente.</td></tr>';
      }
    })
    .catch((err) => {
      console.error("Erreur TX:", err);
      tbody.innerHTML =
        '<tr><td colspan="6" style="text-align: center; color: var(--red);">Erreur lors du chargement des transactions.</td></tr>';
    });
}
