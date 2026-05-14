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

  fetch("http://localhost:8081/admin/finance/overview", {
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

  fetch("http://localhost:8081/admin/finance/transactions", {
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

          const tr = document.createElement("tr");
          tr.innerHTML = `
                    <td style="color: var(--txt-m);">#${tx.id}</td>
                    <td>${dateStr}</td>
                    <td style="font-weight: 700; color: white;">${tx.titre}</td>
                    <td>${new Intl.NumberFormat("fr-FR").format(tx.montant)} €</td>
                    <td style="color: #4ade80; font-weight: bold;">+ ${new Intl.NumberFormat("fr-FR").format(tx.commission)} €</td>
                    <td><span class="badge-stripe">Payé</span></td>
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
