import { Link } from 'expo-router';
import React, { useState } from 'react';
import { Text, TouchableOpacity, TouchableOpacityProps, View, Modal } from 'react-native';


export interface FaqCardProps extends TouchableOpacityProps {
    question?: string;
    answer?: string;
}

const FaqCard = ({ question, answer, ...props }: FaqCardProps) => {
    const [isModalVisible, setIsModalVisible] = useState(false);
    const MAX_LENGTH = 8;
    const toggleModal = () => setIsModalVisible(!isModalVisible);
    return (
        <View>
            <Link href={{ pathname: '/(app)/club/faq/[id]', params: { id: 1 } }} asChild>
                <TouchableOpacity {...props} onPress={toggleModal} className='mr-3'>
                    <View className="bg-white p-4 rounded-lg w-80 h-80 border border-[#CDCBCB]">
                        {question && <Text className="text-sm mb-2 font-semibold">{question}</Text>}
                        {answer && <Text numberOfLines={MAX_LENGTH} className="text-sm">{answer}</Text>}
                        <Text className="pt-5">200 think this is helpful.</Text>
                    </View>
                </TouchableOpacity>
            </Link>
        </View>
    );
};


FaqCard.displayName = 'faqCard';

export { FaqCard };
