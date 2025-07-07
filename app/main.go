package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-bm/app/config"
	"go-bm/app/models"
	"log"
	"net/http"
)

func paymentSummaryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	// Retorna como resposta
	fmt.Fprintf(w, "From: %s\nTo: %s\n", from, to)
	w.Write([]byte("Payment Summary Endpoint"))
}
func paymentsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// Implementar lógica para lidar com pagamentos
	w.Write([]byte("Payments Endpoint"))
}

func main() {
	var dbcon = config.SetupDB()
	if dbcon == nil {
		panic("Erro ao estabelecer conexão com o banco de dados")
	}
	_, err := dbcon.Exec(models.CreateTableQuery)
	if err != nil {
		panic("Erro ao criar tabela: " + err.Error())
	} else {
		println("Tabela criada com sucesso")
	}

	//CRUD operations can be performed here
	router := mux.NewRouter()

	router.HandleFunc("/payments_summary", paymentSummaryHandler).Methods("GET")
	router.HandleFunc("/payments", paymentsHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
	defer dbcon.Close()

}
