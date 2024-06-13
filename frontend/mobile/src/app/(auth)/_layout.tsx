import React from 'react';

import { Stack } from 'expo-router';

import { Arrow } from '../(design-system)';
import { Wordmark } from './components/wordmark';

const Layout = () => {
    return (
        <Stack>
            <Stack.Screen
                name="index"
                options={{
                    headerShown: false,
                    headerTitle: '',
                    headerTransparent: true
                }}
            />
        </Stack>
    );
};

export default Layout;
