import React from 'react';
import { View } from 'react-native';

import { Stack, router } from 'expo-router';

import { Button } from '@/components/button';
import { LeftArrow } from '@/components/left-arrow';
import { Wordmark } from '@/components/wordmark';
import { useAuthStore } from '@/hooks/use-auth';

const Layout = () => {
    const { signOut } = useAuthStore();
    return (
        <Stack initialRouteName="welcome">
            <Stack.Screen
                name="welcome"
                options={{
                    title: '',
                    headerShown: true,
                    animationTypeForReplace: 'pop',
                    animation: 'slide_from_left',
                    statusBarColor: 'dark',
                    headerLeft: () => <Wordmark />,
                    headerBackground: () => (
                        <View className="h-full bg-neutral-500" />
                    )
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
            <Stack.Screen
                name="verification"
                options={{
                    title: '',
                    headerShown: true,
                    headerLeft: () => {
                        return <LeftArrow onPress={() => signOut()} />;
                    },
                    headerBackground: () => <View className="h-full bg-white" />
                }}
            />
            <Stack.Screen
                name="user-interests"
                options={{
                    headerShown: true,
                    title: '',
                    headerLeft: () => <Wordmark color={'primary'} />,
                    headerBackground: () => (
                        <View className="h-full bg-neutral-100" />
                    )
                }}
            />
            <Stack.Screen
                name="user-details"
                options={{
                    headerShown: true,
                    title: '',
                    headerLeft: () => <Wordmark color={'primary'} />,
                    headerBackground: () => (
                        <View className="h-full bg-neutral-100" />
                    )
                }}
            />
        </Stack>
    );
};

export default Layout;
