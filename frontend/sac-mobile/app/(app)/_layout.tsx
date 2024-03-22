import React from 'react';

import { Tabs } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';

const HomeTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="home" size={24} color={color} />
);

import HomeLight from '@/assets/icons/home-light';

const AppLayout = () => {
    return (
        <Tabs>
            <Tabs.Screen
                name="homepage"
                options={{
                    title: 'Home',
                    headerShown: false,
                    tabBarIcon: HomeLight
                }}
            />
            <Tabs.Screen
                name="index"
                options={{
                    title: 'Profile',
                    headerShown: false,
                    tabBarIcon: HomeTabBarIcon
                }}
            />
        </Tabs>
    );
};

export default AppLayout;
