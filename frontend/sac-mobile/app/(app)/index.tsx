import React, { useState } from 'react';
import { ScrollView, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { Button } from '@/components/button';
import ClubCard from '@/components/club-card';
import EventCard from '@/components/event-card';
import FAQCard from '@/components/faq-card';
import { FollowedClubs, HomepageList } from '@/lib/const';
import {useUserFollowing } from '@/hooks/use-user-following';
import {Club} from '@/types/club';
import {useAuthStore} from '@/hooks/use-auth';

const Homepage = () => {
    const {user} = useAuthStore();
    const userId = user?.id;
    console.log(user)
    console.log("####")
    const { data: userFollowingData, isLoading: loadingUserFollowing, error: errorUserFollowing } = useUserFollowing(userId) as { data: Club[], isLoading: boolean, error: Error };

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
                                {FollowedClubs.map((club, index) => (
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
                                {/* {userFollowingData && userFollowingData.map((club, index) => (
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
                                ))} */}
                            </View>
                        </ScrollView>
                        <View className="px-[6%] pt-[3%]">
                            {HomepageList.map((item, index) => {
                                      switch (item.type) {
                                          case 'club':
                                              return (
                                                  <ClubCard
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
                                                  <FAQCard
                                                      key={index}
                                                      faq={item}
                                                  />
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
