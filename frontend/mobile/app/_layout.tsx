import { useEffect } from 'react';
import { View } from 'react-native';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import Spinner from 'react-native-loading-spinner-overlay';
import { Provider } from 'react-redux';

import { useFonts } from 'expo-font';
import { Stack, useRouter, useSegments } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import { StatusBar } from 'expo-status-bar';

import FontAwesome from '@expo/vector-icons/FontAwesome';
import { store } from '@/store/store'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import { useAuthStore } from '@/hooks/use-auth';
export {
    // Catch any errors thrown by the Layout component.
    ErrorBoundary
} from 'expo-router';

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync();

export const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            refetchOnWindowFocus: false,
            retry: false,
            staleTime: 1000 * 60 * 5
        }
    }
});

const InitalLayout = () => {
    const { isLoggedIn, fetchCache, isVerified, user } = useAuthStore();
    const router = useRouter();
    const segments = useSegments();

    useEffect(() => {
        const handleNavigation = () => {
            if (isLoggedIn === null) return fetchCache();

            const inApp = segments[0] === '(app)';

            console.log({ isLoggedIn, isVerified, inApp });

            if (isLoggedIn && isVerified === false) {
                router.push('/(auth)/verification');
            } else if (isLoggedIn && isVerified) {
                router.push('/(app)/(tabs)/');
            } else if (!isLoggedIn) {
                router.push('/(auth)/welcome');
            }
        };

        console.log({ isLoggedIn, isVerified, user });

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

    return (
        <>
            <StatusBar style="light" />
            <Stack>
                <Stack.Screen name="(auth)" options={{ headerShown: false }} />
                <Stack.Screen name="(app)" options={{ headerShown: false }} />
            </Stack>
        </>
    );
};

const RootLayout = () => {
    const [loaded, error] = useFonts({
        SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
        'OpenSans-SemiBold': require('../assets/fonts/OpenSans-SemiBold.ttf'),
        ...FontAwesome.font
    });

    useEffect(() => {
        if (error) throw error;
    }, [error]);
    useEffect(() => {
        if (loaded) SplashScreen.hideAsync();
    }, [loaded]);

    if (!loaded) return null;

    return (
        <GestureHandlerRootView style={{ flex: 1 }}>
            <Provider store={store}>
                <QueryClientProvider client={queryClient}>
                    <StatusBar style="light" />
                    <InitalLayout />
                </QueryClientProvider>
            </Provider>
        </GestureHandlerRootView>
    );
};

export default RootLayout;
