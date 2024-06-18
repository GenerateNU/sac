import { useEffect } from 'react';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import 'react-native-reanimated';

import { useFonts } from 'expo-font';
import { Stack, router } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import { StatusBar } from 'expo-status-bar';

import FontAwesome from '@expo/vector-icons/FontAwesome';
import { ThemeProvider } from '@shopify/restyle';

import usePreview from '../hooks/usePreview';
import StoreProvider from '../store/StoreProvider';
import { useAppSelector } from '../store/store';
import { ClubPreview, EventPreview, theme } from './design-system';

export { ErrorBoundary } from 'expo-router';

SplashScreen.preventAutoHideAsync();

const InitalLayout = () => {
    const { accessToken, loggedIn } = useAppSelector((state) => state.user);
    const {
        eventPreviewRef,
        eventId,
        eventApiError,
        eventApiLoading,
        clubPreviewRef,
        clubId,
        clubApiError,
        clubApiLoading
    } = usePreview();

    useEffect(() => {
        if (accessToken && loggedIn) {
            router.push('/app/');
        } else {
            router.push('/auth/');
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [loggedIn]);

    return (
        <>
            <StatusBar style="light" />
            <Stack>
                <Stack.Screen name="auth" options={{ headerShown: false }} />
                <Stack.Screen name="app" options={{ headerShown: false }} />
            </Stack>
            <EventPreview
                ref={eventPreviewRef}
                eventId={eventId}
                isLoading={eventApiLoading}
                error={eventApiError}
            />
            <ClubPreview
                ref={clubPreviewRef}
                clubId={clubId}
                isLoading={clubApiLoading}
                error={clubApiError}
            />
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
