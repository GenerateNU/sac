import React from 'react';

import { Stack } from 'expo-router';
import { MenuView } from '@react-native-menu/menu';
import { LeftArrow } from '@/components/left-arrow';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { View, Platform, StatusBar } from 'react-native';

const Layout = () => {
    return (
        <Stack>
            <Stack.Screen name="(tabs)" options={{ headerShown: false }} />
            <Stack.Screen
                name="event/[id]"
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
                    headerRight: () => {
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
                    }
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
                    headerRight: () => {
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
                    }
                }}
            />
        </Stack>
    );
};

export default Layout;
