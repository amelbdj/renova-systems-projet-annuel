package com.example.upcycleconnectclient.models

data class Annonce(
    val id: Int,
    val titre: String?,
    val description: String?,
    val type: String?,
    val prix: Double,
    val ville: String?,
    val etat: String?,
    val categorie: String?,
    val image: String?,
    val prenom: String?,
    val nom: String?,
    val statut_vente: String?,
    val is_sponsored: Boolean,
    val plan_abo: String?
)
