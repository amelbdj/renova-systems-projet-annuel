package com.example.upcycleconnectpro.network

import com.google.gson.annotations.SerializedName

// Une ligne de l'historique de facturation du professionnel
data class Invoice(
    @SerializedName("id_commande") val id: Int,
    @SerializedName("date") val date: String,
    @SerializedName("titre") val titre: String,
    @SerializedName("montant") val montant: Double,
    @SerializedName("commission") val commission: Double
)
