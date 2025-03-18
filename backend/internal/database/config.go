package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// ConnectDB estabelece conexão com o banco de dados PostgreSQL
func ConnectDB() (*sql.DB, error) {
	// Obter configurações do banco de dados das variáveis de ambiente
	// ou usar valores padrão para desenvolvimento
	host := getEnv("DB_HOST", "192.168.1.106")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "123456")
	dbname := getEnv("DB_NAME", "sipex_db")
	
	// Construir string de conexão
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
	// Abrir conexão com o banco de dados
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	
	// Verificar conexão
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso")
	return db, nil
}

// getEnv obtém o valor de uma variável de ambiente ou retorna um valor padrão
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// CreateTables cria as tabelas necessárias se não existirem
func CreateTables(db *sql.DB) error {
	// Tabela de usuários com os campos em português
	query := `
	CREATE TABLE IF NOT EXISTS usuarios (
		id SERIAL PRIMARY KEY,
		usuario VARCHAR(50) UNIQUE NOT NULL,
		senha VARCHAR(100) NOT NULL,
		nome VARCHAR(100) NOT NULL,
		cpf VARCHAR(14) UNIQUE NOT NULL,
		matricula VARCHAR(50) UNIQUE,
		telefone VARCHAR(20),
		unidade VARCHAR(100),
		email VARCHAR(100) UNIQUE,
		cargo VARCHAR(100),
		papel VARCHAR(20) NOT NULL,
		foto_url TEXT,
		criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		atualizado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	
	_, err := db.Exec(query)
	return err
}
