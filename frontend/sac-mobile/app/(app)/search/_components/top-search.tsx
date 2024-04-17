import { Club } from "@/types/club";
import { useState, useEffect } from "react";
import { ScrollView, Text, View, ViewProps } from "react-native";

import SearchIcon from "@/assets/images/svg/search.svg";

export interface TopSearchProps extends ViewProps {}

export const TopSearch = (p: TopSearchProps) => {
    return (
        <View {...p} className="elevation bg-white rounded-xl flex flex-row p-3 items-center">
            <SearchIcon className="w-5 h-5"></SearchIcon>
            <Text className="ml-2 text-xs">Top Search</Text>
        </View>
    );
}