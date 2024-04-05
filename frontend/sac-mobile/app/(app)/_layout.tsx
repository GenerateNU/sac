import { Tabs } from 'expo-router';
import React, { useEffect } from 'react';

import { MaterialCommunityIcons } from '@expo/vector-icons';

import { useAuthStore } from '@/hooks/use-auth';

const HomeTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="home" size={24} color={color} />
);

const EventsTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="calendar" size={24} color={color} />
);

const AppLayout = () => {
    const { isLoggedIn, fetchUser } = useAuthStore();

    useEffect(() => {
        if (isLoggedIn) fetchUser();
    }, [isLoggedIn, fetchUser]);

    return (
        <Tabs>
            <Tabs.Screen
                name="index"
                options={{
                    title: 'Home',
                    headerShown: false,
                    tabBarIcon: HomeTabBarIcon
                }}
                redirect={!isLoggedIn}
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
