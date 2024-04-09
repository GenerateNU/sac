import React, { useEffect } from 'react';

import { Tabs } from 'expo-router';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import HomeLight from '@/assets/icons/home-light';
import SearchAltLight from '@/assets/icons/search-alt-light';
import UserLight from '@/assets/icons/user-light';
import { useAuthStore } from '@/hooks/use-auth';

const AppLayout = () => {
    const { isLoggedIn, fetchUser } = useAuthStore();

    useEffect(() => {
        if (isLoggedIn) fetchUser();
    }, [isLoggedIn, fetchUser]);

    return (
            <Tabs
                screenOptions={{
                    tabBarStyle: {
                        bottom: '2.5%',
                        left: '4%',
                        alignContent: 'center',
                        paddingTop: '7%',
                        right: 20,
                        elevation: 0,
                        borderRadius: 50,
                        position: 'absolute',
                        shadowColor: '#000000',
                        shadowOffset: {
                            width: 2,
                            height: 2
                        },
                        shadowOpacity: 0.15,
                        shadowRadius: 3.84
                    }
                }}
            >
                <Tabs.Screen
                    name="index"
                    options={{
                        title: 'Profile',
                        tabBarLabel: () => {
                            return null;
                        },
                        headerShown: false,
                        tabBarIcon: HomeLight
                    }}
                />
                <Tabs.Screen
                    name="explore"
                    options={{
                        title: 'Explore',
                        tabBarLabel: () => {
                            return null;
                        },
                        headerShown: false,
                        tabBarIcon: SearchAltLight
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
                        tabBarIcon: UserLight
                    }}
                    redirect={!isLoggedIn}
                />
            </Tabs>
    );
};

export default AppLayout;
