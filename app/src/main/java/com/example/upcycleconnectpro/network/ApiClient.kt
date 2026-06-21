package com.example.upcycleconnectpro.network

import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory

// Fichier de configuration centralisé : toutes les URLs et clés au même endroit.
object ApiClient {

    // --- CONFIGURATION ---

    // Adresse du serveur Go. 10.0.2.2 pointe vers le localhost de ton ordinateur (WAMP) depuis l'émulateur.
    const val BASE_URL = "http://10.0.2.2:8081"

    // URL de base pour afficher les images uploadées (annonces)
    const val UPLOADS_URL = "$BASE_URL/view-uploads/"

    // Clé PUBLIQUE Stripe (pk_test_). La clé secrète reste UNIQUEMENT sur le serveur Go.
    const val STRIPE_PUBLISHABLE_KEY =
        "pk_test_51TNFHBHbaxF1KOTt7VoQhPe58MOQGG1sjM9jdcOT6BdplxZWON3mziYOhakbYJvRDTOPlYIOEWJlKOqiXNaCGMqR00M9dAiIvC"

    // --- RETROFIT ---

    val apiService: ApiService by lazy {
        Retrofit.Builder()
            .baseUrl(BASE_URL)
            .addConverterFactory(GsonConverterFactory.create())
            .build()
            .create(ApiService::class.java)
    }
}
