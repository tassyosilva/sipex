import { Navigate, Outlet } from 'react-router-dom';
import { authService } from '../services/authService';

interface PrivateRouteProps {
    allowedRoles?: string[];
}

const PrivateRoute: React.FC<PrivateRouteProps> = ({ allowedRoles }) => {
    const isAuthenticated = authService.isAuthenticated();

    // Se não estiver autenticado, redirecionar para login
    if (!isAuthenticated) {
        return <Navigate to="/login" replace />;
    }

    // Se houver restrição de papel e o usuário não tiver o papel necessário
    if (allowedRoles && allowedRoles.length > 0) {
        const usuario = authService.getUsuarioLogado();
        if (!usuario || !allowedRoles.includes(usuario.papel)) {
            return <Navigate to="/acesso-negado" replace />;
        }
    }

    // Se estiver autenticado e tiver as permissões necessárias
    return <Outlet />;
};

export default PrivateRoute;
