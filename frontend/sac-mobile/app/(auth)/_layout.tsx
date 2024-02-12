import React from 'react';

import { Stack } from 'expo-router';
import { SafeAreaView } from 'react-native';

const AuthLayout = () => {
    return (
        <Stack initialRouteName="login">
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
        </Stack>
    );
};

export default AuthLayout;
