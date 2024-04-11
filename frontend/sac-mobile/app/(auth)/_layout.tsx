import React from 'react';

import { Stack } from 'expo-router';

const AuthLayout = () => {
    return (
        <Stack initialRouteName="welcome">
            <Stack.Screen
                name="welcome"
                options={{
                    title: '',
                    headerShown: true,
                    animationTypeForReplace: 'push',
                    statusBarColor: 'dark'
                }}
            />
            <Stack.Screen name="login" options={{ headerShown: false }} />
            <Stack.Screen name="register" options={{ headerShown: false }} />
            <Stack.Screen
                name="verification"
                options={{ headerShown: false }}
            />
            <Stack.Screen
                name="user-details"
                options={{ headerShown: false }}
            />
        </Stack>
    );
};

export default AuthLayout;
