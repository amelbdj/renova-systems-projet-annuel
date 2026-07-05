# 02 - Deroule demo clic par clic

## Preparation avant de commencer

1. Lancer Docker : `docker compose up -d --build`.
2. Verifier : `docker compose ps`.
3. Ouvrir le front Docker : `http://localhost:8088/` si le conteneur expose ce port, sinon verifier la colonne PORTS.
4. Ouvrir aussi un terminal dans le dossier projet.
5. Ouvrir phpMyAdmin ou `db/init.sql` pour la preuve SQL.
6. Preparer des comptes de demo. Les emails existent dans `db/init.sql`, mais les mots de passe sont hashes. Si vous ne connaissez pas les mots de passe, resetter avant la soutenance.

## Comptes a preparer

Comptes visibles dans `db/init.sql` :

- `test.client@renova.test` role Utilisateur
- `test.admin@renova.test` role Administrateur
- `test.salarie@renova.test` role Salarie
- `test.pro@renova.test` role Pro
- `test.user@test.fr` role Utilisateur
- `test.salarie@test.fr` role Salarie
- `test.admin@test.fr` role Administrateur
- `test.pro@test.fr` role Pro

Mot de passe : non lisible dans le dump car hash bcrypt. A preparer avant demo, par exemple `Test123!`.

## Demo principale

1. Ouvrir `http://localhost:8088/`.
   - Montrer : accueil, menu, positionnement UpcycleConnect.
   - Dire : "plateforme metier d'economie circulaire".
   - Eviter : rester trop longtemps sur le design.

2. Aller sur `login.html`.
   - Montrer : choix des espaces.
   - Se connecter avec un compte utilisateur.
   - Page attendue : dashboard particulier.
   - Eviter : creer un compte en live si le temps est serre.

3. Dashboard particulier.
   - Cliquer sur "Mes annonces" ou section annonce.
   - Montrer : annonces existantes, creation annonce.
   - Dire : "l'annonce passe ensuite en validation".
   - Eviter : upload image lourd.

4. Aller sur admin.
   - Se connecter avec compte admin.
   - Ouvrir `admin_validations.html`.
   - Montrer : onglets annonces, evenements, contenus.
   - Cliquer seulement si une donnee de test est prete.
   - Eviter : valider toutes les donnees et vider la demo.

5. Montrer code API.
   - Ouvrir `API/route/annonces.go`.
   - Ouvrir `API/admin/annonce.go`.
   - Ouvrir `API/bdd/annonceReq.go`.
   - Montrer : route -> handler -> requete SQL.

6. Catalogue.
   - Ouvrir `annonceAll.html`.
   - Ouvrir une annonce.
   - Montrer : seuls les contenus valides apparaissent.
   - Eviter : lancer paiement Stripe si non prepare.

7. Logistique.
   - Ouvrir `admin_conteneurs.html`.
   - Montrer : conteneurs, box, statuts.
   - Ouvrir `simulateur.html` si un code de test existe.
   - Dire : "le code simule l'interaction avec une borne".

8. Professionnel.
   - Se connecter compte pro.
   - Ouvrir `espPro.html`.
   - Montrer : annonces, projets, etapes, premium.
   - Eviter : checkout Stripe live sauf deja teste.

9. Evenements et formations.
   - Ouvrir `evenement.html`.
   - Puis salarie : `salarie/salarie_events.html`.
   - Montrer : creation evenement/formation, PDF ressources.
   - Eviter : upload gros PDF.

10. Salarie moderation.
    - Ouvrir `salarie/salarie_forum.html`.
    - Montrer : moderation messages et bannissement possible.
    - Eviter : bannir un compte utile.

11. Docker / Nginx.
    - Terminal : `docker compose ps`.
    - Ouvrir `docker-compose.yml`.
    - Ouvrir `Frontend/nginx.conf`.
    - Dire : "Nginx sert le front et proxy les appels API".

12. Base SQL.
    - Ouvrir `db/init.sql`.
    - Montrer tables : `utilisateur`, `annonce`, `evenement`, `conteneur`, `box`, `translations`, `document`.
    - Dire : "donnees prechargees pour demo fiable".

13. Conclusion.
    - Revenir accueil ou admin dashboard.
    - Dire la phrase de conclusion du script.
