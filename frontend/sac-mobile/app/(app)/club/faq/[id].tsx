import { View, Text, SafeAreaView } from 'react-native'
import React from 'react'
import { useLocalSearchParams } from 'expo-router';

const Faq = () => {
    const { id } = useLocalSearchParams<{ id: string }>();

    return (
        <SafeAreaView className='items-center justify-center flex-1 bg-white'>
            <Text>{id}</ Text>
        </SafeAreaView>
    )
}

export default Faq