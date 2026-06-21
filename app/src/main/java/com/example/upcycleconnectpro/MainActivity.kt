package com.example.upcycleconnectpro

import android.os.Bundle
import android.view.View
import androidx.appcompat.app.AppCompatActivity
import androidx.fragment.app.Fragment
import androidx.lifecycle.lifecycleScope
import com.google.android.material.bottomnavigation.BottomNavigationView
import com.onesignal.OneSignal
import kotlinx.coroutines.launch

class MainActivity : AppCompatActivity() {

    private lateinit var bottomNav: BottomNavigationView

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        // 0. On demande la permission d'envoyer des notifications (Android 13+)
        lifecycleScope.launch {
            OneSignal.Notifications.requestPermission(true)
        }

        // 1. Récupération de la barre de navigation du bas
        bottomNav = findViewById(R.id.bottom_navigation)

        // 2. Écouteur de clics sur les onglets de la barre de navigation
        bottomNav.setOnItemSelectedListener { item ->
            when (item.itemId) {
                // Correspond à android:id="@+id/nav_home"
                R.id.nav_home -> {
                    replaceFragment(AccueilFragment())
                    true
                }
                // Correspond à android:id="@+id/nav_catalogue"
                R.id.nav_catalogue -> {
                    replaceFragment(CatalogueFragment())
                    true
                }
                // Correspond à android:id="@+id/nav_articles"
                R.id.nav_articles -> {
                    replaceFragment(HomeFragment())
                    true
                }
                // Correspond à android:id="@+id/nav_profile"
                R.id.nav_profile -> {
                    replaceFragment(ProfileFragment())
                    true
                }
                // Correspond à android:id="@+id/nav_settings"
                R.id.nav_settings -> {
                    replaceFragment(MenuFragment())
                    true
                }
                else -> false
            }
        }

        // 3. Premier démarrage : Vérification de la session
        if (savedInstanceState == null) {
            val sessionManager = SessionManager(this)

            if (sessionManager.getToken() != null) {
                // L'utilisateur a déjà un token valide -> On affiche la barre
                bottomNav.visibility = View.VISIBLE

                // On associe l'appareil à l'utilisateur pour recevoir ses notifications
                val userId = sessionManager.getUserId()
                if (userId != -1) {
                    OneSignal.login(userId.toString())
                }

                // On sélectionne l'onglet Accueil : cela déclenche AUTOMATIQUEMENT
                // le OnItemSelectedListener au-dessus, chargeant la page proprement une seule fois.
                bottomNav.selectedItemId = R.id.nav_home
            } else {
                // Aucun token trouvé -> On cache la barre et direction l'écran de Login
                bottomNav.visibility = View.GONE
                replaceFragment(LoginFragment())
            }
        }
    }

    /**
     * Fonction permettant de changer le fragment affiché dans le conteneur principal.
     * Gère aussi automatiquement l'affichage ou le masquage de la barre de navigation.
     */
    fun replaceFragment(fragment: Fragment) {
        if (fragment is LoginFragment) {
            bottomNav.visibility = View.GONE
        } else {
            bottomNav.visibility = View.VISIBLE
        }

        supportFragmentManager.beginTransaction()
            .replace(R.id.fragment_container, fragment)
            .commit()
    }
}