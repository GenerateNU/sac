import { create } from 'zustand';

import { Tokens, User } from '@/types/user';
import { completeVerification, loginByEmail, signUp } from '@/services/auth';
import { tokenCache } from '@/lib/utils';
import { getCurrentUser } from '@/services/user';

export type UserSignUp = {
    first_name: string;
    last_name: string;
    email: string;
    password: string;
};

export type AuthStoreState = {
    user: User | null;
    isLoggedIn: boolean | null;
    isVerified: boolean | null;
    tokens: Tokens | null;
};

export type AuthStoreActions = {
    fetchUser: () => void;
    fetchCache: () => Promise<void>;
    setTokens: (tokens: Tokens) => void;
    signIn: (email: string, password: string) => Promise<void>;
    signUp: (userObj: UserSignUp) => Promise<void>;
    completeVerification: (email: string, code: string) => Promise<void>;
    signOut: () => void;
};

export const useAuthStore = create<AuthStoreState & AuthStoreActions>((set) => ({
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
            set({ user: JSON.parse(user), tokens: { accessToken, refreshToken }, isLoggedIn: true });
        } else {
            set({ isLoggedIn: false });
        }
    },
    signIn: async (email: string, password: string) => {
        const { user, tokens } = await loginByEmail(email, password);
        if (user.is_verified === false) {
            set({ user, tokens, isLoggedIn: true, isVerified: false });
        } else {
            set({ user, tokens, isLoggedIn: true, isVerified: true });
        }
        await tokenCache.saveToken('user', JSON.stringify(user));
        await tokenCache.saveToken('accessToken', tokens.accessToken);
        await tokenCache.saveToken('refreshToken', tokens.refreshToken);
    },
    signUp: async (userObj: UserSignUp) => {
        const { user, tokens } = await signUp(userObj.first_name, userObj.last_name, userObj.email, userObj.password);
        set({ user, tokens, isLoggedIn: true, isVerified: false });
        await tokenCache.saveToken('user', JSON.stringify(user));
        await tokenCache.saveToken('accessToken', tokens.accessToken);
        await tokenCache.saveToken('refreshToken', tokens.refreshToken);
    },
    completeVerification: async (email: string, code: string) => {
        await completeVerification(email, code);
        set({ isVerified: true });
    },
    signOut: () => {
        set({ tokens: null, isLoggedIn: false, isVerified: false, user: null });
        tokenCache.deleteToken('accessToken');
        tokenCache.deleteToken('refreshToken');
    },
    fetchUser: async () => {
        const user = await getCurrentUser();
        set({ user, isLoggedIn: true, isVerified: user.is_verified });
    },

}));