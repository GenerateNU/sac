import React from 'react';
import { ScrollView, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { ClubHomepageCard } from '@/components/club-homepage-card';
import { EventCard } from '@/components/event-card';
import { FAQCard } from '@/components/faq-card';
import { useAuthStore } from '@/hooks/use-auth';
import { useUserFollowing } from '@/hooks/use-user-following';
import { HomepageList } from '@/lib/const';
import { Club } from '@/types/club';

const Homepage = () => {
    const { user } = useAuthStore();
    const userId: string = user?.id!;
    const {
        data: userFollowingData,
        isLoading,
        error
    } = useUserFollowing(userId) as {
        data: Club[];
        isLoading: boolean;
        error: Error;
    };

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <SafeAreaView edges={['top']}>
            <ScrollView showsVerticalScrollIndicator={false}>
                <View className="pb-[28%]">
                    <View className="pt-[6%]">
                        <View className="flex-row justify-between items-center mb-[4%] mx-[6%]">
                            <Text className="text-3xl">Followed Clubs</Text>
                        </View>
                        <ScrollView
                            horizontal
                            showsHorizontalScrollIndicator={false}
                        >
                            <View className="flex-row">
                                <View className="pl-7"></View>
                                {userFollowingData &&
                                    userFollowingData.map((club, index) => (
                                        <View
                                            className="flex-col w-20 mr-5 items-center"
                                            key={index}
                                        >
                                            <View className="w-20 h-20 bg-gray-300 rounded-full"></View>
                                            <Text
                                                numberOfLines={1}
                                                ellipsizeMode="tail"
                                                className="pt-2 flex-wrap text-xs"
                                            >
                                                {club.name}
                                            </Text>
                                        </View>
                                    ))}
                            </View>
                        </ScrollView>
                        <View className="px-[6%] pt-[3%]">
                            {HomepageList.map((item, index) => {
                                switch (item.type) {
                                    case 'club':
                                        return (
                                            <ClubHomepageCard
                                                key={index}
                                                club={item}
                                            />
                                        );
                                    case 'event':
                                        return (
                                            <EventCard
                                                key={index}
                                                event={item}
                                            />
                                        );
                                    case 'faq':
                                        return (
                                            <FAQCard key={index} faq={item} />
                                        );
                                }
                            })}
                        </View>
                    </View>
                </View>
            </ScrollView>
        </SafeAreaView>
    );
};

export default Homepage;
