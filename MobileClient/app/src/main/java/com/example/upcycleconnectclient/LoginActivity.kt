package com.example.upcycleconnectclient

import android.content.Intent
import android.os.Bundle
import android.widget.Button
import android.widget.EditText
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.lifecycleScope
import com.example.upcycleconnectclient.models.LoginRequest
import com.example.upcycleconnectclient.network.ApiClient
import kotlinx.coroutines.launch

class LoginActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_login)

        val session = SessionManager(this)

        if (session.getToken() != null) {
            goToHome()
            return
        }

        val emailInput = findViewById<EditText>(R.id.emailInput)
        val passwordInput = findViewById<EditText>(R.id.passwordInput)
        val loginButton = findViewById<Button>(R.id.loginButton)
        val statusText = findViewById<TextView>(R.id.statusText)

        loginButton.setOnClickListener {
            val email = emailInput.text.toString().trim()
            val password = passwordInput.text.toString()

            if (email.isEmpty() || password.isEmpty()) {
                statusText.text = "Veuillez remplir tous les champs."
                return@setOnClickListener
            }

            statusText.text = "Connexion en cours..."

            lifecycleScope.launch {
                try {
                    val response = ApiClient.apiService.login(LoginRequest(email, password))
                    val data = response.body()

                    if (response.isSuccessful && data?.token != null) {
                        session.saveToken(data.token)
                        session.saveUserId(data.id ?: -1)
                        session.saveUserName(data.prenom ?: "")
                        goToHome()
                    } else {
                        statusText.text = "Email ou mot de passe incorrect."
                    }
                } catch (e: Exception) {
                    statusText.text = "Erreur de connexion au serveur."
                }
            }
        }
    }

    private fun goToHome() {
        startActivity(Intent(this, MainActivity::class.java))
        finish()
    }
}
