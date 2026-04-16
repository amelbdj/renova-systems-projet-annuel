document.addEventListener('DOMContentLoaded', () => {
    const token = localStorage.getItem('token');
    if (!token) {
        window.location.href = "login.html";
        return;
    }

    const firstName = localStorage.getItem('userName');
    const score = localStorage.getItem('userScore') || 0;
    const hasSeenTutorial = localStorage.getItem('tutorielVu');
    const tutorialOverlay = document.getElementById('tut');

    const navNameDisplay = document.getElementById('navName');
    const heroNameDisplay = document.getElementById('heroName');

    if (firstName) {
        if (navNameDisplay) navNameDisplay.textContent = firstName;
        if (heroNameDisplay) heroNameDisplay.textContent = firstName;
    }

    const allScoreDisplays = document.querySelectorAll('.user-score');
    allScoreDisplays.forEach(element => {
        element.textContent = score;
    });

    const scoreFill = document.querySelector('.score-lfill');
    if (scoreFill) {
        const percentage = Math.min((score / 1000) * 100, 100);
        scoreFill.style.width = percentage + "%";
    }

    if (hasSeenTutorial === "true") {
        if (tutorialOverlay) tutorialOverlay.style.display = 'none';
    } else {
        console.log("showing the tutorial");
    }
    
    loadMyAnnonces(); 
});

function logout() {
    localStorage.clear();
    window.location.href = "login.html";
}

async function loadMyAnnonces() {
    const userId = localStorage.getItem('userId');
    const annGrid = document.querySelector('.ann-grid');
    if (!annGrid) return;

    try {
        const response = await fetch(`http://localhost:8081/mes-annonces?id=${userId}`);
        const annonces = await response.json();
        
        annGrid.innerHTML = '';

        if (!annonces || annonces.length === 0) {
            annGrid.innerHTML = '<p style="color:var(--txt-m); padding:20px;">Vous n\'avez pas encore d\'annonces.</p>';
        } else {
            annonces.forEach(ann => {
                const card = document.createElement('div');
                card.className = 'ann-card fu';
                const statusClass = ann.statut_validation === 'Validée' ? 't-green' : 't-amber';

                card.innerHTML = `
                    <div class="ann-thumb" style="background:linear-gradient(135deg,#0a0f1e,#101828)">📦</div>
                    <div class="ann-body">
                        <div class="ann-name">${ann.titre}</div>
                        <div class="ann-meta">${ann.prix > 0 ? ann.prix + ' €' : 'Don gratuit'}</div>
                        <div class="ann-foot">
                            <span class="tag ${statusClass}">${ann.statut_validation}</span>
                            <button class="btn btn-xs" 
                                    style="background:rgba(255,92,92,.08); color:var(--red); border:1px solid rgba(255,92,92,.2)" 
                                    onclick="deleteAnnonce(${ann.id})">
                                Supprimer
                            </button>
                            <button class="btn btn-g btn-xs" onclick='openEditForm(${JSON.stringify(ann)})'>Modifier</button>
                        </div>
                    </div>
                `;
                annGrid.appendChild(card);
            });
        }

        const addBox = document.createElement('div');
        addBox.className = 'ann-add';
        addBox.onclick = toggleAnnForm;
        addBox.innerHTML = '<div class="plus">＋</div><span>Ajouter une annonce</span>';
        annGrid.appendChild(addBox);

    } catch (err) {
        console.error("Erreur chargement annonces:", err);
    }
}

async function deleteAnnonce(id) {
    if(!confirm("Voulez-vous supprimez l'annonce ?")) return;
    
    const token = localStorage.getItem('token');
    try {
        const response = await fetch(`http://localhost:8081/admin/annonces/delete/${id}`, {
            method: 'DELETE',
            headers: { 'Authorization': `Bearer ${token}` }
        });

        if (response.ok) {
            loadMyAnnonces();
        }
    } catch (err) {
        console.error("Delete error:", err);
    }
}

function openEditForm(ann) {
    document.querySelector('#annForm input[type="text"]').value = ann.titre;
    document.querySelector('#annForm textarea').value = ann.description;
    document.querySelector('#annForm input[type="number"]').value = ann.prix;
    document.querySelector('#annForm select').value = ann.type;
    
    document.getElementById('editAnnId').value = ann.id;
    
    document.getElementById('formTitle').textContent = "Modifier l'annonce";
    document.querySelector('#annForm .btn-p').textContent = "Enregistrer les modifications";
    
    const f = document.getElementById('annForm');
    f.style.display = 'block';
    f.scrollIntoView({ behavior: 'smooth', block: 'start' });
}