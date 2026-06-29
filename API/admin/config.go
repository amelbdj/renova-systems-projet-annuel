package admin

import "os"

// envOr retourne la variable d'environnement si elle est definie,
// sinon la valeur par defaut (pratique pour garder le dev local fonctionnel).
func envOr(key string, defaut string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaut
}
