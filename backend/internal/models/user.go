package models

import "time"

// Usuario representa um usuário do sistema
type Usuario struct {
	ID          int       `json:"id"`
	Usuario     string    `json:"usuario"`
	Senha       string    `json:"-"` // Não retornar senha nas respostas JSON
	Nome        string    `json:"nome"`
	CPF         string    `json:"cpf"`
	Matricula   string    `json:"matricula,omitempty"`
	Telefone    string    `json:"telefone,omitempty"`
	Unidade     string    `json:"unidade,omitempty"`
	Email       string    `json:"email,omitempty"`
	Cargo       string    `json:"cargo,omitempty"`
	Papel       string    `json:"papel"` // admin, diretor, perito, atendente
	FotoURL     string    `json:"foto_url,omitempty"`
	CriadoEm    time.Time `json:"criado_em"`
	AtualizadoEm time.Time `json:"atualizado_em"`
}

// CredenciaisUsuario representa as credenciais para login
type CredenciaisUsuario struct {
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
}

// RespostaUsuario representa os dados do usuário retornados após autenticação
type RespostaUsuario struct {
	ID          int    `json:"id"`
	Usuario     string `json:"usuario"`
	Nome        string `json:"nome"`
	CPF         string `json:"cpf"`
	Matricula   string `json:"matricula,omitempty"`
	Telefone    string `json:"telefone,omitempty"`
	Unidade     string `json:"unidade,omitempty"`
	Email       string `json:"email,omitempty"`
	Cargo       string `json:"cargo,omitempty"`
	Papel       string `json:"papel"`
	FotoURL     string `json:"foto_url,omitempty"`
	Token       string `json:"token"`
}
