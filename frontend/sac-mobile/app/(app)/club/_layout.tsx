import React from 'react';

import { Stack, Tabs } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';

const HomeTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="home" size={24} color={color} />
);

const ClubTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="car" size={24} color={color} />
);

const ClubLayout = () => {
    return (
        <Stack>
            <Stack.Screen
                name="[id]"
                options={{
                    title: 'Club',
                    headerShown: false,
                }}
            />
        </Stack>
    );
};

export default ClubLayout;
