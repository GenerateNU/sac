import { useEffect } from 'react';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import 'react-native-reanimated';

import { useFonts } from 'expo-font';
import { Stack, router } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import { StatusBar } from 'expo-status-bar';

import FontAwesome from '@expo/vector-icons/FontAwesome';
import { ThemeProvider } from '@shopify/restyle';

import { theme } from './(design-system)';
import StoreProvider from '../store/StoreProvider';
import { useAppSelector } from '../store/store';

export { ErrorBoundary } from 'expo-router';

SplashScreen.preventAutoHideAsync();

const InitalLayout = () => {
    const { accessToken } = useAppSelector((state) => state.user);

    useEffect(() => {
        if (accessToken) {
            router.push('/(app)/');
        } else {
            router.push('/(auth)/');
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

    return (
        <GestureHandlerRootView style={{ flex: 1 }}>
            <StoreProvider>
                <ThemeProvider theme={theme}>
                    <StatusBar style="light" />
                    <InitalLayout />
                </ThemeProvider>
            </StoreProvider>
        </GestureHandlerRootView>
    );
};

export default RootLayout;
