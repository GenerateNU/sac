import { useEffect } from 'react';

import { useFonts } from 'expo-font';
import { Stack, router } from 'expo-router';
import { getItemAsync } from 'expo-secure-store';
import * as SplashScreen from 'expo-splash-screen';

import FontAwesome from '@expo/vector-icons/FontAwesome';

import { useAuthStore } from '@/hooks/use-auth';
import { User } from '@/types/user';

export {
    // Catch any errors thrown by the Layout component.
    ErrorBoundary
} from 'expo-router';

export const unstable_settings = {
    // Ensure that reloading on `/modal` keeps a back button present.
    initialRouteName: '(app)'
};

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync();

export default function RootLayout() {
    const [loaded, error] = useFonts({
        SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
        'OpenSans-SemiBold': require('../assets/fonts/OpenSans-SemiBold.ttf'),
        ...FontAwesome.font
    });

    // Expo Router uses Error Boundaries to catch errors in the navigation tree.
    useEffect(() => {
        if (error) throw error;
    }, [error]);

    useEffect(() => {
        if (loaded) {
            SplashScreen.hideAsync();
        }
    }, [loaded]);

    if (!loaded) {
        return null;
    }

    return <RootLayoutNav />;
}

function RootLayoutNav() {
    const { isLoggedIn, login } = useAuthStore();

    useEffect(() => {
        const checkLoginStatus = async () => {
            const accessToken = await getItemAsync('accessToken');
            const refreshToken = await getItemAsync('refreshToken');
            const savedUser = await getItemAsync('user');

            const user: User = savedUser ? JSON.parse(savedUser) : null;

            if (accessToken && refreshToken) {
                login({ accessToken, refreshToken }, user);
            }
        };

        checkLoginStatus();
    }, [login]);

    useEffect(() => {
        if (isLoggedIn === null) return;

        if (!isLoggedIn) {
            router.push('/(auth)/login');
            return;
        }

        router.push('/(app)/');
    }, [isLoggedIn]);

    return (
        <Stack>
            <Stack.Screen name="(app)" options={{ headerShown: false }} />
            <Stack.Screen name="(auth)" options={{ headerShown: false }} />
        </Stack>
    );
}
