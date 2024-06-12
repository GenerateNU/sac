import { useEffect } from 'react';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import 'react-native-reanimated';

import { useFonts } from 'expo-font';
import { Stack, router } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import { StatusBar } from 'expo-status-bar';

import FontAwesome from '@expo/vector-icons/FontAwesome';
import { ThemeProvider } from '@shopify/restyle';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import StoreProvider from '../store/StoreProvider';
import { useAppSelector } from '../store/store';
import { theme } from './(design-system)';

export { ErrorBoundary } from 'expo-router';

SplashScreen.preventAutoHideAsync();

const InitalLayout = () => {
    const { accessToken } = useAppSelector((state) => state.user);

    useEffect(() => {
        if (!accessToken) {
            router.push('/(app)/');
        } else {
            router.push('/(app)/profile');
        }
    }, [accessToken]);

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

    const queryClient = new QueryClient();

    return (
        <GestureHandlerRootView style={{ flex: 1 }}>
            <StoreProvider>
                <QueryClientProvider client={queryClient}>
                    <ThemeProvider theme={theme}>
                        <StatusBar style="light" />
                        <InitalLayout />
                    </ThemeProvider>
                </QueryClientProvider>
            </StoreProvider>
        </GestureHandlerRootView>
    );
};

export default RootLayout;
