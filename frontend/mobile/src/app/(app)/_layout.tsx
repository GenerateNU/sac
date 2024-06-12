import React from 'react';

import { Stack } from 'expo-router';

import { Arrow } from '../(design-system)';
import { Text } from '../(design-system)';

const Layout = () => {
    return (
        <Stack>
            <Stack.Screen name="(tabs)" options={{ headerShown: false }} />
            <Stack.Screen
                name="event"
                options={{
                    headerTitle: '',
                    headerTransparent: true,
                    headerShown: false,
                    headerTitleStyle: {
                        color: 'white'
                    }
                }}
            />
            <Stack.Screen
                name="user"
                options={{
                    headerTitle: 'Profile',
                    headerTransparent: true,
                    headerShown: false
                }}
            />
        </Stack>
    );
};

export default Layout;
