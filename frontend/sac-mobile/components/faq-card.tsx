import React from 'react';
import { Text, TouchableOpacity, View } from 'react-native';

import Input from '@/components/input';
import { FAQ } from '@/types/item';

const FAQCard = ({ faq }: { faq: FAQ }) => {
    return (
        <TouchableOpacity className="bg-gray-200 rounded-2xl my-[2%] p-[5%]">
            <View className="flex-row">
                <View className="bg-gray-300 rounded-xl w-16 h-16"></View>
                <View className="ml-[5%]">
                    <Text className="text-base leading-6 font-bold ">{faq.clubName}</Text>
                    <Text className="text-sm font-medium leading-5 text-gray-500">Frequently Asked</Text>
                    <Text className="text-sm font-medium leading-5 text-gray-500">Questions</Text>
                </View>
            </View>
            <View className="mt-[4%]">
                <Text className="text-sm font-bold">Questions:</Text>
                <Text>{faq.question}</Text>
                <Text className="text-sm font-bold mt-[3%]">Answer:</Text>
                <Text numberOfLines={2} ellipsizeMode="tail">
                    {faq.answer}
                </Text>
            </View>
        </TouchableOpacity>
    );
};

export default FAQCard;
