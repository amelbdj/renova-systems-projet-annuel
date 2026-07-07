let currentItem = null;


if (localStorage.getItem("userRole") === "Pro") {
  document.documentElement.classList.add("theme-pro");
}

async function loadOneAnnonce() {
  const urlParams = new URLSearchParams(window.location.search);
  const id = urlParams.get("id");
  const paymentStatus = urlParams.get("payment");

  if (!id) {
    document.getElementById("notFound").style.display = "block";
    return;
  }

  if (paymentStatus === "success") {
    console.log("Paiement détecté, mise à jour du statut...");
    const buyerId = urlParams.get("buyer_id");

    try {
      const res = await fetch(
        `${API_BASE_URL}/api/annonces/vendre?id=${id}&buyer_id=${buyerId}`,
        { method: "POST" },
      );
      if (res.ok) {
        if (typeof t === "function") {
          alert(t("oneAnnonce.payment_success"));
        } else {
          alert("Paiement reussi ! L'objet est maintenant a vous.");
        }
        window.location.href = `/annonce?id=${id}`;
      } else {
        console.error("Le serveur a renvoyé une erreur lors de la vente.");
      }
    } catch (e) {
      console.error("Erreur mise à jour statut:", e);
    }
  }

  try {
    const response = await fetch(`${API_BASE_URL}/api/annonces?id=${id}`);
    if (!response.ok) throw new Error("404");

    const ann = await response.json();
    currentItem = ann;
    renderPage(ann);
  } catch (err) {
    console.error("Erreur JS:", err);
    document.getElementById("notFound").style.display = "block";
  }
}

async function openCheckout(type) {
  const buyerId = localStorage.getItem("userId");
  if (!buyerId) {
    alert(t("oneAnnonce.login_required"));
    window.location.href = "/login";
    return;
  }

  if (type === "buy") {
    try {
      const response = await fetch(
        `${API_BASE_URL}/api/payment-annonce?annonce_id=${currentItem.id}&buyer_id=${buyerId}`,
        { method: "POST" },
      );

      if (response.status === 403) {
        alert(t("oneAnnonce.stripe_required"));
        window.location.href = `/profil?id=${buyerId}`;
        return;
      }

      const data = await response.json();
      window.location.href = data.url;
    } catch (err) {
      alert(t("oneAnnonce.payment_error"));
    }
  } else if (type === "reserve") {
    const confirmReserve = confirm(t("oneAnnonce.reserve_confirm"));

    if (!confirmReserve) return;

    try {
      const response = await fetch(
        `${API_BASE_URL}/api/annonces/vendre?id=${currentItem.id}&buyer_id=${buyerId}`,
        { method: "POST" },
      );

      if (response.ok) {
        alert(t("oneAnnonce.reserve_success"));
        window.location.reload();
      } else {
        alert(t("oneAnnonce.reserve_error"));
      }
    } catch (err) {
      console.error("Erreur de réservation:", err);
      alert(t("oneAnnonce.server_error"));
    }
  }
}

function renderPage(item) {
  const isSold =
    item.statut_vente === "VENDU" ||
    item.statut_vente === "EN ATTENTE DEPOT" ||
    item.statut_vente === "EN BOX";
  const isFree = item.prix <= 0 || item.type.toLowerCase() === "don";
  const imgSrc = item.image ? `${API_BASE_URL}${item.image}` : null;

  document.getElementById("bcCat").textContent = item.categorie || "Objet";
  document.getElementById("bcTitle").textContent = item.titre;
  document.title = item.titre + " — UpcycleConnect";

  document.getElementById("pageContent").innerHTML = `
    <div class="page ${isSold ? "is-sold" : ""}">
      <div>
        <div class="img-hero fu" style="background: ${imgSrc ? `url('${imgSrc}') center/cover` : "var(--bg4)"}">
          ${imgSrc ? "" : '<div style="font-size:3rem"><i class="fa-solid fa-box-archive"></i></div>'}
          ${isSold ? '<div class="sold-tag" data-i18n="oneAnnonce.sold_tag">VENDU</div>' : ""}
        </div>

        <div class="detail-section fu">
          <div class="ds-title" data-i18n="oneAnnonce.description">Description</div>
          <div class="ds-text">${item.description || '<span data-i18n="oneAnnonce.no_description">Aucune description fournie.</span>'}</div>
        </div>

        <div class="detail-section fu">
          <div class="ds-title" data-i18n="oneAnnonce.info">Informations</div>
          <p><i class="fas fa-map-marker-alt"></i> <span data-i18n="oneAnnonce.place">Lieu</span> : ${item.ville} (${item.code_postal})</p>
          <p><i class="fas fa-info-circle"></i> <span data-i18n="oneAnnonce.condition">État</span> : ${item.etat || '<span data-i18n="oneAnnonce.unspecified">Non spécifié</span>'}</p>
          <p><i class="fas fa-weight-hanging"></i> <span data-i18n="oneAnnonce.weight">Poids</span> : ${item.poids_kg} kg</p>
        </div>

        <div class="detail-section fu">
          <div class="ds-title" data-i18n="oneAnnonce.seller">Vendeur</div>
          <div class="seller-card">
            <div class="seller-ava"><i class="fas fa-user"></i></div>
            <div>
              <div style="font-weight:700; color:#fff">${item.prenom} ${item.nom}</div>
            </div>
          </div>
        </div>
      </div>

      <div class="buy-box fu">
        <div class="buy-box-inner">
          <div class="bb-type ${isFree ? "don" : "vente"}">${isFree ? '<span data-i18n="annonce.type.donation">Don gratuit</span>' : '<span data-i18n="oneAnnonce.for_sale">À vendre</span>'}</div>
          <h1 class="bb-title">${item.titre}</h1>

          <div class="bb-price-wrap">
            <div class="bb-price ${isFree ? "free" : ""}">${isFree ? '<span data-i18n="annonce.card.free">Gratuit</span>' : item.prix + " €"}</div>
          </div>

          ${!isSold ? '<div id="contact-zone" style="margin-bottom: 15px;"></div>' : ""}

          ${
            isSold
              ? `<button class="btn-main btn-sold" disabled data-i18n="oneAnnonce.sold_btn">
                 Cet objet a été vendu
               </button>`
              : isFree
                ? `<button class="btn-main btn-reserve" onclick="openCheckout('reserve')" data-i18n="oneAnnonce.reserve_btn"> Réserver l'objet</button>`
                : `<button class="btn-main btn-buy-now" onclick="openCheckout('buy')"> <i class="fas fa-shopping-cart"></i> <span data-i18n="oneAnnonce.buy_btn">Acheter maintenant</span></button>`
          }
          
        </div>
      </div>
    </div>`;

  
  if (typeof appliquerTraductions === "function") appliquerTraductions();

  const monUserId = parseInt(localStorage.getItem("userId"));

  const vendeurId = parseInt(item.id_user);

  if (!isSold && monUserId !== vendeurId && !isNaN(vendeurId)) {
    const btnContainer = document.getElementById("contact-zone");
    if (btnContainer) {
      btnContainer.innerHTML = `
              <button id="btn-dynamic-contact" class="btn-main" style="background-color: var(--vi); margin-bottom: 10px;">
                  <i class="fas fa-comment-dots"></i> <span data-i18n="oneAnnonce.contact_seller">Contacter le vendeur</span>
              </button>`;
      if (typeof appliquerTraductions === "function") appliquerTraductions();

      document
        .getElementById("btn-dynamic-contact")
        .addEventListener("click", () => {
          openChat(item.id, vendeurId);
        });
    }
  }
}
document.addEventListener("DOMContentLoaded", loadOneAnnonce);
