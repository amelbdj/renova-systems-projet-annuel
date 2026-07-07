package bdd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const driver = "mysql"

var Db *sql.DB

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
