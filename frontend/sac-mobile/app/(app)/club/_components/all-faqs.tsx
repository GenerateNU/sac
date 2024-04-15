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
        <View className='pb-[10%]'>
            <View className='flex-row justify-between'>
                <Text className="text-black font-bold pb-4">FAQs</Text>
                <TouchableOpacity onPress={() => router.push('/(app)/club/faq/view-all-faqs')}>
                    <Text>View All</Text>
                </TouchableOpacity>
            </View>
            <ScrollView horizontal={true} className="pb-[10%]">
                <View className="flex-row">
                    {faqs && faqs.map((faq: FaqCardProps) => (
                        <FaqCard id={faq.id} question={faq.question} answer={faq.answer} />
                    ))}
                </View>
            </ScrollView>
            <Button variant="outline" onPress={() => router.push("/(app)/club/ask/ask-a-question")}>Ask a Question</Button>
        </View>
    );
}

export { FAQs };