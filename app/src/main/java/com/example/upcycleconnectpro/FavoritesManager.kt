package com.example.upcycleconnectpro

import android.content.Context

// Gestion des favoris en local (sur le téléphone) via SharedPreferences.
// On stocke simplement la liste des IDs d'annonces mises en favori.
class FavoritesManager(context: Context) {

    private val prefs = context.getSharedPreferences("Favoris", Context.MODE_PRIVATE)
    private val key = "favoris_ids"

    // Retourne l'ensemble des IDs favoris
    fun getFavorites(): Set<Int> {
        val stringSet = prefs.getStringSet(key, emptySet()) ?: emptySet()
        return stringSet.mapNotNull { it.toIntOrNull() }.toSet()
    }

    fun isFavorite(annonceId: Int): Boolean {
        return getFavorites().contains(annonceId)
    }

    // Ajoute ou retire le favori, et renvoie le nouvel état (true = désormais favori)
    fun toggle(annonceId: Int): Boolean {
        val current = getFavorites().toMutableSet()
        val nowFavorite: Boolean
        if (current.contains(annonceId)) {
            current.remove(annonceId)
            nowFavorite = false
        } else {
            current.add(annonceId)
            nowFavorite = true
        }
        // On reconvertit en Set<String> pour le stockage
        prefs.edit().putStringSet(key, current.map { it.toString() }.toSet()).apply()
        return nowFavorite
    }
}
