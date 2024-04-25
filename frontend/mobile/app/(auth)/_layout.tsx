import React from 'react';
import { View } from 'react-native';

import { Stack, router } from 'expo-router';

import { Arrow } from '@/components/arrow';
import { Button } from '@/components/button';
import { Wordmark } from '@/components/wordmark';
import { useAuthStore } from '@/hooks/use-auth';

const Layout = () => {
    const { onLogout } = useAuthStore();
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
                    headerStyle: { backgroundColor: 'transparent' },
                    headerLeft: () => <Wordmark color={'primary'} />,
                    headerBackground: () => <View className="h-full" />
                }}
            />
            <Stack.Screen
                name="login"
                options={{
                    title: '',
                    headerShown: true,
                    headerLeft: () => <Wordmark color={'primary'} />,
                    headerBackground: () => (
                        <View className="h-full bg-neutral-300" />
                    ),
                    animationTypeForReplace: 'push',
                    animation: 'slide_from_right'
                }}
            />
            <Stack.Screen
                name="register"
                options={{
                    headerStyle: { backgroundColor: 'transparent' },
                    title: '',
                    headerShown: true,
                    headerLeft: () => <Wordmark color={'primary'} />,
                    headerRight: () => (
                        <Button
                            variant={'secondary'}
                            onPress={() => router.replace('/login')}
                        >
                            Login
                        </Button>
                    ),
                    headerBackground: () => <View className="h-full" />,
                    animationTypeForReplace: 'pop',
                    animation: 'slide_from_left'
                }}
            />
            <Stack.Screen
                name="verification"
                options={{
                    headerTransparent: true,
                    title: '',
                    headerShown: true,
                    headerLeft: () => {
                        return <Arrow onPress={() => onLogout()} />;
                    },
                    headerBackground: () => <View className="h-full" />
                }}
            />
            <Stack.Screen
                name="user-interests"
                options={{
                    headerTransparent: true,
                    headerShown: true,
                    title: '',
                    headerLeft: () => <Wordmark color={'primary'} />,
                    headerBackground: () => <View className="h-full" />,
                    headerRight: () => (
                        <Arrow
                            direction={'right'}
                            onPress={() => router.replace('/(app)/(tabs)/')}
                        />
                    )
                }}
            />
            <Stack.Screen
                name="user-details"
                options={{
                    headerShown: true,
                    title: '',
                    headerLeft: () => <Wordmark color={'primary'} />,
                    headerBackground: () => <View className="h-full " />,
                    headerRight: () => (
                        <Arrow
                            direction={'right'}
                            onPress={() => router.replace('/(app)/(tabs)/')}
                        />
                    )
                }}
            />
        </Stack>
    );
};

export default Layout;
