import { router, useLocalSearchParams } from 'expo-router';
import React from 'react';
import { SafeAreaView, ScrollView, Text, View } from 'react-native';

import { ClubCard } from '@/app/(app)/search/_components/club-card';
import { SearchBar } from '@/app/(app)/search/_components/search-bar';
import { useClubSearch } from '@/hooks/use-club-search';

const ResultsExpanded = () => {
    // TODO: Add grayboxing.
    const search: string = useLocalSearchParams().query as any as string;
    const { data, error, isLoading } = useClubSearch(search);

    const clubCards =
        !isLoading && error == null
            ? data!.map((club) => {
                  return <ClubCard club={club} className="h-28 mb-2" />;
              })
            : [];

    return (
        <SafeAreaView className="w-full h-full bg-[#fafafa] pt-[50px] flex flex-col px-[5%]">
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
                <View className="w-full flex flex-row justify-between items-center">
                    <Text className="font-bold text-base mb-2">
                        Recommended Clubs
                    </Text>
                </View>

                {clubCards}
            </ScrollView>
        </SafeAreaView>
    );
};

export default ResultsExpanded;
