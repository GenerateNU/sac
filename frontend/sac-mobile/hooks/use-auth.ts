import { deleteItemAsync, setItemAsync } from 'expo-secure-store';

import { create } from 'zustand';

import { Tokens, User } from '@/types/user';

export type AuthStore = {
    isLoggedIn: boolean | null;
    accessToken: string | null;
    refreshToken: string | null;
    user: User | null;
    login: (tokens: Tokens, user: User) => void;
    logout: () => void;
    setTokens: (tokens: Tokens) => void;
};

export const useAuthStore = create<AuthStore>((set) => ({
    isLoggedIn: null,
    accessToken: null,
    refreshToken: null,
    user: null,
    login: (tokens: Tokens, user: User) => {
        set({ isLoggedIn: true, ...tokens, user });
        setItemAsync('accessToken', tokens.accessToken);
        setItemAsync('refreshToken', tokens.refreshToken);
        setItemAsync('user', JSON.stringify(user));
    },
    logout: () => {
        set({
            isLoggedIn: false,
            accessToken: null,
            refreshToken: null,
            user: null
        });
        deleteItemAsync('accessToken');
        deleteItemAsync('refreshToken');
        deleteItemAsync('user');
    },
    setTokens: (tokens: Tokens) => set({ ...tokens })
}));
