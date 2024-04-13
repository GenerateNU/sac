import React from 'react';
import { View } from 'react-native';

import { Stack, router } from 'expo-router';

import { Button } from '@/components/button';
import { Wordmark } from '@/components/wordmark';

const Layout = () => {
    return (
        <Stack initialRouteName="welcome">
            <Stack.Screen
                name="welcome"
                options={{
                    title: '',
                    headerShown: true,
                    animationTypeForReplace: 'pop',
                    animation: 'slide_from_left',
                    statusBarColor: 'dark'
                }}
            />
            <Stack.Screen
                name="login"
                options={{
                    title: '',
                    headerShown: true,
                    headerLeft: () => <Wordmark />,
                    headerBackground: () => (
                        <View className="h-full bg-neutral-500" />
                    ),
                    animationTypeForReplace: 'push',
                    animation: 'slide_from_right'
                }}
            />
            <Stack.Screen
                name="register"
                options={{
                    title: '',
                    headerShown: true,
                    headerLeft: () => <Wordmark />,
                    headerRight: () => (
                        <Button
                            variant={'secondary'}
                            onPress={() => router.replace('/login')}
                        >
                            Login
                        </Button>
                    ),
                    headerBackground: () => (
                        <View className="h-full bg-neutral-500" />
                    ),
                    animationTypeForReplace: 'pop',
                    animation: 'slide_from_left'
                }}
            />
            <Stack.Screen name="verification" options={{ headerShown: true }} />
            <Stack.Screen name="user-details" options={{ headerShown: true }} />
        </Stack>
    );
};

export default Layout;
