import React from 'react';
import { Tabs } from 'expo-router';
import { Text } from '@/app/(design-system)';
import Home from '@/assets/svg/home.svg';
import HomeSelected from '@/assets/svg/home-selected.svg';
import Profile from '@/assets/svg/profile.svg';
import ProfileSelected from '@/assets/svg/profile-selected.svg';
import Calendar from '@/assets/svg/calendar.svg';
import CalendarSelected from '@/assets/svg/calendar-selected.svg';
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
                    tabBarIcon: ({ focused }) =>
                        focused ? (
                            <Box paddingTop="m">
                                <HomeSelected />
                            </Box>
                        ) : (
                            <Box paddingTop="m">
                                <Home />
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
                    tabBarIcon: ({ focused }) =>
                        focused ? (
                            <Box paddingTop="m">
                                <CalendarSelected />
                            </Box>
                        ) : (
                            <Box paddingTop="m">
                                <Calendar />
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
                    tabBarIcon: ({ focused }) =>
                        focused ? (
                            <Box paddingTop="m">
                                <ProfileSelected />
                            </Box>
                        ) : (
                            <Box paddingTop="m">
                                <Profile />
                            </Box>
                        )
                }}
            />
        </Tabs>
    );
};

export default Layout;
