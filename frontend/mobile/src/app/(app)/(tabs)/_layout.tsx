import React, { useEffect, useRef, useState } from 'react';

import { Tabs } from 'expo-router';

import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { faCalendarDays } from '@fortawesome/free-solid-svg-icons/faCalendarDays';
import { faHouse } from '@fortawesome/free-solid-svg-icons/faHouse';
import { faUser } from '@fortawesome/free-solid-svg-icons/faUser';

import { EventPreview, Text } from '@/src/app/(design-system)';
import { Box } from '@/src/app/(design-system)';
import { Icon } from '@/src/app/(design-system)/components/Icon/Icon';
import { useAppDispatch, useAppSelector } from '@/src/store/store';
import BottomSheet from '@gorhom/bottom-sheet';
import { setEventShouldPreview } from '@/src/store/slices/eventSlice';

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
    const dispatch = useAppDispatch();
    const eventPreviewRef = useRef<BottomSheet>(null);
    const clubPreviewRef = useRef<BottomSheet>(null);

    const { shouldPreview: eventShouldPreview, id: eventId } = useAppSelector((state) => state.event);

    useEffect(() => {
        dispatch(setEventShouldPreview(false));
    }, [])

    useEffect(() => {
        if (eventShouldPreview) {
            eventPreviewRef.current?.snapToIndex(0);
        }
    }, [eventShouldPreview]);


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
                    name="index"
                    options={{
                        title: 'Home',
                        headerShown: false,
                        tabBarLabel: ({ focused }) =>
                            TabBarLabel({ focused, title: 'Home' }),
                        tabBarIcon: ({ focused }) =>
                            TabBarIcon({ focused, icon: faHouse })
                    }}
                />
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
            <EventPreview eventId={eventId} ref={eventPreviewRef} />
        </>
    );
};

export default Layout;
