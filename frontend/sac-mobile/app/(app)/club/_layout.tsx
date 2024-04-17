import React from 'react';

import { Stack } from 'expo-router';

const Layout = () => {
    return (
        <Stack>
            <Stack.Screen name="[id]" options={{ headerShown: false }} />
            <Stack.Screen
                name="faq/[id]"
                options={{ headerShown: false, presentation: 'modal' }}
            />
            {/* <Stack.Screen
                name="following/following"
                options={{
                    headerShown: true
                }}
            /> */}
        </Stack>
    );
};

export default Layout;
