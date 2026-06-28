# Flux Annonce → Commande → Box (dépôt / retrait)

Documentation du parcours d'achat d'une annonce avec dépôt en conteneur (box)
et récupération par l'acheteur.

## Tables impliquées

| Table | Rôle | Champs clés |
|---|---|---|
| `annonce` | L'objet en vente | `statut_vente`, `prix`, `id_user` (vendeur) |
| `order` (commande) | L'achat | `id_acheteur`, `id_annonce`, `montant_total`, `commission`, `type` |
| `historique_conteneurs` | Le cœur : la transaction box | `acheteur_id`, `vendeur_id`, `annonce_id`, `conteneur_id` (= id de la box), `code_ouverture`, `code_barre_recuperation`, `date_reservation`, `date_depot_effective`, `date_retrait_effective` |
| `box` | Le casier | `numero`, `statut` (libre / reservee / occupee), `code_secret` |
| `conteneur` | Le locker physique | `nom`, `adresse` |
| `paiement` | Suivi Stripe | `id_commande`, `stripe_id`, `statut` |

## Le flux étape par étape

### 1. Achat

Route : `POST /api/annonces/vendre` → handler `ConfirmPaymentAndOrder`

- Crée la commande dans `order` (avec la commission de 5 % et `type = 'annonce'`).
- `ReserveBox` : trouve une box libre, génère les deux codes, crée la ligne
  `historique_conteneurs`, et passe la box en `reservee`.
- `annonce.statut_vente` → `EN ATTENTE DEPOT`.
- Notifie le vendeur et l'acheteur (push OneSignal).

### 2. Dépôt (par le vendeur)

Route : `POST /api/hardware/simulate-deposit` → handler `SimulateHardwareDeposit`

- Cherche la ligne par `code_ouverture` (le PIN du vendeur) et vérifie que
  `date_depot_effective IS NULL`.
- Pose `date_depot_effective`, passe la box en `occupee`.
- `annonce.statut_vente` → `EN ATTENTE DE RECUPERATION`.

### 3. Retrait (par l'acheteur)

Route : `POST /api/hardware/simulate-withdrawal` → handler `SimulateHardwareWithdrawal`

- Cherche la ligne par `code_barre_recuperation` et vérifie que l'objet est
  bien déposé (`date_depot_effective IS NOT NULL`) et pas déjà récupéré
  (`date_retrait_effective IS NULL`).
- Pose `date_retrait_effective`, libère la box (`statut = 'libre'`,
  `code_secret = NULL`).
- `annonce.statut_vente` → `VENDU`.

## Où sont stockés les codes

| Code | Pour qui | Stocké dans | Format | Exemple |
|---|---|---|---|---|
| Code de dépôt (PIN) | Vendeur | `historique_conteneurs.code_ouverture` (copié aussi dans `box.code_secret`) | numérique | `790452` |
| Code-barre de retrait | Acheteur | `historique_conteneurs.code_barre_recuperation` | `UC-<idAnnonce>-<idBox>` | `UC-7-1` |

Les deux codes sont générés au moment de l'achat (dans `ReserveBox`) et vivent
dans la même ligne de `historique_conteneurs`.

## Cycle de vie du statut de l'annonce

```
EN VENTE
  -> (achat)              EN ATTENTE DEPOT
  -> (dépôt PIN)          EN ATTENTE DE RECUPERATION
  -> (retrait code-barre) VENDU
```

## Côté commande / finance

- La commande (`order`) porte le montant et la commission de 5 % : c'est ce que
  lit le dashboard admin → Finances (volume d'affaires + commissions).
- Pour les annonces, aucune ligne `paiement` n'est créée (le statut affiché
  retombe sur « payé »).
- Pour les événements / formations payés via Stripe, une ligne `paiement` suit
  le statut (`pending` → `succeeded`, mis à jour par le webhook Stripe).

## Notes / valeurs d'enum

`annonce.statut_vente` accepte :
`EN VENTE`, `EN ATTENTE DEPOT`, `RESERVEE`, `EN BOX`,
`EN ATTENTE DE RECUPERATION`, `RECUPERE`, `VENDU`.

`order.type` : `annonce` ou `evenement`.
