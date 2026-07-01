package com.example.upcycleconnectclient

import android.content.Intent
import android.os.Bundle
import android.widget.Button
import android.widget.TextView
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import androidx.core.view.ViewCompat
import androidx.core.view.WindowInsetsCompat

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_main)
        ViewCompat.setOnApplyWindowInsetsListener(findViewById(R.id.main)) { v, insets ->
            val systemBars = insets.getInsets(WindowInsetsCompat.Type.systemBars())
            v.setPadding(systemBars.left, systemBars.top, systemBars.right, systemBars.bottom)
            insets
        }

        val session = SessionManager(this)

        findViewById<TextView>(R.id.welcomeText).text =
            "Bonjour ${session.getUserName() ?: ""}"

        findViewById<Button>(R.id.annoncesButton).setOnClickListener {
            startActivity(Intent(this, AnnoncesActivity::class.java))
        }

        findViewById<Button>(R.id.deposerButton).setOnClickListener {
            startActivity(Intent(this, DeposerAnnonceActivity::class.java))
        }

        findViewById<Button>(R.id.catalogueButton).setOnClickListener {
            startActivity(Intent(this, CatalogueActivity::class.java))
        }

        findViewById<Button>(R.id.scoreButton).setOnClickListener {
            startActivity(Intent(this, ScoreActivity::class.java))
        }

        findViewById<Button>(R.id.logoutButton).setOnClickListener {
            session.clearSession()
            startActivity(Intent(this, LoginActivity::class.java))
            finish()
        }
    }
}
