package middleware

import (
	"context"
	"net/http"

	"github.com/tassyosilva/sipex/internal/auth"
)

// Chave para armazenar informações do usuário no contexto
type contextKey string

const UsuarioContextKey contextKey = "usuario"

// AutenticarMiddleware verifica se o usuário está autenticado
func AutenticarMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extrair token do cabeçalho
		tokenString, err := auth.ExtrairTokenDaRequisicao(r)
		if err != nil {
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}

		// Validar token
		claims, err := auth.ValidarToken(tokenString)
		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		// Adicionar informações do usuário ao contexto
		ctx := context.WithValue(r.Context(), UsuarioContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// VerificarPapel verifica se o usuário tem o papel necessário
func VerificarPapel(papelNecessario string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Obter claims do contexto
			claims, ok := r.Context().Value(UsuarioContextKey).(*auth.Claims)
			if !ok {
				http.Error(w, "Não autorizado", http.StatusUnauthorized)
				return
			}

			// Verificar papel
			if claims.Papel != papelNecessario && claims.Papel != "admin" {
				http.Error(w, "Acesso negado", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
