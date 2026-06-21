package com.example.upcycleconnectpro

import android.app.Application
import com.onesignal.OneSignal

class UpcycleApp : Application() {

    companion object {
        // Même App ID que celui utilisé par le serveur Go pour envoyer les notifications
        const val ONESIGNAL_APP_ID = "79a53223-420a-46c8-83d9-1ca162fcb64f"
    }

    override fun onCreate() {
        super.onCreate()
        // Initialisation du SDK OneSignal une seule fois, au lancement de l'app
        OneSignal.initWithContext(this, ONESIGNAL_APP_ID)
    }
}
