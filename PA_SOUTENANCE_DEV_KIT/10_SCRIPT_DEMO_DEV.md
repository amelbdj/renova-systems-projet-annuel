# 10 — Script de démonstration détaillé (15 min, avec répliques mot à mot)

> **Mode d'emploi :** 🎬 = action à faire à l'écran · 🗣️ = **à dire, texte à réciter**.
> Support : **https://upcycleconnect.pro/** (prod). Onglet local `http://localhost:8088` prêt en secours.
> Avant de commencer : être **déconnecté**, avoir les 4 comptes notés, DevTools fermés, cache désactivé si tu montres du local.

---

## SÉQUENCE 0 — Introduction (0:00 → 1:30)

🎬 *Afficher la page d'accueil `https://upcycleconnect.pro/`.*

🗣️ « Bonjour. Je vais vous présenter **UpcycleConnect**, notre plateforme d'économie circulaire, développée pour **Renova Systems**. »

🗣️ « Le principe est simple : les **particuliers** donnent ou vendent des objets, les **professionnels** récupèrent la matière pour la transformer, et le tout transite par un **réseau de conteneurs connectés**. »

🎬 *Pointer le sélecteur de langue, puis la barre d'adresse.*

🗣️ « L'application est **multilingue**, français et anglais, et surtout elle est **déjà déployée en ligne**, en HTTPS, sur une vraie infrastructure Docker — ce n'est pas du localhost. »

🗣️ « Elle gère **quatre rôles**, chacun avec son tableau de bord : le particulier, le professionnel, le salarié, et l'administrateur. Je vais vous les montrer un par un. »

🎬 *Cliquer sur « Connexion » → saisir le compte particulier → se connecter.*

🗣️ « Commençons par le compte le plus courant : le particulier. »

---

## SÉQUENCE 1 — Dashboard PARTICULIER (1:30 → 4:00)

🎬 *La page `/client` s'affiche.*

🗣️ « Voici l'espace particulier. En haut, trois indicateurs : le nombre d'**annonces**, les **dépôts actifs** en conteneur, et le **score écologique**. Ces chiffres sont **calculés en temps réel** depuis la base, ils ne sont pas codés en dur. »

🎬 *Cliquer sur « ＋ Ajouter une annonce ».*

🗣️ « La fonction principale du particulier, c'est de publier une annonce. Je remplis un titre… je choisis une **catégorie** — et ces catégories viennent directement de la base de données, elles sont gérées par l'admin… »

🎬 *Remplir titre, catégorie, type (Vente/Don), prix, description, et surtout ajouter une photo.*

🗣️ « … je précise si c'est un **don** ou une **vente**, j'ajoute une **photo** de l'objet, et je soumets. »

🎬 *Cliquer « Soumettre l'annonce ».*

🗣️ « Point important : l'annonce n'est **pas publiée immédiatement**. Elle passe en **attente de validation** — c'est l'administrateur qui contrôle. On retrouvera cette annonce tout à l'heure dans le back-office. »

🎬 *Aller sur `/annonces`.*

🗣️ « Voici la marketplace. On peut **filtrer par catégorie**, par type, rechercher, trier. Et remarquez la couleur de l'interface : le particulier est en **violet**. Gardez ça en tête, ça va changer avec le professionnel. »

🎬 *Revenir sur `/client`, pointer la section conteneurs et le score.*

🗣️ « Enfin, quand un objet est vendu, le particulier reçoit un **code PIN** pour le déposer dans un casier, et l'acheteur un **code-barres** pour le récupérer. À chaque objet récupéré, l'utilisateur gagne des points sur son **Upcycling Score**. »

🗣️ « Passons maintenant du côté professionnel. »

---

## SÉQUENCE 2 — Dashboard PROFESSIONNEL (4:00 → 6:30)

🎬 *Se déconnecter, se reconnecter avec le compte pro → `/pro`.*

🗣️ « Voici l'espace professionnel. Première chose : l'interface est maintenant en **teal**, plus en violet. L'application **s'adapte au rôle** de l'utilisateur, y compris sur la page des annonces. »

🎬 *Ouvrir `/annonces` pour montrer les couleurs pro, puis revenir sur `/pro`.*

🗣️ « Le pro consulte les mêmes annonces, mais dans son thème. S'il achète, le paiement passe par **Stripe**, et la plateforme prélève une **commission de 5 %** — techniquement via Stripe Connect, l'argent est reversé au vendeur moins la commission. »

🎬 *Pointer l'encart abonnement.*

🗣️ « On a un modèle **Freemium / Premium** : gratuitement, le pro accède aux annonces de base ; en Premium, il débloque des tableaux de bord avancés, des alertes prioritaires et une meilleure visibilité. »

🎬 *Montrer un projet avec photos avant/après.*

🗣️ « Le pro peut aussi documenter ses **projets d'upcycling**, avec des photos **avant / après** et une estimation du **CO₂ évité** — c'est le cœur de la valeur écologique. Et il peut **sponsoriser** une annonce pour la mettre en avant. »

🗣️ « Ces annonces, ces conteneurs, ces événements… tout ça est animé par l'équipe interne. C'est le rôle du salarié. »

---

## SÉQUENCE 3 — Dashboard SALARIÉ (6:30 → 9:00)

🎬 *Se reconnecter avec le compte salarié → `/salarie`.*

🗣️ « Voici l'espace salarié, l'équipe interne d'UpcycleConnect. Il dispose de plusieurs sous-espaces : les **événements**, un **planning**, les **contenus et articles**, et un **forum**. »

🎬 *Aller sur `/salarie/evenements` → ouvrir le formulaire de création.*

🗣️ « Sa fonction clé, c'est de créer des **événements et des formations**. Je renseigne un titre, une date, un lieu, un tarif, une **image**, et je peux joindre un **plan en PDF** et des **ressources pédagogiques**. »

🎬 *Soumettre (ou expliquer sans soumettre).*

🗣️ « Comme pour les annonces, l'événement part **en attente de validation** de l'admin. Et il y a une règle métier importante : pour proposer un événement **payant**, le salarié doit d'abord avoir un **compte Stripe** — cette vérification est faite **côté serveur**, on ne fait jamais confiance au navigateur. »

🎬 *Pointer la cloche de notifications.*

🗣️ « Le salarié est prévenu par une **notification** dès que l'admin valide ou refuse son contenu. On y arrive justement : le back-office administrateur. »

---

## SÉQUENCE 4 — Dashboard ADMINISTRATEUR (9:00 → 13:00) — LE POINT FORT

🎬 *Se reconnecter avec le compte admin → `/admin`.*

🗣️ « Voici le cœur de la plateforme : le back-office. La **Vue d'ensemble** affiche des indicateurs — utilisateurs, revenus du mois, conteneurs — une **activité récente** dynamique, et un **badge** qui compte les éléments à valider. »

🎬 *Cliquer sur « Validations » (`/admin/validations`).*

🗣️ « Et voici l'annonce que j'ai créée tout à l'heure en tant que particulier. Je l'**approuve**… »

🎬 *Cliquer « Approuver » sur l'annonce.*

🗣️ « … et elle **disparaît immédiatement** de la liste des éléments en attente. La mise à jour est en direct, sans recharger la page. L'annonce est maintenant publiée sur la marketplace. »

🎬 *Aller sur `/admin/users`.*

🗣️ « La gestion des utilisateurs : la liste est **paginée**, dix par page, avec une **recherche**. Je peux **valider**, **refuser** ou **bannir** un compte. Par exemple, un professionnel qui s'inscrit passe en attente, et c'est ici que je vérifie son dossier avant de l'activer. »

🎬 *Aller sur `/admin/conteneurs`, ouvrir un conteneur.*

🗣️ « La logistique : je gère les **conteneurs** et l'état de chaque **casier** — libre, occupé, en maintenance. Je peux en créer, en ajouter, et générer un **rapport logistique en PDF**. »

🎬 *Cliquer « 📄 Rapport logistique » → le PDF se télécharge.*

🗣️ « Et voilà, le PDF est généré et téléchargé. »

🎬 *Aller sur `/admin/finances`.*

🗣️ « Les **finances** : le volume d'affaires, les **revenus issus de la commission de 5 %**, et le détail des transactions. »

🎬 *Revenir sur `/admin` et pointer les cartes Catégories / Langues / Notification.*

🗣️ « Et depuis la Vue d'ensemble, l'admin gère aussi les **catégories** — j'en ajoute une, elle apparaît aussitôt dans le filtre des annonces — l'ajout d'une **nouvelle langue** par import d'un fichier de traduction, et l'**envoi de notifications** à une audience. »

🗣️ « Un point technique essentiel : **toutes** ces actions sont protégées **côté serveur** par un contrôle de rôle. Un particulier, même en manipulant le code du navigateur, ne pourra **jamais** accéder à ce back-office. »

---

## SÉQUENCE 5 — Preuve technique & conclusion (13:00 → 15:00)

🎬 *Basculer sur un terminal, taper `docker ps`.*

🗣️ « Côté technique : l'application tourne en **trois conteneurs Docker** — la base MySQL, l'API en Go, et le serveur web Nginx. Le front et le back sont **séparés** et orchestrés par Docker Compose. »

🎬 *Montrer un appel API (Postman ou curl) : login puis une route protégée.*

🗣️ « L'API est écrite en **Go**. Je me connecte, je récupère un **jeton JWT**, et je l'utilise pour appeler une route protégée. Sans jeton valide, l'API répond une erreur 401 : la sécurité est bien **côté serveur**. »

🎬 *Revenir au navigateur, pointer le cadenas HTTPS et l'URL.*

🗣️ « Et pour finir, la preuve que tout ceci est **réellement en ligne** : le site est accessible sur `upcycleconnect.pro`, en **HTTPS**, derrière Nginx, sur une IP publique. »

🗣️ « **Pour résumer** : quatre rôles, quatre tableaux de bord, un parcours complet du dépôt à la récupération, un paiement sécurisé par Stripe, une API Go protégée par JWT, le tout déployé en production avec Docker, Nginx et HTTPS. Je vous remercie, je suis prêt à répondre à vos questions. »

---

## Antisèche minutage
| Séquence | Fin visée |
|---|---|
| Intro + connexion | 1:30 |
| Particulier | 4:00 |
| Professionnel | 6:30 |
| Salarié | 9:00 |
| Admin | 13:00 |
| Technique + conclusion | 15:00 |

## À ÉVITER
- Dérouler un **paiement Stripe complet** en live (le décrire suffit).
- Tester le **push OneSignal en local** (HTTPS requis).
- Montrer des données de test brutes (comptes « Rejeté », annonces `azerty…`).
- Rester bloqué sur un bug : passer à la **capture de secours** et continuer.

## Plan B
- Si la prod tombe : basculer sur `http://localhost:8088` (même parcours).
- Si tout tombe : dérouler avec les **captures d'écran** de secours.
