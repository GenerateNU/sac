import { View, Text } from 'react-native'
import React from 'react'
import { useLocalSearchParams } from 'expo-router'
import { useFAQ } from '@/hooks/use-faqs'

const Modal = () => {
    const { id } = useLocalSearchParams<{ id: string }>()

    const { data: faq, error, isLoading } = useFAQ(id)

    if (isLoading) {
        return <Text>Loading...</Text>
    }

    if (error) {
        return <Text>Error: {error.message}</Text>
    }

    return (
        <View className='flex-1 justify-center items-center bg-black/50'>
            <View className='bg-white opacity-100 rounded-lg p-4 w-80'>
                {faq && <Text className="text-sm mb-2 font-semibold">{faq.question}</Text>}
                {faq && <Text className="text-sm">{faq.answer}</Text>}
            </View>
        </View>
    )
}

export default Modal
