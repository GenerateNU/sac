import React from 'react';
import { Platform, View } from 'react-native';

import { Stack } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';
import { MenuView } from '@react-native-menu/menu';

import { LeftArrow } from '@/components/left-arrow';

const EventDotsVertical = () => {
    return (
        <MenuView
            onPressAction={({ nativeEvent }) => {
                console.warn(JSON.stringify(nativeEvent));
            }}
            actions={[
                {
                    id: 'share',
                    title: 'Share Event',
                    image: Platform.select({
                        ios: 'square.and.arrow.up',
                        android: 'share-variant'
                    })
                },
                {
                    id: 'report',
                    title: 'Report Event',
                    image: Platform.select({
                        ios: 'person.crop.circle.badge.exclamationmark.fill',
                        android: 'person-circle-outline'
                    })
                }
            ]}
        >
            <MaterialCommunityIcons
                name="dots-vertical"
                size={24}
                color="white"
            />
        </MenuView>
    );
};

const ClubDotsVertical = () => {
    return (
        <MenuView
            onPressAction={({ nativeEvent }) => {
                console.warn(JSON.stringify(nativeEvent));
            }}
            actions={[
                {
                    id: 'share',
                    title: 'Share Club',
                    image: Platform.select({
                        ios: 'square.and.arrow.up',
                        android: 'share-variant'
                    })
                },
                {
                    id: 'report',
                    title: 'Report Club',
                    image: Platform.select({
                        ios: 'person.crop.circle.badge.exclamationmark.fill',
                        android: 'person-circle-outline'
                    })
                }
            ]}
        >
            <MaterialCommunityIcons
                name="dots-vertical"
                size={24}
                color="white"
            />
        </MenuView>
    );
};

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
                    headerRight: () => <EventDotsVertical />
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
                    headerRight: () => <ClubDotsVertical />
                }}
            />
            <Stack.Screen
                name="following/index"
                options={{
                    title: 'Followed Clubs',
                    headerTitleStyle: {
                        fontSize: 20
                    },
                    headerBackground: () => {
                        return <View className="h-full bg-gray-100" />;
                    },
                    headerLeft: () => <LeftArrow />
                }}
            />
        </Stack>
    );
};

export default Layout;
