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
        </Stack>
    );
};

export default Layout;
