import React from 'react';

import { Stack } from 'expo-router';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            staleTime: 1000 * 60 * 5,
            refetchOnWindowFocus: false,
        },
    },
});

const AuthLayout = () => {
    return (
        <QueryClientProvider client={queryClient}>
        <Stack initialRouteName="welcome">
            <Stack.Screen
                name="welcome"
                options={{
                    title: 'Welcome',
                    headerShown: false
                }}
            />
            <Stack.Screen
                name="login"
                options={{
                    title: 'Login',
                    headerShown: false
                }}
            />
            <Stack.Screen
                name="register"
                options={{
                    title: 'Register',
                    headerShown: false
                }}
            />
            <Stack.Screen
                name="verification"
                options={{
                    title: 'Verification',
                    headerShown: false
                }}
            />
            <Stack.Screen
                name="user-details"
                options={{
                    title: 'User Details',
                    headerShown: false
                }}
            />
            <Stack.Screen
                name="user-interests"
                options={{
                    title: 'User Interests',
                    headerShown: false
                }}
            />
        </Stack>
        </QueryClientProvider>
    );
};

export default AuthLayout;
