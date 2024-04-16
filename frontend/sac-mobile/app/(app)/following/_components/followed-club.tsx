import { useState } from 'react';
import { Pressable, Text, View } from 'react-native';

import { router } from 'expo-router';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';
import {
    useCreateClubFollowing,
    useDeleteClubFollowing,
    useUserFollowing
} from '@/hooks/use-user';
import { Club } from '@/types/club';

const FollowedClub = ({ club }: { club: Club }) => {
    const { user } = useAuthStore();

    const {
        data: following,
        isLoading,
        error
    } = useUserFollowing(user?.id || '');

    const followClub = useCreateClubFollowing();
    const unfollowClub = useDeleteClubFollowing();

    const [buttonClicked, setButtonClicked] = useState('Following');

    const handlePressFollow = () => {
        if (following?.some((c) => c.id === club.id)) {
            console.log('Follow Club');
            followClub.mutateAsync({
                userId: user?.id || '',
                clubId: club.id
            });
            setButtonClicked('Following');
        } else {
            console.log('Unfollow Club');
            unfollowClub.mutateAsync({
                userId: user?.id || '',
                clubId: club.id
            });
            setButtonClicked('Follow');
        }
    };

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View className="flex-row items-center w-full px-[4%] justify-between">
            <Pressable
                className="flex-row items-center"
                onPress={() => router.push(`/club/${club.id}`)}
            >
                <View className="w-16 h-16 bg-gray-400 rounded-full" />
                <Text className="pl-[4%] font-bold text-lg">{club.name}</Text>
            </Pressable>
            <Button
                className={`h-8 rounded-lg px-7 py-[1%] ${buttonClicked === 'Following' ? 'bg-gray-300 text-black text-md' : 'bg-gray-500 text-white text-md'}`}
                onPress={handlePressFollow}
            >
                {buttonClicked === 'Following' ? 'Following' : 'Follow'}
            </Button>
        </View>
    );
};

export { FollowedClub };
