// db.go contém todas as funções relacionadas à conexão
// com o banco de dados
package utils

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	//"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConectarBD() *sqlx.DB {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Conecta ao banco de dados
	db, err := sqlx.Connect("postgres", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}

	// Checa conexão com o banco de dados
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
