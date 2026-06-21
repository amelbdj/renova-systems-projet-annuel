package com.example.upcycleconnectpro

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import com.example.upcycleconnectpro.network.Invoice
import com.google.android.material.button.MaterialButton

class FacturationAdapter(
    private val factures: List<Invoice>,
    private val onTelecharger: (Invoice) -> Unit
) : RecyclerView.Adapter<FacturationAdapter.ViewHolder>() {

    class ViewHolder(view: View) : RecyclerView.ViewHolder(view) {
        val tvTitre: TextView = view.findViewById(R.id.tv_facture_titre)
        val tvMontant: TextView = view.findViewById(R.id.tv_facture_montant)
        val tvDate: TextView = view.findViewById(R.id.tv_facture_date)
        val tvRef: TextView = view.findViewById(R.id.tv_facture_ref)
        val btnTelecharger: MaterialButton = view.findViewById(R.id.btn_telecharger_facture)
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {
        val view = LayoutInflater.from(parent.context)
            .inflate(R.layout.item_facture, parent, false)
        return ViewHolder(view)
    }

    override fun onBindViewHolder(holder: ViewHolder, position: Int) {
        val facture = factures[position]
        holder.tvTitre.text = facture.titre
        holder.tvMontant.text = String.format("%.2f €", facture.montant)
        holder.tvDate.text = facture.date
        holder.tvRef.text = "Facture n° ${facture.id}"

        holder.btnTelecharger.setOnClickListener {
            onTelecharger(facture)
        }
    }

    override fun getItemCount(): Int = factures.size
}
