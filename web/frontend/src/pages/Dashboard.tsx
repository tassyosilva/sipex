import { Box, Typography, Grid, Paper } from '@mui/material';
import { authService } from '../services/authService';

const Dashboard = () => {
    const usuario = authService.getUsuarioLogado();

    return (
        <Box>
            <Typography variant="h4" gutterBottom sx={{ color: 'primary.main', fontWeight: 'medium' }}>
                Dashboard
            </Typography>

            <Typography variant="subtitle1" gutterBottom sx={{ mb: 3 }}>
                Bem-vindo, {usuario?.nome}!
            </Typography>

            <Grid container spacing={3} sx={{ mt: 2 }}>
                <Grid item xs={12} md={6} lg={4}>
                    <Paper
                        elevation={3}
                        sx={{
                            p: 3,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 200,
                            justifyContent: 'center',
                            alignItems: 'center',
                            borderRadius: 2,
                            transition: 'transform 0.2s, box-shadow 0.2s',
                            '&:hover': {
                                transform: 'translateY(-5px)',
                                boxShadow: '0 10px 20px rgba(0,0,0,0.1)',
                            },
                        }}
                    >
                        <Typography variant="h6" gutterBottom color="text.secondary">
                            Requisições Pendentes
                        </Typography>
                        <Typography variant="h3" color="primary.main">0</Typography>
                    </Paper>
                </Grid>

                <Grid item xs={12} md={6} lg={4}>
                    <Paper
                        elevation={3}
                        sx={{
                            p: 3,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 200,
                            justifyContent: 'center',
                            alignItems: 'center',
                            borderRadius: 2,
                            transition: 'transform 0.2s, box-shadow 0.2s',
                            '&:hover': {
                                transform: 'translateY(-5px)',
                                boxShadow: '0 10px 20px rgba(0,0,0,0.1)',
                            },
                        }}
                    >
                        <Typography variant="h6" gutterBottom color="text.secondary">
                            Laudos em Elaboração
                        </Typography>
                        <Typography variant="h3" color="primary.main">0</Typography>
                    </Paper>
                </Grid>

                <Grid item xs={12} md={6} lg={4}>
                    <Paper
                        elevation={3}
                        sx={{
                            p: 3,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 200,
                            justifyContent: 'center',
                            alignItems: 'center',
                            borderRadius: 2,
                            transition: 'transform 0.2s, box-shadow 0.2s',
                            '&:hover': {
                                transform: 'translateY(-5px)',
                                boxShadow: '0 10px 20px rgba(0,0,0,0.1)',
                            },
                        }}
                    >
                        <Typography variant="h6" gutterBottom color="text.secondary">
                            Laudos Concluídos
                        </Typography>
                        <Typography variant="h3" color="secondary.main">0</Typography>
                    </Paper>
                </Grid>
            </Grid>
        </Box>
    );
};

export default Dashboard;
