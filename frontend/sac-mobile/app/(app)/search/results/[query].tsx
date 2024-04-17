import React, {  } from 'react';
import { SafeAreaView, ScrollView, Text, TouchableOpacity, View } from 'react-native';
import { ClubCard, ClubCardNoTags } from '../_components/club-card';
import { SearchBar } from '../_components/search-bar';
import { EventCard, HorizontalEventCard } from '../_components/event-card';
import { TopSearch } from '../_components/top-search';
import { router, useLocalSearchParams } from 'expo-router';
import { queryClubs } from '@/services/club';
import { useClubSearch } from '@/hooks/use-club-search';

const Results = () => {
    // TODO: Add grayboxing.   
    const search: string = (useLocalSearchParams().query as any) as string
    const { data, error, isLoading } = useClubSearch(search)

    const clubCards = !isLoading && error == null ? data!.slice(0,2).map((club, ind, arr) => {
        return (<ClubCard club={club} className="h-28 mb-2"/>);
    }) : [];

    return (
        <SafeAreaView className="w-full h-full bg-[#fafafa] pt-[50px] flex flex-col px-[5%]">
             <SearchBar
                searchBoxEditable={false} 
                searchBoxText={search}  
                cancelOrBackButton='Back'/>
            <ScrollView className="flex flex-col mt-6">
                <View className="w-full flex flex-row justify-between items-center">
                    <Text className="font-bold text-base mb-2">Recommended Clubs</Text>
                    <TouchableOpacity
                        onPress={() => {
                            router.push(`/search/results/expanded/${search}`)
                        }}>
                        <Text className="text-xs text-zinc-600">
                            View All
                        </Text>
                    </TouchableOpacity>
                </View>

                {clubCards}

                <View className="w-full flex flex-row justify-between items-center">
                    <Text className="font-bold text-base mb-2">Recommended Events</Text>
                    <TouchableOpacity>
                        <Text className="text-xs text-zinc-600">
                            View All
                        </Text>
                    </TouchableOpacity>
                </View>
                    
                <View className="w-full h-56 flex flex-row mb-2 justify-between">
                    <EventCard className="w-1/2"/>
                    <EventCard className="w-1/2"/>    
                </View>
                
                <View className="w-full h-56 flex flex-row mb-2 justify-between">
                    <EventCard className="w-1/2"/>
                    <EventCard className="w-1/2"/>    
                </View>
            </ScrollView>
        </SafeAreaView>
    );
};

export default Results;
