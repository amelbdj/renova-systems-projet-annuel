# 09 - Timing repetition

Objectif final : 14 min 30, pas 15 min.

## Repartition cible

| Personne | Temps total cible | Parties |
|---|---:|---|
| Faty | 5 min 00 | Accueil, login, client, catalogue, conclusion |
| Amel | 5 min 00 | Validation, API Go, logistique, base SQL |
| Ndoya | 4 min 30 | Evenements, salarie, Docker/Nginx |

## Repetition 1 - Lecture lente

But : comprendre le script.

- Lire sans cliquer.
- Chronometrer chaque personne.
- Couper les phrases trop longues.
- Objectif : moins de 17 min.

## Repetition 2 - Clics reels

But : caler les onglets.

- Ouvrir les pages dans l'ordre du fichier `02_DEROULE_DEMO_CLIC_PAR_CLIC.md`.
- Verifier que chaque page charge sans erreur rouge bloquante.
- Noter les temps morts.
- Objectif : moins de 15 min 30.

## Repetition 3 - Version soutenance

But : tenir le timing.

- Une seule personne controle la souris ou bien chaque personne controle sa partie, mais il faut choisir avant.
- Interdiction d'improviser une fonctionnalite externe.
- Si une action prend plus de 15 secondes, passer au plan B.
- Objectif : 14 min 30.

## Signaux de transition

- Faty vers Amel : "Une fois l'annonce soumise, elle passe par le back-office."
- Amel vers Faty : "Apres validation, l'annonce est visible cote catalogue."
- Faty vers Ndoya : "La plateforme ne s'arrete pas au particulier : elle sert aussi les professionnels et l'equipe interne."
- Ndoya vers Amel : "Je termine sur le deploiement, puis Amel montre la base SQL."
- Amel vers conclusion : "On revient sur la valeur globale livree."

## Regles de parole

- Ne pas dire "mon code" : dire "notre API", "notre front", "notre workflow".
- Ne pas dire "normalement" : dire "pour securiser la demo".
- Ne pas dire "c'est juste une page" : dire "cet ecran permet a l'utilisateur de...".
- Ne pas s'excuser si une integration externe n'est pas lancee : expliquer que le risque est maitrise.

## Checklist juste avant passage

- [ ] Navigateur ouvert sur accueil Docker.
- [ ] Zoom navigateur a 90 ou 100 %, pas devtools visible.
- [ ] Comptes de test notes.
- [ ] Captures de secours ouvertes dans un dossier.
- [ ] Terminal pret avec `docker compose ps`.
- [ ] IDE ouvert sur `API/route/annonces.go`.
- [ ] phpMyAdmin ou `db/init.sql` pret.
- [ ] Chacune connait sa premiere phrase.
- [ ] Chacune connait sa transition.

## Chrono de secours

Si a 07:30 vous n'avez pas fini admin validation, couper :
- forum ;
- paiement Stripe live ;
- PDF live ;
- creation evenement live.

Si a 12:00 vous n'avez pas commence Docker, passer directement a :
- `docker compose ps`;
- `nginx.conf`;
- conclusion.
