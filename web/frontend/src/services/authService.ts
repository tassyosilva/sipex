import api from '../config/api';
import { CredenciaisUsuario, RespostaUsuario } from '../types/usuario';

export const authService = {
    async login(credenciais: CredenciaisUsuario): Promise<RespostaUsuario> {
        try {
            const response = await api.post<RespostaUsuario>('/auth/login', credenciais);
            // Armazenar token e dados do usuário no sessionStorage
            sessionStorage.setItem('token', response.data.token);
            sessionStorage.setItem('usuario', JSON.stringify(response.data));
            return response.data;
        } catch (error) {
            // Apenas propaga o erro para o componente que chamou esta função
            // Não modifica o comportamento existente, apenas garante que o erro seja propagado
            throw error;
        }
    },

    logout(): void {
        sessionStorage.removeItem('token');
        sessionStorage.removeItem('usuario');
    },

    getUsuarioLogado(): RespostaUsuario | null {
        const usuario = sessionStorage.getItem('usuario');
        return usuario ? JSON.parse(usuario) : null;
    },

    isAuthenticated(): boolean {
        return !!sessionStorage.getItem('token');
    },

    hasRole(role: string): boolean {
        const usuario = this.getUsuarioLogado();
        return usuario ? usuario.papel === role : false;
    }
};

export default authService;