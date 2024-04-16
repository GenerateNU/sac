import React from 'react';
import {
    FlatList,
    Pressable,
    Text,
    TouchableOpacity,
    View
} from 'react-native';

import { Stack, router } from 'expo-router';

import { useUserFollowing } from '@/hooks/use-user';
import { Club } from '@/types/club';

const ViewAll = ({ isFollowingAny }: { isFollowingAny: boolean }) => {
    return (
        <>
            {isFollowingAny && (
                <View className="pr-6 mt-4 ">
                    <TouchableOpacity
                        onPress={() => router.push('/(app)/following/')}
                        className="flex-row items-center"
                    >
                        <Text className="text-sm font-bold">View all</Text>
                    </TouchableOpacity>
                </View>
            )}
        </>
    );
};

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
            <Pressable
                className="flex-col items-center w-20 mr-5"
                key={item.id}
                onPress={() => router.push(`/club/${item.id}`)}
            >
                <View className="w-20 h-20 bg-gray-300 rounded-full" />
                <Text
                    numberOfLines={1}
                    ellipsizeMode="tail"
                    className="flex-wrap pt-2 text-xs"
                >
                    {item.name}
                </Text>
            </Pressable>
        );
    };

    if (!userFollowingData) {
        return null;
    }

    return (
        <>
            <Stack.Screen
                options={{
                    title: '',
                    headerRight: () => {
                        return (
                            <ViewAll
                                isFollowingAny={userFollowingData.length > 0}
                            />
                        );
                    }
                }}
            />
            <View className="flex-row">
                <View className="pl-7" />
                <FlatList
                    horizontal
                    data={userFollowingData}
                    renderItem={renderFollowing}
                    keyExtractor={(item) => item.id.toString()}
                />
            </View>
        </>
    );
};

export { FollowingHeader };
