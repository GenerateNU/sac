import { useState, useEffect } from "react";
import { View, ScrollView, Text, ViewProps, TouchableOpacity } from "react-native";
import { TopSearch } from "./top-search";
import { ClubCardNoTags } from "./club-card";
import { HorizontalEventCard } from "./event-card";

export interface SearchHomeViewProps extends ViewProps {

}

export const SearchHomeView = (p: SearchHomeViewProps) => {
    return (
        <ScrollView className="flex flex-col mt-6">
            <Text className="font-bold text-base mb-2">Top Searches</Text>
            {/* FIXME: this and discover clubs have the same layout, make like a 2d grid scroll view elem */}
            <ScrollView horizontal showsHorizontalScrollIndicator={false} className="overflow-visible mb-6" contentContainerStyle={{
                display: "flex",
                flexDirection: "column"
            }}>
                <View className="flex flex-row">
                    <TopSearch className="mr-2 w-40"/>
                    <TopSearch className="mr-2 w-40"/>
                    <TopSearch className="mr-2 w-40"/>
                </View>
                <View className="mt-2 flex flex-row">
                    <TopSearch className="mr-2 w-40"/>
                    <TopSearch className="mr-2 w-40"/>
                    <TopSearch className="mr-2 w-40"/>
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
            
            <ScrollView horizontal showsHorizontalScrollIndicator={false} className="overflow-visible" contentContainerStyle={{
                display: "flex",
                flexDirection: "column"
            }}>
                <HorizontalEventCard className="w-full mb-2"/>
                <HorizontalEventCard className="w-full mb-2"/>
                <HorizontalEventCard className="w-full mb-2"/>
                <HorizontalEventCard className="w-full mb-2"/>
            </ScrollView>

        </ScrollView>
    );
}

export default SearchHomeView;