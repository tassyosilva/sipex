package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tassyosilva/sipex/internal/database"
)

func main() {
	// Conectar ao banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()
	
	// Criar tabelas se não existirem
	err = database.CreateTables(db)
	if err != nil {
		log.Fatalf("Erro ao criar tabelas: %v", err)
	}
	
	// Configurar rotas
	r := mux.NewRouter()
	
	// Rota de teste
	r.HandleFunc("/api/saude", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API está funcionando!"))
	}).Methods("GET")
	
	// Definir porta
	port := "8080"
	
	// Iniciar servidor
	fmt.Printf("Servidor rodando na porta %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
