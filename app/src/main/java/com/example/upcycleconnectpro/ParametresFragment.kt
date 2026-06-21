package com.example.upcycleconnectpro

import android.os.Bundle
import android.util.Patterns
import android.view.View
import android.widget.EditText
import android.widget.ImageButton
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import com.example.upcycleconnectpro.network.ApiClient
import com.example.upcycleconnectpro.network.UpdatePasswordRequest
import com.example.upcycleconnectpro.network.UpdateProfileRequest
import com.google.android.material.button.MaterialButton
import kotlinx.coroutines.launch

class ParametresFragment : Fragment(R.layout.fragment_parametres) {

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        val btnBack = view.findViewById<ImageButton>(R.id.btn_back_parametres)

        // Champs profil
        val etPrenom = view.findViewById<EditText>(R.id.et_param_prenom)
        val etNom = view.findViewById<EditText>(R.id.et_param_nom)
        val etEmail = view.findViewById<EditText>(R.id.et_param_email)
        val btnEnregistrer = view.findViewById<MaterialButton>(R.id.btn_enregistrer_profil)

        // Champs mot de passe
        val etAncien = view.findViewById<EditText>(R.id.et_ancien_mdp)
        val etNouveau = view.findViewById<EditText>(R.id.et_nouveau_mdp)
        val etConfirm = view.findViewById<EditText>(R.id.et_confirm_mdp)
        val btnChanger = view.findViewById<MaterialButton>(R.id.btn_changer_mdp)

        btnBack.setOnClickListener {
            parentFragmentManager.popBackStack()
        }

        val sessionManager = SessionManager(requireContext())
        val token = sessionManager.getToken()
        val userId = sessionManager.getUserId()

        // 1. Pré-remplissage des infos personnelles
        if (token != null && userId != -1) {
            viewLifecycleOwner.lifecycleScope.launch {
                try {
                    val user = ApiClient.apiService.getUserById(userId, "Bearer $token")
                    etPrenom.setText(user.Prenom)
                    etNom.setText(user.Nom)
                    etEmail.setText(user.Email)
                } catch (e: Exception) {
                    Toast.makeText(requireContext(), "Impossible de charger votre profil", Toast.LENGTH_SHORT).show()
                }
            }
        }

        // 2. Enregistrement des modifications du profil
        btnEnregistrer.setOnClickListener {
            val prenom = etPrenom.text.toString().trim()
            val nom = etNom.text.toString().trim()
            val email = etEmail.text.toString().trim()

            if (prenom.isEmpty() || nom.isEmpty() || email.isEmpty()) {
                Toast.makeText(requireContext(), "Veuillez remplir tous les champs", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }
            if (!Patterns.EMAIL_ADDRESS.matcher(email).matches()) {
                Toast.makeText(requireContext(), "Adresse email invalide", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }
            if (token == null || userId == -1) {
                Toast.makeText(requireContext(), "Erreur : non connecté", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }

            btnEnregistrer.isEnabled = false
            btnEnregistrer.text = "Enregistrement..."

            viewLifecycleOwner.lifecycleScope.launch {
                try {
                    val response = ApiClient.apiService.updateProfile(
                        token = "Bearer $token",
                        request = UpdateProfileRequest(nom = nom, prenom = prenom, email = email)
                    )
                    if (response.isSuccessful) {
                        Toast.makeText(requireContext(), "Profil mis à jour avec succès !", Toast.LENGTH_LONG).show()
                    } else {
                        Toast.makeText(
                            requireContext(),
                            ErrorMessages.fromServer(response.errorBody()?.string()),
                            Toast.LENGTH_LONG
                        ).show()
                    }
                } catch (e: Exception) {
                    Toast.makeText(requireContext(), ErrorMessages.fromNetwork(), Toast.LENGTH_LONG).show()
                } finally {
                    btnEnregistrer.isEnabled = true
                    btnEnregistrer.text = "Enregistrer les modifications"
                }
            }
        }

        // 3. Changement de mot de passe
        btnChanger.setOnClickListener {
            val ancien = etAncien.text.toString().trim()
            val nouveau = etNouveau.text.toString().trim()
            val confirm = etConfirm.text.toString().trim()

            if (ancien.isEmpty() || nouveau.isEmpty() || confirm.isEmpty()) {
                Toast.makeText(requireContext(), "Veuillez remplir tous les champs", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }
            if (nouveau.length < 6) {
                Toast.makeText(requireContext(), "Le nouveau mot de passe doit faire au moins 6 caractères", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }
            if (nouveau != confirm) {
                Toast.makeText(requireContext(), "Les deux mots de passe ne correspondent pas", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }
            if (token == null || userId == -1) {
                Toast.makeText(requireContext(), "Erreur : non connecté", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }

            btnChanger.isEnabled = false
            btnChanger.text = "Mise à jour..."

            viewLifecycleOwner.lifecycleScope.launch {
                try {
                    val response = ApiClient.apiService.updatePassword(
                        token = "Bearer $token",
                        request = UpdatePasswordRequest(
                            id = userId,
                            old_password = ancien,
                            new_password = nouveau
                        )
                    )
                    if (response.isSuccessful) {
                        Toast.makeText(requireContext(), "Mot de passe mis à jour avec succès !", Toast.LENGTH_LONG).show()
                        etAncien.text.clear()
                        etNouveau.text.clear()
                        etConfirm.text.clear()
                    } else {
                        Toast.makeText(
                            requireContext(),
                            ErrorMessages.fromServer(response.errorBody()?.string()),
                            Toast.LENGTH_LONG
                        ).show()
                    }
                } catch (e: Exception) {
                    Toast.makeText(requireContext(), ErrorMessages.fromNetwork(), Toast.LENGTH_LONG).show()
                } finally {
                    btnChanger.isEnabled = true
                    btnChanger.text = "Mettre à jour le mot de passe"
                }
            }
        }
    }
}
