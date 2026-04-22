let currentItem = null;

async function loadOneAnnonce() {
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');
    const paymentStatus = urlParams.get('payment');

    if (!id) {
        document.getElementById('notFound').style.display = 'block';
        return;
    }

    // --- GESTION DU RETOUR DE PAIEMENT ---
    if (paymentStatus === 'success') {
      console.log("Paiement détecté, mise à jour du statut...");

       
        try {
            const res = await fetch(`http://localhost:8081/api/annonces/vendre?id=${id}`, { 
                method: 'PUT' 
            });
           if (res.ok) {
            alert("🎉 Paiement réussi ! L'objet est maintenant à vous.");
            window.location.href = `oneAnnonce.html?id=${id}`;
        } else {
            console.error("Le serveur a renvoyé une erreur lors de la vente.");
        }
        } catch (e) {
            console.error("Erreur mise à jour statut:", e);
        }
    }

    try {
        const response = await fetch(`http://localhost:8081/api/annonces?id=${id}`);
        if (!response.ok) throw new Error("404");
        
        const ann = await response.json();
        currentItem = ann;
        renderPage(ann);

    } catch (err) {
        console.error("Erreur JS:", err);
        document.getElementById('notFound').style.display = 'block';
    }
}

async function openCheckout(type) {
    const buyerId = localStorage.getItem('userId');
    if (!buyerId) {
        alert("Veuillez vous connecter pour continuer.");
        window.location.href = "login.html";
        return;
    }

    if (type === 'buy') {
        try {
            const response = await fetch(`http://localhost:8081/api/payment-annonce?annonce_id=${currentItem.id}&buyer_id=${buyerId}`, {
                method: 'POST'
            });

            if (response.status === 403) {
                alert(" Configurez d'abord votre Stripe ID dans votre profil !");
                window.location.href = `profil.html?id=${buyerId}`;
                return;
            }

            const data = await response.json();
            window.location.href = data.url; 
        } catch (err) {
            alert("Erreur lors de la création de la session de paiement.");
        }
    }
}

function renderPage(item) {
    const isSold = item.statut_vente === 'VENDU';
    const isFree = item.prix <= 0 || item.type.toLowerCase() === 'don';
    const imgSrc = item.image ? `http://localhost:8081${item.image}` : null;

    document.getElementById('bcCat').textContent = item.categorie || 'Objet';
    document.getElementById('bcTitle').textContent = item.titre;
    document.title = item.titre + ' — ReNova';

    document.getElementById('pageContent').innerHTML = `
    <div class="page ${isSold ? 'is-sold' : ''}">
      <div>
        <div class="img-hero fu" style="background: ${imgSrc ? `url('${imgSrc}') center/cover` : 'var(--bg4)'}">
          ${imgSrc ? '' : '<div style="font-size:3rem"><i class="fa-solid fa-box-archive"></i></div>'}
          ${isSold ? '<div class="sold-tag">VENDU</div>' : ''}
        </div>

        <div class="detail-section fu">
          <div class="ds-title">Description</div>
          <div class="ds-text">${item.description || "Aucune description fournie."}</div>
        </div>

        <div class="detail-section fu">
          <div class="ds-title">Informations</div>
          <p><i class="fas fa-map-marker-alt"></i> Lieu : ${item.ville} (${item.code_postal})</p>
          <p><i class="fas fa-info-circle"></i> État : ${item.etat || 'Non spécifié'}</p>
          <p><i class="fas fa-weight-hanging"></i> Poids : ${item.poids_kg} kg</p>
        </div>

        <div class="detail-section fu">
          <div class="ds-title">Vendeur</div>
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
          <div class="bb-type ${isFree ? 'don' : 'vente'}">${isFree ? ' Don gratuit' : ' À vendre'}</div>
          <h1 class="bb-title">${item.titre}</h1>
          
          <div class="bb-price-wrap">
            <div class="bb-price ${isFree ? 'free' : ''}">${isFree ? 'Gratuit' : item.prix + ' €'}</div>
            
          </div>

          ${isSold 
            ? `<button class="btn-main btn-sold" disabled>
                 Cet objet a été vendu
               </button>`
            : (isFree 
                ? `<button class="btn-main btn-reserve" onclick="openCheckout('reserve')"> Réserver l'objet</button>`
                : `<button class="btn-main btn-buy-now" onclick="openCheckout('buy')"> <i class="fas fa-shopping-cart"></i> Acheter maintenant</button>`)
          }
          
          
        </div>
      </div>
    </div>`;
}

document.addEventListener('DOMContentLoaded', loadOneAnnonce);