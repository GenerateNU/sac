import React from 'react';

import { Stack } from 'expo-router';

import { Arrow } from '../(design-system)';
import { Wordmark } from './components/wordmark';

const Layout = () => {
    return (
        <Stack>
            <Stack.Screen
                name="index"
                options={{
                    headerShown: false,
                    headerTitle: '',
                    headerTransparent: true
                }}
            />
            <Stack.Screen
                name="login"
                options={{
                    headerShown: true,
                    headerTitle: '',
                    headerTransparent: true,
                    headerLeft: () => <Wordmark color="white" />
                }}
            />
            <Stack.Screen
                name="signup"
                options={{
                    headerShown: true,
                    headerTitle: '',
                    headerTransparent: true,
                    headerLeft: () => <Wordmark color="white" />
                }}
            />
            <Stack.Screen
                name="email-verification"
                options={{
                    headerShown: true,
                    headerTitle: '',
                    headerTransparent: true,
                    headerLeft: () => <Arrow color="black" />
                }}
            />
            <Stack.Screen
                name="forgot-password"
                options={{
                    headerShown: true,
                    headerTitle: '',
                    headerTransparent: true,
                    headerLeft: () => <Arrow color="black" />
                }}
            />
            <Stack.Screen
                name="reset-password"
                options={{
                    headerShown: true,
                    headerTitle: '',
                    headerTransparent: true,
                    headerLeft: () => <Wordmark color="white" />
                }}
            />
            <Stack.Screen
                name="user-details"
                options={{
                    headerShown: true,
                    headerTitle: '',
                    headerTransparent: true,
                    headerLeft: () => <Arrow color="black" />
                }}
            />
        </Stack>
    );
};

export default Layout;
