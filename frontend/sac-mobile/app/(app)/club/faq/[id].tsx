import React from 'react';
import { SafeAreaView, Text } from 'react-native';

import { useLocalSearchParams } from 'expo-router';
import { useClubFAQ } from '@/hooks/use-club';

const Faq = () => {
    // refetch data based on the id

    const { id } = useLocalSearchParams<{ id: string }>();

    const { data: faq, isLoading, error } = useClubFAQ(id);

    return (
        <SafeAreaView className="items-center flex-1 bg-white px-5">
            <Text className=''>{id}</Text>
            <Text className='font-semibold text-2xl px-5'>{faq?.question}</Text>
            <Text className='px-5 text-lg'>{faq?.answer}</Text>
        </SafeAreaView>
    );
};

export default Faq;

// import { View, Text } from 'react-native'
// import React from 'react'
// import { useLocalSearchParams } from 'expo-router'
// import { useFAQ } from '@/hooks/use-faqs'

// const FAQ = ({ question, answer }: { question: string; answer: string }) => {

//     return (
//         <View className='flex-1 justify-center items-center bg-black/50'>
//             <View className='bg-white opacity-100 rounded-lg p-4 w-80'>
//                 <Text className="text-sm mb-2 font-semibold">{question}</Text>
//                 <Text className="text-sm">{answer}</Text>
//             </View>
//         </View>
//     )
// }

// export default FAQ
