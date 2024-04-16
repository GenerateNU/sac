import axios from 'axios';
import { createContext, useContext, useEffect, useState } from 'react';
import * as SecureStore from 'expo-secure-store';
import { User } from '@/types/user';
import { extractTokens, tokenCache } from '@/lib/utils';
import { API_BASE_URL } from '@/lib/const';
import { Tokens } from '@/types/auth';

type AuthProps = {
    user: User | null;
    token: Tokens | null;
    initialized: boolean;
    verified: boolean;
    getUser: () => void;
    onRegister: (registerObj: UserSignUp) => Promise<any>;
    onLogin: (email: string, password: string) => Promise<any>;
    onLogout: () => Promise<void>;
    onVerify: (email: string, code: string) => Promise<any>;
};

type UserSignUp = {
    first_name: string;
    last_name: string;
    email: string;
    password: string;
};

const AuthContext = createContext<Partial<AuthProps>>({});

export function useAuth() {
    return useContext(AuthContext);
}

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
    const [verified, setVerified] = useState(false);
    const [user, setUser] = useState<User | null>(null);
    const [initialized, setInitialized] = useState(false);
    const [token, setToken] = useState<Tokens | null>(null);

    useEffect(() => {
        const loadToken = async () => {
            const storedUser = await tokenCache.getToken('user');
            const user = storedUser ? JSON.parse(storedUser) as User : null;
            const refreshToken = await tokenCache.getToken('refreshToken');
            const accessToken = await tokenCache.getToken('accessToken');

            if (accessToken && refreshToken) {
                setToken({ accessToken, refreshToken });
            }

            if (user && user.is_verified) {
                setUser(user);
                setVerified(true);
            }

            setInitialized(true);
        };
        loadToken();
    }, []);

    const getUser = async () => {
        try {
            const response = await axios.get(`${API_BASE_URL}/users/me`,
                { headers: { Authorization: `Bearer ${token}` } }
            );
            setUser(response.data);
        } catch (error: any) {
            console.error(error);
        }
    }

    const handleVerify = async (email: string, code: string) => {
        try {
            const response = await axios.post(`${API_BASE_URL}/auth/verify-email`, { email, token: code });
            const user = response.data;

            setUser(user);
            setVerified(true);

            return response;
        } catch (error: any) {
            return { error: true, msg: error.response.data.msg };
        }
    }

    const handleLogin = async (email: string, password: string) => {
        try {
            const response = await axios.post(`${API_BASE_URL}/auth/login`, { email, password });
            const tokens = extractTokens(
                response.headers['set-cookie'],
                response.headers
            );
            const user = response.data;

            setToken(tokens.accessToken);
            setUser(user);
            if (user.is_verified) {
                setVerified(true);
            }

            await tokenCache.saveToken('user', JSON.stringify(user));
            await tokenCache.saveToken('accessToken', tokens.accessToken);
            await tokenCache.saveToken('refreshToken', tokens.refreshToken);

            return response;
        } catch (error: any) {
            return { error: true, msg: error.response.data.msg };
        }
    };

    const handleRegister = async (registerObj: UserSignUp) => {
        try {
            const response = await axios.post(`${API_BASE_URL}/users/`, { ...registerObj });
            const tokens = extractTokens(
                response.headers['set-cookie'],
                response.headers
            );
            const user = response.data;

            setToken(tokens.accessToken);
            setUser(user);

            await tokenCache.saveToken('user', JSON.stringify(user));
            await tokenCache.saveToken('accessToken', tokens.accessToken);
            await tokenCache.saveToken('refreshToken', tokens.refreshToken);

            return response;
        } catch (error: any) {
            return { error: true, msg: error.response.data.msg };
        }
    };

    const handleLogout = async () => {
        // TODO: fetch call to logout
        setToken(null);
        setUser(null);
        setVerified(false);

        await tokenCache.deleteToken('user');
        await tokenCache.deleteToken('accessToken');
        await tokenCache.deleteToken('refreshToken');
    };

    const value = {
        user,
        token,
        initialized,
        verified,
        getUser: getUser,
        onRegister: handleRegister,
        onLogin: handleLogin,
        onLogout: handleLogout,
        onVerify: handleVerify,
    };

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
};