import React from 'react';

import { Stack } from 'expo-router';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            staleTime: 1000 * 60 * 5,
            refetchOnWindowFocus: false
        }
    }
});

const AuthLayout = () => {
    return (
        <QueryClientProvider client={queryClient}>
            <Stack initialRouteName="welcome">
                <Stack.Screen
                    name="welcome"
                    options={{
                        title: 'Wordmark',
                        headerShown: false,
                        animationTypeForReplace: 'push',
                        statusBarColor: 'dark'
                    }}
                />
                <Stack.Screen name="login" options={{ headerShown: false }} />
                <Stack.Screen
                    name="register"
                    options={{ headerShown: false }}
                />
                <Stack.Screen
                    name="verification"
                    options={{ headerShown: false }}
                />
                <Stack.Screen
                    name="user-details"
                    options={{ headerShown: false }}
                />
                <Stack.Screen
                    name="user-interests"
                    options={{ headerShown: false }}
                />
            </Stack>
        </QueryClientProvider>
    );
};

export default AuthLayout;
