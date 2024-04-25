import { create } from 'zustand';

import { tokenCache } from '@/lib/utils';
import { completeVerification, login, logout, signUp } from '@/services/auth';
import { getCurrentUser } from '@/services/user';
import { Tokens } from '@/types/auth';
import { User } from '@/types/user';

// import { login, logout, register, verifyEmail } from '@/services/auth';

type UserSignUp = {
    first_name: string;
    last_name: string;
    email: string;
    password: string;
};

type AuthStoreState = {
    user: User | null;
    isLoggedIn: boolean | null;
    isVerified: boolean | null;
    tokens: Tokens | null;
};

type AuthStoreActions = {
    fetchUser: () => void;
    fetchCache: () => Promise<void>;
    setTokens: (tokens: Tokens) => void;
    onLogin: (email: string, password: string) => Promise<void>;
    onRegister: (userObj: UserSignUp) => Promise<void>;
    onVerify: (email: string, code: string) => Promise<void>;
    onLogout: () => void;
};

const useAuthStore = create<AuthStoreState & AuthStoreActions>((set) => ({
    user: null,
    isLoggedIn: null,
    isVerified: null,
    tokens: null,
    setTokens: (tokens: Tokens) => {
        set({ tokens });
        tokenCache.saveToken('accessToken', tokens.accessToken);
        tokenCache.saveToken('refreshToken', tokens.refreshToken);
    },
    fetchCache: async () => {
        const accessToken = await tokenCache.getToken('accessToken');
        const refreshToken = await tokenCache.getToken('refreshToken');
        const user = await tokenCache.getToken('user');

        if (accessToken && refreshToken && user) {
            set({
                user: JSON.parse(user),
                tokens: { accessToken, refreshToken },
                isLoggedIn: true
            });
        } else {
            set({ isLoggedIn: false });
        }
    },
    onLogin: async (email: string, password: string) => {
        const { user, tokens } = await login(email, password);
        if (user.is_verified === false) {
            set({ user, tokens, isLoggedIn: true, isVerified: false });
        } else {
            set({ user, tokens, isLoggedIn: true, isVerified: true });
        }
        await tokenCache.saveToken('user', JSON.stringify(user));
        await tokenCache.saveToken('accessToken', tokens.accessToken);
        await tokenCache.saveToken('refreshToken', tokens.refreshToken);
    },
    onRegister: async (userObj: UserSignUp) => {
        const { user, tokens } = await signUp(
            userObj.first_name,
            userObj.last_name,
            userObj.email,
            userObj.password
        );
        set({ user, tokens, isLoggedIn: true, isVerified: false });
        await tokenCache.saveToken('user', JSON.stringify(user));
        await tokenCache.saveToken('accessToken', tokens.accessToken);
        await tokenCache.saveToken('refreshToken', tokens.refreshToken);
    },
    onVerify: async (email: string, code: string) => {
        await completeVerification(email, code);
        set({ isVerified: true });
    },
    onLogout: async () => {
        await logout();
        set({
            tokens: null,
            isLoggedIn: false,
            isVerified: false,
            user: null
        });
        tokenCache.deleteToken('accessToken');
        tokenCache.deleteToken('refreshToken');
    },
    fetchUser: async () => {
        const user = await getCurrentUser();
        set({ user, isLoggedIn: true, isVerified: user.is_verified });
    }
}));

export { UserSignUp, AuthStoreState, AuthStoreActions, useAuthStore };
