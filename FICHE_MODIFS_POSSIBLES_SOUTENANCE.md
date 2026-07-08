# Fiche modifs possibles en soutenance

Ce fichier liste des petites modifications qu'un professeur peut demander pendant la soutenance.

Le but est d'avoir :

- la demande possible ;
- l'endroit ou modifier ;
- une reponse simple a donner ;
- un exemple de code facile a expliquer.

## 1. Connexion avec la touche Entree

### Demande possible

"Quand je suis sur la page de connexion, je veux appuyer sur Entree au lieu de cliquer sur le bouton."

### Reponse a donner

"Oui, on peut ajouter un ecouteur d'evenement sur les champs email et mot de passe. Si la touche appuyee est Entree, on appelle la fonction `submitLogin()`."

### Fichier a modifier

```txt
Frontend/script/login.js
```

### Code a ajouter en bas du fichier

```js
document.addEventListener("DOMContentLoaded", function () {
  var email = document.getElementById("loginEmail");
  var password = document.getElementById("loginPwd");

  function connexionAvecEntree(event) {
    if (event.key === "Enter") {
      submitLogin();
    }
  }

  if (email) {
    email.addEventListener("keydown", connexionAvecEntree);
  }

  if (password) {
    password.addEventListener("keydown", connexionAvecEntree);
  }
});
```

## 2. Afficher une erreur dans la page au lieu d'un alert

### Demande possible

"Au lieu d'avoir une alerte navigateur, est-ce qu'on peut afficher l'erreur directement dans le formulaire ?"

### Reponse a donner

"Oui, on ajoute une petite zone de message dans le HTML, puis le JavaScript remplit cette zone en cas d'erreur."

### Fichier HTML a modifier

```txt
Frontend/login.html
```

### Code HTML a ajouter sous le bouton connexion

```html
<div id="loginMsg" style="margin-top: 10px; color: #ff5a5a; font-size: 13px"></div>
```

### Fichier JS a modifier

```txt
Frontend/script/login.js
```

### Exemple dans `submitLogin()`

Remplacer :

```js
alert("Email ou mot de passe incorrect.");
```

par :

```js
var msg = document.getElementById("loginMsg");
if (msg) {
  msg.textContent = "Email ou mot de passe incorrect.";
}
```

## 3. Desactiver le bouton pendant la connexion

### Demande possible

"Si l'utilisateur clique plusieurs fois sur connexion, est-ce qu'on peut eviter plusieurs appels API ?"

### Reponse a donner

"Oui, on desactive le bouton pendant le chargement, puis on le reactive a la fin."

### Fichier a modifier

```txt
Frontend/login.html
```

### Modifier le bouton connexion

Ajouter un `id` :

```html
<button
  id="btnLogin"
  class="btn-primary"
  onclick="submitLogin()"
  data-i18n="login.form.submit"
>
  Se connecter →
</button>
```

### Fichier JS

```txt
Frontend/script/login.js
```

### Code dans `submitLogin()`

Au debut de la fonction :

```js
var btn = document.getElementById("btnLogin");
if (btn) {
  btn.disabled = true;
  btn.textContent = "Connexion...";
}
```

A la fin, dans un bloc `finally` :

```js
finally {
  if (btn) {
    btn.disabled = false;
    btn.textContent = "Se connecter →";
  }
}
```

## 4. Nettoyer la fonction en double dans la messagerie

### Demande possible

"Pourquoi la fonction `appendMessageToUI` existe deux fois ?"

### Reponse a donner

"C'est un doublon. Le navigateur garde la derniere fonction declaree, donc ca marche, mais on peut nettoyer en gardant seulement la deuxieme version."

### Fichier a modifier

```txt
Frontend/script/messagerie/messagerie.js
```

### Action a faire

Supprimer la premiere version de :

```js
function appendMessageToUI(texte, isMe) {
  ...
}
```

Garder la deuxieme version, celle qui utilise :

```js
messageRow.className = `message-row ${isMe ? "me" : "them"}`;
```

## 5. Recharger les conversations apres envoi d'un message

### Demande possible

"Quand j'envoie un message, est-ce que la liste des conversations peut se mettre a jour automatiquement ?"

### Reponse a donner

"Oui, apres l'envoi du message, on peut appeler `loadMyMessages()` si la fonction existe."

### Fichier a modifier

```txt
Frontend/script/messagerie/messagerie.js
```

### Code a ajouter dans `sendChatMessage()`

Juste apres :

```js
input.value = "";
```

Ajouter :

```js
if (typeof loadMyMessages === "function") {
  loadMyMessages();
}
```

## 6. Trier les conversations par dernier message

### Demande possible

"Les conversations pourraient-elles etre triees par la plus recente ?"

### Reponse a donner

"Oui, il faut modifier la requete SQL pour recuperer la date du dernier message et trier dessus."

### Fichier a modifier

```txt
API/bdd/chatReq.go
```

### Fonction concernee

```go
func GetUserConversations(userID int) ([]map[string]interface{}, error)
```

### Idee de modification SQL

Ajouter :

```sql
MAX(m.date_envoi) AS dernier_message
```

Puis :

```sql
ORDER BY dernier_message DESC
```

Exemple simplifie :

```go
query := `
    SELECT
        CASE WHEN expediteur_id = ? THEN destinataire_id ELSE expediteur_id END as contact_id,
        u.nom,
        u.prenom,
        a.titre as annonce_titre,
        a.id as annonce_id,
        MAX(m.date_envoi) as dernier_message
    FROM pa2026.message m
    JOIN pa2026.utilisateur u ON u.id = (CASE WHEN m.expediteur_id = ? THEN m.destinataire_id ELSE m.expediteur_id END)
    JOIN pa2026.annonce a ON a.id = m.annonce_id
    WHERE m.expediteur_id = ? OR m.destinataire_id = ?
    GROUP BY contact_id, u.nom, u.prenom, a.titre, a.id
    ORDER BY dernier_message DESC
`
```

## 7. Utiliser l'ID du token au lieu de l'ID envoye par le front

### Demande possible

"Pourquoi le front envoie `expediteur_id` ? Un utilisateur pourrait modifier l'id."

### Reponse a donner

"C'est vrai. Pour securiser, le backend doit recuperer l'utilisateur connecte depuis le token JWT, et ne pas faire confiance a l'id envoye par le front."

### Fichier a modifier

```txt
API/admin/chat.go
```

### Fonction concernee

```go
func SendMessageHandler(w http.ResponseWriter, r *http.Request)
```

### Code a ajouter apres le decode JSON

```go
userID, ok := r.Context().Value("userID").(int)
if !ok || userID == 0 {
    http.Error(w, "Utilisateur non connecte", http.StatusUnauthorized)
    return
}

msg.ExpediteurID = userID
```

### Explication

Le front peut envoyer un mauvais `expediteur_id`, mais le backend l'ecrase avec l'id du token.

## 8. Ajouter une date dans les messages

### Demande possible

"Est-ce qu'on peut afficher l'heure d'envoi des messages ?"

### Reponse a donner

"Oui, l'API renvoie deja `date_envoi`. Il faut l'afficher dans la bulle de message."

### Fichier a modifier

```txt
Frontend/script/messagerie/messagerie.js
```

### Modifier l'appel

Dans `openChat()`, remplacer :

```js
appendMessageToUI(msg.contenu, isMe);
```

par :

```js
appendMessageToUI(msg.contenu, isMe, msg.date_envoi);
```

### Modifier la fonction

```js
function appendMessageToUI(texte, isMe, dateEnvoi) {
  var heure = "";

  if (dateEnvoi) {
    var date = new Date(dateEnvoi);
    heure = date.toLocaleTimeString("fr-FR", {
      hour: "2-digit",
      minute: "2-digit",
    });
  }

  ...
}
```

Puis ajouter dans la bulle :

```js
bubble.innerHTML = texte + "<br><small>" + heure + "</small>";
```

## 9. Ajouter un compteur de messages non lus

### Demande possible

"Est-ce qu'on peut voir combien de messages ne sont pas lus ?"

### Reponse a donner

"Oui, la table `message` a deja une colonne `lu`. Il faut compter les messages ou `destinataire_id` est l'utilisateur connecte et `lu = 0`."

### Fichier backend possible

```txt
API/bdd/chatReq.go
```

### Requete SQL possible

```sql
SELECT COUNT(*)
FROM message
WHERE destinataire_id = ?
AND lu = 0
```

### Route possible

```txt
GET /api/chat/unread?userId=...
```

## 10. Marquer les messages comme lus

### Demande possible

"Quand j'ouvre une conversation, est-ce que les messages peuvent passer en lus ?"

### Reponse a donner

"Oui, quand on ouvre l'historique, le backend peut mettre `lu = 1` pour les messages recus par l'utilisateur."

### Fichier a modifier

```txt
API/bdd/chatReq.go
```

### Requete possible

```go
_, err = Db.Exec(`
    UPDATE pa2026.message
    SET lu = 1
    WHERE annonce_id = ?
    AND expediteur_id = ?
    AND destinataire_id = ?
`, annonceID, user2, user1)
```

## 11. Verifier qu'un utilisateur ne contacte pas lui-meme

### Demande possible

"Est-ce qu'un utilisateur peut s'envoyer un message a lui-meme ?"

### Reponse a donner

"Normalement le bouton est cache si l'annonce appartient a l'utilisateur connecte. On peut aussi ajouter une verification backend."

### Fichier frontend

```txt
Frontend/script/annonces/oneAnnonce.js
```

### Code deja present

```js
if (!isSold && monUserId !== vendeurId && !isNaN(vendeurId)) {
  ...
}
```

### Verification backend possible

Dans :

```txt
API/admin/chat.go
```

Ajouter :

```go
if msg.ExpediteurID == msg.DestinataireID {
    http.Error(w, "Impossible de s'envoyer un message a soi-meme", http.StatusBadRequest)
    return
}
```

## 12. Ajouter un bouton retour dans une modale

### Demande possible

"Quand j'ouvre une fenetre, je veux pouvoir revenir facilement."

### Reponse a donner

"On peut ajouter un bouton qui appelle simplement la fonction de fermeture de la modale."

### Exemple messagerie

Fichier :

```txt
Frontend/oneAnnonce.html
```

La fonction existe deja :

```js
closeChat()
```

Bouton possible :

```html
<button onclick="closeChat()" type="button">Retour</button>
```

## 13. Ajouter une recherche dans une liste

### Demande possible

"Est-ce qu'on peut rechercher dans les annonces / utilisateurs / conversations ?"

### Reponse a donner

"Oui, on ajoute un input, puis on filtre le tableau cote front avec `includes()`."

### Exemple simple JS

```js
function filtrerListe() {
  var recherche = document.getElementById("search").value.toLowerCase();
  var items = document.querySelectorAll(".conversation-item");

  items.forEach(function (item) {
    if (item.textContent.toLowerCase().includes(recherche)) {
      item.style.display = "block";
    } else {
      item.style.display = "none";
    }
  });
}
```

### HTML

```html
<input id="search" oninput="filtrerListe()" placeholder="Rechercher..." />
```

## 14. Ajouter une confirmation avant suppression

### Demande possible

"Avant de supprimer, peut-on demander confirmation ?"

### Reponse a donner

"Oui, on utilise `confirm()` avant d'appeler l'API de suppression."

### Exemple

```js
if (!confirm("Voulez-vous vraiment supprimer ?")) {
  return;
}
```

## 15. Remplacer les alert par des messages propres

### Demande possible

"Les alertes navigateur ne sont pas tres propres. Peut-on faire mieux ?"

### Reponse a donner

"Oui, on peut utiliser une div de message dans la page."

### Exemple HTML

```html
<div id="messagePage"></div>
```

### Exemple JS

```js
function afficherMessage(texte, couleur) {
  var msg = document.getElementById("messagePage");
  if (!msg) return;

  msg.textContent = texte;
  msg.style.color = couleur;
}
```

Utilisation :

```js
afficherMessage("Operation reussie", "#00c97a");
```

## 16. Questions possibles du professeur

### "Pourquoi avoir retire le WebSocket ?"

Reponse :

"En local, le WebSocket fonctionnait. Mais en production, il etait bloque par la configuration du proxy. Pour garantir une messagerie fiable pendant la demo, l'envoi passe par HTTP avec `/api/chat/send`. C'est plus simple et plus stable."

### "Est-ce que les messages sont sauvegardes ?"

Reponse :

"Oui, ils sont sauvegardes dans la table `message`. Donc meme si on ferme la page, on peut retrouver l'historique."

### "Est-ce que c'est securise ?"

Reponse :

"Les routes sont protegees avec le token. Pour aller plus loin, on peut recuperer l'ID utilisateur depuis le token au lieu de l'envoyer depuis le front."

### "Comment tester rapidement ?"

Reponse :

"J'ouvre une annonce avec un compte qui n'est pas le vendeur, je clique sur contacter, j'envoie un message, puis je verifie dans le dashboard ou dans phpMyAdmin table `message`."

