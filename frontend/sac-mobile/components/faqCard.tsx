import React from 'react';
import { Text, View } from 'react-native';

import Like from '@/assets/images/svg/like.svg';

type FaqCardProps = {
    question?: string;
    answer?: string;
};

const FaqCard = ({ question, answer }: FaqCardProps) => {
    const MAX_LENGTH = 8;

    return (
        <View className="bg-white mr-3 p-4 rounded-2xl w-72 h-80 border flex-col justify-between border-[#CDCBCB]">
            <View className='flex gap-y-2'>
                {question && (
                    <Text className="w-10/12 text-lg font-semibold text-black">
                        {question}
                    </Text>
                )}
                {answer && (
                    <Text
                        numberOfLines={MAX_LENGTH}
                        className="font-normal tracking-tight text-md text-neutral-500"
                    >
                        {answer}
                    </Text>
                )}
            </View>
            <View className="flex-row items-center">
                <Like />
                <Text className="text-gray-400">
                    {`${Math.round(Math.random() * 200)} think this is helpful.`}
                </Text>
            </View>
        </View>
    );
};

export { FaqCard };
