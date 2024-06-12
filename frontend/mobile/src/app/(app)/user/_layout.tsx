import React from 'react';

import { Stack } from 'expo-router';

const Layout = () => {
    return (
        <Stack>
            <Stack.Screen name="interests" options={{ headerShown: false }} />
            <Stack.Screen
                name="details"
                options={{
                    headerTitle: '',
                    headerTransparent: true,
                    headerShown: false
                }}
            />
        </Stack>
    );
};

export default Layout;
