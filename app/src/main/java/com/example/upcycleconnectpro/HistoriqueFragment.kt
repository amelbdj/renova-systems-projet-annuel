package com.example.upcycleconnectpro

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
import com.example.upcycleconnectpro.network.Reservation
import kotlinx.coroutines.launch

class HistoriqueFragment : Fragment(R.layout.fragment_historique) {

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        // 1. Récupération des vues
        val btnBack = view.findViewById<ImageButton>(R.id.btn_back_historique)
        val rvDepots = view.findViewById<RecyclerView>(R.id.rv_depots)
        val rvRetraits = view.findViewById<RecyclerView>(R.id.rv_retraits)
        val tvEmptyDepots = view.findViewById<TextView>(R.id.tv_empty_depots)
        val tvEmptyRetraits = view.findViewById<TextView>(R.id.tv_empty_retraits)

        rvDepots.layoutManager = LinearLayoutManager(context)
        rvRetraits.layoutManager = LinearLayoutManager(context)

        btnBack.setOnClickListener {
            parentFragmentManager.popBackStack()
        }

        // 2. Utilisateur connecté
        val realUserId = SessionManager(requireContext()).getUserId()
        if (realUserId == -1) {
            Toast.makeText(context, "Erreur : utilisateur non connecté.", Toast.LENGTH_LONG).show()
            return
        }

        // 3. Chargement des deux listes
        viewLifecycleOwner.lifecycleScope.launch {
            // --- Liste "À déposer" (réservations en tant que vendeur) ---
            try {
                val depots = ApiClient.apiService.getUserReservations(userId = realUserId)
                afficherListe(rvDepots, tvEmptyDepots, depots, PinRetraitFragment.MODE_DEPOT)
            } catch (e: Exception) {
                tvEmptyDepots.visibility = View.VISIBLE
                tvEmptyDepots.text = "Erreur de chargement des dépôts."
            }

            // --- Liste "À retirer" (lots déposés en tant qu'acheteur) ---
            try {
                val retraits = ApiClient.apiService.getUserPickups(userId = realUserId)
                afficherListe(rvRetraits, tvEmptyRetraits, retraits, PinRetraitFragment.MODE_RETRAIT)
            } catch (e: Exception) {
                tvEmptyRetraits.visibility = View.VISIBLE
                tvEmptyRetraits.text = "Erreur de chargement des retraits."
            }
        }
    }

    // Remplit un RecyclerView et gère l'état vide. Au clic, ouvre l'écran PIN avec le bon mode.
    private fun afficherListe(
        recyclerView: RecyclerView,
        emptyView: TextView,
        liste: List<Reservation>,
        mode: String
    ) {
        if (liste.isEmpty()) {
            emptyView.visibility = View.VISIBLE
            recyclerView.visibility = View.GONE
            return
        }

        emptyView.visibility = View.GONE
        recyclerView.visibility = View.VISIBLE
        recyclerView.adapter = HistoriqueAdapter(liste) { reservation ->
            val pinFragment = PinRetraitFragment()
            val bundle = Bundle()
            bundle.putSerializable("RESERVATION", reservation)
            bundle.putString("MODE", mode)
            pinFragment.arguments = bundle

            parentFragmentManager.beginTransaction()
                .replace(R.id.fragment_container, pinFragment)
                .addToBackStack(null)
                .commit()
        }
    }
}
