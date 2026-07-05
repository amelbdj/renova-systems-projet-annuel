# Seed users

Petit script local pour ajouter beaucoup d'utilisateurs de test dans la table `utilisateur`.

Depuis le dossier `API` :

```powershell
go run ./cmd/seed_users
```

Par defaut, ca ajoute 15 000 utilisateurs avec le mot de passe :

```txt
Test123!
```

Pour choisir le nombre :

```powershell
go run ./cmd/seed_users -count 1000
```

Pour choisir le mot de passe :

```powershell
go run ./cmd/seed_users -count 15000 -password Test123!
```

Pour supprimer les utilisateurs de test :

```powershell
go run ./cmd/seed_users -delete
```

Les emails crees ressemblent a :

```txt
user_seed_00001@renova.test
user_seed_00002@renova.test
```

Le script utilise la meme configuration de base que l'API :

- `DB_HOST`, defaut `localhost`
- `DB_PORT`, defaut `3306`
- `DB_USER`, defaut `root`
- `DB_PASS`, defaut `root`
- `DB_NAME`, defaut `pa2026`
