package main

import (
    "fmt"
    "log"
    "net/http"
    
    "github.com/gorilla/mux"
    "github.com/tassyosilva/sipex/internal/database"
    "github.com/tassyosilva/sipex/internal/handlers"
    "github.com/tassyosilva/sipex/internal/middleware"
    "github.com/tassyosilva/sipex/internal/repository"
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
    
    // Inicializar repositórios
    usuarioRepo := repository.NovoUsuarioRepository(db)
    
    // Inicializar usuário admin se necessário
    err = usuarioRepo.InicializarAdmin()
    if err != nil {
        log.Fatalf("Erro ao inicializar usuário admin: %v", err)
    }
    
    // Inicializar handlers
    authHandler := handlers.NovoAuthHandler(usuarioRepo)
    usuarioHandler := handlers.NovoUsuarioHandler(usuarioRepo)
    
    // Configurar rotas
    r := mux.NewRouter()
    
    // Rotas públicas
    r.HandleFunc("/api/saude", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("API está funcionando!"))
    }).Methods("GET")
    
    // Rotas de autenticação
    r.HandleFunc("/api/auth/login", authHandler.Login).Methods("POST")
    
    // API Router com middleware de autenticação
    api := r.PathPrefix("/api").Subrouter()
    api.Use(middleware.AutenticarMiddleware)
    
    // Rotas protegidas
    // Rota de teste para verificar autenticação
    api.HandleFunc("/perfil", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Usuário autenticado!"))
    }).Methods("GET")
    
    // Rotas de usuários (protegidas)
    api.HandleFunc("/usuarios", usuarioHandler.Listar).Methods("GET")
    api.HandleFunc("/usuarios/{id}", usuarioHandler.ObterPorID).Methods("GET")
    api.HandleFunc("/usuarios/{id}", usuarioHandler.Atualizar).Methods("PUT")
    api.HandleFunc("/usuarios/{id}/senha", usuarioHandler.AtualizarSenha).Methods("PUT")
    
    // Rotas de administração (apenas para admin)
    admin := r.PathPrefix("/api/admin").Subrouter()
    admin.Use(middleware.AutenticarMiddleware)
    admin.Use(middleware.VerificarPapel("admin"))
    
    // Rota para registrar novos usuários (apenas admin pode criar)
    admin.HandleFunc("/usuarios", authHandler.Registrar).Methods("POST")
    admin.HandleFunc("/usuarios/{id}", usuarioHandler.Excluir).Methods("DELETE")
    
    // Configurar CORS
    handler := corsMiddleware(r)
    
    // Definir porta
    port := "8080"
    
    // Iniciar servidor
    fmt.Printf("Servidor rodando na porta %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, handler))
}

// corsMiddleware adiciona os cabeçalhos CORS necessários
func corsMiddleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        handler.ServeHTTP(w, r)
    })
}