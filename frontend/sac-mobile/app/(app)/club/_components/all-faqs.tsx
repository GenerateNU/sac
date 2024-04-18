import { Button } from '@/components/button';
import { FaqCard, FaqCardProps } from '@/components/faqCard';
import { useClubFAQs } from '@/hooks/use-club';
import { Link, router } from 'expo-router';
import { ScrollView, Text, TouchableOpacity, View } from 'react-native';

const FAQs = () => {
    const { data: faqs, isLoading, error } = useClubFAQs("sac");

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View>
            <View className="flex-row items-center justify-between">
                <Text className="text-lg font-semibold text-black">FAQ</Text>
                <Text className='text-neutral-400 text-[15px] font-normal'>View All</Text>
            </View>
            <ScrollView className="pt-[2%] pb-[10%]" horizontal={true}>
                {faqs && faqs.slice(0, 2).map((faq: FaqCardProps, index) => (
                    <FaqCard {...faq} key={index} />
                ))}
                <View className="flex-col">
                    <View className='px-2'>
                        <Button className='bg-[#f8c45c]' variant="secondary" size="screenwide">View All Events</Button>
                    </View>
                </View>
            </ScrollView>
        </View>

        // <View className='pb-[10%]'>
        //     <View className='flex-row justify-between'>
        //         <Text className="text-lg font-semibold text-black">FAQ</Text>
        //         <TouchableOpacity onPress={() => router.push('/(app)/club/faq/view-all-faqs')}>
        //             <Text>View All</Text>
        //         </TouchableOpacity>
        //     </View>
        //     <ScrollView horizontal={true} className="pb-[10%]">
        //         <View className="flex-row">
        //             {faqs && faqs.map((faq: FaqCardProps) => (
        //                 <FaqCard id={faq.id} question={faq.question} answer={faq.answer} />
        //             ))}
        //         </View>
        //     </ScrollView>
        //     <Button variant="outline" onPress={() => router.push("/(app)/club/ask/ask-a-question")}>Ask a Question</Button>
        // </View>
    );
}

export { FAQs };