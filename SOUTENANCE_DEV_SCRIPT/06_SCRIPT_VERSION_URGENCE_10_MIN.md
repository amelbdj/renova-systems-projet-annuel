# 06 - Script version urgence 10 minutes

## 00:00 - 01:00 - Faty

Ecran : accueil.  
Dire : "UpcycleConnect est une plateforme metier pour connecter particuliers, professionnels et equipe interne autour de l'upcycling."  
Montrer : menu, promesse, acces.

## 01:00 - 02:00 - Faty

Ecran : login.  
Dire : "Le site gere plusieurs roles et redirige chacun vers son espace."  
Montrer : connexion utilisateur.

## 02:00 - 03:30 - Faty

Ecran : dashboard particulier.  
Dire : "Le particulier peut creer et suivre ses annonces. Une annonce n'est pas publiee directement : elle passe par validation."  
Montrer : formulaire annonce ou annonce existante.

## 03:30 - 05:00 - Amel

Ecran : admin validations.  
Dire : "Le back-office controle les annonces, evenements et contenus avant publication."  
Montrer : annonce en attente, boutons valider/refuser.

## 05:00 - 06:00 - Amel

Ecran : code API.  
Dire : "La logique est portee par une API Go. Les routes sont separees des handlers et des requetes SQL."  
Montrer : `route/annonces.go`, `admin/annonce.go`, `bdd/annonceReq.go`.

## 06:00 - 07:00 - Faty

Ecran : catalogue annonces.  
Dire : "Une fois validee, l'annonce est visible dans le catalogue."  
Montrer : liste + detail.

## 07:00 - 08:00 - Ndoya

Ecran : dashboard pro.  
Dire : "Le professionnel dispose d'un espace dedie pour ses annonces, ses projets d'upcycling et les options premium."  
Montrer : `espPro.html`.

## 08:00 - 09:00 - Ndoya

Ecran : Docker/Nginx.  
Dire : "Le projet se deploie avec Docker Compose : MySQL, backend Go et frontend Nginx avec proxy API."  
Montrer : `docker compose ps`, `nginx.conf`.

## 09:00 - 10:00 - Amel

Ecran : SQL.  
Dire : "La base est pre-remplie avec utilisateurs, annonces, evenements, box et traductions pour garantir une demo stable."  
Conclusion : "Nous avons livre une plateforme multi-roles fonctionnelle, avec workflows metier, API Go, base SQL et deploiement Docker."
