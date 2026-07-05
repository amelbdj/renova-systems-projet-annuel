# 07 - Questions post presentation

1. Pourquoi avoir choisi Go pour l'API ?  
Reponse : Go est simple a compiler, rapide, et adapte pour exposer des routes HTTP claires. Dans notre projet, les routes sont dans `API/route`, les handlers dans `API/admin`, et les requetes SQL dans `API/bdd`.

2. Pourquoi separer front et back ?  
Reponse : Le front gere l'affichage et l'experience utilisateur, tandis que l'API centralise la logique metier, la securite, les roles et la base de donnees.

3. Ou est la logique metier ?  
Reponse : Principalement dans `API/admin` pour les handlers et `API/bdd` pour les operations SQL, par exemple validation annonce, box, score, paiement.

4. Comment sont geres les roles ?  
Reponse : Le role est stocke dans le JWT et dans le localStorage cote front. L'API verifie les roles avec `VerifyRoleMiddleware`.

5. Pourquoi verifier les roles cote API ?  
Reponse : Le front peut etre modifie par l'utilisateur. La vraie protection doit donc etre cote serveur.

6. Comment fonctionne le workflow annonce ?  
Reponse : Une annonce est creee avec un statut, puis l'admin ou salarie peut la valider ou la refuser via les routes admin.

7. Comment eviter qu'une annonce non validee soit visible ?  
Reponse : Le catalogue utilise les annonces validees, notamment via les requetes de `GetValidatedAnnonces`.

8. Comment fonctionne l'Upcycling Score ?  
Reponse : Le score est calcule lors de certaines actions logistiques dans `boxReq.go`, puis ajoute a l'utilisateur selon l'objet et la recuperation.

9. Comment sont calculees les commissions ?  
Reponse : Les commandes et paiements sont stockes en base. Les vues finances utilisent `order.go` pour calculer volume et revenu.

10. Comment fonctionne la base de donnees ?  
Reponse : MySQL contient les tables utilisateurs, annonces, evenements, box, paiements, documents et traductions. Le dump est dans `db/init.sql`.

11. Comment importer la base ?  
Reponse : En Docker, `db/init.sql` est monte dans `/docker-entrypoint-initdb.d/init.sql` au premier demarrage du conteneur MySQL.

12. Comment exporter la base ?  
Reponse : Avec phpMyAdmin ou `mysqldump`, puis on versionne un dump propre si necessaire.

13. Comment lancer le projet en local Docker ?  
Reponse : `docker compose up -d --build`, puis ouvrir le port du conteneur frontend visible dans `docker compose ps`.

14. Comment fonctionne Docker Compose ?  
Reponse : Il lance trois services : MySQL, backend Go et frontend Nginx, connectes sur le reseau `uc_net`.

15. Comment fonctionne Nginx ?  
Reponse : Nginx sert les fichiers front et redirige `/api` vers le backend avec `proxy_pass http://backend:8081`.

16. Comment prouver que ce n'est pas seulement du localhost WAMP ?  
Reponse : Montrer `docker compose ps`, ouvrir l'URL Docker, puis montrer que les appels passent par `/api`.

17. Comment ajouter une route API ?  
Reponse : Ajouter le handler dans `API/admin`, la route dans `API/route`, et la fonction SQL dans `API/bdd` si besoin.

18. Comment proteger une route ?  
Reponse : Entourer le handler avec `auth.VerifyTokenMiddleware` ou `auth.VerifyRoleMiddleware`.

19. Comment ajouter un champ a une annonce ?  
Reponse : Modifier la table SQL, le modele Go, la requete `CreateAnnonce/UpdateAnnonce`, puis le formulaire front.

20. Que faire si la base ne repond pas ?  
Reponse : Verifier `docker compose ps`, les logs MySQL, les variables `DB_HOST`, `DB_USER`, `DB_PASS`.

21. Que faire si l'API renvoie 500 ?  
Reponse : Lire les logs backend, identifier la route appelee, verifier la requete SQL et les donnees envoyees.

22. Que se passe-t-il si un non-admin appelle une route admin ?  
Reponse : Les routes sensibles utilisent `VerifyRoleMiddleware`, donc l'API doit retourner une erreur d'acces.

23. Pourquoi certains endpoints ne sont pas proteges ?  
Reponse : Certains endpoints publics servent le catalogue ou les traductions. Les actions sensibles doivent etre protegees.

24. Comment fonctionne Stripe ?  
Reponse : Le backend cree des sessions checkout ou payment intent avec les cles Stripe. Les webhooks peuvent mettre a jour les statuts.

25. Pourquoi ne pas faire le paiement en live ?  
Reponse : Stripe depend d'un service externe. En soutenance courte, on montre le code et les ecrans, sauf si l'environnement test est valide juste avant.

26. Comment sont generees les factures PDF ?  
Reponse : `API/admin/pdf.go` genere les factures ou contrats et enregistre le chemin dans la table `document`.

27. Comment fonctionne le multilingue ?  
Reponse : Les textes sont stockes dans `translations`, servis par `/api/translations`, puis appliques dans le front.

28. Comment ajouter une langue ?  
Reponse : Ajouter une ligne dans `languages`, ajouter les traductions correspondantes et utiliser l'interface admin de traduction.

29. Comment fonctionnent les box ?  
Reponse : Les box appartiennent a des conteneurs. Les reservations et retraits sont traces dans `historique_conteneurs`.

30. Ou sont stockes les codes PIN ?  
Reponse : Dans l'historique des conteneurs, notamment les champs de code lies au depot ou a la recuperation.

31. Pourquoi un simulateur hardware ?  
Reponse : Pour representer le comportement d'une borne sans avoir de materiel physique en soutenance.

32. Comment fonctionnent les evenements ?  
Reponse : Les salaries peuvent creer evenements/formations, les utilisateurs peuvent les consulter et s'inscrire.

33. Comment sont gerees les ressources PDF de formation ?  
Reponse : Le front salarie envoie les fichiers, le backend les stocke et les relie a l'evenement.

34. Comment fonctionne le forum ?  
Reponse : Les utilisateurs creent des sujets/messages, et les salaries peuvent moderer les messages.

35. Comment fonctionne le bannissement ?  
Reponse : L'API met a jour l'utilisateur en base, cree une notification et peut envoyer un mail de bannissement.

36. Comment fonctionne le mot de passe oublie ?  
Reponse : L'utilisateur demande un reset, l'API cree un token, envoie un lien par mail, puis `reset-password.html` renvoie le nouveau mot de passe.

37. Est-ce que le SMTP est obligatoire ?  
Reponse : Oui pour recevoir le mail en vrai. Sans SMTP, le parcours peut etre montre mais l'envoi reel ne doit pas etre promis.

38. Quelles fonctionnalites sont les plus stables ?  
Reponse : Accueil, login, roles, annonces, validation admin, catalogue, SQL, Docker/Nginx, multilingue.

39. Quelles fonctionnalites sont plus risquees ?  
Reponse : Stripe, SMTP, OneSignal, PDF si les dossiers ne sont pas montes, et les codes PIN si la donnee n'est pas preparee.

40. Qu'est-ce qui manque ?  
Reponse : Swagger/OpenAPI n'est pas trouve. Le QR code complet n'est pas clairement implemente. Certaines integrations externes demandent une configuration.

41. Comment ameliorer le projet apres soutenance ?  
Reponse : Ajouter documentation OpenAPI, tests automatises, seeds de comptes demo, journal d'audit admin, et stabiliser les integrations externes.

42. Pourquoi utiliser des donnees pre-remplies ?  
Reponse : Pour garantir une demonstration fiable dans un temps limite et montrer le workflow sans attendre la creation de chaque etape.
