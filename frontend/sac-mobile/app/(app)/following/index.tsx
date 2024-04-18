import React from 'react';
import { ActivityIndicator, FlatList, Text, View } from 'react-native';

import { FollowedClub } from '@/app/(app)/following/_components/followed-club';
import { useAuthStore } from '@/hooks/use-auth';
import { useUserFollowing } from '@/hooks/use-user';
import { Club } from '@/types/club';

const FollowedClubsPage = () => {
    const { user } = useAuthStore();

    const {
        data: following,
        isLoading,
        error
    } = useUserFollowing(user?.id ?? '');

    if (isLoading) {
        return <ActivityIndicator size={'large'} />;
    }

    if (error) {
        return <Text>Error loading followed clubs</Text>;
    }

    if (!following) {
        return <Text>No clubs followed</Text>;
    }

    const renderFollowing = ({ item }: { item: Club }) => {
        return <FollowedClub club={item} />;
    };

    return (
        <View className="flex-col items-center mt-5">
            <FlatList
                data={following}
                renderItem={renderFollowing}
                keyExtractor={(item) => item.id.toString()}
                scrollEnabled={following.length >= 7}
            />
        </View>
    );
};

export default FollowedClubsPage;
