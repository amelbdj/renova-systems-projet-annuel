package com.example.upcycleconnectpro

// Centralise la traduction des erreurs (serveur ou réseau) en messages clairs pour l'utilisateur.
// Couvre notamment les cas du descriptif : PIN expiré, Box indisponible, Échec de réservation.
object ErrorMessages {

    // Erreur renvoyée par le serveur (corps de la réponse HTTP)
    fun fromServer(raw: String?): String {
        val msg = (raw ?: "").lowercase()
        return when {
            // PIN expiré / invalide / déjà utilisé
            msg.contains("pin") && (msg.contains("expir") || msg.contains("invalide") ||
                    msg.contains("incorrect")) ->
                "Code PIN invalide ou expiré. Vérifiez le code affiché, ou demandez-en un nouveau."

            // Objet déjà déposé / déjà récupéré
            msg.contains("déjà") || msg.contains("deja") ->
                "Cette opération a déjà été effectuée."

            // Box / casier indisponible
            (msg.contains("casier") || msg.contains("box")) && msg.contains("disponible") ->
                "Aucune box n'est disponible pour le moment. Réessayez un peu plus tard."

            // Échec de réservation
            msg.contains("réserv") || msg.contains("reserv") || msg.contains("acheteur") ->
                "La réservation a échoué. L'objet n'est peut-être plus disponible."

            // Ancien mot de passe incorrect
            msg.contains("ancien mot de passe") ->
                "L'ancien mot de passe est incorrect."

            // Cas vide
            msg.isBlank() ->
                "Une erreur est survenue. Veuillez réessayer."

            // Par défaut : on renvoie le message serveur tel quel
            else -> raw ?: "Une erreur est survenue."
        }
    }

    // Erreur réseau (exception : serveur injoignable, pas de connexion...)
    fun fromNetwork(): String {
        return "Impossible de joindre le serveur. Vérifiez votre connexion internet et réessayez."
    }
}
