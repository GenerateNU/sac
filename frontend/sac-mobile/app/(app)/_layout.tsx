import React from 'react';
import { View } from 'react-native';

import { Stack } from 'expo-router';

import { Arrow } from '@/components/arrow';

import { ClubKebab } from './_components/club-kebab';
import { EventKabab } from './_components/event-kebab';

const Layout = () => {
    return (
        <Stack>
            <Stack.Screen name="(tabs)" options={{ headerShown: false }} />
            <Stack.Screen
                name="event"
                options={{
                    headerTitle: '',
                    headerTransparent: true,
                    headerShown: true,
                    headerTitleStyle: {
                        color: 'white'
                    },
                    headerLeft: () => <Arrow />,
                    headerRight: () => <EventKabab />
                }}
            />
            <Stack.Screen
                name="club"
                options={{
                    headerTitle: '',
                    headerTransparent: true,
                    headerShown: true,
                    headerTitleStyle: {
                        color: 'white'
                    },
                    headerLeft: () => <Arrow />,
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
                    headerLeft: () => <Arrow />
                }}
            />
            <Stack.Screen
                name="search"
                options={{
                    headerShown: false
                }}/>
        </Stack>
    );
};

export default Layout;
