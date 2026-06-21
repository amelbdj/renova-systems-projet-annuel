package com.example.upcycleconnectpro.network

import com.google.gson.annotations.SerializedName

data class LoginRequest(
    @SerializedName("email")
    val email: String,

    @SerializedName("mot_de_passe")
    val motDePasse: String
)


    data class LoginResponse(
val token: String? = null,
val id: Int? = null,          // C'est lui qu'on veut !
val role: String? = null,
val prenom: String? = null,
val message: String? = null
)


// On crée ce petit modèle au cas où le Go renvoie "user": {"id": 1}
data class UserInfo(
    val id: Int
)