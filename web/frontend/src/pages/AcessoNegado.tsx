import { Box, Typography, Button } from '@mui/material';
import { useNavigate } from 'react-router-dom';

const AcessoNegado = () => {
    const navigate = useNavigate();

    return (
        <Box
            sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'center',
                height: '100vh',
            }}
        >
            <Typography variant="h4" gutterBottom>
                Acesso Negado
            </Typography>
            <Typography variant="body1" gutterBottom>
                Você não tem permissão para acessar esta página.
            </Typography>
            <Button
                variant="contained"
                sx={{ mt: 3 }}
                onClick={() => navigate('/dashboard')}
            >
                Voltar para o Dashboard
            </Button>
        </Box>
    );
};

export default AcessoNegado;