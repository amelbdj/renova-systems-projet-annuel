# 05 - Checklist points a montrer

## A. Obligatoire a montrer

- [ ] Site public : `landing.html`
- [ ] Connexion : `login.html`
- [ ] Au moins deux roles : Utilisateur + Admin, idealement aussi Salarie ou Pro
- [ ] Dashboard particulier : `espClient.html`
- [ ] Creation ou formulaire annonce
- [ ] Catalogue annonce : `annonceAll.html`
- [ ] Workflow / statut : annonce en attente puis validation admin
- [ ] Back-office admin : `admin_dashboard.html` ou `admin_validations.html`
- [ ] Utilisateurs / roles : `admin_users.html`
- [ ] Validation / moderation : admin validations ou salarie forum
- [ ] API Go : `API/route/*.go`, `API/admin/*.go`, endpoint `/api/languages`
- [ ] Base SQL : `db/init.sql`, tables principales
- [ ] Docker : `docker compose ps`
- [ ] Nginx : `Frontend/nginx.conf`
- [ ] Site Docker : `http://localhost:8088/` ou port affiche par Docker

## B. Bonus a montrer si stable

- [ ] Upcycling Score : dashboard client / `API/bdd/boxReq.go`
- [ ] Commission : admin finances
- [ ] Facture PDF : document deja genere
- [ ] Code PIN : simulateur depot/retrait
- [ ] Conteneur : admin conteneurs
- [ ] Notifications : cloche ou panneau notifications
- [ ] Multilingue : FR/EN
- [ ] Paiement : ecran Stripe ou capture, pas forcement live
- [ ] Statistiques : admin dashboard / finances / salarie forum
- [ ] Evenements : catalogue public
- [ ] Formations : creation salarie + ressources PDF
- [ ] Projets d'upcycling : dashboard pro

## A ne pas promettre

- [ ] Swagger/OpenAPI complet : non trouve
- [ ] QR code complet : non trouve clairement dans le code
- [ ] Paiement reel sans preparation Stripe
- [ ] Envoi mail SMTP sans test juste avant
