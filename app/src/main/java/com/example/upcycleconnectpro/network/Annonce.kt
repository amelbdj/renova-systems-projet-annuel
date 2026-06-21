package com.example.upcycleconnectpro.network

import com.google.gson.annotations.SerializedName

data class Annonce(
    @SerializedName("id") val id: Int,
    @SerializedName("titre") val titre: String,
    @SerializedName("description") val description: String,
    @SerializedName("prix") val prix: Double,
    @SerializedName("image_url") val imageUrl: String? = null
)