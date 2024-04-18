import { useClubTags } from "@/hooks/use-club-tags";
import { Club } from "@/types/club";
import { router } from "expo-router";
import { useState, useEffect } from "react";
import { Pressable, ScrollView, Text, View, ViewProps } from "react-native";

export interface ClubCardProps extends ViewProps {
    club?: Club
    showTags?: boolean,
}


// @refresh reset
export const ClubCard = (p: ClubCardProps) => {
    const { data: tags, isLoading, error } = useClubTags(p.club ? p.club.id : "tester");

    const tagComponents = !isLoading && error == null ? tags!.map((val, ind, arr) => {
        return (
            <View key={val.name} className="border-zinc-200 border py-1 px-2 rounded-full mr-1">
                <Text className="text-base font-light">{val.name}</Text>
            </View>
        );
    }) : [];

    return (
        <Pressable
            onPress={() => {
                router.push(`/(app)/club/${p.club?.id}`)
            }}>
            <View {...p} className="elevation bg-white rounded-xl flex flex-col p-3">
                <View className="w-full h-1/2 flex flex-row items-center">
                    <View className="w-12 h-12 rounded-full bg-zinc-200"></View>
                    <Text className="ml-4 text-base">{p.club ? p.club.name : "Club"}</Text>
                </View>
                {p.showTags ? <ScrollView horizontal className="w-full h-1/2 flex flex-row mt-2" showsHorizontalScrollIndicator={false}>
                    {tagComponents}
                    </ScrollView> : <></>}
            </View>
        </Pressable>
        
    );
}

export const ClubCardNoTags = (p: ClubCardProps) => {
    return (
        <View {...p} className="elevation bg-white rounded-xl flex flex-row p-3 items-center">
            <View className="w-12 h-12 rounded-full bg-zinc-200"></View>
            <Text className="ml-4 text-base">Club Name</Text>
        </View>
    );
}