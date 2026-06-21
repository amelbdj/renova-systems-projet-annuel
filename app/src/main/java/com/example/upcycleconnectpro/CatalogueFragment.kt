package com.example.upcycleconnectpro

import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.util.Log
import android.view.View
import android.widget.EditText
import android.widget.ImageButton
import android.widget.TextView
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import androidx.recyclerview.widget.GridLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.example.upcycleconnectpro.network.ApiClient
import com.example.upcycleconnectpro.network.Annonce
import kotlinx.coroutines.launch

class CatalogueFragment : Fragment(R.layout.fragment_catalogue) {

    private var recyclerView: RecyclerView? = null
    private var tvEmpty: TextView? = null

    // On garde la liste complète en mémoire pour pouvoir filtrer sans rappeler le serveur
    private var toutesAnnonces: List<Annonce> = emptyList()

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        val recyclerView = view.findViewById<RecyclerView>(R.id.rv_catalogue)
        this.recyclerView = recyclerView
        tvEmpty = view.findViewById(R.id.tv_empty_catalogue)
        val etRecherche = view.findViewById<EditText>(R.id.et_recherche)
        val btnEffacer = view.findViewById<ImageButton>(R.id.btn_effacer_recherche)

        // Affichage en grille 2 colonnes
        recyclerView.layoutManager = GridLayoutManager(context, 2)

        // Filtrage à chaque caractère tapé
        etRecherche.addTextChangedListener(object : TextWatcher {
            override fun afterTextChanged(s: Editable?) {
                val texte = s.toString()
                filtrer(texte)

                // On affiche le bouton "effacer" seulement quand il y a du texte
                if (texte.isEmpty()) {
                    btnEffacer.visibility = View.GONE
                } else {
                    btnEffacer.visibility = View.VISIBLE
                }
            }
            override fun beforeTextChanged(s: CharSequence?, start: Int, count: Int, after: Int) {}
            override fun onTextChanged(s: CharSequence?, start: Int, before: Int, count: Int) {}
        })

        // Bouton "effacer" : on vide le champ de recherche
        btnEffacer.setOnClickListener {
            etRecherche.setText("")
        }

        // Appel réseau pour récupérer les annonces
        viewLifecycleOwner.lifecycleScope.launch {
            try {
                toutesAnnonces = ApiClient.apiService.getAllAnnonces()
                afficher(toutesAnnonces)
            } catch (e: Exception) {
                Log.e("CatalogueFragment", "Erreur réseau", e)
                Toast.makeText(requireContext(), "Erreur de chargement du catalogue", Toast.LENGTH_SHORT).show()
            }
        }
    }

    // Garde uniquement les annonces dont le titre contient le texte recherché
    private fun filtrer(recherche: String) {
        val texte = recherche.trim().lowercase()
        val resultat = if (texte.isEmpty()) {
            toutesAnnonces
        } else {
            toutesAnnonces.filter { annonce ->
                // On cherche dans le titre OU dans la description
                annonce.titre.lowercase().contains(texte) ||
                        annonce.description.lowercase().contains(texte)
            }
        }
        afficher(resultat)
    }

    // Affiche une liste d'annonces (et le message "aucun résultat" si vide)
    private fun afficher(liste: List<Annonce>) {
        val rv = recyclerView ?: return

        if (liste.isEmpty()) {
            tvEmpty?.visibility = View.VISIBLE
        } else {
            tvEmpty?.visibility = View.GONE
        }

        rv.adapter = AnnonceAdapter(liste) { clickedAnnonce ->
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

    override fun onResume() {
        super.onResume()
        recyclerView?.adapter?.notifyDataSetChanged()
    }

    override fun onDestroyView() {
        super.onDestroyView()
        recyclerView = null
        tvEmpty = null
    }
}
