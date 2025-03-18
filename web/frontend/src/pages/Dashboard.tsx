import { Box, Typography, Grid, Paper } from '@mui/material';
import { authService } from '../services/authService';

const Dashboard = () => {
    const usuario = authService.getUsuarioLogado();

    return (
        <Box>
            <Typography variant="h4" gutterBottom>
                Dashboard
            </Typography>

            <Typography variant="subtitle1" gutterBottom>
                Bem-vindo, {usuario?.nome}!
            </Typography>

            <Grid container spacing={3} sx={{ mt: 2 }}>
                <Grid item xs={12} md={6} lg={4}>
                    <Paper
                        sx={{
                            p: 3,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 200,
                            justifyContent: 'center',
                            alignItems: 'center',
                        }}
                    >
                        <Typography variant="h6" gutterBottom>
                            Requisições Pendentes
                        </Typography>
                        <Typography variant="h3">0</Typography>
                    </Paper>
                </Grid>

                <Grid item xs={12} md={6} lg={4}>
                    <Paper
                        sx={{
                            p: 3,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 200,
                            justifyContent: 'center',
                            alignItems: 'center',
                        }}
                    >
                        <Typography variant="h6" gutterBottom>
                            Laudos em Elaboração
                        </Typography>
                        <Typography variant="h3">0</Typography>
                    </Paper>
                </Grid>

                <Grid item xs={12} md={6} lg={4}>
                    <Paper
                        sx={{
                            p: 3,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 200,
                            justifyContent: 'center',
                            alignItems: 'center',
                        }}
                    >
                        <Typography variant="h6" gutterBottom>
                            Laudos Concluídos
                        </Typography>
                        <Typography variant="h3">0</Typography>
                    </Paper>
                </Grid>
            </Grid>
        </Box>
    );
};

export default Dashboard;
