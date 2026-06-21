package com.example.upcycleconnectpro

import android.os.Bundle
import android.view.View
import android.widget.TextView
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.example.upcycleconnectpro.network.ApiClient
import kotlinx.coroutines.launch

class AccueilFragment : Fragment(R.layout.fragment_accueil) {

    // Référence gardée pour rafraîchir les cœurs au retour sur l'écran
    private var rvAnnonces: RecyclerView? = null

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        val tvScore = view.findViewById<TextView>(R.id.tv_score)
        val tvObjetsSauves = view.findViewById<TextView>(R.id.tv_objets_sauves)
        val tvDechetsEvites = view.findViewById<TextView>(R.id.tv_dechets_evites)
        val rvAnnonces = view.findViewById<RecyclerView>(R.id.rv_accueil_annonces)
        this.rvAnnonces = rvAnnonces

        rvAnnonces.layoutManager = LinearLayoutManager(context)

        val sessionManager = SessionManager(requireContext())
        val userId = sessionManager.getUserId()

        if (userId != -1) {
            viewLifecycleOwner.lifecycleScope.launch {
                try {
                    // 1. Charger les statistiques écologiques
                    val stats = ApiClient.apiService.getEcoStats(userId)
                    tvScore.text = String.format("%.0f", stats.score)
                    tvObjetsSauves.text = stats.objets_donnes.toString()
                    tvDechetsEvites.text = String.format("%.1f kg", stats.dechets_evites)

                    // 2. Charger les dernières annonces avec la méthode existante
                    val annonces = ApiClient.apiService.getAllAnnonces()

                    if (annonces.isNotEmpty()) {
                        // On prend les 10 dernières annonces
                        val dernieresAnnonces = annonces.takeLast(10).reversed()

                        // 3. Configuration de l'adaptateur avec la logique de clic du catalogue
                        rvAnnonces.adapter = AnnonceAdapter(dernieresAnnonces) { clickedAnnonce ->
                            val detailFragment = DetailFragment()

                            val bundle = Bundle()
                            bundle.putInt("ANNONCE_ID", clickedAnnonce.id)
                            detailFragment.arguments = bundle

                            parentFragmentManager.beginTransaction()
                                .replace(R.id.fragment_container, detailFragment)
                                .addToBackStack(null)
                                .commit()
                        }
                    }

                } catch (e: Exception) {
                    Toast.makeText(context, "Erreur réseau : ${e.message}", Toast.LENGTH_SHORT).show()
                }
            }
        }
    }

    // Au retour sur l'écran, on re-affiche les cartes pour mettre à jour l'état des cœurs
    override fun onResume() {
        super.onResume()
        rvAnnonces?.adapter?.notifyDataSetChanged()
    }

    override fun onDestroyView() {
        super.onDestroyView()
        rvAnnonces = null
    }
}