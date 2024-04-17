import { useState, useEffect } from "react";
import { View, Text, ViewProps, TouchableOpacity } from "react-native";

import SearchIcon from '@/assets/images/svg/search.svg';
import { Button } from "@/components/button";
import { MaterialCommunityIcons } from "@expo/vector-icons";

export interface SearchRecentsProps extends ViewProps {

}

export const SearchRecents = (p: SearchRecentsProps) => {
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
        <View className="flex flex-col mt-6">
            <Text className="font-bold text-base mb-2">Recents</Text>
            {searchRecents}
        </View>
    );
}

export default SearchRecents;