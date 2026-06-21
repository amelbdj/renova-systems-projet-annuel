package com.example.upcycleconnectpro

import android.content.Intent
import android.os.Bundle
import android.view.View
import android.widget.ImageButton
import android.widget.TextView
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.example.upcycleconnectpro.network.ApiClient
import com.example.upcycleconnectpro.network.Invoice
import kotlinx.coroutines.launch

class FacturationFragment : Fragment(R.layout.fragment_facturation) {

    // Nom du pro, utilisé pour personnaliser la facture PDF
    private var nomClient: String = "Client professionnel"

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        val btnBack = view.findViewById<ImageButton>(R.id.btn_back_facturation)
        val rvFacturation = view.findViewById<RecyclerView>(R.id.rv_facturation)
        val tvEmpty = view.findViewById<TextView>(R.id.tv_empty_facturation)

        rvFacturation.layoutManager = LinearLayoutManager(context)

        btnBack.setOnClickListener {
            parentFragmentManager.popBackStack()
        }

        val sessionManager = SessionManager(requireContext())
        val token = sessionManager.getToken()
        val userId = sessionManager.getUserId()

        if (token == null || userId == -1) {
            Toast.makeText(context, "Erreur : non connecté", Toast.LENGTH_SHORT).show()
            return
        }

        viewLifecycleOwner.lifecycleScope.launch {
            // 1. Nom du client (pour la facture)
            try {
                val user = ApiClient.apiService.getUserById(userId, "Bearer $token")
                nomClient = "${user.Prenom} ${user.Nom}"
            } catch (_: Exception) {
                // Pas bloquant : on garde le nom par défaut
            }

            // 2. Liste des factures
            try {
                val factures = ApiClient.apiService.getProInvoices("Bearer $token")

                if (factures.isEmpty()) {
                    tvEmpty.visibility = View.VISIBLE
                    rvFacturation.visibility = View.GONE
                    return@launch
                }

                tvEmpty.visibility = View.GONE
                rvFacturation.visibility = View.VISIBLE
                rvFacturation.adapter = FacturationAdapter(factures) { facture ->
                    telechargerEtOuvrir(facture)
                }

            } catch (e: Exception) {
                Toast.makeText(context, ErrorMessages.fromNetwork(), Toast.LENGTH_LONG).show()
            }
        }
    }

    // Génère le PDF en local puis propose de l'ouvrir
    private fun telechargerEtOuvrir(facture: Invoice) {
        try {
            val fichier = InvoicePdfGenerator.generate(requireContext(), facture, nomClient)
            val uri = InvoicePdfGenerator.getShareableUri(requireContext(), fichier)

            val intent = Intent(Intent.ACTION_VIEW).apply {
                setDataAndType(uri, "application/pdf")
                addFlags(Intent.FLAG_GRANT_READ_URI_PERMISSION)
            }

            try {
                startActivity(intent)
            } catch (e: Exception) {
                // Aucune app PDF installée : on confirme au moins que le fichier est enregistré
                Toast.makeText(
                    requireContext(),
                    "Facture enregistrée : ${fichier.name}\n(Aucune application PDF trouvée pour l'ouvrir)",
                    Toast.LENGTH_LONG
                ).show()
            }

        } catch (e: Exception) {
            Toast.makeText(requireContext(), "Erreur lors de la génération de la facture", Toast.LENGTH_SHORT).show()
        }
    }
}
