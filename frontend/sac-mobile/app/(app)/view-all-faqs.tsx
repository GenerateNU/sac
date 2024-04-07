import { View, Text, SafeAreaView, TextInput, TouchableOpacity, ScrollView } from 'react-native'
import React from 'react'
import { Button } from '@/components/button'
import { MaterialCommunityIcons } from '@expo/vector-icons'
import { router } from 'expo-router'
import { ViewAllFaqCard, ViewAllFaqCardProps } from '@/components/view-all-faqs-card'
import { useFAQs } from '@/hooks/use-faqs'


const AllFAQs = () => {
    const { data: faqs, isLoading, error } = useFAQs("sac");
    return (
        <SafeAreaView className="bg-white h-[100%]">
            <ScrollView className="flex-1">
                <View>
                    <Button size={'icon'} variant="outline" className='flex justify-center items-center absolute top-1 left-6 z-10 bg-gray-300 rounded-full' onPress={() => router.back()}>
                        <MaterialCommunityIcons name="chevron-left" size={24} color="rgb(75 85 99)" />
                    </Button>
                    <Text className='text-center font-bold text-xl pt-2 pb-5'>Club Name</Text>
                    <View className='flex-row border border-[#CDCBCB] pl-3 rounded-lg mx-[5%] items-center'>
                        <MaterialCommunityIcons name="magnify" size={30} color="rgb(75 85 99)" />
                        <TextInput className='w-[90%] p-4' placeholder='Search' />
                    </View>
                    <View className='flex-row px-[5%] py-[4%] justify-between'>
                        <Text className='text-center font-medium text-3xl'>FAQ</Text>
                        <TouchableOpacity onPress={() => router.push('/(app)/ask-a-question')}>
                            <MaterialCommunityIcons name="plus-circle-outline" size={40} color="black" />
                        </TouchableOpacity>
                    </View>
                    <View className='border-b mx-6'></View>
                    <View className="">
                        {faqs && faqs.map((faq: ViewAllFaqCardProps) => (
                            <ViewAllFaqCard key={faq.question} question={faq.question} answer={faq.answer} />
                        ))}
                    </View>
                </View>
            </ScrollView>
        </SafeAreaView >
    )
}

export default AllFAQs