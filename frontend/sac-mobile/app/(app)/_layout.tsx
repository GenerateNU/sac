import React from 'react';
import { View } from 'react-native';

import { Stack } from 'expo-router';

import { LeftArrow } from '@/components/left-arrow';

import { ClubKebab } from './_components/club-kebab';
import { EventKabab } from './_components/event-kebab';

const Layout = () => {
    return (
        <Stack>
            <Stack.Screen name="(tabs)" options={{ headerShown: false }} />
            <Stack.Screen
                name="event/[id]"
                options={{
                    headerShown: true,
                    headerTitleStyle: {
                        color: 'white'
                    },
                    headerBackground: () => (
                        <View className="h-full bg-gray-600" />
                    ),
                    headerLeft: () => <LeftArrow />,
                    headerRight: () => <EventKabab />
                }}
            />
            <Stack.Screen
                name="club"
                options={{
                    animationTypeForReplace: 'push',
                    animation: 'slide_from_right',
                    headerShown: true,
                    headerTitleStyle: {
                        color: 'white'
                    },
                    headerBackground: () => (
                        <View className="h-full bg-gray-600" />
                    ),
                    headerLeft: () => <LeftArrow />,
                    headerRight: () => <ClubKebab />
                }}
            />
            <Stack.Screen
                name="following/index"
                options={{
                    title: 'Followed Clubs',
                    headerTitleStyle: {
                        fontSize: 20
                    },
                    headerBackground: () => (
                        <View className="h-full bg-gray-100" />
                    ),
                    headerLeft: () => <LeftArrow />
                }}
            />
        </Stack>
    );
};

export default Layout;
