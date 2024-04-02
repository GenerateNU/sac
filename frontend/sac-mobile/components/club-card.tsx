import { Image, Text, View } from 'react-native';

import { Club } from '@/types/club';

import ClubContacts from './club-contacts';

const ClubCard = ({ club }: { club: Club }) => {
    return (
        <View className="h-36 w-64 rounded-xl border border-gray-300">
            <View className="h-full p-4 flex flex-col justify-between">
                <View className="flex flex-row items-center gap-2">
                    {club.logo ? (
                        <Image
                            source={{ uri: club.logo }}
                            className="w-4 h-4"
                        />
                    ) : (
                        <View className="w-4 h-4 rounded-sm bg-gray-300" />
                    )}
                    <Text className="text-md font-bold">{club.name}</Text>
                </View>
                {/* TODO: Melody's Readmoreable */}
                <Text className="text-sm">{club.preview}</Text>
                <ClubContacts clubID={club.id} />
            </View>
        </View>
    );
};

export default ClubCard;
