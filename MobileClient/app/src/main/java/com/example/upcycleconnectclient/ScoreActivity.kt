package com.example.upcycleconnectclient

import android.os.Bundle
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.lifecycleScope
import com.example.upcycleconnectclient.network.ApiClient
import kotlinx.coroutines.launch

class ScoreActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_score)

        val session = SessionManager(this)
        val scoreValue = findViewById<TextView>(R.id.scoreValue)
        val objetsValue = findViewById<TextView>(R.id.objetsValue)
        val dechetsValue = findViewById<TextView>(R.id.dechetsValue)
        val statusText = findViewById<TextView>(R.id.statusText)

        lifecycleScope.launch {
            try {
                val response = ApiClient.apiService.getEcoStats(session.getUserId())
                val stats = response.body()

                if (stats != null) {
                    scoreValue.text = stats.score.toString()
                    objetsValue.text = stats.objets_donnes.toString()
                    dechetsValue.text = "${stats.dechets_evites} kg"
                } else {
                    statusText.text = "Impossible de charger le score."
                }
            } catch (e: Exception) {
                statusText.text = "Erreur de chargement du score."
            }
        }
    }
}
