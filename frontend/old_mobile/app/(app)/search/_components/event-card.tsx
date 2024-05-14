import { ScrollView, Text, View, ViewProps } from 'react-native';

import { MaterialCommunityIcons } from '@expo/vector-icons';

export interface EventCardProps extends ViewProps {}

export const EventCard = (p: EventCardProps) => {
    return (
        <View
            {...p}
            className="elevation bg-white rounded-xl flex flex-col justify-end p-3"
        >
            <View className="h-16 bg-zinc-200 rounded-xl mb-2" />
            <Text className="text-xs text-zinc-600">Club Name</Text>
            <Text className="text-base mb-4">Tryouts</Text>
            <Text className="text-base">8pm - 10:30pm</Text>
            <Text className="text-xs text-zinc-600">April 23, 2024</Text>
            <View className="w-full flex flex-row justify-start">
                <MaterialCommunityIcons
                    className="w-4 h-4 text-zinc-600"
                    name="map-marker"
                />
                <Text className="text-xs text-zinc-600 ml-1 mb-4">
                    Hayden 010A
                </Text>
            </View>
        </View>
    );
};

export const HorizontalEventCard = (p: EventCardProps) => {
    const tagComponents = ['Free Food', 'Open To All'].map((val) => {
        return (
            <View
                key={val}
                className="border-zinc-200 border py-1 px-2 rounded-full mr-1"
            >
                <Text className="text-base font-light">{val}</Text>
            </View>
        );
    });

    return (
        <View
            {...p}
            className="elevation bg-white rounded-xl flex flex-row p-3 justify-start items-start"
        >
            <View className="w-20 h-20 bg-zinc-200  rounded-xl" />
            <View className="flex flex-grow flex-col justify-start ml-3">
                <Text className="text-xs text-zinc-600">
                    April 23, 2024 - 8-9pm
                </Text>
                <Text className="text-base text-black">
                    Applications Close!
                </Text>
                <Text className="text-xs text-zinc-600">
                    Hosted by NUImpact
                </Text>
                <ScrollView
                    horizontal
                    className="w-full flex flex-row mt-2"
                    showsHorizontalScrollIndicator={false}
                >
                    {tagComponents}
                </ScrollView>
            </View>
        </View>
    );
};
