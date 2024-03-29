import React from 'react';

import { Stack, router } from 'expo-router';
import { Button } from '@/components/button';

const AuthLayout = () => {

    return (
        <Stack initialRouteName="welcome">
            <Stack.Screen
                name="welcome"
                options={{
                    title: 'Welcome',
                    headerShown: false,
                }}
            />
            <Stack.Screen
                name="login"
                options={{
                    title: 'Login',
                    headerShown: false,
                }}
            />
            <Stack.Screen
                name="register"
                options={{
                    title: 'Register',
                    headerShown: false,
                }}
            />
            <Stack.Screen
                name="verification"
                options={{
                    title: 'Verification Code',
                    headerShown: false,
                }}
            />
            <Stack.Screen
                name="user-details"
                options={{
                    title: 'User Details',
                    headerShown: false,
                }}
            />
        </Stack>
    );
};

export default AuthLayout;
