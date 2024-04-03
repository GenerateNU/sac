import { useEffect } from 'react';
import { View } from 'react-native';
import Spinner from 'react-native-loading-spinner-overlay';

import { useFonts } from 'expo-font';
import { Slot, useRouter, useSegments } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';

import FontAwesome from '@expo/vector-icons/FontAwesome';

import { useAuthStore } from '@/hooks/use-auth';

export {
    // Catch any errors thrown by the Layout component.
    ErrorBoundary
} from 'expo-router';

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync();

const InitalLayout = () => {
    const { isLoggedIn, fetchCache, isVerified } = useAuthStore();
    const router = useRouter();
    const segments = useSegments();

    useEffect(() => {
        const handleNavigation = () => {
            if (isLoggedIn === null) return;

            const inApp = segments[0] === '(app)';

            console.log({ isLoggedIn, isVerified, inApp });

            if (isLoggedIn && isVerified === false) {
                router.push('/(auth)/verification');
            } else if (isLoggedIn && isVerified) {
                router.push('/(app)/');
            } else if (!isLoggedIn) {
                router.push('/(auth)/welcome');
            }
        };

        fetchCache();
        handleNavigation();

        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [isLoggedIn, isVerified]);

    if (isLoggedIn === null)
        return (
            <View className="items-center justify-center flex-1">
                <Spinner
                    visible
                    textContent="Welcome to SAC Mobile"
                    // eslint-disable-next-line react-native/no-inline-styles
                    textStyle={{ color: '#FFF' }}
                />
            </View>
        );

    return <Slot />;
};

const RootLayout = () => {
    const [loaded, error] = useFonts({
        SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
        ...FontAwesome.font
    });

    useEffect(() => {
        if (error) throw error;
    }, [error]);
    useEffect(() => {
        if (loaded) SplashScreen.hideAsync();
    }, [loaded]);

    if (!loaded) return null;

    return <InitalLayout />;
};

export default RootLayout;
