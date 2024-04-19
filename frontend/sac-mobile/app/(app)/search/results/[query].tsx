import React from 'react';
import {
    SafeAreaView,
    ScrollView,
    Text,
    TouchableOpacity,
    View
} from 'react-native';

import { router, useLocalSearchParams } from 'expo-router';

import { ClubCard } from '@/app/(app)/search/_components/club-card';
import { EventCard } from '@/app/(app)/search/_components/event-card';
import { SearchBar } from '@/app/(app)/search/_components/search-bar';
import { useClubSearch } from '@/hooks/use-club-search';
import { ClubHomePageCard } from '../../(tabs)/_components/club-homepage-card';
import { EventHomePageCard } from '../../(tabs)/_components/event-homepage-card';
import { useEvents } from '@/hooks/use-event';

const Results = () => {
    // TODO: Add grayboxing.
    const search: string = useLocalSearchParams().query as any as string;
    const { data, error, isLoading } = useClubSearch(search);
    const { data: events } = useEvents();

    const clubCards =
        !isLoading && error == null
            ? data!.slice(0, 2).map((club) => {
                  return (
                    <View className='mb-3'>
                        <ClubHomePageCard club={club} partial />
                    </View>
                  );
              })
            : [];

    return (
        <SafeAreaView className="w-full h-full bg-[#fafafa] pt-[50px] flex flex-col px-[5%]">
            <View className='px-[5%]'>
                <SearchBar
                    searchBoxEditable={false}
                    searchBoxText={search}
                    cancelOrBackButton="Back"
                />
                <ScrollView className="flex flex-col mt-6">
                    <View className="w-full flex flex-row justify-between items-center mb-3">
                        <Text className="text-xl font-light leading-relaxed text-zinc-900">
                            Search Results
                        </Text>
                        <TouchableOpacity
                            onPress={() => {
                                router.push(`/search/results/expanded/${search}`);
                            }}
                        >
                            <Text className="text-xs text-zinc-600">View All</Text>
                        </TouchableOpacity>
                    </View>

                    {clubCards}

                    <View className="w-full flex flex-row justify-between items-center mb-3">
                        <Text className="text-xl font-light leading-relaxed text-zinc-900">
                            Recommended Events
                        </Text>
                        <TouchableOpacity>
                            <Text className="text-xs text-zinc-600">View All</Text>
                        </TouchableOpacity>
                    </View>

                    <View className="w-full h-56 flex flex-row mb-2 justify-between">
                        {events?.slice(0,2)?.map((event) => {
                            return (
                                <EventHomePageCard event={event}/>
                            )
                        })}
                    </View>

                    <View className="w-full h-56 flex flex-row mb-12 justify-between">
                        {events?.slice(2,4)?.map((event) => {
                                return (
                                    <EventHomePageCard event={event}/>
                                )
                        })}
                    </View>
                </ScrollView>
            </View>
        </SafeAreaView>
    );
};

export default Results;
