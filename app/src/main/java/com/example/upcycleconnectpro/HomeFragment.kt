package com.example.upcycleconnectpro

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.example.upcycleconnectpro.network.ApiClient
import kotlinx.coroutines.launch

class HomeFragment : Fragment() {

    private lateinit var recyclerView: RecyclerView
    private lateinit var articleAdapter: ArticleAdapter

    // 1. On "gonfle" (charge) le design XML de la page des articles
    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        return inflater.inflate(R.layout.fragment_home, container, false)
    }

    // 2. Une fois la vue chargée, on la configure
    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        recyclerView = view.findViewById(R.id.recycler_view_articles)
        recyclerView.layoutManager = LinearLayoutManager(requireContext())

        // --- C'EST ICI QUE SE TROUVE LA NOUVELLE LOGIQUE DE CLIC ---
        // On initialise l'Adapter et on lui dit quoi faire au clic
        articleAdapter = ArticleAdapter(emptyList()) { clickedArticle ->

            // On prépare la page de destination (assure-toi que ce nom correspond à ton fichier !)
            val detailFragment = DetailArticleFragment()

            // On glisse l'article COMPLET dans le colis
            val bundle = Bundle()
            bundle.putSerializable("ARTICLE_COMPLET", clickedArticle)
            detailFragment.arguments = bundle

            // On fait la transition d'écran
            parentFragmentManager.beginTransaction()
                .replace(R.id.fragment_container, detailFragment)
                .addToBackStack(null) // Permet à la flèche retour de fonctionner
                .commit()
        }

        recyclerView.adapter = articleAdapter

        // On lance la requête vers ton serveur Go
        fetchArticles()
    }

    // 3. La fonction qui va chercher les données en fond (Coroutine)
    private fun fetchArticles() {
        viewLifecycleOwner.lifecycleScope.launch {
            try {
                // Appel API
                val response = ApiClient.apiService.getArticles("")

                // Le filtre des articles validés
                val articlesValides = response.filter { article ->
                    val statut = article.statut?.lowercase() ?: ""
                    statut == "valide" || statut == "publié" || statut == "publie" || statut == "en ligne"
                }

                // On envoie les données filtrées à l'Adapter
                articleAdapter.updateData(articlesValides)

            } catch (e: Exception) {
                Toast.makeText(requireContext(), "Erreur : ${e.localizedMessage}", Toast.LENGTH_LONG).show()
            }
        }
    }
}