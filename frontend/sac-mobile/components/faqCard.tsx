import { Link, router } from 'expo-router';
import React, { useState } from 'react';
import { Text, TouchableOpacity, TouchableOpacityProps, View, Modal, Pressable } from 'react-native';
import Like from '@/assets/images/svg/like.svg';

export interface FaqCardProps extends TouchableOpacityProps {
    id?: string;
    question?: string;
    answer?: string;
}

const FaqCard = ({ id, question, answer, ...props }: FaqCardProps) => {
    const MAX_LENGTH = 8;
    return (
        // <Link href={{ pathname: '/(app)/club/faq/[id]', params: { id: 1 } }} asChild>
        <Pressable onPress={() => router.push(`/(app)/club/faq/${id}`)} className='mr-3'>
            <View className="bg-white p-4 rounded-2xl w-80 h-80 border border-[#CDCBCB] flex gap-y-2">
                {question && <Text className="w-10/12 text-lg font-semibold text-black">{question}</Text>}
                {answer && <Text numberOfLines={MAX_LENGTH} className="font-normal tracking-tight text-md text-neutral-500">{answer}</Text>}
                <View className="flex-row items-center flex-1">
                    <Like />
                    <Text className="text-[#F8C45C]">200 think this is helpful.</Text>
                </View>
            </View>
        </Pressable>
        // </Link>
    );
};


FaqCard.displayName = 'faqCard';

export { FaqCard };
