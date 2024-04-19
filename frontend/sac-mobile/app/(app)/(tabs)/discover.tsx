import React from 'react';
import {
    SafeAreaView,
    ScrollView,
    Text,
    TouchableOpacity,
    View
} from 'react-native';

import { router } from 'expo-router';

import { SearchBar } from '@/app/(app)/search/_components/search-bar';
import { TopSearch } from '@/app/(app)/search/_components/top-search';
import { useClubs } from '@/hooks/use-club';
import { useEvents } from '@/hooks/use-event';

import { ClubHomePageCard } from './_components/club-homepage-card';
import { EventHomePageCardHorizontal } from './_components/event-homepage-card';

const Discover = () => {
    const { data: events, isLoading: eIsLoading, error: eError } = useEvents();
    const { data: clubs, isLoading: cIsLoading, error: cError } = useClubs();

    if (eIsLoading || cIsLoading) {
        return <Text>Loading...</Text>;
    }

    if (eError || cError) {
        return <Text>Error: {eError?.message || cError?.message}</Text>;
    }

    return (
        <SafeAreaView className="w-full h-full bg-[#fafafa]">
            <View className="px-[5%] mb-12">
                <SearchBar
                    searchBoxEditable={false}
                    cancelOrBackButton="Back"
                    onSearchBoxPressed={() => {
                        router.push('/(app)/search/recents');
                        return undefined;
                    }}
                />
                <ScrollView
                    className="flex flex-col mt-6"
                    showsVerticalScrollIndicator={false}
                >
                    <Text className="text-xl mb-3 font-light leading-relaxed text-zinc-900">
                        Top Searches
                    </Text>
                    {/* FIXME: this and discover clubs have the same layout, make like a 2d grid scroll view elem */}
                    <ScrollView
                        horizontal
                        showsHorizontalScrollIndicator={false}
                        className="overflow-visible mb-6"
                        contentContainerStyle={{
                            display: 'flex',
                            flexDirection: 'column'
                        }}
                    >
                        <View className="flex flex-row">
                            <TopSearch
                                search={'COExist'}
                                className="mr-2 w-40"
                            />
                            <TopSearch
                                search={'Volleyball'}
                                className="mr-2 w-40"
                            />
                            <TopSearch search={'iGEM'} className="mr-2 w-40" />
                        </View>
                        <View className="mt-2 flex flex-row">
                            <TopSearch
                                search={'Jiu Jitsu'}
                                className="mr-2 w-40"
                            />
                            <TopSearch search={'Music'} className="mr-2 w-40" />
                            <TopSearch
                                search={'Kappa Sigma'}
                                className="mr-2 w-40"
                            />
                        </View>
                    </ScrollView>

                    <View className="w-full flex flex-row justify-between items-center mb-3">
                        <Text className="text-xl font-light leading-relaxed text-zinc-900">
                            Discover Clubs
                        </Text>
                        <TouchableOpacity>
                            <Text className="text-xs text-zinc-600">
                                View All
                            </Text>
                        </TouchableOpacity>
                    </View>

                    <ScrollView
                        horizontal
                        showsHorizontalScrollIndicator={false}
                        className="overflow-visible mb-6"
                        contentContainerStyle={{
                            display: 'flex',
                            flexDirection: 'column'
                        }}
                    >
                        <View className="flex flex-row" style={{ gap: 10 }}>
                            {clubs?.map((club) => {
                                return <ClubHomePageCard club={club} partial />;
                            })}
                        </View>
                    </ScrollView>

                    <View className="w-full flex flex-row justify-between items-center mb-3">
                        <Text className="text-xl font-light leading-relaxed text-zinc-900">
                            Discover Events
                        </Text>
                        <TouchableOpacity>
                            <Text className="text-xs text-zinc-600">
                                View All
                            </Text>
                        </TouchableOpacity>
                    </View>
                    <View className="flex flex-col mb-5" style={{ gap: 15 }}>
                        {events?.map((event) => {
                            return (
                                <EventHomePageCardHorizontal event={event} />
                            );
                        })}
                    </View>
                </ScrollView>
            </View>
        </SafeAreaView>
    );
};

export default Discover;
