package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tassyosilva/sipex/internal/auth"
	"github.com/tassyosilva/sipex/internal/middleware"
	"github.com/tassyosilva/sipex/internal/models"
	"github.com/tassyosilva/sipex/internal/repository"
)

// UsuarioHandler gerencia as requisições relacionadas a usuários
type UsuarioHandler struct {
	UsuarioRepo *repository.UsuarioRepository
}

// NovoUsuarioHandler cria uma nova instância do handler de usuários
func NovoUsuarioHandler(usuarioRepo *repository.UsuarioRepository) *UsuarioHandler {
	return &UsuarioHandler{
		UsuarioRepo: usuarioRepo,
	}
}

// Listar retorna todos os usuários
func (h *UsuarioHandler) Listar(w http.ResponseWriter, r *http.Request) {
	usuarios, err := h.UsuarioRepo.Listar()
	if err != nil {
		http.Error(w, "Erro ao listar usuários: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// ObterPorID retorna um usuário pelo ID
func (h *UsuarioHandler) ObterPorID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	usuario, err := h.UsuarioRepo.ObterPorID(id)
	if err != nil {
		http.Error(w, "Erro ao obter usuário: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

// Atualizar atualiza os dados de um usuário
func (h *UsuarioHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Verificar se o usuário existe
	_, err = h.UsuarioRepo.ObterPorID(id)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	var usuario models.Usuario
	err = json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Erro ao decodificar requisição", http.StatusBadRequest)
		return
	}

	usuario.ID = id
	err = h.UsuarioRepo.Atualizar(usuario)
	if err != nil {
		http.Error(w, "Erro ao atualizar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"mensagem": "Usuário atualizado com sucesso"}`))
}

// AtualizarSenha atualiza a senha de um usuário
func (h *UsuarioHandler) AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Verificar se o usuário existe
	_, err = h.UsuarioRepo.ObterPorID(id)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	// Verificar se o usuário atual tem permissão para alterar a senha
	claims, ok := r.Context().Value(middleware.UsuarioContextKey).(*auth.Claims)
	if !ok || (claims.UsuarioID != id && claims.Papel != "admin") {
		http.Error(w, "Acesso negado", http.StatusForbidden)
		return
	}

	var senhaData struct {
		NovaSenha string `json:"nova_senha"`
	}

	err = json.NewDecoder(r.Body).Decode(&senhaData)
	if err != nil {
		http.Error(w, "Erro ao decodificar requisição", http.StatusBadRequest)
		return
	}

	if senhaData.NovaSenha == "" {
		http.Error(w, "Nova senha não pode ser vazia", http.StatusBadRequest)
		return
	}

	err = h.UsuarioRepo.AtualizarSenha(id, senhaData.NovaSenha)
	if err != nil {
		http.Error(w, "Erro ao atualizar senha: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"mensagem": "Senha atualizada com sucesso"}`))
}

// Excluir remove um usuário
func (h *UsuarioHandler) Excluir(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Verificar se o usuário existe
	_, err = h.UsuarioRepo.ObterPorID(id)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	err = h.UsuarioRepo.Excluir(id)
	if err != nil {
		http.Error(w, "Erro ao excluir usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"mensagem": "Usuário excluído com sucesso"}`))
}
