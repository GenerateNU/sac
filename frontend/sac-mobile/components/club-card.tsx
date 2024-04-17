import { Image, Text, TouchableOpacity, View } from 'react-native';

import { useAssets } from 'expo-asset';
import { router } from 'expo-router';

import { Club } from '@/types/club';

import { ClubContacts } from './club-contacts';

export interface ClubCardProps extends ViewProps {
    club: Club
}
const ClubCard = (p: ClubCardProps) => {
    const [asset, error] = useAssets(
        require('@/assets/images/placeholder_club_logo.png')
    );

    if (error) {
        console.error(error);
    }

    return (
        <TouchableOpacity onPress={() => router.push(`/club/${club.id}`)}>
            <View className="inline-flex p-5 my-2 mr-3 bg-white border-2 rounded-[22px] border-zinc-200">
                <View className="flex-row items-center gap-2">
                    {asset && (
                        <Image
                            source={{ uri: asset[0].uri }}
                            className="w-5 h-5"
                            resizeMode="contain"
                        />
                    )}

                    <Text className="text-xl font-normal leading-snug text-zinc-900">
                        Club Name
                    </Text>
                </View>
                <Text className="flex-wrap w-[254px] mt-2 text-sm font-light truncate text-zinc-900">
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
                    do eiusmod tempor incididunt ut labore et dolore magna
                    aliqua...
                </Text>
                <Text className="w-[254px] text-zinc-900 text-xs font-extralight">
                    Read more...
                </Text>
                <ClubContacts clubID={club.id} />
            </View>
        </TouchableOpacity>
    );
};

export { ClubCard };
