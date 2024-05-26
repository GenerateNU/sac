import React from 'react';
import { Tabs } from 'expo-router';
import HomeIcon from '@/assets/svg/home-icon.svg';
import HomeSelectedIcon from '@/assets/svg/calendar-icon.svg';
import ProfileIcon from '@/assets/svg/calendar-icon.svg';
import ProfileSelectedIcon from '@/assets/svg/calendar-icon.svg';
import CalendarIcon from '@/assets/svg/calendar-icon.svg';
import CalendarSelectedIcon from '@/assets/svg/calendar-selected-icon.svg';
import { View } from 'react-native';
import { Text } from '@/app/(design-system)';

const Layout = () => {
    return (
        <Tabs screenOptions={{
            tabBarShowLabel: true,
            tabBarStyle: {
                borderTopWidth: 1,
                borderTopColor: '#00000066',
            },
        }}>
            <Tabs.Screen
                name="index"
                options={{
                    title: 'Home',
                    headerShown: false,
                    tabBarLabel: ({ focused }) => (
                        <Text style={{ fontSize: 12, color: focused ? 'black' : '#C3C9D0' }}>
                            Home
                        </Text>
                    ),
                    // tabBarIcon: ({ focused }) =>
                    //     focused ? (
                    //         <View>
                    //             <HomeSelectedIcon />
                    //         </View>
                    //     ) : (
                    //         <View>
                    //             <HomeIcon />
                    //         </View>
                    //     )
                }}
            />
            <Tabs.Screen
                name="calendar"
                options={{
                    title: 'Calendar',
                    headerShown: false,
                    tabBarLabel: ({ focused }) => (
                        <Text style={{ fontSize: 12, color: focused ? 'black' : '#C3C9D0' }}>
                            Calendar
                        </Text>
                    ),
                    // tabBarIcon: ({ focused }) =>
                    //     focused ? (
                    //         <View>
                    //             <CalendarSelectedIcon />
                    //         </View>
                    //     ) : (
                    //         <View>
                    //             <CalendarIcon />
                    //         </View>
                    //     )
                }}
            />
            <Tabs.Screen
                name="profile"
                options={{
                    title: 'Profile',
                    headerShown: false,
                    tabBarLabel: ({ focused }) => (
                        <Text style={{ fontSize: 12, color: focused ? 'black' : '#C3C9D0' }}>
                            Profile
                        </Text>
                    ),
                    // tabBarIcon: ({ focused }) =>
                    //     focused ? (
                    //         <View>
                    //             <ProfileSelectedIcon />
                    //         </View>
                    //     ) : (
                    //         <View>
                    //             <ProfileIcon />
                    //         </View>
                    //     )
                }}
            />
        </Tabs>
    );
};

export default Layout;
