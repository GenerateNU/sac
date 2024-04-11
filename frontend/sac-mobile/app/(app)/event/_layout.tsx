import React from 'react';
import { Platform, View } from 'react-native';

import { Stack } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';
import { MenuView } from '@react-native-menu/menu';

import { LeftArrow } from '@/components/left-arrow';

const EventLayout = () => {
    return (
        <Stack>
            <Stack.Screen
                name="[id]"
                options={{
                    animationTypeForReplace: 'push',
                    animation: 'slide_from_right',
                    headerShown: true,
                    headerTitle: 'Event',
                    headerTitleAlign: 'center',
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
        </Stack>
    );
};

export default EventLayout;
