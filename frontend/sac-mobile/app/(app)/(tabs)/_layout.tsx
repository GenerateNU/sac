import React, { useEffect } from 'react';
import { Text, View } from 'react-native';

import { Link, Tabs } from 'expo-router';

import HomeLight from '@/assets/icons/home-light';
import SearchAltLight from '@/assets/icons/search-alt-light';
import UserLight from '@/assets/icons/user-light';
import { useAuthStore } from '@/hooks/use-auth';

const Layout = () => {
    const { isLoggedIn, fetchUser } = useAuthStore();

    useEffect(() => {
        if (isLoggedIn) fetchUser();
    }, [isLoggedIn, fetchUser]);

    return (
        <Tabs
            screenOptions={{
                tabBarStyle: {
                    bottom: '2.5%',
                    left: '4%',
                    alignContent: 'center',
                    paddingTop: '7%',
                    right: 20,
                    elevation: 0,
                    borderRadius: 50,
                    position: 'absolute',
                    shadowColor: '#000000',
                    shadowOffset: {
                        width: 2,
                        height: 2
                    },
                    shadowOpacity: 0.15,
                    shadowRadius: 3.84
                }
            }}
        >
            <Tabs.Screen
                name="index"
                options={{
                    title: '',
                    headerTitleStyle: {
                        fontSize: 30
                    },
                    tabBarLabel: () => {
                        return null;
                    },
                    headerBackground: () => {
                        return <View className="h-full bg-gray-100" />;
                    },
                    headerShown: true,
                    tabBarIcon: HomeLight,
                    headerRight: () => {
                        return (
                            <View className="pr-6 mt-4 ">
                                <Link href={'/(app)/following/'} asChild>
                                    <Text className="text-sm font-bold">
                                        View all
                                    </Text>
                                </Link>
                            </View>
                        );
                    },
                    headerLeft: () => {
                        return (
                            <Text className="pl-6 mt-4 text-2xl font-bold">
                                SAC
                            </Text>
                        );
                    }
                }}
            />
            <Tabs.Screen
                name="search"
                options={{
                    title: 'Search',
                    tabBarLabel: () => {
                        return null;
                    },
                    headerShown: false,
                    tabBarIcon: SearchAltLight
                }}
            />
            <Tabs.Screen
                name="profile"
                options={{
                    title: 'profile',
                    tabBarLabel: () => {
                        return null;
                    },
                    headerShown: false,
                    tabBarIcon: UserLight
                }}
                redirect={!isLoggedIn}
            />
        </Tabs>
    );
};

export default Layout;
