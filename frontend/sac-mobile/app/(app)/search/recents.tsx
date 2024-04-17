import React, {  } from 'react';
import { SafeAreaView, ScrollView, Text, TouchableOpacity, View } from 'react-native';
import { ClubCardNoTags } from './_components/club-card';
import { SearchTopView } from './_components/top-view';
import { HorizontalEventCard } from './_components/event-card';
import { TopSearch } from './_components/top-search';
import { router } from 'expo-router';
import MaterialCommunityIcons from '@expo/vector-icons/build/MaterialCommunityIcons';

import SearchIcon from "@/assets/images/svg/search.svg";

const Recents = () => {    
    const searchRecents = [
        "UI/UX Design",
        "Women's Club Volleyball",
        "NU Italian Club",
        "Generate"
    ].map((val, ind, arr) => {
        return (
            <View className="w-full flex flex-row mb-4 items-center">
                <SearchIcon className="w-5 h-5 mr-2"></SearchIcon>
                <Text className="text-base flex-grow">{val}</Text>
                <TouchableOpacity>
                    <MaterialCommunityIcons
                        name="close" size={20} className="w-5 h-5">
                    </MaterialCommunityIcons>
                </TouchableOpacity>
            </View>
        )
    })
    return (
        <SafeAreaView className="w-full h-full bg-[#fafafa] pt-[50px] flex flex-col px-[5%]">
             <SearchTopView 
                active={true} 
                onSearchBoxGainFocus={() => { console.log("XXX") }}
                onSearchBoxLoseFocus={() => { router.push("/search/results")}}/>
            <View className="flex flex-col mt-6">
                <Text className="font-bold text-base mb-2">Recents</Text>
                {searchRecents}
            </View>
        </SafeAreaView>
    );
};

export default Recents;
