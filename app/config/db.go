package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
	err  error
)

const (
	host     = "localhost"
	port     = 5434
	user     = "postgres"
	password = "postgres"
	dbname   = "rinha-2025-go"
)

func SetupDB() *sql.DB {
	once.Do(func() {
		psqlInfo :=
			"host=" + host +
				" port=" + strconv.Itoa(port) +
				" user=" + user +
				" password=" + password +
				" dbname=" + dbname +
				" sslmode=disable"

		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatalf("Erro ao abrir conexão com o banco: %v", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatalf("Erro ao testar conexão com o banco: %v", err)
		}

		log.Println("Conexão com o banco de dados estabelecida com sucesso")
	})

	return db
}
