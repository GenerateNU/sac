import { useEffect } from 'react';
import { usePathname, Redirect } from 'expo-router';

import FontAwesome from '@expo/vector-icons/FontAwesome';
import { useFonts } from 'expo-font';
import { Stack, router } from 'expo-router';
import { deleteItemAsync, getItemAsync } from 'expo-secure-store';
import * as SplashScreen from 'expo-splash-screen';

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
    const previousRoute = usePathname();

    useEffect(() => {
        const checkLoginStatus = async () => {
            // deleteItemAsync('accessToken');
            // deleteItemAsync('refreshToken');
            // deleteItemAsync('user');
            const accessToken = await getItemAsync('accessToken');
            const refreshToken = await getItemAsync('refreshToken');
            const savedUser = await getItemAsync('user');

            const user: User = savedUser ? JSON.parse(savedUser) : null;

            if (accessToken && refreshToken) {
                // Update the Zustand store on successful login
                login({ accessToken, refreshToken }, user);
            }

            console.log('accessToken', accessToken);
        };

        checkLoginStatus();
    }, [login]); // Only depend on login

    useEffect(() => {
        if (!isLoggedIn) {
            // Store the previous route before redirecting to the login page
            //     const currentRoute = navigationRef.getCurrentRoute();
            //     console.log('currentRoute', currentRoute);
            router.push('/(auth)/login');
            // } else if (previousRoute.current) {
            //     // Redirect back to the previous route if available
            //     router.push(previousRoute.current);
        }
    }, [isLoggedIn]);


    // Log out console to check if the isLoggedIn state is updating correctly
    console.log('isLoggedIn', isLoggedIn);

    return (
        <Stack>
            <Stack.Screen name="(app)" options={{ headerShown: false }} />
            <Stack.Screen name="(auth)" options={{ headerShown: false }} />
        </Stack>
    );
}
