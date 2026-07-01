package com.example.upcycleconnectclient

import android.os.Bundle
import android.widget.ArrayAdapter
import android.widget.Button
import android.widget.EditText
import android.widget.Spinner
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.lifecycleScope
import com.example.upcycleconnectclient.network.ApiClient
import kotlinx.coroutines.launch
import okhttp3.MediaType.Companion.toMediaTypeOrNull
import okhttp3.RequestBody
import okhttp3.RequestBody.Companion.toRequestBody

class DeposerAnnonceActivity : AppCompatActivity() {

    private val types = listOf("Don gratuit", "Vente")
    private val categories = listOf("Mobilier", "Outillage", "Textile", "Électronique", "Déco", "Livres")

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_deposer_annonce)

        val session = SessionManager(this)

        val titreInput = findViewById<EditText>(R.id.titreInput)
        val typeSpinner = findViewById<Spinner>(R.id.typeSpinner)
        val categorieSpinner = findViewById<Spinner>(R.id.categorieSpinner)
        val prixInput = findViewById<EditText>(R.id.prixInput)
        val descriptionInput = findViewById<EditText>(R.id.descriptionInput)
        val publierButton = findViewById<Button>(R.id.publierButton)
        val statusText = findViewById<TextView>(R.id.statusText)

        typeSpinner.adapter = ArrayAdapter(this, android.R.layout.simple_spinner_dropdown_item, types)
        categorieSpinner.adapter = ArrayAdapter(this, android.R.layout.simple_spinner_dropdown_item, categories)

        publierButton.setOnClickListener {
            val titre = titreInput.text.toString().trim()
            val description = descriptionInput.text.toString().trim()

            if (titre.isEmpty()) {
                statusText.text = "Le titre est obligatoire."
                return@setOnClickListener
            }

            val type = types[typeSpinner.selectedItemPosition]
            val categorieId = (categorieSpinner.selectedItemPosition + 1).toString()

            var prix = prixInput.text.toString()
            if (type == "Don gratuit" || prix.isEmpty()) {
                prix = "0"
            }

            statusText.text = "Publication en cours..."

            lifecycleScope.launch {
                try {
                    val fields = mapOf(
                        "titre" to texte(titre),
                        "type" to texte(type),
                        "id_categorie" to texte(categorieId),
                        "prix" to texte(prix),
                        "description" to texte(description),
                        "id_user" to texte(session.getUserId().toString()),
                        "code_postal" to texte("75000"),
                        "ville" to texte("Paris"),
                        "etat" to texte("Bon état"),
                        "poids_kg" to texte("1.0"),
                        "quantite" to texte("1")
                    )

                    val response = ApiClient.apiService.creerAnnonce(
                        "Bearer ${session.getToken()}",
                        fields
                    )

                    if (response.isSuccessful) {
                        statusText.text = "Annonce publiée ! En attente de validation."
                        titreInput.text.clear()
                        prixInput.text.clear()
                        descriptionInput.text.clear()
                    } else {
                        val err = response.errorBody()?.string() ?: ""
                        if (err.contains("STRIPE")) {
                            statusText.text = "Configurez vos paiements sur le site web d'abord."
                        } else {
                            statusText.text = "Erreur lors de la publication."
                        }
                    }
                } catch (e: Exception) {
                    statusText.text = "Erreur de connexion au serveur."
                }
            }
        }
    }

    private fun texte(valeur: String): RequestBody {
        return valeur.toRequestBody("text/plain".toMediaTypeOrNull())
    }
}
