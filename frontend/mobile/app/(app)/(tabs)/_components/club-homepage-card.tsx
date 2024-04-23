import React from 'react';
import { Pressable, Text, View } from 'react-native';

import { router } from 'expo-router';

import EventDemoDay from '@/assets/images/svg/event-demo-day.svg';
import { useClubTags } from '@/hooks/use-club';
import { pluralize } from '@/lib/utils';
import { Club } from '@/types/club';

import { HomePageTagList } from './event-homepage-card';

const ClubHomePageCard = ({
    club,
    partial
}: {
    club: Club;
    partial?: boolean;
}) => {
    return (
        <Pressable
            className="flex-row items-center px-4 py-5 bg-white shadow-sm rounded-xl"
            style={{ height: 'auto' }}
            onPress={() => router.push(`/club/${club.id}`)}
        >
            <View className="flex-col items-start">
                <View className="flex-row items-center">
                    <EventDemoDay
                        width={64}
                        height={64}
                        className="rounded-xl"
                    />
                    <View className="flex-col ml-5">
                        <Text
                            style={{
                                maxWidth: 180 // Adjust the value as needed
                            }}
                            numberOfLines={2}
                            className="mb-1 overflow-hidden text-lg font-light leading-5"
                        >
                            {club.name}
                        </Text>
                        <View className="flex-row items-center gap-x-2">
                            <Text className="text-xs font-light text-gray-500">
                                {club.num_members}{' '}
                                {pluralize(
                                    club.num_members,
                                    'member',
                                    'members'
                                )}
                            </Text>
                        </View>
                    </View>
                </View>
                {!partial && (
                    <HomePageTagList
                        id={club.id}
                        tagQueryResult={useClubTags}
                    />
                )}
            </View>
        </Pressable>
    );
};

export { ClubHomePageCard };
