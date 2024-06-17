import React from 'react';

import { Stack } from 'expo-router';

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
