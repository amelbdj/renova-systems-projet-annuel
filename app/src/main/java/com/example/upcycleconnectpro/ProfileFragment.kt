package com.example.upcycleconnectpro

import android.os.Bundle
import android.util.Base64
import android.util.Log
import android.view.View
import android.widget.Button
import android.widget.TextView
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import com.example.upcycleconnectpro.network.ApiClient
import kotlinx.coroutines.launch
import org.json.JSONObject
import java.util.* // Import nécessaire pour uppercase(Locale.ROOT)

class ProfileFragment : Fragment(R.layout.fragment_profile) {

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        val sessionManager = SessionManager(requireContext())
        val token = sessionManager.getToken()

        // Récupération des vues
        val tvInitials = view.findViewById<TextView>(R.id.tv_profile_initials) // LA NOUVELLE VUE
        val tvName = view.findViewById<TextView>(R.id.tv_full_name)
        val tvRole = view.findViewById<TextView>(R.id.tv_role)
        val tvEntreprise = view.findViewById<TextView>(R.id.tv_entreprise)
        val tvSiret = view.findViewById<TextView>(R.id.tv_siret)
        val tvScore = view.findViewById<TextView>(R.id.tv_score)
        val btnLogout = view.findViewById<Button>(R.id.btn_logout)

        if (token != null) {
            val userId = extractIdFromToken(token)

            if (userId != null) {
                viewLifecycleOwner.lifecycleScope.launch {
                    try {
                        val user = ApiClient.apiService.getUserById(userId, "Bearer $token")

                        // --- LOGIQUE DES INITIALES (PDP) ---
                        // On récupère la première lettre du Prénom et du Nom, ou "" si vide
                        val initialPrenom = user.Prenom.take(1)
                        val initialNom = user.Nom.take(1)

                        // On assemble et on met en majuscule proprement
                        val initials = "$initialPrenom$initialNom".uppercase(Locale.ROOT)

                        // On applique au TextView de l'avatar
                        tvInitials.text = initials


                        // Affichage des autres données
                        tvName.text = "${user.Prenom} ${user.Nom}"
                        tvRole.text = user.Role
                        tvEntreprise.text = "Entreprise: ${user.NomEntreprise ?: "Non renseigné"}"
                        tvSiret.text = "SIRET: ${user.Siret ?: "Non renseigné"}"
                        tvScore.text = user.Score.toString()

                    } catch (e: Exception) {
                        Log.e("UpcycleProfil", "Erreur API", e)
                        Toast.makeText(context, "Erreur réseau", Toast.LENGTH_SHORT).show()
                        tvInitials.text = "!" // Signe d'erreur dans l'avatar
                    }
                }
            }
        }

        btnLogout.setOnClickListener {
            sessionManager.clearSession()

            // On dissocie l'appareil de l'utilisateur : il ne recevra plus ses notifications
            com.onesignal.OneSignal.logout()

            (requireActivity() as MainActivity).replaceFragment(LoginFragment())
            Toast.makeText(context, "Déconnexion réussie", Toast.LENGTH_SHORT).show()
        }
    }

    // --- FONCTION DE DÉCODAGE DU TOKEN (Inchangée) ---
    private fun extractIdFromToken(token: String): Int? {
        try {
            val parts = token.split(".")
            if (parts.size == 3) {
                val payload = String(Base64.decode(parts[1], Base64.URL_SAFE))
                val jsonObject = JSONObject(payload)
                return if (jsonObject.has("id")) {
                    jsonObject.getInt("id")
                } else if (jsonObject.has("user_id")) {
                    jsonObject.getInt("user_id")
                } else {
                    null
                }
            }
        } catch (e: Exception) {
            Log.e("UpcycleProfil", "Erreur lors du décodage du token", e)
        }
        return null
    }
}