package com.example.upcycleconnectpro.network

import okhttp3.MultipartBody
import okhttp3.RequestBody
import okhttp3.ResponseBody
import retrofit2.Response
import retrofit2.http.Body
import retrofit2.http.GET
import retrofit2.http.Header
import retrofit2.http.Multipart
import retrofit2.http.POST
import retrofit2.http.PUT
import retrofit2.http.Part
import retrofit2.http.Path
import retrofit2.http.Query

// La structure de réponse pour la réservation
data class ReserveResponse(
    val status: String?,
    val message: String?
)

data class Reservation(
    val lieu: String,
    val numero_box: String,
    val code_pin: String,
    val barcode: String,
    val objet: String,
    val date: String,
    val etat: String
) : java.io.Serializable

data class ValiderRetraitResponse(
    val message: String?,
    val status: String?
)

// Corps de requête pour valider un dépôt (route POST /api/box/deposit)
data class DepositRequest(
    val pin_code: String
)

// Corps de requête pour changer le mot de passe (route POST /api/user/update-password)
data class UpdatePasswordRequest(
    val id: Int,
    val old_password: String,
    val new_password: String
)

// Corps de requête pour modifier les données personnelles (route PUT /api/user/profile)
data class UpdateProfileRequest(
    val nom: String,
    val prenom: String,
    val email: String
)

data class EcoStatsResponse(
    val score: Double,
    val objets_donnes: Int,
    val dechets_evites: Double
)

data class PaymentIntentResponse(
    val client_secret: String
)

// L'interface unique contenant toutes tes routes
interface ApiService {

    @GET("/admin/articles")
    suspend fun getArticles(
        @Header("Authorization") token: String
    ): List<Article>

    @GET("/api/user/ecostats")
    suspend fun getEcoStats(@Query("user_id") userId: Int): EcoStatsResponse

    @POST("/admin/login")
    suspend fun login(@Body request: LoginRequest): LoginResponse

    @GET("/api/annonces/all")
    suspend fun getAllAnnonces(): List<Annonce>

    @GET("/admin/users/{id}")
    suspend fun getUserById(
        @Path("id") userId: Int,
        @Header("Authorization") token: String
    ): User

    @GET("/api/annonces")
    suspend fun getAnnonceById(@Query("id") id: Int): Annonce

    // La route de réservation
    @POST("/api/annonces/vendre")
    suspend fun reserverAnnonce(
        @Header("Authorization") token: String,
        @Query("id") annonceId: Int,
        @Query("buyer_id") buyerId: Int
    ): Response<ReserveResponse>

    @GET("/api/user/boxes")
    suspend fun getUserReservations(
        @Query("user_id") userId: Int
    ): List<Reservation>

    // Lots déposés en attente d'être retirés par le pro (acheteur)
    @GET("/api/user/pickups/{id}")
    suspend fun getUserPickups(
        @Path("id") userId: Int
    ): List<Reservation>

    // La route Stripe
    @POST("/api/mobile/payment-intent")
    suspend fun getPaymentIntent(@Query("annonce_id") annonceId: Int): PaymentIntentResponse

    // Validation du retrait via code PIN
    @POST("/api/boxes/valider-retrait")
    suspend fun validerRetrait(
        @Header("Authorization") token: String,
        @Query("code_pin") codePin: String
    ): Response<ValiderRetraitResponse>

    // Validation du dépôt via code PIN
    @POST("/api/box/deposit")
    suspend fun validerDepot(
        @Body request: DepositRequest
    ): Response<ValiderRetraitResponse>

    // Changement de mot de passe
    @POST("/api/user/update-password")
    suspend fun updatePassword(
        @Header("Authorization") token: String,
        @Body request: UpdatePasswordRequest
    ): Response<ValiderRetraitResponse>

    // Modification des données personnelles
    @PUT("/api/user/profile")
    suspend fun updateProfile(
        @Header("Authorization") token: String,
        @Body request: UpdateProfileRequest
    ): Response<ValiderRetraitResponse>

    // Historique & Facturation du professionnel
    @GET("/api/pro/invoices")
    suspend fun getProInvoices(
        @Header("Authorization") token: String
    ): List<Invoice>

    // Création d'une annonce (envoi multipart : champs texte + image)
    @Multipart
    @POST("/admin/annonces/add")
    suspend fun createAnnonce(
        @Header("Authorization") token: String,
        @Part("titre") titre: RequestBody,
        @Part("description") description: RequestBody,
        @Part("type") type: RequestBody,
        @Part("prix") prix: RequestBody,
        @Part("ville") ville: RequestBody,
        @Part("code_postal") codePostal: RequestBody,
        @Part("etat") etat: RequestBody,
        @Part("poids_kg") poidsKg: RequestBody,
        @Part("quantite") quantite: RequestBody,
        @Part("id_categorie") idCategorie: RequestBody,
        @Part("id_user") idUser: RequestBody,
        @Part image: MultipartBody.Part?
    ): Response<ResponseBody>
}