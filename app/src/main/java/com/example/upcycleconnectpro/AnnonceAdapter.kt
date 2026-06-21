package com.example.upcycleconnectpro

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageButton
import android.widget.ImageView
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import coil.load
import coil.transform.RoundedCornersTransformation
import com.example.upcycleconnectpro.network.Annonce
import com.example.upcycleconnectpro.network.ApiClient

class AnnonceAdapter(
    private val annonces: List<Annonce>,
    private val onItemClick: (Annonce) -> Unit
) : RecyclerView.Adapter<AnnonceAdapter.AnnonceViewHolder>() {

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): AnnonceViewHolder {
        val view = LayoutInflater.from(parent.context)
            .inflate(R.layout.item_annonce, parent, false)
        return AnnonceViewHolder(view)
    }

    override fun onBindViewHolder(holder: AnnonceViewHolder, position: Int) {
        val annonce = annonces[position]
        holder.bind(annonce, onItemClick)
    }

    override fun getItemCount(): Int {
        return annonces.size
    }

    class AnnonceViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView) {
        private val ivImage: ImageView = itemView.findViewById(R.id.iv_annonce_image)
        private val tvTitre: TextView = itemView.findViewById(R.id.tv_annonce_titre)
        private val tvPrix: TextView = itemView.findViewById(R.id.tv_annonce_prix)
        private val btnFavori: ImageButton = itemView.findViewById(R.id.btn_favori)

        fun bind(annonce: Annonce, onItemClick: (Annonce) -> Unit) {
            tvTitre.text = annonce.titre
            tvPrix.text = "${annonce.prix} €"

            // Construction de l'URL avec imageUrl (base centralisée dans ApiClient)
            val finalImageUrl = ApiClient.UPLOADS_URL + (annonce.imageUrl ?: "")

            // Chargement de l'image avec Coil
            ivImage.load(finalImageUrl) {
                crossfade(true)
                placeholder(R.drawable.ic_launcher_background)
                error(R.drawable.ic_launcher_background)
                transformations(RoundedCornersTransformation(16f))
            }

            // --- GESTION DU FAVORI (local) ---
            val favoritesManager = FavoritesManager(itemView.context)
            // On affiche le bon cœur selon l'état actuel
            majIconeFavori(favoritesManager.isFavorite(annonce.id))

            btnFavori.setOnClickListener {
                val nowFavorite = favoritesManager.toggle(annonce.id)
                majIconeFavori(nowFavorite)
            }

            itemView.setOnClickListener {
                onItemClick(annonce)
            }
        }

        private fun majIconeFavori(estFavori: Boolean) {
            if (estFavori) {
                btnFavori.setImageResource(R.drawable.ic_heart)
            } else {
                btnFavori.setImageResource(R.drawable.ic_heart_border)
            }
        }
    }
}
