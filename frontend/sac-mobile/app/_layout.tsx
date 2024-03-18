import { useEffect } from 'react';

import { useFonts } from 'expo-font';
import { Slot, Stack, useRouter, useSegments } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';

import FontAwesome from '@expo/vector-icons/FontAwesome';

import { ClerkProvider, useAuth } from '@clerk/clerk-expo';
import * as SecureStore from 'expo-secure-store';

export {
    // Catch any errors thrown by the Layout component.
    ErrorBoundary
} from 'expo-router';

const CLERK_PUBLISHABLE_KEY = process.env.EXPO_PUBLIC_CLERK_PUBLISHABLE_KEY

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync();

const tokenCache = {
    async getToken(key: string) {
        try {
            return SecureStore.getItemAsync(key);
        } catch (error) {
            console.error('[RootLayoutNav] Error retrieving token:', error);
            return null;
        }
    },
    async saveToken(key: string, value: string) {
        try {
            return SecureStore.setItemAsync(key, value);
        } catch (error) {
            console.error('[RootLayoutNav] Error setting token:', error);
        }
    }
};

const InitalLayout = () => {
    const { isLoaded, isSignedIn } = useAuth();
    const router = useRouter();
    const segments = useSegments();

    useEffect(() => {
        if (!isLoaded) return;

        const inApp = segments[0] === "(app)";

        if (isSignedIn && !inApp) {
            router.push("/(app)/");
        } else if (!isSignedIn) {
            router.push("/(auth)/login");
        }

        console.log({ isSignedIn, inApp });
    }, [isSignedIn]);


    return <Slot />;
}

const RootLayout = () => {
    const [loaded, error] = useFonts({
        SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
        ...FontAwesome.font
    });

    useEffect(() => { if (error) throw error }, [error]);
    useEffect(() => { if (loaded) SplashScreen.hideAsync() }, [loaded]);

    if (!loaded) return null;

    return (
        <ClerkProvider publishableKey={CLERK_PUBLISHABLE_KEY!} tokenCache={tokenCache}>
            <InitalLayout />
        </ClerkProvider>
    );
}

export default RootLayout;