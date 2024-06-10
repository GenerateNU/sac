import { useEffect } from 'react';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import 'react-native-reanimated';
import { Provider } from 'react-redux';

import { useFonts } from 'expo-font';
import { Stack, router } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import { StatusBar } from 'expo-status-bar';

import FontAwesome from '@expo/vector-icons/FontAwesome';
import { ThemeProvider } from '@shopify/restyle';

import { store } from '@/store/store';

import { AuthProvider, useAuth } from './(auth)/context';
import { theme } from './(design-system)';

export { ErrorBoundary } from 'expo-router';

SplashScreen.preventAutoHideAsync();

const InitalLayout = () => {
    const { loading, authToken } = useAuth();

    useEffect(() => {
        if (!loading) {
            if (authToken) {
                router.push('/(app)/');
            } else {
                router.push('/(auth)/');
            }
        }
    }, [authToken, loading]);

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
                <ThemeProvider theme={theme}>
                    <AuthProvider>
                        <StatusBar style="light" />
                        <InitalLayout />
                    </AuthProvider>
                </ThemeProvider>
            </Provider>
        </GestureHandlerRootView>
    );
};

export default RootLayout;
