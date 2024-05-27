import React from 'react';
import { Tabs } from 'expo-router';
import { Text } from '@/app/(design-system)';
import Home from '@/assets/svg/home.svg';
import Profile from '@/assets/svg/profile.svg';
import Calendar from '@/assets/svg/calendar.svg';
import { Box } from '@/app/(design-system)';

const Layout = () => {
    return (
        <Tabs screenOptions={{
            tabBarShowLabel: true,
            tabBarStyle: {
                borderTopWidth: 1,
                borderTopColor: '#00000066',
                height: 85,
            },
        }}>
            <Tabs.Screen
                name="index"
                options={{
                    title: 'Home',
                    headerShown: false,
                    tabBarLabel: ({ focused }) => (
                        <Text variant="caption-1" color={focused ? 'black' : 'gray'}>
                            Home
                        </Text>
                    ),
                    tabBarIcon: ({ focused }) => (
                        <Box paddingTop="m">
                            <Home fill={focused? 'black' : 'gray'}/>
                        </Box>
                    )
                }}
            />
            <Tabs.Screen
                name="calendar"
                options={{
                    title: 'Calendar',
                    headerShown: false,
                    tabBarLabel: ({ focused }) => (
                        <Text variant="caption-1" color={focused ? 'black' : 'gray'}>
                            Calendar
                        </Text>
                    ),
                    tabBarIcon: ({ focused }) => (
                        <Box paddingTop="m">
                            <Calendar fill={focused? 'black' : 'gray'}/>
                        </Box>
                    )
                }}
            />
            <Tabs.Screen
                name="profile"
                options={{
                    title: 'Profile',
                    headerShown: false,
                    tabBarLabel: ({ focused }) => (
                        <Text variant="caption-1" color={focused ? 'black' : 'gray'}>
                            Profile
                        </Text>
                    ),
                    tabBarIcon: ({ focused }) => (
                        <Box paddingTop="m">
                            <Profile fill={focused? 'black' : 'gray'}/>
                        </Box>
                    )
                }}
            />
        </Tabs>
    );
};

export default Layout;
