// Configuration commune du front - adresse de l'API Go.
//
// Detection automatique de l'environnement :
//  - WAMP / dev local : le site est servi sous .../Frontend/...  -> API Go en direct sur :8081
//  - Docker / VM      : le site est servi a la racine            -> API via /api (proxy Nginx, meme origine)
//
// Tu peux forcer une valeur en definissant window.API_BASE_URL AVANT ce script.
(function () {
  if (window.API_BASE_URL) return;

  var p = window.location;
  // On n'est en WAMP (backend Go direct sur :8081) QUE si :
  //  - on est sur la machine locale (localhost / 127.0.0.1)
  //  - ET l'URL contient le dossier du projet (.../Frontend/...)
  // Sur la VM (IP publique), meme via /Frontend/, on passe toujours par /api.
  var estMachineLocale =
    p.hostname === "localhost" || p.hostname === "127.0.0.1";
  var estWampLocal = estMachineLocale && p.pathname.indexOf("/Frontend/") !== -1;

  window.API_BASE_URL = estWampLocal
    ? "http://localhost:8081" // dev local (WAMP) : backend Go direct
    : p.origin + "/api"; // Docker / VM : meme origine, proxy Nginx vers le backend
})();

var API_BASE_URL = window.API_BASE_URL;

// Adresse WebSocket (messagerie temps reel) deduite de l'API.
window.WS_BASE_URL = window.WS_BASE_URL || API_BASE_URL.replace(/^http/, "ws");
var WS_BASE_URL = window.WS_BASE_URL;

function getLoginPage() {
  if (window.location.pathname.indexOf("/salarie/") !== -1) {
    return "../login.html";
  }
  return "login.html";
}

function logout() {
  localStorage.clear();
  window.location.href = getLoginPage();
}
