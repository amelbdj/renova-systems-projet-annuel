package com.example.upcycleconnectpro

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import com.google.android.material.card.MaterialCardView

class MenuFragment : Fragment(R.layout.fragment_menu) {

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        // 1. Récupération des cartes (boutons) du menu
        val cardCreerAnnonce = view.findViewById<MaterialCardView>(R.id.card_creer_annonce)
        val cardHistorique = view.findViewById<MaterialCardView>(R.id.card_historique)
        val cardFacturation = view.findViewById<MaterialCardView>(R.id.card_facturation)
        val cardFavoris = view.findViewById<MaterialCardView>(R.id.card_favoris)
        val cardSupport = view.findViewById<MaterialCardView>(R.id.card_support)
        val cardParametres = view.findViewById<MaterialCardView>(R.id.card_parametres)

        // Publier une annonce
        cardCreerAnnonce.setOnClickListener {
            ouvrirFragment(CreerAnnonceFragment())
        }

        // 2. Action pour l'historique : dépôts à faire + retraits à effectuer
        cardHistorique.setOnClickListener {
            ouvrirFragment(HistoriqueFragment())
        }

        // Historique & Facturation (récupérations passées + téléchargement des factures)
        cardFacturation.setOnClickListener {
            ouvrirFragment(FacturationFragment())
        }

        // 3. Mes matériaux favoris
        cardFavoris.setOnClickListener {
            ouvrirFragment(FavorisFragment())
        }

        // 4. Contacter le support
        cardSupport.setOnClickListener {
            ouvrirFragment(SupportFragment())
        }

        // 5. Paramètres du compte
        cardParametres.setOnClickListener {
            ouvrirFragment(ParametresFragment())
        }
    }

    // Petite fonction utilitaire pour éviter de répéter le code de navigation
    private fun ouvrirFragment(fragment: Fragment) {
        parentFragmentManager.beginTransaction()
            .replace(R.id.fragment_container, fragment)
            .addToBackStack(null)
            .commit()
    }
}