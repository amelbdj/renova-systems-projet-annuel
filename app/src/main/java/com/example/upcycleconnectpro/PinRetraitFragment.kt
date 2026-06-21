package com.example.upcycleconnectpro

import android.os.Bundle
import android.view.View
import android.widget.ImageButton
import android.widget.TextView
import android.widget.Toast
import androidx.core.content.ContextCompat
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import com.example.upcycleconnectpro.network.ApiClient
import com.example.upcycleconnectpro.network.DepositRequest
import com.example.upcycleconnectpro.network.Reservation
import com.google.android.material.button.MaterialButton
import kotlinx.coroutines.launch

class PinRetraitFragment : Fragment(R.layout.fragment_pin_retrait) {

    companion object {
        const val MODE_DEPOT = "DEPOT"
        const val MODE_RETRAIT = "RETRAIT"
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        // 1. Récupération de la réservation + du mode (dépôt ou retrait)
        val reservation = arguments?.getSerializable("RESERVATION") as? Reservation
        val mode = arguments?.getString("MODE") ?: MODE_RETRAIT

        if (reservation == null) {
            Toast.makeText(requireContext(), "Erreur : opération introuvable", Toast.LENGTH_SHORT).show()
            parentFragmentManager.popBackStack()
            return
        }

        // 2. Récupération des vues
        val btnBack = view.findViewById<ImageButton>(R.id.btn_back_pin)
        val tvTitreEcran = view.findViewById<TextView>(R.id.tv_pin_titre_ecran)
        val tvObjet = view.findViewById<TextView>(R.id.tv_pin_objet)
        val tvEtat = view.findViewById<TextView>(R.id.tv_pin_etat)
        val tvLieu = view.findViewById<TextView>(R.id.tv_pin_lieu)
        val tvBox = view.findViewById<TextView>(R.id.tv_pin_box)
        val tvCodeLabel = view.findViewById<TextView>(R.id.tv_pin_code_label)
        val tvCode = view.findViewById<TextView>(R.id.tv_pin_code)
        val tvDate = view.findViewById<TextView>(R.id.tv_pin_date)
        val btnValider = view.findViewById<MaterialButton>(R.id.btn_valider_retrait)

        // 3. Adaptation des textes selon le mode
        val estDepot = mode == MODE_DEPOT
        if (estDepot) {
            tvTitreEcran.text = "Dépôt du matériau"
            tvCodeLabel.text = "CODE DE DÉPÔT"
            btnValider.text = "Valider le dépôt"
        } else {
            tvTitreEcran.text = "Retrait du matériau"
            tvCodeLabel.text = "CODE DE RETRAIT"
            btnValider.text = "Valider le retrait"
        }

        // 4. Remplissage des données
        tvObjet.text = reservation.objet
        tvLieu.text = reservation.lieu
        tvBox.text = reservation.numero_box
        tvCode.text = reservation.code_pin
        tvDate.text = reservation.date
        tvEtat.text = reservation.etat.uppercase()

        // 5. Bouton retour
        btnBack.setOnClickListener {
            parentFragmentManager.popBackStack()
        }

        // 6. Bouton de validation (dépôt ou retrait)
        btnValider.setOnClickListener {
            btnValider.isEnabled = false
            btnValider.text = "Validation en cours..."

            viewLifecycleOwner.lifecycleScope.launch {
                try {
                    val response = if (estDepot) {
                        ApiClient.apiService.validerDepot(DepositRequest(reservation.code_pin))
                    } else {
                        val token = SessionManager(requireContext()).getToken()
                        if (token.isNullOrEmpty()) {
                            Toast.makeText(requireContext(), "Erreur : non connecté", Toast.LENGTH_SHORT).show()
                            btnValider.isEnabled = true
                            btnValider.text = "Valider le retrait"
                            return@launch
                        }
                        ApiClient.apiService.validerRetrait("Bearer $token", reservation.code_pin)
                    }

                    if (response.isSuccessful) {
                        val message = if (estDepot) "Dépôt validé avec succès !" else "Retrait validé avec succès !"
                        Toast.makeText(requireContext(), message, Toast.LENGTH_LONG).show()

                        tvEtat.text = if (estDepot) "DÉPOSÉ" else "RETIRÉ"
                        tvEtat.setTextColor(
                            ContextCompat.getColor(requireContext(), android.R.color.holo_green_light)
                        )
                        btnValider.text = if (estDepot) "Dépôt effectué" else "Retrait effectué"
                        btnValider.alpha = 0.5f
                    } else {
                        Toast.makeText(
                            requireContext(),
                            ErrorMessages.fromServer(response.errorBody()?.string()),
                            Toast.LENGTH_LONG
                        ).show()
                        btnValider.isEnabled = true
                        btnValider.text = if (estDepot) "Valider le dépôt" else "Valider le retrait"
                    }

                } catch (e: Exception) {
                    Toast.makeText(requireContext(), ErrorMessages.fromNetwork(), Toast.LENGTH_LONG).show()
                    btnValider.isEnabled = true
                    btnValider.text = if (estDepot) "Valider le dépôt" else "Valider le retrait"
                }
            }
        }
    }
}
