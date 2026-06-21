package com.example.upcycleconnectpro

import android.os.Bundle
import android.view.View
import android.widget.ImageButton
import android.widget.TextView
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import androidx.recyclerview.widget.GridLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.example.upcycleconnectpro.network.ApiClient
import kotlinx.coroutines.launch

class FavorisFragment : Fragment(R.layout.fragment_favoris) {

    private var rvFavoris: RecyclerView? = null
    private var tvEmpty: TextView? = null

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        val btnBack = view.findViewById<ImageButton>(R.id.btn_back_favoris)
        rvFavoris = view.findViewById(R.id.rv_favoris)
        tvEmpty = view.findViewById(R.id.tv_empty_favoris)

        rvFavoris?.layoutManager = GridLayoutManager(context, 2)

        btnBack.setOnClickListener {
            parentFragmentManager.popBackStack()
        }
        // Le chargement réel se fait dans onResume (pour rafraîchir au retour sur l'écran)
    }

    // Rechargé à chaque fois qu'on revient sur l'écran : un favori retiré disparaît aussitôt
    override fun onResume() {
        super.onResume()
        chargerFavoris()
    }

    private fun chargerFavoris() {
        val recyclerView = rvFavoris ?: return
        val emptyView = tvEmpty ?: return

        val favorisIds = FavoritesManager(requireContext()).getFavorites()

        // Si aucun favori, on affiche le message et on s'arrête là
        if (favorisIds.isEmpty()) {
            emptyView.visibility = View.VISIBLE
            recyclerView.visibility = View.GONE
            recyclerView.adapter = null
            return
        }

        // On charge toutes les annonces puis on garde uniquement les favorites
        viewLifecycleOwner.lifecycleScope.launch {
            try {
                val toutesAnnonces = ApiClient.apiService.getAllAnnonces()
                val annoncesFavorites = toutesAnnonces.filter { favorisIds.contains(it.id) }

                if (annoncesFavorites.isEmpty()) {
                    emptyView.visibility = View.VISIBLE
                    recyclerView.visibility = View.GONE
                    recyclerView.adapter = null
                    return@launch
                }

                emptyView.visibility = View.GONE
                recyclerView.visibility = View.VISIBLE
                recyclerView.adapter = AnnonceAdapter(annoncesFavorites) { clickedAnnonce ->
                    val detailFragment = DetailFragment()
                    val bundle = Bundle()
                    bundle.putInt("ANNONCE_ID", clickedAnnonce.id)
                    detailFragment.arguments = bundle

                    parentFragmentManager.beginTransaction()
                        .replace(R.id.fragment_container, detailFragment)
                        .addToBackStack(null)
                        .commit()
                }

            } catch (e: Exception) {
                Toast.makeText(context, "Erreur de chargement des favoris", Toast.LENGTH_SHORT).show()
            }
        }
    }

    override fun onDestroyView() {
        super.onDestroyView()
        rvFavoris = null
        tvEmpty = null
    }
}
