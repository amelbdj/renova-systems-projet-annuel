package com.example.upcycleconnectclient

import android.os.Bundle
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.lifecycleScope
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.example.upcycleconnectclient.network.ApiClient
import kotlinx.coroutines.launch

class AnnoncesActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_annonces)

        val session = SessionManager(this)
        val recycler = findViewById<RecyclerView>(R.id.annoncesRecycler)
        val statusText = findViewById<TextView>(R.id.statusText)

        recycler.layoutManager = LinearLayoutManager(this)

        lifecycleScope.launch {
            try {
                val response = ApiClient.apiService.getAnnonces(session.getUserId())
                val annonces = response.body() ?: emptyList()

                if (annonces.isEmpty()) {
                    statusText.text = "Aucune annonce disponible."
                } else {
                    statusText.text = ""
                    recycler.adapter = AnnonceAdapter(annonces)
                }
            } catch (e: Exception) {
                statusText.text = "Erreur de chargement des annonces."
            }
        }
    }
}
