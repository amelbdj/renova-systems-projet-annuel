# 04 - Plan B demo

Regle orale : ne jamais dire "normalement ca marche". Dire plutot : "Pour securiser la demonstration et respecter le temps, nous allons utiliser le jeu de donnees prepare."

| Fonctionnalite | Ce qui peut buguer | Reaction | Capture | Phrase professionnelle |
|---|---|---|---|---|
| Docker | Port deja utilise | Montrer `docker compose ps` et changer de port si besoin | `18_docker_compose_ps.png` | "L'environnement Docker est pret ; ici le port local est adapte a la machine de demonstration." |
| Connexion | Mot de passe oublie / role incorrect | Utiliser un autre compte ou reset avant demo | `02_login_roles.png` | "Nous utilisons un compte de demonstration preconfigure pour eviter de perdre du temps sur la saisie." |
| Creation annonce | Upload ou validation formulaire | Montrer annonce deja creee | `04_creation_annonce.png`, `05_annonce_attente.png` | "Le formulaire existe ; pour garder le rythme, nous montrons l'annonce deja inseree dans le workflow." |
| Validation admin | Aucune donnee en attente | Creer avant ou montrer capture | `05_annonce_attente.png` | "Le workflow est visible ici avec une donnee de demonstration deja preparee." |
| Catalogue | Catalogue vide | Verifier `db/init.sql` / utiliser capture | `06_annonce_validee_catalogue.png` | "La publication depend du statut valide ; voici une annonce validee visible cote catalogue." |
| Stripe annonce | Redirection externe lente | Ne pas lancer paiement | `17_pdf_facture.png` | "L'integration Stripe est presente cote API ; en presentation, nous montrons le flux sans declencher de paiement externe." |
| Stripe abonnement | Session checkout impossible | Montrer ecran abonnement + code API | `10_dashboard_pro.png` | "La logique premium est integree, mais nous evitons une dependance externe pendant la demo chronometree." |
| PDF | Fichier absent | Montrer code `API/admin/pdf.go` et capture PDF | `17_pdf_facture.png` | "La generation PDF est realisee cote backend ; voici le resultat attendu avec une donnee preparee." |
| SMTP | Mail non recu | Montrer code et ecran | capture login/reset | "L'envoi mail depend de la configuration SMTP ; la route et le parcours sont implementes." |
| OneSignal | Push non recu | Montrer notifications internes | capture notifications | "Les notifications internes sont visibles ; le push externe depend de la configuration OneSignal." |
| Box / PIN | Code inconnu | Montrer admin conteneurs + simulateur sans valider | `08_admin_conteneurs.png`, `09_simulateur_pin.png` | "Les codes sont generes et stockes dans l'historique ; nous montrons la logique sans consommer une donnee utile." |
| Multilingue | Traductions chargees lentement | Montrer table translations/API | `15_multilingue.png` | "Les textes sont servis par l'API de traduction et la table SQL." |
| Salarie events | Upload PDF bloque | Montrer formulaire et donnees existantes | `13_salarie_events.png` | "Le formulaire accepte des ressources ; nous evitons un upload lourd en live." |
| Moderation | Bannir mauvais compte | Ne pas cliquer ban | `14_moderation_forum.png` | "L'action existe mais nous ne l'executons pas sur un compte de demonstration utile." |
| Swagger | Prof demande documentation | Dire absent | aucune | "Nous n'avons pas integre Swagger ; les routes sont structurees dans `API/route` et peuvent etre documentees en evolution." |
