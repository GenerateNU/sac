import { Club } from "@/types/club";
import { useState, useEffect } from "react";
import { ScrollView, Text, TouchableOpacity, View, ViewProps } from "react-native";

import SearchIcon from "@/assets/images/svg/search.svg";
import { router } from "expo-router";

export interface TopSearchProps extends ViewProps {
    search: string
}

export const TopSearch = (p: TopSearchProps) => {
    return (
        <TouchableOpacity
            onPress={() => {
                router.push(`/(app)/search/results/${p.search}`)
            }}>
            <View {...p} className="elevation bg-white rounded-xl flex flex-row p-3 items-center">
                <SearchIcon className="w-5 h-5"></SearchIcon>
                <Text className="ml-2 text-xs">{p.search}</Text>
            </View>
        </TouchableOpacity>
        
    );
}