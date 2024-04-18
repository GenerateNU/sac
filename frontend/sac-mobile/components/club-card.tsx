import { Text, TouchableOpacity, View } from 'react-native';

import { router } from 'expo-router';

import { Club } from '@/types/club';
import { Button } from './button';

const ClubCard = ({ club }: { club: Club }) => {


    return (
        <TouchableOpacity 
            className='w-auto'
            onPress={() => router.push(`/club/${club.id}`)}
        >
            <View className="flex flex-row justify-between p-4 bg-white border-[1px] rounded-xl border-zinc-200">
                <View className="flex-row items-center gap-2">
                    <View className="h-10 w-10 border-2 border-slate-200 bg-slate-50 rounded-full"/>
                    <View>
                        <Text className="text-base font-normal leading-snug text-zinc-900">
                            {club.name}
                        </Text>
                        <Text className="text-sm font-normal leading-snug text-zinc-900">
                            112 Followers
                        </Text>
                    </View>
                </View>
                <Button>
                    Follow
                </Button>
            </View>
        </TouchableOpacity>
    );
};

export { ClubCard };
