package com.example.upcycleconnectpro

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import com.example.upcycleconnectpro.network.Reservation

class HistoriqueAdapter(
    private val reservations: List<Reservation>,
    private val onItemClick: (Reservation) -> Unit
) : RecyclerView.Adapter<HistoriqueAdapter.ViewHolder>() {

    class ViewHolder(view: View) : RecyclerView.ViewHolder(view) {
        val tvObjet: TextView = view.findViewById(R.id.tv_res_objet)
        val tvEtat: TextView = view.findViewById(R.id.tv_res_etat)
        val tvLieu: TextView = view.findViewById(R.id.tv_res_lieu)
        val tvBox: TextView = view.findViewById(R.id.tv_res_box)
        val tvPin: TextView = view.findViewById(R.id.tv_res_pin)
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {
        val view = LayoutInflater.from(parent.context)
            .inflate(R.layout.item_reservation, parent, false)
        return ViewHolder(view)
    }

    override fun onBindViewHolder(holder: ViewHolder, position: Int) {
        val res = reservations[position]
        holder.tvObjet.text = res.objet
        holder.tvEtat.text = res.etat
        holder.tvLieu.text = res.lieu
        holder.tvBox.text = res.numero_box
        holder.tvPin.text = res.code_pin

        // Clic sur la carte → ouvre l'écran PIN
        holder.itemView.setOnClickListener {
            onItemClick(res)
        }
    }

    override fun getItemCount(): Int = reservations.size
}