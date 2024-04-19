import React from 'react';
import { SafeAreaView, ScrollView, Text, View } from 'react-native';

import { router, useLocalSearchParams } from 'expo-router';

import { ClubCard } from '@/app/(app)/search/_components/club-card';
import { SearchBar } from '@/app/(app)/search/_components/search-bar';
import { useClubSearch } from '@/hooks/use-club-search';
import { ClubHomePageCard } from '@/app/(app)/(tabs)/_components/club-homepage-card';

const ResultsExpanded = () => {
    // TODO: Add grayboxing.
    const search: string = useLocalSearchParams().query as any as string;
    const { data, error, isLoading } = useClubSearch(search);

    const clubCards =
        !isLoading && error == null
            ? data!.map((club) => {
                  return (
                    <View className='mb-3'>
                        <ClubHomePageCard club={club}/>
                    </View>
                  )
              })
            : [];

    return (
        <SafeAreaView className="w-full h-full bg-[#fafafa]">
            <View className='px-[5%]'>
                <SearchBar
                    searchBoxEditable={false}
                    searchBoxText={search}
                    cancelOrBackButton="Back"
                    onSearchBoxPressed={() => {
                        router.back();
                        router.back();
                        return 'Focus';
                    }}
                />
                <ScrollView className="flex flex-col mt-6">
                    <View className="w-full flex flex-row justify-between items-center mb-3">
                        <Text className="text-xl font-light leading-relaxed text-zinc-900">
                            Search Results
                        </Text>
                    </View>

                    {clubCards}
                </ScrollView>
            </View>
        </SafeAreaView>
    );
};

export default ResultsExpanded;
