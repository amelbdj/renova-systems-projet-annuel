package com.example.upcycleconnectpro.network

import com.google.gson.annotations.SerializedName
import java.io.Serializable // 1. AJOUTER CET IMPORT

data class Article(
    @SerializedName("id") val id: Int = 0,
    @SerializedName("titre") val titre: String = "",
    @SerializedName("contenu") val contenu: String = "",
    @SerializedName("categorie") val categorie: String? = null,
    @SerializedName("statut") val statut: String? = null
) : Serializable // 2. AJOUTER CECI À LA FIN