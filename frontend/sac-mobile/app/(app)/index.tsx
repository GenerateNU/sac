import React, { useState } from 'react';
import { Pressable, ScrollView, Text, View } from 'react-native';
import { SafeAreaView} from 'react-native-safe-area-context';

import { Button } from '@/components/button';
import ClubCard from '@/components/club-card';
import EventCard from '@/components/event-card';
import FAQCard from '@/components/faq-card';
import { ChronologicalList, FollowedClubs, HomepageList } from '@/lib/const';

const Homepage = () => {
    const [selectedTab, setSelectedTab] = useState('Relevance');

    const handleTabPress = (tab: string) => {
        setSelectedTab(tab);
    };

    return (
        <SafeAreaView edges={['top']}>
            <ScrollView showsVerticalScrollIndicator={false}>
                <View className="pt-[6%]">
                    <View className="flex-row justify-between items-center mb-[4%] mx-[6%]">
                        <Text className="text-3xl">Followed Clubs</Text>
                        <Text className="text-sm font-bold">View all</Text>
                    </View>
                    <ScrollView
                        horizontal
                        showsHorizontalScrollIndicator={false}
                    >
                        <View className="flex-row pl-[3%] pr-[20%]">
                            {FollowedClubs.map((club, index) => (
                                <View
                                    className="flex-col mr-3 items-center"
                                    key={index}
                                >
                                    <View className="w-16 h-16 bg-gray-300 rounded-full"></View>
                                    <Text className="mt-2 flex-wrap">
                                        {club.name}
                                    </Text>
                                </View>
                            ))}
                        </View>
                    </ScrollView>
                    <View className="flex-row justify-center mt-[8%]">
                        <Button
                            variant={
                                'Relevance' === selectedTab
                                    ? 'underline'
                                    : 'menu'
                            }
                            size="sm"
                            onPress={() => handleTabPress('Relevance')}
                        >
                            {' '}
                            Relevance
                        </Button>
                        <Button
                            variant={
                                'Chronological' === selectedTab
                                    ? 'underline'
                                    : 'menu'
                            }
                            size="sm"
                            onPress={() => handleTabPress('Chronological')}
                        >
                            {' '}
                            Chronological
                        </Button>
                    </View>
                    <View className="px-[6%] pt-[3%]">
                        {selectedTab === 'Relevance'
                            ? HomepageList.map((item, index) => {
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
                                              <FAQCard key={index} faq={item} />
                                          );
                                  }
                              })
                            : ChronologicalList.map((item, index) => {
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
                                              <FAQCard key={index} faq={item} />
                                          );
                                  }
                              })}
                    </View>
                </View>
            </ScrollView>
        </SafeAreaView>
    );
};

export default Homepage;
