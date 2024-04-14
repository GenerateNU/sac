import React from 'react';
import {
    FlatList,
    ScrollView,
    Text,
    TouchableOpacity,
    View
} from 'react-native';

import { Link, Stack } from 'expo-router';

import { useUserFollowing } from '@/hooks/use-user';
import { Club } from '@/types/club';

const FollowingHeader = ({ id }: { id: string }) => {
    const { data: userFollowingData, isLoading, error } = useUserFollowing(id);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    const renderFollowing = ({ item }: { item: Club }) => {
        return (
            <Link
                key={item.id}
                href={{
                    pathname: '/(app)/club/[id]',
                    params: { id: item.id }
                }}
                asChild
            >
                <TouchableOpacity
                    className="flex-col w-20 mr-5 items-center"
                    key={item.id}
                >
                    <View className="w-20 h-20 bg-gray-300 rounded-full" />
                    <Text
                        numberOfLines={1}
                        ellipsizeMode="tail"
                        className="pt-2 flex-wrap text-xs"
                    >
                        {item.name}
                    </Text>
                </TouchableOpacity>
            </Link>
        );
    };
    return (
        <>
            <Stack.Screen
                options={{
                    title: '',
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
                    }
                }}
            />
            <ScrollView horizontal showsHorizontalScrollIndicator={false}>
                <View className="flex-row">
                    <View className="pl-7" />
                    <FlatList
                        horizontal
                        data={userFollowingData}
                        renderItem={renderFollowing}
                        keyExtractor={(item) => item.id.toString()}
                    />
                </View>
            </ScrollView>
        </>
    );
};

export { FollowingHeader };
