package com.example.upcycleconnectpro

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.EditText
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import com.example.upcycleconnectpro.network.ApiClient
import com.example.upcycleconnectpro.network.LoginRequest
import kotlinx.coroutines.launch

class LoginFragment : Fragment() {

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        return inflater.inflate(R.layout.fragment_login, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        val emailInput = view.findViewById<EditText>(R.id.et_email)
        val passwordInput = view.findViewById<EditText>(R.id.et_password)
        val loginButton = view.findViewById<Button>(R.id.btn_login)

        loginButton.setOnClickListener {
            val email = emailInput.text.toString().trim()
            val password = passwordInput.text.toString().trim()

            if (email.isNotEmpty() && password.isNotEmpty()) {
                performLogin(email, password)
            } else {
                Toast.makeText(requireContext(), "Veuillez remplir tous les champs", Toast.LENGTH_SHORT).show()
            }
        }
    }

    private fun performLogin(email: String, mdp: String) {
        viewLifecycleOwner.lifecycleScope.launch {
            try {
                // 1. Appel API vers ton serveur Go
                val response = ApiClient.apiService.login(LoginRequest(email, mdp))

                if (!response.token.isNullOrEmpty()) {

                    // 1.bis CONTRÔLE D'ACCÈS : seuls les professionnels/artisans sont autorisés.
                    // On accepte les différentes orthographes possibles, sans tenir compte de la casse.
                    val role = response.role?.trim()?.lowercase() ?: ""
                    val rolesAutorises = listOf("prestataire", "professionnel", "pro", "artisan")

                    if (!rolesAutorises.contains(role)) {
                        Toast.makeText(
                            requireContext(),
                            "Accès refusé (rôle reçu : \"${response.role}\"). Application réservée aux professionnels et artisans.",
                            Toast.LENGTH_LONG
                        ).show()
                        return@launch
                    }

                    // 2. Initialisation du SessionManager
                    val sessionManager = SessionManager(requireContext())

                    // 3. Sauvegarde du Token
                    sessionManager.saveToken(response.token)

                    // 4. Sauvegarde du véritable ID Utilisateur (Le correctif est ici !)
                    val userId = response.id

                    if (userId != null && userId != 0) {
                        sessionManager.saveUserId(userId)

                        // On associe l'appareil à cet utilisateur pour les notifications OneSignal
                        com.onesignal.OneSignal.login(userId.toString())

                        // Toast temporaire pour te prouver que ça marche :
                        Toast.makeText(requireContext(), "Connexion OK ! (ID: $userId)", Toast.LENGTH_SHORT).show()
                    } else {
                        Toast.makeText(requireContext(), "Attention : ID introuvable", Toast.LENGTH_LONG).show()
                    }

                    // 5. Navigation vers la page principale de l'application
                    (activity as MainActivity).replaceFragment(CatalogueFragment())

                } else {
                    Toast.makeText(requireContext(), "Erreur : ${response.message ?: "Identifiants incorrects"}", Toast.LENGTH_LONG).show()
                }
            } catch (e: Exception) {
                Toast.makeText(requireContext(), "Erreur réseau : ${e.localizedMessage}", Toast.LENGTH_LONG).show()
            }
        }
    }
}