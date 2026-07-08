# Documentation messagerie

## Objectif

La messagerie permet a deux utilisateurs de discuter autour d'une annonce.

Exemple :

- un client regarde une annonce ;
- il clique sur "Contacter le vendeur" ;
- il envoie un message ;
- le vendeur retrouve la conversation dans son dashboard, dans la partie "Messages".

## Fonctionnement simple

La messagerie utilise principalement la base de donnees.

Quand un utilisateur envoie un message :

1. Le front recupere l'annonce concernee.
2. Le front recupere l'utilisateur connecte avec `localStorage.getItem("userId")`.
3. Le front recupere le vendeur de l'annonce avec `id_user`.
4. Le front envoie le message au backend.
5. Le backend sauvegarde le message dans la table `message`.
6. Quand on ouvre la section "Messages", le front recharge les conversations depuis la BDD.

## Fichiers importants

### Frontend

Fichier principal :

```txt
Frontend/script/messagerie/messagerie.js
```

Ce fichier gere :

- l'ouverture de la fenetre de chat ;
- le chargement de l'historique ;
- l'envoi d'un message ;
- l'affichage des conversations.

Pages qui utilisent la messagerie :

```txt
Frontend/oneAnnonce.html
Frontend/espClient.html
Frontend/espPro.html
Frontend/article.html
Frontend/evenement.html
Frontend/forum.html
```

### Backend

Routes :

```txt
API/route/chat.go
```

Handlers :

```txt
API/admin/chat.go
```

Requetes SQL :

```txt
API/bdd/chatReq.go
```

Modele :

```txt
API/models/message.go
```

## Routes API

### Envoyer un message

```txt
POST /api/chat/send
```

Cette route sauvegarde un message dans la BDD.

Corps JSON attendu :

```json
{
  "annonce_id": 1,
  "expediteur_id": 18,
  "destinataire_id": 1,
  "contenu": "Bonjour, votre annonce est-elle disponible ?"
}
```

### Recuperer l'historique d'une conversation

```txt
GET /api/chat/history?annonce_id=1&user1=18&user2=1
```

Cette route recupere tous les messages entre deux utilisateurs pour une annonce precise.

### Recuperer les conversations d'un utilisateur

```txt
GET /api/chat/conversations?userId=18
```

Cette route permet d'afficher la liste des conversations dans le dashboard.

## Table SQL utilisee

La table principale est :

```txt
message
```

Colonnes importantes :

```txt
id
annonce_id
expediteur_id
destinataire_id
contenu
lu
date_envoi
```

Chaque ligne correspond a un message envoye.

## Pourquoi on n'utilise plus le WebSocket pour envoyer

Au debut, la messagerie utilisait le WebSocket :

```txt
wss://upcycleconnect.pro/api/ws/chat
```

Mais en production, le WebSocket peut etre bloque par la configuration du proxy ou du serveur.

Probleme vu en console :

```txt
WebSocket connection failed
```

Donc pour rendre la messagerie plus fiable, l'envoi passe maintenant par une route HTTP classique :

```txt
POST /api/chat/send
```

C'est plus simple et plus stable pour la demo.

Le WebSocket n'est plus indispensable pour que les messages fonctionnent.

## Parcours utilisateur

### Depuis une annonce

1. L'utilisateur va sur une annonce.
2. Il clique sur "Contacter le vendeur".
3. Une fenetre de discussion s'ouvre.
4. Il ecrit son message.
5. Le message est envoye au backend.
6. Le message est sauvegarde dans la table `message`.

### Depuis le dashboard

1. L'utilisateur ouvre son dashboard.
2. Il va dans la partie "Messages".
3. Le front appelle :

```txt
GET /api/chat/conversations?userId=...
```

4. Les conversations sont affichees.
5. Quand il clique sur une conversation, l'historique est charge.

## Verification rapide

Dans le navigateur, ouvrir DevTools puis l'onglet Network.

Quand on envoie un message, on doit voir une requete :

```txt
/api/chat/send
```

ou en prod avec le proxy :

```txt
/api/api/chat/send
```

Le statut doit etre :

```txt
200
```

Ensuite, dans phpMyAdmin, verifier la table :

```txt
message
```

Le dernier message doit apparaitre avec :

- l'id de l'annonce ;
- l'id de l'expediteur ;
- l'id du destinataire ;
- le contenu du message.

## Points a dire en soutenance

La messagerie est liee aux annonces. On ne discute pas dans le vide : chaque conversation concerne une annonce precise.

Les messages sont stockes en base, donc ils restent visibles meme si l'utilisateur ferme la page.

Pour eviter les problemes en production, l'envoi est fait avec une route HTTP classique. C'est plus simple, plus stable, et plus facile a expliquer.

## Limites actuelles

La messagerie est fonctionnelle, mais elle reste simple.

Limites possibles :

- pas encore de compteur de messages non lus ;
- pas encore de notification temps reel fiable en production ;
- l'ID utilisateur est encore envoye depuis le front ;
- pas de piece jointe dans les messages.

Pour une version plus avancee, on pourrait :

- utiliser l'ID utilisateur depuis le token JWT ;
- ajouter un compteur de messages non lus ;
- securiser davantage les conversations ;
- remettre le WebSocket quand le proxy prod sera bien configure.

