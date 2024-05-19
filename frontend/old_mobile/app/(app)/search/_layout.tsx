import React from 'react';

import { Stack } from 'expo-router';

// @refresh reset
const SearchLayout = () => {
    return (
        <Stack screenOptions={{ headerShown: false }}>
            <Stack.Screen
                name="results/[query]"
                options={{ headerShown: false }}
            />
            <Stack.Screen name="recents" options={{ headerShown: false }} />
            <Stack.Screen
                name="results/expanded/[query]"
                options={{ headerShown: false }}
            />
        </Stack>
    );
};

export default SearchLayout;
