package com.example.upcycleconnectclient

import android.os.Bundle
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.lifecycleScope
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.example.upcycleconnectclient.network.ApiClient
import kotlinx.coroutines.launch

class CatalogueActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_catalogue)

        val session = SessionManager(this)
        val recycler = findViewById<RecyclerView>(R.id.catalogueRecycler)
        val statusText = findViewById<TextView>(R.id.statusText)

        recycler.layoutManager = LinearLayoutManager(this)

        lifecycleScope.launch {
            try {
                val response = ApiClient.apiService.getEvenements("Bearer ${session.getToken()}")
                val tous = response.body() ?: emptyList()

                val evenements = tous.filter {
                    (it.statut_validation ?: "").lowercase().startsWith("valid")
                }

                if (evenements.isEmpty()) {
                    statusText.text = "Aucun événement pour le moment."
                } else {
                    statusText.text = ""
                    recycler.adapter = EvenementAdapter(evenements)
                }
            } catch (e: Exception) {
                statusText.text = "Erreur de chargement des événements."
            }
        }
    }
}
