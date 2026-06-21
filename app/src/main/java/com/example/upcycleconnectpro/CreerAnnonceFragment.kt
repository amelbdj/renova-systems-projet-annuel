package com.example.upcycleconnectpro

import android.net.Uri
import android.os.Bundle
import android.view.View
import android.widget.ArrayAdapter
import android.widget.EditText
import android.widget.ImageButton
import android.widget.ImageView
import android.widget.Spinner
import android.widget.Toast
import androidx.activity.result.contract.ActivityResultContracts
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import com.example.upcycleconnectpro.network.ApiClient
import com.google.android.material.button.MaterialButton
import kotlinx.coroutines.launch
import okhttp3.MediaType.Companion.toMediaTypeOrNull
import okhttp3.MultipartBody
import okhttp3.RequestBody
import okhttp3.RequestBody.Companion.toRequestBody

class CreerAnnonceFragment : Fragment(R.layout.fragment_creer_annonce) {

    // Catégories : l'index +1 correspond à l'id en base (1=Textile, 2=Bois, 3=Plastique, 4=Métal)
    private val categories = listOf("Textile", "Bois", "Plastique", "Métal")
    private val etats = listOf("Neuf", "Très bon état", "Bon état", "État moyen", "Usé")

    // Image choisie par l'utilisateur
    private var imageUri: Uri? = null
    private lateinit var ivApercu: ImageView

    // Sélecteur d'image (galerie)
    private val pickImage = registerForActivityResult(ActivityResultContracts.GetContent()) { uri ->
        if (uri != null) {
            imageUri = uri
            ivApercu.setImageURI(uri)
        }
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        val btnBack = view.findViewById<ImageButton>(R.id.btn_back_creer)
        ivApercu = view.findViewById(R.id.iv_apercu_image)
        val btnChoisirImage = view.findViewById<MaterialButton>(R.id.btn_choisir_image)
        val etTitre = view.findViewById<EditText>(R.id.et_titre)
        val etDescription = view.findViewById<EditText>(R.id.et_description)
        val etPrix = view.findViewById<EditText>(R.id.et_prix)
        val spCategorie = view.findViewById<Spinner>(R.id.sp_categorie)
        val spEtat = view.findViewById<Spinner>(R.id.sp_etat)
        val etVille = view.findViewById<EditText>(R.id.et_ville)
        val etCodePostal = view.findViewById<EditText>(R.id.et_code_postal)
        val etPoids = view.findViewById<EditText>(R.id.et_poids)
        val etQuantite = view.findViewById<EditText>(R.id.et_quantite)
        val btnPublier = view.findViewById<MaterialButton>(R.id.btn_publier)

        // Remplissage des listes déroulantes
        spCategorie.adapter = ArrayAdapter(requireContext(), android.R.layout.simple_spinner_dropdown_item, categories)
        spEtat.adapter = ArrayAdapter(requireContext(), android.R.layout.simple_spinner_dropdown_item, etats)

        btnBack.setOnClickListener {
            parentFragmentManager.popBackStack()
        }

        btnChoisirImage.setOnClickListener {
            pickImage.launch("image/*")
        }

        btnPublier.setOnClickListener {
            val titre = etTitre.text.toString().trim()
            val description = etDescription.text.toString().trim()
            val prix = etPrix.text.toString().trim()
            val ville = etVille.text.toString().trim()
            val codePostal = etCodePostal.text.toString().trim()
            val poids = etPoids.text.toString().trim()
            val quantite = etQuantite.text.toString().trim()

            // Vérifications de base
            if (titre.isEmpty() || description.isEmpty() || prix.isEmpty() ||
                ville.isEmpty() || codePostal.isEmpty() || poids.isEmpty() || quantite.isEmpty()
            ) {
                Toast.makeText(requireContext(), "Veuillez remplir tous les champs", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }

            val sessionManager = SessionManager(requireContext())
            val token = sessionManager.getToken()
            val userId = sessionManager.getUserId()
            if (token == null || userId == -1) {
                Toast.makeText(requireContext(), "Erreur : non connecté", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }

            val idCategorie = (spCategorie.selectedItemPosition + 1).toString()
            val etat = etats[spEtat.selectedItemPosition]
            val typeMateriau = categories[spCategorie.selectedItemPosition] // on réutilise la catégorie comme "type"

            publier(
                token, titre, description, typeMateriau, prix, ville, codePostal,
                etat, poids, quantite, idCategorie, userId.toString(), btnPublier
            )
        }
    }

    private fun publier(
        token: String, titre: String, description: String, type: String, prix: String,
        ville: String, codePostal: String, etat: String, poids: String, quantite: String,
        idCategorie: String, idUser: String, btnPublier: MaterialButton
    ) {
        btnPublier.isEnabled = false
        btnPublier.text = "Publication..."

        viewLifecycleOwner.lifecycleScope.launch {
            try {
                // Préparation de l'image (optionnelle)
                var imagePart: MultipartBody.Part? = null
                if (imageUri != null) {
                    val inputStream = requireContext().contentResolver.openInputStream(imageUri!!)
                    val bytes = inputStream?.readBytes()
                    inputStream?.close()
                    if (bytes != null) {
                        val reqFile = bytes.toRequestBody("image/*".toMediaTypeOrNull())
                        imagePart = MultipartBody.Part.createFormData("image", "photo_${System.currentTimeMillis()}.jpg", reqFile)
                    }
                }

                val response = ApiClient.apiService.createAnnonce(
                    token = "Bearer $token",
                    titre = textPart(titre),
                    description = textPart(description),
                    type = textPart(type),
                    prix = textPart(prix),
                    ville = textPart(ville),
                    codePostal = textPart(codePostal),
                    etat = textPart(etat),
                    poidsKg = textPart(poids),
                    quantite = textPart(quantite),
                    idCategorie = textPart(idCategorie),
                    idUser = textPart(idUser),
                    image = imagePart
                )

                if (response.isSuccessful) {
                    Toast.makeText(
                        requireContext(),
                        "Annonce publiée ! Elle sera visible après validation par un administrateur.",
                        Toast.LENGTH_LONG
                    ).show()
                    parentFragmentManager.popBackStack()
                } else {
                    val erreur = response.errorBody()?.string() ?: ""
                    // Cas spécifique : le compte de paiement n'est pas configuré
                    if (response.code() == 403 || erreur.contains("STRIPE_NOT_CONFIGURED")) {
                        Toast.makeText(
                            requireContext(),
                            "Vous devez d'abord configurer votre compte de paiement (Stripe) pour publier une annonce.",
                            Toast.LENGTH_LONG
                        ).show()
                    } else {
                        Toast.makeText(requireContext(), ErrorMessages.fromServer(erreur), Toast.LENGTH_LONG).show()
                    }
                    btnPublier.isEnabled = true
                    btnPublier.text = "Publier l'annonce"
                }

            } catch (e: Exception) {
                Toast.makeText(requireContext(), ErrorMessages.fromNetwork(), Toast.LENGTH_LONG).show()
                btnPublier.isEnabled = true
                btnPublier.text = "Publier l'annonce"
            }
        }
    }

    // Transforme un texte en "morceau" multipart pour Retrofit
    private fun textPart(value: String): RequestBody {
        return value.toRequestBody("text/plain".toMediaTypeOrNull())
    }
}
