# Seed demo

Ajoute des donnees utiles pour la soutenance :

- Stripe ID de demo sur les comptes principaux ;
- compte pro marque premium ;
- annonces `[DEMO]` avec plusieurs statuts ;
- annonce sponsorisee pour montrer le tri premium.

Depuis le dossier `API` :

```powershell
go run ./cmd/seed_demo
```

Pour nettoyer :

```powershell
go run ./cmd/seed_demo -delete
```

Comptes modifies :

- `test.client@renova.test`
- `test.pro@renova.test`
- `test.salarie@renova.test`

Important : les Stripe ID sont des valeurs de demo. Elles servent a debloquer l'affichage et les parcours de demonstration. Pour un vrai paiement Stripe, il faut des IDs Stripe reels generes par Stripe.
