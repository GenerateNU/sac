import React from 'react';
import { Text, TouchableOpacity, View } from 'react-native';

import { Button } from '@/components/button';
import Input from '@/components/input';
import { FAQ } from '@/types/item';

const FAQCard = ({ faq }: { faq: FAQ }) => {
    const length = () => {
        if (faq.clubName.length > 11) {
            return faq.clubName.substring(0, 11) + '...';
        } else {
            return faq.clubName;
        }
    };
    return (
        <TouchableOpacity className="bg-gray-200 rounded-2xl my-[2%] p-[5%]">
            <View className="flex-row">
                <View className="bg-gray-300 rounded-xl w-16 h-16"></View>
                <View className="ml-[5%]">
                    <Text className="text-base leading-6 font-bold ">
                        {faq.clubName}
                    </Text>
                    <Text className="text-sm font-medium leading-5 text-gray-500">
                        Frequently Asked
                    </Text>
                    <Text className="text-sm font-medium leading-5 text-gray-500">
                        Questions
                    </Text>
                </View>
            </View>
            <View className="mt-[7%]">
                <Text className="text-sm font-bold">Questions:</Text>
                <Text>{faq.question}</Text>
                <Text className="text-sm font-bold mt-[4%]">Answer:</Text>
                <Text numberOfLines={2} ellipsizeMode="tail">
                    {faq.answer}
                </Text>
                <View className="mt-[6%] mb-[1.5%]">
                    <Input
                        variant="faq"
                        placeholder={'Submit a question for ' + length()}
                    />
                </View>
            </View>
        </TouchableOpacity>
    );
};

export default FAQCard;
