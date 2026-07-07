






(function () {
  if (window.API_BASE_URL) return;

  var p = window.location;
  
  
  
  
  var estMachineLocale =
    p.hostname === "localhost" || p.hostname === "127.0.0.1";
  var estWampLocal = estMachineLocale && p.pathname.indexOf("/Frontend/") !== -1;

  window.API_BASE_URL = estWampLocal
    ? "http://localhost:8081" 
    : p.origin + "/api"; 
})();

var API_BASE_URL = window.API_BASE_URL;


window.WS_BASE_URL = window.WS_BASE_URL || API_BASE_URL.replace(/^http/, "ws");
var WS_BASE_URL = window.WS_BASE_URL;

function getLoginPage() {
  if (window.location.pathname.indexOf("/salarie/") !== -1) {
    return "/login";
  }
  return "/login";
}

function logout() {
  localStorage.clear();
  window.location.href = getLoginPage();
}
