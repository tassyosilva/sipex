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
        if (error.response && error.response.status === 401) {
            // Redirecionar para login se o token expirou ou é inválido
            sessionStorage.removeItem('token');
            sessionStorage.removeItem('usuario');
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);

export default api;
