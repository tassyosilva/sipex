import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import {
    Box,
    Button,
    TextField,
    Typography,
    Paper,
    Alert,
    CircularProgress,
    CssBaseline
} from '@mui/material';
import { authService } from '../services/authService';
import { CredenciaisUsuario } from '../types/usuario';
import logo from '../assets/logo.png';

const Login = () => {
    const navigate = useNavigate();
    const [credenciais, setCredenciais] = useState<CredenciaisUsuario>({ usuario: '', senha: '' });
    const [erro, setErro] = useState<string | null>(null);
    const [carregando, setCarregando] = useState(false);

    // Adicionando um efeito para aplicar estilos diretamente ao body
    useEffect(() => {
        // Salvar os estilos originais
        const originalStyle = {
            margin: document.body.style.margin,
            padding: document.body.style.padding,
            overflow: document.body.style.overflow,
            display: document.body.style.display,
            height: document.body.style.height,
            background: document.body.style.background
        };

        // Aplicar estilos para garantir centralização
        document.body.style.margin = '0';
        document.body.style.padding = '0';
        document.body.style.overflow = 'hidden';
        document.body.style.display = 'flex';
        document.body.style.height = '100vh';
        document.body.style.width = '100vw';
        document.body.style.background = '#f0f0f0';
        document.body.style.justifyContent = 'center';
        document.body.style.alignItems = 'center';

        // Limpar ao desmontar o componente
        return () => {
            document.body.style.margin = originalStyle.margin;
            document.body.style.padding = originalStyle.padding;
            document.body.style.overflow = originalStyle.overflow;
            document.body.style.display = originalStyle.display;
            document.body.style.height = originalStyle.height;
            document.body.style.background = originalStyle.background;
        };
    }, []);

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
        <>
            <CssBaseline /> {/* Reset de CSS do Material UI */}
            <Box
                sx={{
                    position: 'absolute',
                    left: '50%',
                    top: '50%',
                    transform: 'translate(-50%, -50%)',
                    width: '100%',
                    maxWidth: '400px',
                    height: 'auto',
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    justifyContent: 'center',
                    margin: 0,
                    padding: 0,
                }}
            >
                <Paper
                    elevation={6}
                    sx={{
                        padding: 4,
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                        width: '100%',
                        borderRadius: 2,
                        backgroundColor: '#ffffff',
                        boxShadow: '0 8px 24px rgba(0, 0, 0, 0.12)',
                    }}
                >
                    <Box
                        component="img"
                        src={logo}
                        alt="SIPEX Logo"
                        sx={{
                            width: 120,
                            height: 'auto',
                            mb: 3,
                            objectFit: 'contain'
                        }}
                    />

                    <Typography
                        component="h1"
                        variant="h5"
                        sx={{
                            mb: 1,
                            fontWeight: 'bold',
                            color: '#000000',
                            textAlign: 'center'
                        }}
                    >
                        SIPEX
                    </Typography>

                    <Typography
                        variant="subtitle1"
                        sx={{
                            mb: 3,
                            color: '#333333',
                            textAlign: 'center'
                        }}
                    >
                        Sistema Integrado de Perícias e Exames
                    </Typography>

                    {erro && (
                        <Alert severity="error" sx={{ width: '100%', mb: 2 }}>
                            {erro}
                        </Alert>
                    )}

                    <Box component="form" onSubmit={handleSubmit} sx={{ width: '100%' }}>
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
                            variant="outlined"
                            sx={{ mb: 2 }}
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
                            variant="outlined"
                            sx={{ mb: 3 }}
                        />
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            sx={{
                                mt: 2,
                                mb: 2,
                                py: 1.5,
                                backgroundColor: '#D4AF37',
                                color: '#000000',
                                fontWeight: 'bold',
                                '&:hover': {
                                    backgroundColor: '#C5A028',
                                }
                            }}
                            disabled={carregando}
                        >
                            {carregando ? <CircularProgress size={24} /> : 'Entrar'}
                        </Button>
                    </Box>

                    <Typography variant="body2" color="text.secondary" align="center" sx={{ mt: 2 }}>
                        © {new Date().getFullYear()} SIPEX - Todos os direitos reservados
                    </Typography>
                </Paper>
            </Box>
        </>
    );
};

export default Login;