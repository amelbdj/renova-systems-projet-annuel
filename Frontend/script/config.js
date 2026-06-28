// Configuration commune du front
// Change cette valeur une seule fois si l API Go change d adresse.
window.API_BASE_URL = window.API_BASE_URL || "http://localhost:8081";
var API_BASE_URL = window.API_BASE_URL;

// Adresse utilisee pour la messagerie en temps reel.
window.WS_BASE_URL = window.WS_BASE_URL || window.API_BASE_URL.replace("http", "ws");
var WS_BASE_URL = window.WS_BASE_URL;