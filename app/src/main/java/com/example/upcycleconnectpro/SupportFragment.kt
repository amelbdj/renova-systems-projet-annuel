package com.example.upcycleconnectpro

import android.content.Intent
import android.net.Uri
import android.os.Bundle
import android.view.View
import android.widget.EditText
import android.widget.ImageButton
import android.widget.Toast
import androidx.fragment.app.Fragment
import com.google.android.material.button.MaterialButton

class SupportFragment : Fragment(R.layout.fragment_support) {

    // Adresse email du support
    private val emailSupport = "support@upcycleconnect.fr"

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        val btnBack = view.findViewById<ImageButton>(R.id.btn_back_support)
        val etSujet = view.findViewById<EditText>(R.id.et_support_sujet)
        val etMessage = view.findViewById<EditText>(R.id.et_support_message)
        val btnEnvoyer = view.findViewById<MaterialButton>(R.id.btn_envoyer_support)

        btnBack.setOnClickListener {
            parentFragmentManager.popBackStack()
        }

        btnEnvoyer.setOnClickListener {
            val sujet = etSujet.text.toString().trim()
            val message = etMessage.text.toString().trim()

            if (sujet.isEmpty() || message.isEmpty()) {
                Toast.makeText(requireContext(), "Veuillez remplir le sujet et le message", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }

            // On ouvre l'application email du téléphone, pré-remplie
            val intent = Intent(Intent.ACTION_SENDTO).apply {
                data = Uri.parse("mailto:")
                putExtra(Intent.EXTRA_EMAIL, arrayOf(emailSupport))
                putExtra(Intent.EXTRA_SUBJECT, sujet)
                putExtra(Intent.EXTRA_TEXT, message)
            }

            try {
                startActivity(Intent.createChooser(intent, "Envoyer via..."))
            } catch (e: Exception) {
                Toast.makeText(requireContext(), "Aucune application email trouvée", Toast.LENGTH_LONG).show()
            }
        }
    }
}
