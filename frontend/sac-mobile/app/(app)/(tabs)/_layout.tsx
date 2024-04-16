import React, { useEffect } from 'react';
import { Text, View } from 'react-native';

import { Link, Tabs, useSegments } from 'expo-router';

import HomeSelectedIcon from '@/assets/images/svg/home-selected.svg';
import HomeIcon from '@/assets/images/svg/home.svg';
import ProfileSelectedIcon from '@/assets/images/svg/profile-selected.svg';
import ProfileIcon from '@/assets/images/svg/profile.svg';
import SearchSelectedIcon from '@/assets/images/svg/search-selected.svg';
import SearchIcon from '@/assets/images/svg/search.svg';
import { useAuthStore } from '@/hooks/use-auth';
import { Wordmark } from '@/components/wordmark';

const Layout = () => {
    const { isLoggedIn, fetchUser } = useAuthStore();

    useEffect(() => {
        if (isLoggedIn) fetchUser();
    }, [isLoggedIn, fetchUser]);

    return (
        <Tabs screenOptions={{ tabBarShowLabel: false }}>
            <Tabs.Screen
                name="index"
                options={{
                    title: '',
                    headerBackground: () => (
                        <View className="h-full bg-gray-100" />
                    ),
                    headerLeft: () => (
                        <View className='pl-6'>
                            <Wordmark color='primary' />
                        </View>
                    ),
                    headerShown: true,
                    tabBarIcon: ({ focused }) =>
                        focused ? <HomeSelectedIcon /> : <HomeIcon />
                }}
            />
            <Tabs.Screen
                name="search"
                options={{
                    title: 'search',
                    headerShown: false,
                    tabBarIcon: ({ focused }) =>
                        focused ? <SearchSelectedIcon /> : <SearchIcon />
                }}
            />
            <Tabs.Screen
                name="profile"
                options={{
                    title: 'profile',
                    headerShown: false,
                    tabBarIcon: ({ focused }) =>
                        focused ? <ProfileSelectedIcon /> : <ProfileIcon />
                }}
                redirect={!isLoggedIn}
            />
        </Tabs>
    );
};

export default Layout;
