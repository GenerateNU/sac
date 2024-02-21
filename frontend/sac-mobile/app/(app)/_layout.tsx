import React from 'react';

import { Tabs } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';

const HomeTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="home" size={24} color={color} />
);

const AppLayout = () => {
    return (
        <Tabs>
            <Tabs.Screen
                name="index"
                options={{
                    title: 'Home',
                    headerShown: false,
                    tabBarIcon: HomeTabBarIcon
                }}
            />
        </Tabs>
    );
};

export default AppLayout;
