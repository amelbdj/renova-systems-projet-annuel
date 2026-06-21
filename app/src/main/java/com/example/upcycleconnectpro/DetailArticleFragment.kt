package com.example.upcycleconnectpro

import android.os.Bundle
import android.view.View
import android.widget.ImageButton
import android.widget.TextView
import androidx.fragment.app.Fragment
import com.example.upcycleconnectpro.network.Article

class DetailArticleFragment : Fragment(R.layout.fragment_detail_article) {

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        // Flèche retour
        view.findViewById<ImageButton>(R.id.btn_back_article).setOnClickListener {
            parentFragmentManager.popBackStack()
        }

        // 1. Récupération des vues depuis ton XML
        val tvTitre = view.findViewById<TextView>(R.id.tv_detail_titre)
        val tvContenu = view.findViewById<TextView>(R.id.tv_detail_contenu)

        // 2. On récupère l'article complet envoyé par le HomeFragment
        // "ARTICLE_COMPLET" doit être exactement le même nom que dans ton HomeFragment
        val article = arguments?.getSerializable("ARTICLE_COMPLET") as? Article

        // 3. Si l'article existe, on affiche ses infos
        if (article != null) {
            tvTitre.text = article.titre
            tvContenu.text = article.contenu
        }
    }
}