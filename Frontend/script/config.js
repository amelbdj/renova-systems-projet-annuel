// Configuration commune du front - adresse de l'API Go.
//
// - En LOCAL (WAMP, dev) : on tape directement le backend Go sur le port 8081.
// - En DEPLOIEMENT (VM / Docker) : l'API est servie derriere /api par Nginx
//   (meme origine que le site), donc pas besoin d'IP en dur.
//
// On detecte le contexte automatiquement, mais tu peux toujours forcer
// une valeur en definissant window.API_BASE_URL avant ce script.
(function () {
  var loc = window.location;
  var estLocalDev =
    loc.hostname === "localhost" ||
    loc.hostname === "127.0.0.1" ||
    loc.port === "8081";

  if (!window.API_BASE_URL) {
    window.API_BASE_URL = estLocalDev
      ? "http://localhost:8081" // dev local (WAMP)
      : loc.origin + "/api"; // VM / Docker (proxy Nginx)
  }
})();

var API_BASE_URL = window.API_BASE_URL;

// Adresse WebSocket (messagerie temps reel) deduite de l'API.
window.WS_BASE_URL = window.WS_BASE_URL || API_BASE_URL.replace(/^http/, "ws");
var WS_BASE_URL = window.WS_BASE_URL;
