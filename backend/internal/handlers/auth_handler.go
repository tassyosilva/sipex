package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tassyosilva/sipex/internal/auth"
	"github.com/tassyosilva/sipex/internal/models"
	"github.com/tassyosilva/sipex/internal/repository"
)

// AuthHandler gerencia as requisições de autenticação
type AuthHandler struct {
	UsuarioRepo *repository.UsuarioRepository
}

// NovoAuthHandler cria uma nova instância do handler de autenticação
func NovoAuthHandler(usuarioRepo *repository.UsuarioRepository) *AuthHandler {
	return &AuthHandler{
		UsuarioRepo: usuarioRepo,
	}
}

// Login autentica um usuário e retorna um token JWT
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credenciais models.CredenciaisUsuario

	// Decodificar o corpo da requisição
	err := json.NewDecoder(r.Body).Decode(&credenciais)
	if err != nil {
		http.Error(w, "Erro ao decodificar requisição", http.StatusBadRequest)
		return
	}

	// Verificar credenciais
	usuario, err := h.UsuarioRepo.VerificarCredenciais(credenciais)
	if err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Gerar token JWT
	token, err := auth.GerarToken(usuario)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	// Preparar resposta
	resposta := models.RespostaUsuario{
		ID:        usuario.ID,
		Usuario:   usuario.Usuario,
		Nome:      usuario.Nome,
		CPF:       usuario.CPF,
		Matricula: usuario.Matricula,
		Telefone:  usuario.Telefone,
		Unidade:   usuario.Unidade,
		Email:     usuario.Email,
		Cargo:     usuario.Cargo,
		Papel:     usuario.Papel,
		FotoURL:   usuario.FotoURL,
		Token:     token,
	}

	// Enviar resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta)
}

// Registrar cria um novo usuário
func (h *AuthHandler) Registrar(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario

	// Decodificar o corpo da requisição
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Erro ao decodificar requisição", http.StatusBadRequest)
		return
	}

	// Validar dados do usuário
	if usuario.Usuario == "" || usuario.Senha == "" || usuario.Nome == "" || usuario.CPF == "" || usuario.Papel == "" {
		http.Error(w, "Campos obrigatórios não preenchidos", http.StatusBadRequest)
		return
	}

	// Criar usuário
	id, err := h.UsuarioRepo.Criar(usuario)
	if err != nil {
		http.Error(w, "Erro ao criar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Preparar resposta
	resposta := map[string]interface{}{
		"id":      id,
		"mensagem": "Usuário criado com sucesso",
	}

	// Enviar resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resposta)
}
