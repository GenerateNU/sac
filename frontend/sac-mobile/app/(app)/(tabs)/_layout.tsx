import React, { useEffect } from 'react';

import { Tabs } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';

import { useAuthStore } from '@/hooks/use-auth';

const HomeTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="home" size={24} color={color} />
);

const Layout = () => {
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
        </Tabs>
    );
};

export default Layout;
