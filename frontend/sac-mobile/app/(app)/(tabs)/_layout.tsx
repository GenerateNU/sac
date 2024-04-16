import React, { useEffect } from 'react';

import { Tabs } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';

import { useAuthStore } from '@/hooks/use-auth';
import { useAuth } from '@/context/AuthContext';

const HomeTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="home" size={24} color={color} />
);

const ProfileTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="account" size={24} color={color} />
);

const SearchTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="magnify" size={24} color={color} />
);

const Layout = () => {
    // const { isLoggedIn, fetchUser } = useAuthStore();
    const { initialized, getUser } = useAuth();

    useEffect(() => {
        if (initialized) getUser!();
    }, [initialized]);

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
                name="search"
                options={{
                    title: 'Search',
                    headerShown: false,
                    tabBarIcon: SearchTabBarIcon
                }}
                redirect={!isLoggedIn}
            />
            <Tabs.Screen
                name="profile"
                options={{
                    title: 'Profile',
                    headerShown: false,
                    tabBarIcon: ProfileTabBarIcon
                }}
                // redirect={!isLoggedIn}
            />
        </Tabs>
    );
};

export default Layout;
