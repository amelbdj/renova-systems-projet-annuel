package admin

import "os"

func envOr(key string, defaut string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaut
}
