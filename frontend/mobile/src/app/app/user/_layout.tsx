import React from 'react';

import { Stack } from 'expo-router';

const Layout = () => {
    return (
        <Stack>
            <Stack.Screen
                name="interest"
                options={{
                    headerShown: false
                }}
            />
            <Stack.Screen
                name="detail"
                options={{
                    headerShown: false
                }}
            />
            <Stack.Screen
                name="following"
                options={{
                    headerShown: false
                }}
            />
        </Stack>
    );
};

export default Layout;
