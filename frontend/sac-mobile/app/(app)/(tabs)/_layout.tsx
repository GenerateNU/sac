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

const Layout = () => {
    const { isLoggedIn, fetchUser } = useAuthStore();

    useEffect(() => {
        if (isLoggedIn) fetchUser();
    }, [isLoggedIn, fetchUser]);

    return (
        <Tabs
            screenOptions={
                {
                    // tabBarStyle: {
                    //     bottom: '2.5%',
                    //     left: '4%',
                    //     alignContent: 'center',
                    //     paddingTop: '7%',
                    //     right: 20,
                    //     elevation: 0,
                    //     borderRadius: 50,
                    //     position: 'absolute',
                    //     shadowColor: '#000000',
                    //     shadowOffset: {
                    //         width: 2,
                    //         height: 2
                    //     },
                    //     shadowOpacity: 0.15,
                    //     shadowRadius: 3.84
                    // }
                }
            }
        >
            <Tabs.Screen
                name="index"
                options={{
                    title: '',
                    headerTitleStyle: {
                        fontSize: 30
                    },
                    tabBarLabel: () => {
                        return null;
                    },
                    headerBackground: () => {
                        return <View className="h-full bg-gray-100" />;
                    },
                    headerShown: true,
                    tabBarIcon: ({ focused }) => {
                        return focused ? <HomeSelectedIcon /> : <HomeIcon />;
                    },
                    headerLeft: () => {
                        return (
                            <Text className="pl-6 mt-4 text-2xl font-bold">
                                SAC
                            </Text>
                        );
                    }
                }}
            />
            <Tabs.Screen
                name="search"
                options={{
                    title: 'Search',
                    tabBarLabel: () => {
                        return null;
                    },
                    headerShown: false,
                    tabBarIcon: ({ focused }) => {
                        return focused ? (
                            <SearchSelectedIcon />
                        ) : (
                            <SearchIcon />
                        );
                    }
                }}
            />
            <Tabs.Screen
                name="profile"
                options={{
                    title: 'profile',
                    tabBarLabel: () => {
                        return null;
                    },
                    headerShown: false,
                    tabBarIcon: ({ focused }) => {
                        return focused ? (
                            <ProfileSelectedIcon />
                        ) : (
                            <ProfileIcon />
                        );
                    }
                }}
                redirect={!isLoggedIn}
            />
        </Tabs>
    );
};

export default Layout;
