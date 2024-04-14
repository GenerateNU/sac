import { View, Text, SafeAreaView, TextInput, TouchableOpacity } from 'react-native'
import React from 'react'
import { Button } from '@/components/button';
import { router } from 'expo-router';
import { MaterialCommunityIcons } from '@expo/vector-icons';

const AskAQuestion = () => {
    return (
        <SafeAreaView className="bg-white h-[100%]">
            <View>
                <Button size={'icon'} variant="outline" className='flex justify-center items-center absolute top-1 left-6 z-10 bg-gray-300 rounded-full' onPress={() => router.back()}>
                    <MaterialCommunityIcons name="chevron-left" size={24} color="rgb(75 85 99)" />
                </Button>
                <Text className='text-center font-bold text-xl'>FAQ</Text>
                <Text className='p-[5%] text-left font-bold text-lg'>Question</Text>
                <TextInput multiline={true} numberOfLines={4} className='border border-[#CDCBCB] h-40 rounded-lg w-[90%]  mx-[5%] p-4' />
            </View>
        </SafeAreaView >
    )
}

export default AskAQuestion;