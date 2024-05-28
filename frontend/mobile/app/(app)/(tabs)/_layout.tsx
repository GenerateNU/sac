import React from 'react';

import { Tabs } from 'expo-router';

import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { faCalendarDays } from '@fortawesome/free-solid-svg-icons/faCalendarDays';
import { faHouse } from '@fortawesome/free-solid-svg-icons/faHouse';
import { faUser } from '@fortawesome/free-solid-svg-icons/faUser';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Text } from '@/app/(design-system)';
import { Box, Colors } from '@/app/(design-system)';

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
        <FontAwesomeIcon
            size={20}
            color={focused ? 'black' : `${Colors.gray}`}
            icon={icon}
        />
    </Box>
);

const Layout = () => {
    return (
        <Tabs
            screenOptions={{
                tabBarShowLabel: true,
                tabBarStyle: {
                    borderTopWidth: 1,
                    borderTopColor: '#00000066',
                    height: 85
                }
            }}
        >
            <Tabs.Screen
                name="index"
                options={{
                    title: 'Home',
                    headerShown: false,
                    tabBarLabel: ({ focused }) => (
                        <TabBarLabel focused={focused} title="Home" />
                    ),
                    tabBarIcon: ({ focused }) => (
                        <TabBarIcon focused={focused} icon={faHouse} />
                    )
                }}
            />
            <Tabs.Screen
                name="calendar"
                options={{
                    title: 'Calendar',
                    headerShown: false,
                    tabBarLabel: ({ focused }) => (
                        <TabBarLabel focused={focused} title="Calendar" />
                    ),
                    tabBarIcon: ({ focused }) => (
                        <TabBarIcon focused={focused} icon={faCalendarDays} />
                    )
                }}
            />
            <Tabs.Screen
                name="profile"
                options={{
                    title: 'Profile',
                    headerShown: false,
                    tabBarLabel: ({ focused }) => (
                        <TabBarLabel focused={focused} title="Profile" />
                    ),
                    tabBarIcon: ({ focused }) => (
                        <TabBarIcon focused={focused} icon={faUser} />
                    )
                }}
            />
        </Tabs>
    );
};

export default Layout;
