import { MaterialCommunityIcons } from '@expo/vector-icons';
import React, { useState } from 'react';
import { Text, TouchableOpacity, TouchableOpacityProps, View, Modal } from 'react-native';


export interface ViewAllFaqCardProps {
    question?: string;
    answer?: string;
}

const ViewAllFaqCard = ({ question, answer }: ViewAllFaqCardProps) => {
    const [isModalVisible, setIsModalVisible] = useState(false);
    const MAX_LENGTH = 8;
    const toggleModal = () => setIsModalVisible(!isModalVisible);
    return (
        <View>
            <Text className='px-[5%] py-[3%] text-left font-medium'>{question}</Text>
            <Text className='px-[5%] pb-[3%] text-left '>{answer}</Text>
            <View className='px-[5%] pb-[3%] flex-row items-center'>
                <MaterialCommunityIcons className='px-[5%]' name="heart" size={24} color="grey" />
                <Text className='px-[3%]'>200 people think this is helpful.</Text>
            </View>
            <View className='border-b mx-6 border-gray-400'></View>
        </View>
    );
};


ViewAllFaqCard.displayName = 'viewAllFaqCard';

export { ViewAllFaqCard };
