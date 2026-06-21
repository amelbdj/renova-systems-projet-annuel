package com.example.upcycleconnectpro

import android.content.Context
import android.content.SharedPreferences

// 1. Le "private val" ici est indispensable pour que le context soit utilisable dans les fonctions !
class SessionManager(private val context: Context) {

    // 2. On initialise les préférences UNE SEULE FOIS pour toute la classe (le tiroir "AppPrefs")
    private val prefs: SharedPreferences = context.getSharedPreferences("AppPrefs", Context.MODE_PRIVATE)

    // --- GESTION DU TOKEN ---

    fun saveToken(token: String) {
        prefs.edit().putString("auth_token", token).apply()
    }

    fun getToken(): String? {
        return prefs.getString("auth_token", null)
    }

    // --- GESTION DE L'ID UTILISATEUR ---

    fun saveUserId(id: Int) {
        // On utilise "prefs" directement au lieu de recréer un accès
        prefs.edit().putInt("USER_ID", id).apply()
    }

    fun getUserId(): Int {
        return prefs.getInt("USER_ID", -1)
    }

    // --- DÉCONNEXION ---

    fun clearSession() {
        // clear() efface tout d'un coup (le token ET l'ID utilisateur), c'est parfait pour une déconnexion complète
        prefs.edit().clear().apply()
    }
}