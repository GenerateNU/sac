import React, { createContext, useContext, useEffect, useState } from 'react';

import * as SecureStore from 'expo-secure-store';

/**
 * Storage interaction methods:
 */
const saveToken = async (token: string) => {
    await SecureStore.setItemAsync('authToken', token);
};

const getToken = async () => {
    return await SecureStore.getItemAsync('authToken');
};

const deleteToken = async () => {
    await SecureStore.deleteItemAsync('authToken');
};

type AuthContextProps = {
    authToken?: string | null;
    login?: (token: string) => void;
    logout?: () => void;
    loading?: boolean;
};

const AuthContext = createContext<AuthContextProps>({});

type AuthProviderProps = {
    children: React.ReactNode;
};

export const AuthProvider = ({ children }: AuthProviderProps) => {
    const [authToken, setAuthToken] = useState<string | null>(null);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        const fetchToken = async () => {
            const token = await getToken();
            setAuthToken(token);
            setLoading(false);
        };
        fetchToken();
    }, []);

    const login = async (token: string) => {
        await saveToken(token);
        setAuthToken(token);
    };

    const logout = async () => {
        await deleteToken();
        setAuthToken(null);
    };

    return (
        <AuthContext.Provider value={{ authToken, login, logout, loading }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => useContext(AuthContext);
