export interface Usuario {
    id: number;
    usuario: string;
    nome: string;
    cpf: string;
    matricula?: string;
    telefone?: string;
    unidade?: string;
    email?: string;
    cargo?: string;
    papel: string; // admin, diretor, perito, atendente
    foto_url?: string;
    criado_em: string;
    atualizado_em: string;
}

export interface CredenciaisUsuario {
    usuario: string;
    senha: string;
}

export interface RespostaUsuario {
    id: number;
    usuario: string;
    nome: string;
    cpf: string;
    matricula?: string;
    telefone?: string;
    unidade?: string;
    email?: string;
    cargo?: string;
    papel: string;
    foto_url?: string;
    token: string;
}
