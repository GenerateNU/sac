import { useState } from 'react';
import { Text, View } from 'react-native';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';
import {
    useCreateClubFollowing,
    useDeleteClubFollowing
} from '@/hooks/use-user';
import { Club } from '@/types/club';

const FollowedClub = ({ club }: { club: Club }) => {
    const { user } = useAuthStore();

    const followClub = useCreateClubFollowing();
    const unfollowClub = useDeleteClubFollowing();

    const [buttonClicked, setButtonClicked] = useState('Following');

    const handlePressFollow = async () => {
        if (buttonClicked === 'Following') {
            await unfollowClub.mutateAsync({
                userID: user?.id || '',
                clubID: club.id
            });
            setButtonClicked('Follow');
        } else {
            await followClub.mutateAsync({
                userID: user?.id || '',
                clubID: club.id
            });
            setButtonClicked('Following');
        }
    };

    return (
        <View style={{ flexDirection: 'row', alignItems: 'center', gap: 8 }}>
            <View style={{ width: 50 }}>
                <View
                    style={{
                        backgroundColor: 'black',
                        height: 50,
                        width: 50,
                        borderRadius: 25
                    }}
                />
            </View>
            <View style={{ marginRight: 8, width: 200 }}>
                <Text numberOfLines={6}>{club.name}</Text>
            </View>
            <Button
                className={`h-8 rounded-lg px-7 py-[1%] ${buttonClicked === 'Following' ? 'bg-gray-500 text-black text-md' : 'bg-gray-500 text-white text-md'}`}
                onPress={handlePressFollow}
                variant={'none'}
            >
                {buttonClicked === 'Following' ? 'Following' : 'Follow'}
            </Button>
        </View>
    );
};

export { FollowedClub };
