let currentItem = null;

async function loadOneAnnonce() {
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');

    if (!id) {
        document.getElementById('notFound').style.display = 'block';
        return;
    }

    try {
        // Fetch depuis ton API Go
        const response = await fetch(`http://localhost:8081/api/annonces?id=${id}`);
        if (!response.ok) throw new Error("404");
        
        const ann = await response.json();
        currentItem = ann; // Sauvegarde globale
        renderPage(ann);

    } catch (err) {
        console.error("Erreur JS:", err);
        document.getElementById('notFound').style.display = 'block';
    }
}

function renderPage(item) {
    const isFree = item.prix <= 0 || item.type.toLowerCase() === 'don';
    const imgSrc = item.image ? `http://localhost:8081${item.image}` : null;

    // Mise à jour textes fixes
    document.getElementById('bcCat').textContent = item.categorie || 'Objet';
    document.getElementById('bcTitle').textContent = item.titre;
    document.title = item.titre + ' — ReNova';

    document.getElementById('pageContent').innerHTML = `
    <div class="page">
      <div>
        <div class="img-hero fu" style="background: ${imgSrc ? `url('${imgSrc}')` : 'var(--bg4)'}">
          ${imgSrc ? '' : '📦'}
        </div>

        <div class="detail-section fu">
          <div class="ds-title">Description</div>
          <div class="ds-text">${item.description || "Aucune description fournie."}</div>
        </div>

        <div class="detail-section fu">
          <div class="ds-title">Informations</div>
          <p>📍 Lieu : ${item.ville} (${item.code_postal})</p>
          <p>✨ État : ${item.etat || 'Non spécifié'}</p>
          <p>⚖️ Poids : ${item.poids_kg} kg</p>
        </div>

        <div class="detail-section fu">
          <div class="ds-title">Vendeur</div>
          <div class="seller-card">
            <div class="seller-ava">👤</div>
            <div>
              <div style="font-weight:700; color:#fff">${item.prenom} ${item.nom}</div>
              <div style="font-size:12px; color:var(--txt-m)">Membre ReNova vérifié ✓</div>
            </div>
          </div>
        </div>
      </div>

      <div class="buy-box fu">
        <div class="buy-box-inner">
          <div class="bb-type ${isFree ? 'don' : 'vente'}">${isFree ? '🎁 Don gratuit' : '💰 À vendre'}</div>
          <h1 class="bb-title">${item.titre}</h1>
          
          <div class="bb-price-wrap">
            <div class="bb-price ${isFree ? 'free' : ''}">${isFree ? 'Gratuit' : item.prix + ' €'}</div>
            <div style="font-size:12px; color:var(--txt-m)">
              ${isFree ? 'Récupération en point relais ou main propre.' : 'Paiement sécurisé via Stripe.'}
            </div>
          </div>

          ${isFree 
            ? `<button class="btn-main btn-reserve" onclick="openCheckout('reserve')">🎁 Réserver l'objet</button>`
            : `<button class="btn-main btn-buy-now" onclick="openCheckout('buy')">💳 Acheter maintenant</button>`
          }
          
          <p style="font-size:11px; color:var(--txt-d); margin-top:15px; text-align:center">
             Transaction protégée par le système ReNova.
          </p>
        </div>
      </div>
    </div>`;
}

document.addEventListener('DOMContentLoaded', loadOneAnnonce);