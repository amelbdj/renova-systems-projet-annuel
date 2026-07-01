package com.example.upcycleconnectclient

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import com.example.upcycleconnectclient.models.Evenement

class EvenementAdapter(private val evenements: List<Evenement>) :
    RecyclerView.Adapter<EvenementAdapter.EvenementViewHolder>() {

    class EvenementViewHolder(view: View) : RecyclerView.ViewHolder(view) {
        val titre: TextView = view.findViewById(R.id.eventTitre)
        val date: TextView = view.findViewById(R.id.eventDate)
        val lieu: TextView = view.findViewById(R.id.eventLieu)
        val animateur: TextView = view.findViewById(R.id.eventAnimateur)
        val prix: TextView = view.findViewById(R.id.eventPrix)
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): EvenementViewHolder {
        val view = LayoutInflater.from(parent.context)
            .inflate(R.layout.item_evenement, parent, false)
        return EvenementViewHolder(view)
    }

    override fun getItemCount(): Int = evenements.size

    override fun onBindViewHolder(holder: EvenementViewHolder, position: Int) {
        val e = evenements[position]

        holder.titre.text = e.titre ?: ""
        holder.date.text = "${e.date_debut ?: ""}"
        holder.lieu.text = "${e.type ?: ""} · ${e.lieu ?: ""}"
        holder.animateur.text = "Animé par ${e.prenomSalarie ?: ""} ${e.nomSalarie ?: ""}"
        holder.prix.text = if (e.prix <= 0) "Gratuit" else "${e.prix} €"
    }
}
