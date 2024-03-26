import React from 'react';

import { Tabs } from 'expo-router';

import HomeLight from '@/assets/icons/home-light';
import SearchAltLight from '@/assets/icons/search-alt-light';
import UserLight from '@/assets/icons/user-light';


const AppLayout = () => {
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
            />
        </Tabs>
    );
};

export default AppLayout;
