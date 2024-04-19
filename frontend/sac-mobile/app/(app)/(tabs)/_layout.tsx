import React, { useEffect } from 'react';
import { View, Text } from 'react-native';

import { Tabs } from 'expo-router';

import HomeSelectedIcon from '@/assets/images/svg/home-selected.svg';
import HomeIcon from '@/assets/images/svg/home.svg';
import ProfileSelectedIcon from '@/assets/images/svg/profile-selected.svg';
import ProfileIcon from '@/assets/images/svg/profile.svg';
import NotificationIcon from '@/assets/images/svg/notification.svg';
import SearchSelectedIcon from '@/assets/images/svg/search-selected.svg';
import SearchIcon from '@/assets/images/svg/search.svg';
import { useAuthStore } from '@/hooks/use-auth';

const Layout = () => {
    const { isLoggedIn, fetchUser, user } = useAuthStore();

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
                        <View className="pl-6">
                            <Text className='h-8 text-black text-[32px] font-normal'>Hi, {user?.first_name}</Text>
                        </View>
                    ),
                    headerRight: () => (
                        <View className="pr-6">
                            <NotificationIcon />
                        </View>
                    ),
                    headerShown: true,
                    tabBarIcon: ({ focused }) =>
                        focused ? <HomeSelectedIcon /> : <HomeIcon />
                }}
                redirect={!isLoggedIn}
            />
            <Tabs.Screen
                name="search"
                options={{
                    title: 'search',
                    headerShown: false,
                    tabBarIcon: ({ focused }) =>
                        focused ? <SearchSelectedIcon /> : <SearchIcon />
                }}
                redirect={!isLoggedIn}
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
