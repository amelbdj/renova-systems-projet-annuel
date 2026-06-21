package com.example.upcycleconnectpro.network

import com.google.gson.annotations.SerializedName

data class User(
    @SerializedName("id") val id: Int,
    @SerializedName("role") val Role: String,
    @SerializedName("nom") val Nom: String,
    @SerializedName("prenom") val Prenom: String,
    @SerializedName("email") val Email: String,
    @SerializedName("score") val Score: Int,
    @SerializedName("nom_entreprise") val NomEntreprise: String?,
    @SerializedName("siret") val Siret: String?
)