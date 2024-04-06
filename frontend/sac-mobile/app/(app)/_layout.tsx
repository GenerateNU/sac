import { Tabs, router } from 'expo-router';
import React, { useEffect } from 'react';

import { MaterialCommunityIcons } from '@expo/vector-icons';

import { useAuthStore } from '@/hooks/use-auth';
import { Button } from '@/components/button';
import { View } from 'react-native';

const HomeTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="home" size={24} color={color} />
);

const EventsTabBarIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="calendar" size={24} color={color} />
);

const LeftButton = () => {
    return (
        <Button
            onPress={() => {
                console.log('Back Button Pressed');
                router.back();
            }}
            size={'icon'}
            variant={'icon'}
            className="flex items-center justify-center bg-gray-300 rounded-full"
        >
            <MaterialCommunityIcons
                name="chevron-left"
                size={28}
                color="rgb(75 85 99)"
            />
        </Button>
    );
}


const AppLayout = () => {
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
            <Tabs.Screen
                name="event"
                options={{
                    title: '',
                    headerShown: true,
                    headerLeft: LeftButton,
                    headerLeftContainerStyle: {
                        paddingLeft: 10
                    },
                    headerBackground: () => (<View className='h-full bg-gray-600' />),
                    tabBarIcon: EventsTabBarIcon,
                    tabBarLabel: 'Events'
                }}
            />
        </Tabs>
    );
};

export default AppLayout;
