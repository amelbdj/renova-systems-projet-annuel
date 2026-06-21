package com.example.upcycleconnectpro

import android.os.Bundle
import android.util.Log
import android.view.View
import android.widget.Button
import android.widget.ImageButton
import android.widget.ImageView
import android.widget.TextView
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import coil.load
import com.example.upcycleconnectpro.network.ApiClient
import com.stripe.android.PaymentConfiguration
import com.stripe.android.paymentsheet.PaymentSheet
import com.stripe.android.paymentsheet.PaymentSheetResult
import kotlinx.coroutines.launch

class DetailFragment : Fragment(R.layout.fragment_detail) {

    private lateinit var paymentSheet: PaymentSheet
    // Variable pour garder l'ID en mémoire pendant le paiement Stripe
    private var currentAnnonceId: Int = -1

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        // 1. Initialiser Stripe avec ta clé PUBLIQUE (centralisée dans ApiClient)
        PaymentConfiguration.init(
            requireContext(),
            ApiClient.STRIPE_PUBLISHABLE_KEY
        )

        // 2. Préparer la fenêtre Stripe
        paymentSheet = PaymentSheet(this, ::onPaymentSheetResult)

        // 3. Récupération des vues
        val tvTitre = view.findViewById<TextView>(R.id.tv_detail_titre)
        val tvPrix = view.findViewById<TextView>(R.id.tv_detail_prix)
        val tvDesc = view.findViewById<TextView>(R.id.tv_detail_description)
        val ivDetailImage = view.findViewById<ImageView>(R.id.iv_detail_image)
        val btnBack = view.findViewById<ImageButton>(R.id.btn_back)
        val btnReserver = view.findViewById<Button>(R.id.btn_action_reserver)

        btnBack.setOnClickListener {
            parentFragmentManager.popBackStack()
        }

        // 4. Récupération de l'ID de l'annonce
        currentAnnonceId = arguments?.getInt("ANNONCE_ID") ?: -1

        if (currentAnnonceId != -1) {
            // Chargement des détails depuis le serveur
            viewLifecycleOwner.lifecycleScope.launch {
                try {
                    val annonce = ApiClient.apiService.getAnnonceById(currentAnnonceId)

                    tvTitre.text = annonce.titre
                    tvPrix.text = "${annonce.prix} €"
                    tvDesc.text = annonce.description

                    ivDetailImage.load(ApiClient.UPLOADS_URL + (annonce.imageUrl ?: "")) {
                        crossfade(true)
                        placeholder(R.drawable.ic_launcher_background)
                        error(R.drawable.ic_launcher_background)
                    }

                } catch (e: Exception) {
                    Toast.makeText(context, "Erreur lors du chargement : ${e.message}", Toast.LENGTH_SHORT).show()
                    tvDesc.text = "Impossible de charger les détails."
                }
            }
        } else {
            Toast.makeText(context, "Erreur : ID introuvable", Toast.LENGTH_SHORT).show()
        }

        // 5. Action du bouton : On déclenche Stripe au lieu de la réservation directe
        btnReserver.setOnClickListener {
            if (currentAnnonceId == -1) return@setOnClickListener

            btnReserver.text = "Chargement Stripe..."
            btnReserver.isEnabled = false

            viewLifecycleOwner.lifecycleScope.launch {
                try {
                    // On demande le ticket à Go
                    val response = ApiClient.apiService.getPaymentIntent(currentAnnonceId)
                    val clientSecret = response.client_secret

                    // On configure et ouvre Stripe
                    val configuration = PaymentSheet.Configuration("UpcycleConnect Pro")
                    paymentSheet.presentWithPaymentIntent(clientSecret, configuration)

                } catch (e: Exception) {
                    Toast.makeText(context, "Erreur serveur : impossible d'initier le paiement", Toast.LENGTH_SHORT).show()
                    btnReserver.text = "Réserver"
                    btnReserver.isEnabled = true
                }
            }
        }
    }

    // 6. Fin de la transaction : Validation et vraie réservation en base de données
    private fun onPaymentSheetResult(paymentSheetResult: PaymentSheetResult) {
        val btnReserver = view?.findViewById<Button>(R.id.btn_action_reserver)
        btnReserver?.text = "Réserver"
        btnReserver?.isEnabled = true

        when (paymentSheetResult) {
            is PaymentSheetResult.Completed -> {
                Toast.makeText(context, "Paiement validé ! Création de la box...", Toast.LENGTH_SHORT).show()

                val sessionManager = SessionManager(requireContext())
                val token = sessionManager.getToken()
                val realUserId = sessionManager.getUserId()

                if (token != null && realUserId != -1 && currentAnnonceId != -1) {
                    viewLifecycleOwner.lifecycleScope.launch {
                        try {
                            val response = ApiClient.apiService.reserverAnnonce(
                                token = "Bearer $token",
                                annonceId = currentAnnonceId,
                                buyerId = realUserId
                            )

                            if (response.isSuccessful) {
                                Toast.makeText(context, "Réservation confirmée avec succès !", Toast.LENGTH_LONG).show()
                                parentFragmentManager.popBackStack() // Retour au catalogue
                            } else {
                                Toast.makeText(
                                    context,
                                    ErrorMessages.fromServer(response.errorBody()?.string()),
                                    Toast.LENGTH_LONG
                                ).show()
                            }

                        } catch (e: Exception) {
                            Log.e("DetailFragment", "Erreur réseau réservation", e)
                            Toast.makeText(context, ErrorMessages.fromNetwork(), Toast.LENGTH_LONG).show()
                        }
                    }
                }
            }
            is PaymentSheetResult.Canceled -> {
                Toast.makeText(context, "Paiement annulé", Toast.LENGTH_SHORT).show()
            }
            is PaymentSheetResult.Failed -> {
                Toast.makeText(context, "Échec du paiement", Toast.LENGTH_LONG).show()
            }
        }
    }
}