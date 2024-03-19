import React from 'react';

import { Stack } from 'expo-router';

const AuthLayout = () => {
    return (
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
                    headerShown: true
                }}
            />
            <Stack.Screen
                name="tags"
                options={{
                    title: 'Tags',
                    headerShown: false
                }}
            />
        </Stack>
    );
};

export default AuthLayout;
