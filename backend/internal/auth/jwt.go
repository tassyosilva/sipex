package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tassyosilva/sipex/internal/models"
)

var jwtKey = []byte("sua_chave_secreta_aqui") // Em produção, use variáveis de ambiente

// Claims representa as claims do JWT
type Claims struct {
	UsuarioID int    `json:"usuario_id"`
	Usuario   string `json:"usuario"`
	Papel     string `json:"papel"`
	jwt.RegisteredClaims
}

// GerarToken gera um token JWT para o usuário
func GerarToken(usuario models.Usuario) (string, error) {
	tempoExpiracao := time.Now().Add(24 * time.Hour)
	
	claims := &Claims{
		UsuarioID: usuario.ID,
		Usuario:   usuario.Usuario,
		Papel:     usuario.Papel,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tempoExpiracao),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	
	return tokenString, err
}

// ValidarToken valida um token JWT
func ValidarToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	
	if err != nil {
		return nil, err
	}
	
	if !token.Valid {
		return nil, errors.New("token inválido")
	}
	
	return claims, nil
}

// ExtrairTokenDaRequisicao extrai o token do cabeçalho Authorization
func ExtrairTokenDaRequisicao(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("cabeçalho de autorização não encontrado")
	}
	
	// Formato esperado: "Bearer {token}"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("formato de autorização inválido")
	}
	
	return parts[1], nil
}
