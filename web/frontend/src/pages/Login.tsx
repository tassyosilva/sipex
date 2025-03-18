import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {
    Box,
    Button,
    Container,
    TextField,
    Typography,
    Paper,
    Alert,
    CircularProgress
} from '@mui/material';
import { authService } from '../services/authService';
import { CredenciaisUsuario } from '../types/usuario';

const Login = () => {
    const navigate = useNavigate();
    const [credenciais, setCredenciais] = useState<CredenciaisUsuario>({ usuario: '', senha: '' });
    const [erro, setErro] = useState<string | null>(null);
    const [carregando, setCarregando] = useState(false);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setCredenciais(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setErro(null);
        setCarregando(true);

        try {
            await authService.login(credenciais);
            navigate('/dashboard');
        } catch (error: any) {
            setErro(error.response?.data?.message || 'Erro ao realizar login. Verifique suas credenciais.');
        } finally {
            setCarregando(false);
        }
    };

    return (
        <Container component="main" maxWidth="xs">
            <Box
                sx={{
                    marginTop: 8,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <Paper
                    elevation={3}
                    sx={{
                        padding: 4,
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                        width: '100%',
                    }}
                >
                    <Typography component="h1" variant="h5" sx={{ mb: 3 }}>
                        SIPEX - Sistema de Gestão de Requisições Periciais
                    </Typography>

                    <Typography component="h2" variant="h6" sx={{ mb: 3 }}>
                        Login
                    </Typography>

                    {erro && (
                        <Alert severity="error" sx={{ width: '100%', mb: 2 }}>
                            {erro}
                        </Alert>
                    )}

                    <Box component="form" onSubmit={handleSubmit} sx={{ mt: 1, width: '100%' }}>
                        <TextField
                            margin="normal"
                            required
                            fullWidth
                            id="usuario"
                            label="Usuário"
                            name="usuario"
                            autoComplete="username"
                            autoFocus
                            value={credenciais.usuario}
                            onChange={handleChange}
                        />
                        <TextField
                            margin="normal"
                            required
                            fullWidth
                            name="senha"
                            label="Senha"
                            type="password"
                            id="senha"
                            autoComplete="current-password"
                            value={credenciais.senha}
                            onChange={handleChange}
                        />
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                            disabled={carregando}
                        >
                            {carregando ? <CircularProgress size={24} /> : 'Entrar'}
                        </Button>
                    </Box>
                </Paper>
            </Box>
        </Container>
    );
};

export default Login;
