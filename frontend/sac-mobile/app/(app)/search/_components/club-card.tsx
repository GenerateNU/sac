import { Club } from "@/types/club";
import { useState, useEffect } from "react";
import { ScrollView, Text, View, ViewProps } from "react-native";

export interface ClubCardProps extends ViewProps {}


// @refresh reset
export const ClubCard = (p: ClubCardProps) => {
    const tags = [
        "Club Sport",
        "Competitive",
        "Yearly",
        "Coed",
        "Beginners Welcomed",
    ].map((val, ind, arr) => {
        return (
            <View key={val} className="border-zinc-200 border py-1 px-2 rounded-full mr-1">
                <Text className="text-base font-light">{val}</Text>
            </View>
        );
    });
    return (
        <View {...p} className="elevation bg-white rounded-xl flex flex-col p-3">
            <View className="w-full h-1/2 flex flex-row items-center">
                <View className="w-12 h-12 rounded-full bg-zinc-200"></View>
                <Text className="ml-4 text-base">Club Name</Text>
            </View>
            <ScrollView horizontal className="w-full h-1/2 flex flex-row mt-2" showsHorizontalScrollIndicator={false}>
                {tags}
            </ScrollView>
        </View>
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