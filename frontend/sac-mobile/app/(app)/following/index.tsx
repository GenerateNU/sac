import React from 'react';
import { ActivityIndicator, FlatList, Text, View } from 'react-native';

import { FollowedClub } from '@/app/(app)/following/_components/followed-club';
import { useAuthStore } from '@/hooks/use-auth';
import { useUserFollowing } from '@/hooks/use-user';
import { Club } from '@/types/club';

const FollowedClubsPage = () => {
    const { user } = useAuthStore();
    const userId: string = user?.id!;
    const {
        data: following,
        isLoading,
        error
    } = useUserFollowing(userId) as {
        data: Club[];
        isLoading: boolean;
        error: Error;
    };

    if (isLoading) {
        return <ActivityIndicator size={'large'} />;
    }

    if (error) {
        return <Text>Loading...</Text>;
    }

    const renderFollowing = ({ item }: { item: Club }) => {
        return <FollowedClub club={item} />;
    };

    const renderSeparator = () => {
        return (
            <View style={{ borderBottomWidth: 1, borderBottomColor: 'gray' }} />
        );
    };

    return (
        <View className="flex-col items-center mt-5">
            <FlatList
                data={following}
                renderItem={renderFollowing}
                keyExtractor={(item) => item.id.toString()}
                ItemSeparatorComponent={renderSeparator}
                scrollEnabled={following.length >= 7}
            />
        </View>
    );
};

export default FollowedClubsPage;
