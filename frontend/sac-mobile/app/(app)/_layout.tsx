import React from 'react';

import { MaterialCommunityIcons } from '@expo/vector-icons';
import { Tabs } from 'expo-router';

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
