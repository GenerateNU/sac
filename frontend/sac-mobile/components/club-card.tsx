import { Image, Text, View, ViewProps } from 'react-native';

import { useAssets } from 'expo-asset';

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
        <View {...p} className="inline-flex p-5 bg-white border-2 rounded-[22px] border-zinc-200">
            <View className="flex-row items-center gap-2">
                {asset && (
                    <Image
                        source={{ uri: asset[0].uri }}
                        className="w-10 h-10"
                        resizeMode="contain"
                    />
                )}

                <Text className="text-xl font-normal leading-snug text-zinc-900">
                    {p.club.name}
                </Text>
            </View>
            <Text className="flex-wrap w-[254px] mt-2 text-sm font-light truncate text-zinc-900">
                {p.club.description.substring(0,120)}
            </Text>
            <Text className="w-[254px] text-zinc-900 text-xs font-extralight">
                Read more...
            </Text>
            <ClubContacts clubID={p.club.id} />
        </View>
    );
};

export { ClubCard };
