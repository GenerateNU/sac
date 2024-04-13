import { Pressable, Text, View } from 'react-native';

import { Link } from 'expo-router';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';
import {
    useCreateClubFollowing,
    useDeleteClubFollowing,
    useUserFollowing
} from '@/hooks/use-user';
import { Club } from '@/types/club';

export const FollowedClub = ({ club }: { club: Club }) => {
    // const { user } = useAuthStore();
    // const userId: string = user?.id!;
    // const {
    //     data: following,
    //     isLoading,
    //     error
    // } = useUserFollowing(userId) as {
    //     data: Club[];
    //     isLoading: boolean;
    //     error: Error;
    // };

    // const followClub = useCreateClubFollowing();
    // const unfollowClub = useDeleteClubFollowing();

    // const isFollowing = following?.some((c) => c.id === club.id);

    // const handlePressFollow = () => {
    //     if (isFollowing) {
    //         unfollowClub.mutateAsync({
    //             userId: userId,
    //             clubId: club.id,
    //         });
    //     } else {
    //         followClub.mutateAsync({
    //             userId: userId,
    //             clubId: club.id,
    //         });
    //     }
    // };

    return (
        <View className="flex-row items-center w-full px-[7%] justify-between">
            <Link
                key={club.id}
                href={{
                    pathname: '/(app)/club/[id]',
                    params: { id: club.id }
                }}
                asChild
            >
                <Pressable className="flex-row items-center">
                    <View className="w-16 h-16 bg-gray-400 rounded-full" />
                    <Text className="pl-[4%] font-bold text-lg">
                        {club.name}
                    </Text>
                </Pressable>
            </Link>
            <Button
                size="follow"
                variant={
                    'follow'
                    // isFollowing ? "follow" : "unfollow"
                }
                // onPress={handlePressFollow}
            >
                {/* {isFollowing ? "Follow" : "Following"} */}
                Following
            </Button>
        </View>
    );
};
