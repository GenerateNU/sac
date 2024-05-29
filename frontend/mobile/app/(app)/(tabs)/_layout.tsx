import React from 'react';

import { Tabs } from 'expo-router';

const Layout = () => {
    return (
        <Tabs>
            <Tabs.Screen
                name="index"
                options={{
                    title: 'home',
                    headerShown: false
                }}
            />
            <Tabs.Screen
                name="discover"
                options={{
                    title: 'search',
                    headerShown: false
                }}
            />
            <Tabs.Screen
                name="profile"
                options={{
                    title: 'profile',
                    headerShown: false
                }}
            />
        </Tabs>
    );
};

export default Layout;
