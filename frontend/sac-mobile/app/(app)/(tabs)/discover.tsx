import React, {  } from 'react';
import { SafeAreaView, ScrollView, Text, TouchableOpacity, View } from 'react-native';
import { ClubCardNoTags } from '../search/_components/club-card';
import { SearchBar } from '../search/_components/search-bar';
import { HorizontalEventCard } from '../search/_components/event-card';
import { TopSearch } from '../search/_components/top-search';
import { router } from 'expo-router';


const Discover = () => {    
    return (
        <SafeAreaView className="w-full h-full bg-[#fafafa] pt-[50px] flex flex-col px-[5%]">
            <SearchBar 
                searchBoxEditable={false}
                cancelOrBackButton="Back"
                onSearchBoxPressed={() => {
                    router.push("/(app)/search/recents")
                    return undefined;
                }}/>
            <ScrollView className="flex flex-col mt-6" showsVerticalScrollIndicator={false}>
            <Text className="font-bold text-base mb-2">Top Searches</Text>
            {/* FIXME: this and discover clubs have the same layout, make like a 2d grid scroll view elem */}
            <ScrollView horizontal showsHorizontalScrollIndicator={false} className="overflow-visible mb-6" contentContainerStyle={{
                display: "flex",
                flexDirection: "column"
            }}>
                <View className="flex flex-row">
                    <TopSearch search={"COExist"} className="mr-2 w-40"/>
                    <TopSearch search={"Volleyball"} className="mr-2 w-40"/>
                    <TopSearch search={"iGEM"} className="mr-2 w-40"/>
                </View>
                <View className="mt-2 flex flex-row">
                    <TopSearch search={"Jiu Jitsu"} className="mr-2 w-40"/>
                    <TopSearch search={"Music"} className="mr-2 w-40"/>
                    <TopSearch search={"Kappa Sigma"} className="mr-2 w-40"/>
                </View>
            </ScrollView>
            
            <View className="w-full flex flex-row justify-between items-center">
                <Text className="font-bold text-base mb-2">Discover Clubs</Text>
                <TouchableOpacity>
                    <Text className="text-xs text-zinc-600">
                        View All
                    </Text>
                </TouchableOpacity>
            </View>

            <ScrollView horizontal showsHorizontalScrollIndicator={false} className="overflow-visible mb-6" contentContainerStyle={{
                display: "flex",
                flexDirection: "column"
            }}>
                <View className="flex flex-row">
                    <ClubCardNoTags className="mr-2 h-15 w-40" />
                    <ClubCardNoTags className="mr-2 h-15 w-40" />
                    <ClubCardNoTags className="mr-2 h-15 w-40" />
                </View>
                <View className="mt-2 flex flex-row">
                    <ClubCardNoTags className="mr-2 h-15 w-40" />
                    <ClubCardNoTags className="mr-2 h-15 w-40" />
                    <ClubCardNoTags className="mr-2 h-15 w-40" />
                </View>
            </ScrollView>

            <View className="w-full flex flex-row justify-between items-center">
                <Text className="font-bold text-base mb-2">Discover Events</Text>
                <TouchableOpacity>
                    <Text className="text-xs text-zinc-600">
                        View All
                    </Text>
                </TouchableOpacity>
            </View>
            
            <ScrollView horizontal showsHorizontalScrollIndicator={false} className="w-full overflow-visible" contentContainerStyle={{
                display: "flex",
                flexDirection: "column",
                width: "100%"
            }}>
                <HorizontalEventCard className="w-fulll mb-2"/>
                <HorizontalEventCard className="w-full mb-2"/>
                <HorizontalEventCard className="w-full mb-2"/>
                <HorizontalEventCard className="w-full mb-2"/>
            </ScrollView>

        </ScrollView>
        </SafeAreaView>
    );
};

export default Discover;
