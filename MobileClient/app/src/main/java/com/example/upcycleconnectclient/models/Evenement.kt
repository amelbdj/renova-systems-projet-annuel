package com.example.upcycleconnectclient.models

data class Evenement(
    val id: Int,
    val titre: String?,
    val description: String?,
    val date_debut: String?,
    val date_fin: String?,
    val nb_places: Int,
    val statut_validation: String?,
    val format: String?,
    val lieu: String?,
    val type: String?,
    val prix: Double,
    val nomSalarie: String?,
    val prenomSalarie: String?,
    val image_url: String?
)
