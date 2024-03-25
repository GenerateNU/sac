import React, { useState } from 'react';
import { Text, TouchableOpacity, TouchableOpacityProps, View, Modal } from 'react-native';


export interface FaqCardProps extends TouchableOpacityProps {
    question?: string;
    answer?: string;
}

const FaqCard = ({ question, answer, ...props }: FaqCardProps) => {
    const [isModalVisible, setIsModalVisible] = useState(false);
    const MAX_LENGTH = 4;
    const toggleModal = () => setIsModalVisible(!isModalVisible);
    return (
        <View>
            <TouchableOpacity {...props} onPress={toggleModal}>
                <View className="bg-white p-4 rounded-lg w-80 h-48 pb-12 border border-[#CDCBCB]">
                    {question && <Text className="text-sm mb-2 font-semibold">{question}</Text>}
                    {answer && <Text numberOfLines={MAX_LENGTH} className="text-sm">{answer}</Text>}
                </View>
            </TouchableOpacity>
            <Modal visible={isModalVisible} transparent={true}>
                <View className='flex-1 justify-center items-center bg-black/50'>
                    <View className='bg-white opacity-100 rounded-lg p-4 w-80'>
                        {question && <Text className="text-sm mb-2 font-semibold">{question}</Text>}
                        {answer && <Text className="text-sm">{answer}</Text>}
                        <TouchableOpacity onPress={toggleModal} style={{ marginTop: 10 }}>
                            <Text className="text-blue-500 underline">Close</Text>
                        </TouchableOpacity>
                    </View>
                </View>
            </Modal>
        </View>
    );
};


FaqCard.displayName = 'faqCard';

export { FaqCard };
