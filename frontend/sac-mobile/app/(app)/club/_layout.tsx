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
            <Stack.Screen
                name="faq/view-all-faqs"
                options={{ headerShown: false, presentation: 'modal' }}
            />
        </Stack>
    );
};

export default Layout;
