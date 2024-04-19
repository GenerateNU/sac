import React from 'react';
import {
    SafeAreaView,
    ScrollView,
    Text,
    TextInput,
    TouchableOpacity,
    View
} from 'react-native';

import { router } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';

import { ViewAllFaqCard } from '@/components/view-all-faqs-card';
import { useClubFAQs } from '@/hooks/use-club';

const AllFAQs = () => {
    const { data: faqs, isLoading, error } = useClubFAQs('sac');

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error fetching FAQs</Text>;
    }

    return (
        <SafeAreaView className="bg-white h-[100%]">
            {/* <Text>Hello</Text> */}
            <ScrollView className="flex-1">
                <View>
                    <Text className="pt-2 pb-5 text-xl font-bold text-center">
                        Club Name
                    </Text>
                    <View className="flex-row border border-[#CDCBCB] pl-3 rounded-lg mx-[5%] items-center">
                        <MaterialCommunityIcons
                            name="magnify"
                            size={30}
                            color="rgb(75 85 99)"
                        />
                        <TextInput
                            className="w-[90%] p-4"
                            placeholder="Search"
                        />
                    </View>
                    <View className="flex-row px-[5%] py-[4%] justify-between">
                        <Text className="text-3xl font-medium text-center">
                            FAQ
                        </Text>
                        <TouchableOpacity
                            onPress={() =>
                                router.push('/(app)/club/ask-a-question')
                            }
                        >
                            <MaterialCommunityIcons
                                name="plus-circle-outline"
                                size={40}
                                color="black"
                            />
                        </TouchableOpacity>
                    </View>
                    <View className="mx-6 border-b" />
                    <View className="">
                        {faqs &&
                            faqs.map((faq) => (
                                <ViewAllFaqCard
                                    key={faq.question}
                                    question={faq.question}
                                    answer={faq.answer}
                                />
                            ))}
                    </View>
                </View>
            </ScrollView>
        </SafeAreaView>
    );
};

export default AllFAQs;
