package com.example.upcycleconnectclient

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import coil.load
import com.example.upcycleconnectclient.models.Annonce
import com.example.upcycleconnectclient.network.ApiClient

class AnnonceAdapter(private val annonces: List<Annonce>) :
    RecyclerView.Adapter<AnnonceAdapter.AnnonceViewHolder>() {

    class AnnonceViewHolder(view: View) : RecyclerView.ViewHolder(view) {
        val image: ImageView = view.findViewById(R.id.annonceImage)
        val titre: TextView = view.findViewById(R.id.annonceTitre)
        val meta: TextView = view.findViewById(R.id.annonceMeta)
        val vendeur: TextView = view.findViewById(R.id.annonceVendeur)
        val prix: TextView = view.findViewById(R.id.annoncePrix)
        val badge: TextView = view.findViewById(R.id.annonceBadge)
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): AnnonceViewHolder {
        val view = LayoutInflater.from(parent.context)
            .inflate(R.layout.item_annonce, parent, false)
        return AnnonceViewHolder(view)
    }

    override fun getItemCount(): Int = annonces.size

    override fun onBindViewHolder(holder: AnnonceViewHolder, position: Int) {
        val a = annonces[position]

        holder.titre.text = a.titre ?: ""
        holder.meta.text = "${a.ville ?: ""} · ${a.categorie ?: ""}"
        holder.vendeur.text = "Par ${a.prenom ?: ""} ${a.nom ?: ""}"

        val gratuit = a.prix <= 0 || a.type == "Don gratuit" || a.type == "don"
        holder.prix.text = if (gratuit) "Gratuit" else "${a.prix} €"

        if (!a.image.isNullOrEmpty()) {
            holder.image.load(ApiClient.SERVER + a.image)
        } else {
            holder.image.setImageDrawable(null)
        }

        when {
            a.is_sponsored -> {
                holder.badge.visibility = View.VISIBLE
                holder.badge.text = "⭐ Sponsorisé"
            }
            a.plan_abo == "plus" || a.plan_abo == "pro" -> {
                holder.badge.visibility = View.VISIBLE
                holder.badge.text = "⚡ Prioritaire"
            }
            else -> holder.badge.visibility = View.GONE
        }
    }
}
