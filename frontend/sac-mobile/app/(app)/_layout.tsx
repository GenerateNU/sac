import { Tabs } from 'expo-router';
import React from 'react';

import { MaterialCommunityIcons } from '@expo/vector-icons';

const HomeTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="home" size={24} color={color} />
);

const EventsTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="calendar" size={24} color={color} />
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
            <Tabs.Screen
                name="event"
                options={{
                    title: 'Event',
                    headerShown: false,
                    tabBarIcon: EventsTabBarIcon
                }}
            />
        </Tabs>
    );
};

export default AppLayout;
