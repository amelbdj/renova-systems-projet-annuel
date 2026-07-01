package com.example.upcycleconnectclient.models

data class LoginRequest(
    val email: String,
    val mot_de_passe: String
)

data class LoginResponse(
    val token: String?,
    val role: String?,
    val id: Int?,
    val prenom: String?,
    val score: Int?,
    val validation: String?
)
