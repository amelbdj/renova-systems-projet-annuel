package bdd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const driver = "mysql"

var Db *sql.DB

// env retourne la variable d'environnement si elle existe, sinon la valeur par defaut.
// -> En local (WAMP) les valeurs par defaut s'appliquent : rien a changer.
// -> En Docker, docker-compose fournit DB_HOST, DB_USER, etc.
func env(cle string, defaut string) string {
	if v := os.Getenv(cle); v != "" {
		return v
	}
	return defaut
}

func NewDB() *sql.DB {
	host := env("DB_HOST", "localhost")
	port := env("DB_PORT", "3306")
	user := env("DB_USER", "root")
	pass := env("DB_PASS", "root")
	dbname := env("DB_NAME", "pa2026")

	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4", user, pass, host, port, dbname)
	conn, err := sql.Open(driver, sqlInfo)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connexion à la bdd reussie")
	return conn
}
