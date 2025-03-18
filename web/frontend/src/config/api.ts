import axios from 'axios';

// Determina a URL base da API com base no ambiente
const baseURL = import.meta.env.DEV
    ? 'http://localhost:8080/api'
    : '/api';

const api = axios.create({
    baseURL,
    headers: {
        'Content-Type': 'application/json',
    },
});

// Interceptor para adicionar o token JWT a todas as requisições
api.interceptors.request.use(
    (config) => {
        const token = sessionStorage.getItem('token');
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => Promise.reject(error)
);

// Interceptor para tratar erros de autenticação
api.interceptors.response.use(
    (response) => response,
    (error) => {
        // Verificar se é uma tentativa de login (não redirecionar neste caso)
        const isLoginRequest = error.config &&
            error.config.url &&
            error.config.url.includes('/auth/login');

        // Apenas redirecionar para login se NÃO for uma tentativa de login
        // e o status for 401 (não autorizado)
        if (error.response &&
            error.response.status === 401 &&
            !isLoginRequest) {
            // Limpar dados de autenticação
            sessionStorage.removeItem('token');
            sessionStorage.removeItem('usuario');

            // Redirecionar para login se o token expirou ou é inválido
            // e não estivermos já na página de login
            if (window.location.pathname !== '/login') {
                window.location.href = '/login';
            }
        }

        // Sempre propagar o erro para que os componentes possam tratá-lo
        return Promise.reject(error);
    }
);

export default api;