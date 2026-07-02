# 11 — Checklist avant soutenance

## La veille
- [ ] `git add -A && git commit -m "Sauvegarde avant soutenance"` (sauvegarde du projet).
- [ ] Créer la branche de secours : `git checkout -b soutenance-modifs` puis revenir sur `amel`.
- [ ] Exporter la base : `docker exec uc_mysql mysqldump -uupcycle -pupcyclePass123 pa2026 > backup_soutenance.sql`.
- [ ] Vérifier que la **prod** est à jour (images `frontend:v32` / `backend:v8` déployées) : `docker compose -f docker-compose-prod.yml pull && up -d` sur la VM.
- [ ] Générer le **PDF du kit** (voir `KIT_SOUTENANCE_DEV_COMPLET.md`) et l'avoir hors-ligne.
- [ ] Faire des **captures d'écran de secours** de chaque page clé (au cas où le réseau tombe).

## Le jour J (30 min avant)
- [ ] **Projet lancé** : `docker compose up -d` → `docker ps` montre 3 conteneurs **Up**.
- [ ] **MySQL healthy** : `docker ps` (statut `healthy`).
- [ ] **API OK** : `curl -s -o /dev/null -w "%{http_code}" http://localhost:8081/admin/users` → `401`.
- [ ] **Front OK** : ouvrir `http://localhost:8088/` → landing s'affiche.
- [ ] **Base OK** : `docker exec uc_mysql mysql -uupcycle -pupcyclePass123 pa2026 -e "SELECT COUNT(*) FROM utilisateur;"`.
- [ ] **Comptes de test OK** : se connecter avec 1 compte de chaque rôle (voir `09_COMPTES_TEST.md`, mots de passe remplis).
- [ ] **Site externe accessible** : ouvrir `https://upcycleconnect.pro` (preuve non-localhost).
- [ ] **Postman/curl prêt** : collection avec login + 2-3 routes protégées + le token en variable.
- [ ] **Exports DB prêts** : `backup_soutenance.sql` accessible.
- [ ] **Cache navigateur** : ouvrir DevTools → « Désactiver le cache » pour les modifs en direct.

## Onglets à préparer
- [ ] `https://upcycleconnect.pro/` (prod)
- [ ] `http://localhost:8088/` (local, pour modif)
- [ ] Un terminal à la racine du projet
- [ ] L'éditeur de code ouvert sur `API/` et `Frontend/script/`
- [ ] Ce kit (PDF) ouvert

## Réflexes anti-panique
- [ ] Savoir dire où se trouve **chaque fichier** (voir `01_ARCHITECTURE`).
- [ ] Avoir répété **2-3 exercices de modif** (voir `12_EXERCICES`).
- [ ] Connaître les 4 commandes clés : `docker compose up -d --build`, `docker compose logs -f backend`, `docker ps`, `docker exec -it uc_mysql mysql ...`.
