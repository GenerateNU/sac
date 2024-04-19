import React from 'react';
import { Text, View } from 'react-native';

import { MaterialCommunityIcons } from '@expo/vector-icons';

type ViewAllFaqCardProps = {
    question?: string;
    answer?: string;
};

const ViewAllFaqCard = ({ question, answer }: ViewAllFaqCardProps) => {
    return (
        <View>
            <Text className="px-[5%] py-[3%] text-left font-medium">
                {question}
            </Text>
            <Text className="px-[5%] pb-[3%] text-left ">{answer}</Text>
            <View className="px-[5%] pb-[3%] flex-row items-center">
                <MaterialCommunityIcons
                    className="px-[5%]"
                    name="heart"
                    size={24}
                    color="grey"
                />
                <Text className="px-[3%]">
                    200 people think this is helpful.
                </Text>
            </View>
            <View className="mx-6 border-b border-gray-400" />
        </View>
    );
};

export { ViewAllFaqCard };
