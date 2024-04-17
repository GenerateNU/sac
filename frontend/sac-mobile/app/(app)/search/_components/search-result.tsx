import { useState, useEffect } from "react";
import { View, ScrollView, Text, ViewProps, TouchableOpacity } from "react-native";
import { ClubCard } from "./club-card";
import { EventCard } from "./event-card";

export interface SearchResultProps extends ViewProps {
    onClubViewAllPressed: () => void
}

export const SearchResult = (p: SearchResultProps) => {
    return (
        <ScrollView className="flex flex-col mt-6">
            <View className="w-full flex flex-row justify-between items-center">
                <Text className="font-bold text-base mb-2">Recommended Clubs</Text>
                <TouchableOpacity
                    onPress={() => {
                        p.onClubViewAllPressed();
                    }}>
                    <Text className="text-xs text-zinc-600">
                        View All
                    </Text>
                </TouchableOpacity>
            </View>

            <ClubCard className="h-28 mb-2"/>
            <ClubCard className="h-28 mb-2"/>

            <View className="w-full flex flex-row justify-between items-center">
                <Text className="font-bold text-base mb-2">Recommended Events</Text>
                <TouchableOpacity>
                    <Text className="text-xs text-zinc-600">
                        View All
                    </Text>
                </TouchableOpacity>
            </View>
            
            <EventCard className="w-full mb-2"/>
            <EventCard className="w-full mb-2"/>
            <EventCard className="w-full mb-2"/>
            
        </ScrollView>
    );
}

export default SearchResult;