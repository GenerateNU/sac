import React from 'react';

import { Tabs } from 'expo-router';

import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { faCalendarDays } from '@fortawesome/free-solid-svg-icons/faCalendarDays';
import { faUser } from '@fortawesome/free-solid-svg-icons/faUser';

import { Text } from '@/src/app/design-system';
import { Box } from '@/src/app/design-system';
import { Icon } from '@/src/app/design-system/components/Icon/Icon';

interface TabBarLabelProps {
    focused: boolean;
    title: string;
}

const TabBarLabel: React.FC<TabBarLabelProps> = ({ focused, title }) => (
    <Text variant="caption-1" color={focused ? 'black' : 'gray'}>
        {title}
    </Text>
);

interface TabBarIconProps {
    focused: boolean;
    icon: IconDefinition;
}

const TabBarIcon: React.FC<TabBarIconProps> = ({ focused, icon }) => (
    <Box paddingTop="m">
        <Icon size="md" color={focused ? 'black' : `gray`} icon={icon} />
    </Box>
);

const Layout = () => {
    return (
        <>
            <Tabs
                screenOptions={{
                    tabBarShowLabel: false,
                    tabBarStyle: {
                        borderTopWidth: 1,
                        borderTopColor: 'gray',
                        height: 85
                    }
                }}
                sceneContainerStyle={{
                    backgroundColor: 'white'
                }}
            >
                <Tabs.Screen
                    name="calendar"
                    options={{
                        title: 'Calendar',
                        headerShown: false,
                        tabBarLabel: ({ focused }) =>
                            TabBarLabel({ focused, title: 'Calendar' }),
                        tabBarIcon: ({ focused }) =>
                            TabBarIcon({ focused, icon: faCalendarDays })
                    }}
                />
                <Tabs.Screen
                    name="profile"
                    options={{
                        title: 'Profile',
                        headerShown: true,
                        headerTransparent: true,
                        tabBarLabel: ({ focused }) =>
                            TabBarLabel({ focused, title: 'Profile' }),
                        tabBarIcon: ({ focused }) =>
                            TabBarIcon({ focused, icon: faUser })
                    }}
                />
            </Tabs>
        </>
    );
};

export default Layout;
