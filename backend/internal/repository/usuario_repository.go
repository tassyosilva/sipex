package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/tassyosilva/sipex/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// UsuarioRepository define as operações do repositório de usuários
type UsuarioRepository struct {
	DB *sql.DB
}

// NovoUsuarioRepository cria uma nova instância do repositório de usuários
func NovoUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

// InicializarAdmin cria um usuário administrador se não existir nenhum
func (r *UsuarioRepository) InicializarAdmin() error {
    // Verificar se já existe algum usuário admin
    query := `SELECT COUNT(*) FROM usuarios WHERE papel = 'admin'`
    var count int
    err := r.DB.QueryRow(query).Scan(&count)
    if err != nil {
        return err
    }

    // Se já existe pelo menos um admin, não faz nada
    if count > 0 {
        return nil
    }

    // Criar hash da senha 'admin'
    senhaHash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    // Inserir usuário admin
    query = `
        INSERT INTO usuarios (
            usuario, senha, nome, cpf, matricula, telefone, unidade, email, cargo, papel, foto_url, criado_em, atualizado_em
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
        )
    `
    agora := time.Now()
    _, err = r.DB.Exec(
        query,
        "admin",
        string(senhaHash),
        "Administrador do Sistema",
        "00000000000", // CPF fictício
        "000000",      // Matrícula fictícia
        "",            // Telefone vazio
        "SIPEX",       // Unidade padrão
        "admin@sipex.gov.br", // Email padrão
        "Administrador", // Cargo
        "admin",       // Papel
        "",            // Sem foto
        agora,
        agora,
    )
    return err
}

// Criar insere um novo usuário no banco de dados
func (r *UsuarioRepository) Criar(usuario models.Usuario) (int, error) {
	// Hash da senha
	senhaHash, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	query := `
	INSERT INTO usuarios (
		usuario, senha, nome, cpf, matricula, telefone, unidade, email, cargo, papel, foto_url, criado_em, atualizado_em
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
	) RETURNING id
	`

	var id int
	agora := time.Now()

	err = r.DB.QueryRow(
		query,
		usuario.Usuario,
		string(senhaHash),
		usuario.Nome,
		usuario.CPF,
		usuario.Matricula,
		usuario.Telefone,
		usuario.Unidade,
		usuario.Email,
		usuario.Cargo,
		usuario.Papel,
		usuario.FotoURL,
		agora,
		agora,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// ObterPorID busca um usuário pelo ID
func (r *UsuarioRepository) ObterPorID(id int) (models.Usuario, error) {
	query := `
	SELECT id, usuario, senha, nome, cpf, matricula, telefone, unidade, email, cargo, papel, foto_url, criado_em, atualizado_em
	FROM usuarios
	WHERE id = $1
	`

	var usuario models.Usuario
	err := r.DB.QueryRow(query, id).Scan(
		&usuario.ID,
		&usuario.Usuario,
		&usuario.Senha,
		&usuario.Nome,
		&usuario.CPF,
		&usuario.Matricula,
		&usuario.Telefone,
		&usuario.Unidade,
		&usuario.Email,
		&usuario.Cargo,
		&usuario.Papel,
		&usuario.FotoURL,
		&usuario.CriadoEm,
		&usuario.AtualizadoEm,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Usuario{}, errors.New("usuário não encontrado")
		}
		return models.Usuario{}, err
	}

	return usuario, nil
}

// ObterPorUsuario busca um usuário pelo nome de usuário
func (r *UsuarioRepository) ObterPorUsuario(usuario string) (models.Usuario, error) {
	query := `
	SELECT id, usuario, senha, nome, cpf, matricula, telefone, unidade, email, cargo, papel, foto_url, criado_em, atualizado_em
	FROM usuarios
	WHERE usuario = $1
	`

	var u models.Usuario
	err := r.DB.QueryRow(query, usuario).Scan(
		&u.ID,
		&u.Usuario,
		&u.Senha,
		&u.Nome,
		&u.CPF,
		&u.Matricula,
		&u.Telefone,
		&u.Unidade,
		&u.Email,
		&u.Cargo,
		&u.Papel,
		&u.FotoURL,
		&u.CriadoEm,
		&u.AtualizadoEm,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Usuario{}, errors.New("usuário não encontrado")
		}
		return models.Usuario{}, err
	}

	return u, nil
}

// Listar retorna todos os usuários
func (r *UsuarioRepository) Listar() ([]models.Usuario, error) {
	query := `
	SELECT id, usuario, nome, cpf, matricula, telefone, unidade, email, cargo, papel, foto_url, criado_em, atualizado_em
	FROM usuarios
	ORDER BY nome
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []models.Usuario
	for rows.Next() {
		var u models.Usuario
		err := rows.Scan(
			&u.ID,
			&u.Usuario,
			&u.Nome,
			&u.CPF,
			&u.Matricula,
			&u.Telefone,
			&u.Unidade,
			&u.Email,
			&u.Cargo,
			&u.Papel,
			&u.FotoURL,
			&u.CriadoEm,
			&u.AtualizadoEm,
		)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usuarios, nil
}

// Atualizar atualiza os dados de um usuário
func (r *UsuarioRepository) Atualizar(usuario models.Usuario) error {
	query := `
	UPDATE usuarios
	SET nome = $1, cpf = $2, matricula = $3, telefone = $4, unidade = $5, email = $6, cargo = $7, papel = $8, foto_url = $9, atualizado_em = $10
	WHERE id = $11
	`

	_, err := r.DB.Exec(
		query,
		usuario.Nome,
		usuario.CPF,
		usuario.Matricula,
		usuario.Telefone,
		usuario.Unidade,
		usuario.Email,
		usuario.Cargo,
		usuario.Papel,
		usuario.FotoURL,
		time.Now(),
		usuario.ID,
	)

	return err
}

// AtualizarSenha atualiza a senha de um usuário
func (r *UsuarioRepository) AtualizarSenha(id int, novaSenha string) error {
	// Hash da nova senha
	senhaHash, err := bcrypt.GenerateFromPassword([]byte(novaSenha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
	UPDATE usuarios
	SET senha = $1, atualizado_em = $2
	WHERE id = $3
	`

	_, err = r.DB.Exec(query, string(senhaHash), time.Now(), id)
	return err
}

// Excluir remove um usuário do banco de dados
func (r *UsuarioRepository) Excluir(id int) error {
	query := `DELETE FROM usuarios WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

// VerificarCredenciais verifica se as credenciais do usuário são válidas
func (r *UsuarioRepository) VerificarCredenciais(credenciais models.CredenciaisUsuario) (models.Usuario, error) {
	usuario, err := r.ObterPorUsuario(credenciais.Usuario)
	if err != nil {
		return models.Usuario{}, errors.New("credenciais inválidas")
	}

	// Verificar senha
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(credenciais.Senha))
	if err != nil {
		return models.Usuario{}, errors.New("credenciais inválidas")
	}

	return usuario, nil
}