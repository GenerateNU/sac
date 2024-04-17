import React, {  } from 'react';
import { SafeAreaView, ScrollView, Text, TouchableOpacity, View } from 'react-native';
import { ClubCard, ClubCardNoTags } from './_components/club-card';
import { SearchTopView } from './_components/top-view';
import { HorizontalEventCard } from './_components/event-card';
import { TopSearch } from './_components/top-search';
import { router } from 'expo-router';


const ResultsExpanded = () => {    
    return (
        <SafeAreaView className="w-full h-full bg-[#fafafa] pt-[50px] flex flex-col px-[5%]">
             <SearchTopView 
                active={false} 
                onSearchBoxGainFocus={() => { console.log("XXX") }}
                onSearchBoxLoseFocus={() => { console.log("YYY")}}/>
            <ScrollView className="flex flex-col mt-6">
                <View className="w-full flex flex-row justify-between items-center">
                    <Text className="font-bold text-base mb-2">Recommended Clubs</Text>
                </View>
                <ClubCard className="h-28 mb-2"/>
                <ClubCard className="h-28 mb-2"/>
                <ClubCard className="h-28 mb-2"/>
                <ClubCard className="h-28 mb-2"/>
                <ClubCard className="h-28 mb-2"/>
                <ClubCard className="h-28 mb-2"/>
                <ClubCard className="h-28 mb-2"/>

                <ClubCard className="h-28 mb-2"/>

            </ScrollView>
        </SafeAreaView>
    );
};

export default ResultsExpanded;
