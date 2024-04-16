import { useEffect } from 'react';
import { View } from 'react-native';
import Spinner from 'react-native-loading-spinner-overlay';

import { useFonts } from 'expo-font';
import { Slot, Stack, useRouter, useSegments } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import { StatusBar } from 'expo-status-bar';

import FontAwesome from '@expo/vector-icons/FontAwesome';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import { useAuthStore } from '@/hooks/use-auth';
import { AuthProvider, useAuth } from '@/context/AuthContext';
import * as SecureStore from 'expo-secure-store';
import { GestureHandlerRootView } from 'react-native-gesture-handler';

export {
    // Catch any errors thrown by the Layout component.
    ErrorBoundary
} from 'expo-router';

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync();

const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            refetchOnWindowFocus: false,
            retry: false,
            staleTime: 1000 * 60 * 5
        }
    }
});

const InitalLayout = () => {
    // const { isLoggedIn, fetchCache, isVerified } = useAuthStore();
    const [loaded, error] = useFonts({
        SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
        ...FontAwesome.font
    });

    const { token, initialized, verified } = useAuth();
    const router = useRouter();
    const segments = useSegments();

    useEffect(() => {
        if (error) throw error;
    }, [error]);
    useEffect(() => {
        if (loaded) SplashScreen.hideAsync();
    }, [loaded]);

    useEffect(() => {
        // SecureStore.deleteItemAsync('accessToken');
        // SecureStore.deleteItemAsync('refreshToken');
        // SecureStore.deleteItemAsync('user');

        if (!initialized) return;

        const inAuthGroup = segments[0] === '(app)';
        console.log({ token, inAuthGroup, verified, initialized });

        if (token && inAuthGroup && verified === false) {
            router.push('/(auth)/verification');
        } else if (token && !inAuthGroup) {
            router.push('/(app)/(tabs)/');
        } else if (!token && inAuthGroup) {
            router.push('/(auth)/welcome');
        }
    }, [token, initialized, verified]);

    if (!loaded || !initialized) {
        // return <ActivityIndicator size={'large'} />;
        return <Slot />;
    }


    return (
        <>
            <StatusBar style="light" />
            <Stack>
                <Stack.Screen name="(auth)" options={{ headerShown: false }} />
                <Stack.Screen name="(app)" options={{ headerShown: false }} />
            </Stack>
        </>
    )

    // useEffect(() => {
    //     const handleNavigation = () => {
    //         if (isLoggedIn === null) return fetchCache();

    //         const inApp = segments[0] === '(app)';

    //         console.log({ isLoggedIn, isVerified, inApp });

    //         if (isLoggedIn && isVerified === false) {
    //             router.push('/(auth)/verification');
    //         } else if (isLoggedIn && isVerified) {
    //             router.push('/(app)/(tabs)/');
    //         } else if (!isLoggedIn) {
    //             router.push('/(auth)/welcome');
    //         }
    //     };

    //     console.log({ isLoggedIn, isVerified });

    //     handleNavigation();

    //     // eslint-disable-next-line react-hooks/exhaustive-deps
    // }, [isLoggedIn, isVerified]);

    // if (isLoggedIn === null)
    //     return (
    //         <View className="items-center justify-center flex-1">
    //             <Spinner
    //                 visible
    //                 textContent="Welcome to SAC Mobile"
    //                 // eslint-disable-next-line react-native/no-inline-styles
    //                 textStyle={{ color: '#FFF' }}
    //             />
    //         </View>
    //     );
};

const RootLayout = () => {

    return (
        <GestureHandlerRootView style={{ flex: 1 }}>
            <AuthProvider>
                <QueryClientProvider client={queryClient}>
                    <StatusBar style="light" />
                    <InitalLayout />
                </QueryClientProvider>
            </AuthProvider>
        </GestureHandlerRootView>
    );
};

export default RootLayout;
