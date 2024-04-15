import { useState, useEffect } from "react";
import { Button, Text, TouchableOpacity, View, ViewProps } from "react-native";

export interface EventCardProps extends ViewProps {

}

export const EventCard = (p: EventCardProps) => {
    return (
        <View {...p} className="elevation bg-white rounded-xl flex flex-col justify-end p-3">
            <Text className="text-xs text-zinc-600">Club Name</Text>
            <Text className="text-base mb-4">Tryouts</Text>
            <Text className="text-base">8pm - 10:30pm</Text>
            <Text className="text-xs text-zinc-600">April 23, 2024</Text>
            <Text className="text-xs text-zinc-600 mb-4">Hayden 010A</Text>

            <TouchableOpacity className="bg-black w-full flex flex-row justify-center items-center h-7 rounded-md">
                <Text className="text-white text-xs font-extralight">Register</Text>
            </TouchableOpacity>
        </View>
    );
}

export const HorizontalEventCard = (p: EventCardProps) => {
    return (
        <View {...p} className="elevation bg-white rounded-xl flex flex-row p-3 justify-between items-end">
            <View className="flex flex-col justify-end">
                <Text className="text-xs text-zinc-600">Club Name</Text>
                <Text className="text-base">Tryouts</Text>
            </View>

            <View className="flex flex-row items-end">
                <View className="flex flex-col justify-end mr-3">
                    <Text className="text-base">8pm - 10:30pm</Text>
                    <Text className="text-xs text-zinc-600">April 23, 2024</Text>
                    <Text className="text-xs text-zinc-600">Hayden 010A</Text>
                </View>

                <TouchableOpacity className="bg-black w-20 flex flex-row justify-center items-center h-7 rounded-md">
                    <Text className="text-white text-xs font-extralight">Register</Text>
                </TouchableOpacity>
            </View>
            
        </View>
    );
}