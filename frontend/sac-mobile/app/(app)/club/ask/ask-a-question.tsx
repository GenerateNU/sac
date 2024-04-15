import { View, Text, SafeAreaView, TextInput, TouchableOpacity } from 'react-native'
import React, { useState } from 'react'
import { Button } from '@/components/button';
import { router } from 'expo-router';
import { MaterialCommunityIcons } from '@expo/vector-icons';

const AskAQuestion = () => {
    const [question, setQuestion] = useState('');
    const handleSubmit = () => {
        // Perform the submission logic here
        // For now, logging a message
        console.log('Submitted Question:', question);
        router.back();
    };


    return (
        <SafeAreaView className="bg-white h-[100%]">
            <View>
                <Text className='text-center font-bold text-xl'>Ask a Question</Text>
                <Text className='p-[5%] text-left font-bold text-lg'>Question</Text>
                <TextInput multiline={true} numberOfLines={4} onChangeText={(text) => setQuestion(text)}
                    value={question} className='border border-[#CDCBCB] h-40 rounded-lg w-[90%]  mx-[5%] p-4' />
            </View>
            <View className='flex-row justify-between p-4'>
                <Button variant='outline' onPress={() => router.back()}>Cancel</Button>
                <Button onPress={() => handleSubmit()}>Submit</Button>
            </View>
        </SafeAreaView >
    )
}

export default AskAQuestion;