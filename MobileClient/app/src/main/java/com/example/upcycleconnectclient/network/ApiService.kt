package com.example.upcycleconnectclient.network

import com.example.upcycleconnectclient.models.Annonce
import com.example.upcycleconnectclient.models.EcoStats
import com.example.upcycleconnectclient.models.Evenement
import com.example.upcycleconnectclient.models.LoginRequest
import com.example.upcycleconnectclient.models.LoginResponse
import okhttp3.RequestBody
import okhttp3.ResponseBody
import retrofit2.Response
import retrofit2.http.Body
import retrofit2.http.GET
import retrofit2.http.Header
import retrofit2.http.Multipart
import retrofit2.http.POST
import retrofit2.http.PartMap
import retrofit2.http.Query

interface ApiService {

    @POST("admin/login")
    suspend fun login(@Body body: LoginRequest): Response<LoginResponse>

    @GET("api/annonces/all")
    suspend fun getAnnonces(@Query("id") userId: Int): Response<List<Annonce>>

    @Multipart
    @POST("admin/annonces/add")
    suspend fun creerAnnonce(
        @Header("Authorization") token: String,
        @PartMap fields: Map<String, @JvmSuppressWildcards RequestBody>
    ): Response<ResponseBody>

    @GET("admin/evenements")
    suspend fun getEvenements(@Header("Authorization") token: String): Response<List<Evenement>>

    @GET("api/user/stats")
    suspend fun getEcoStats(@Query("user_id") userId: Int): Response<EcoStats>
}
