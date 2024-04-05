import { Button } from '@/components/button';
import { FaqCard, FaqCardProps } from '@/components/faqCard';
import { useFAQs } from '@/hooks/use-faqs';
import { router } from 'expo-router';
import { ScrollView, Text, View } from 'react-native';

const FAQs = () => {
    const { data: faqs, isLoading, error } = useFAQs("sac");

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View className='pb-[30%]'>
            <Text className="text-black font-bold pb-4">FAQs</Text>
            <ScrollView horizontal={true} className="pb-[10%]">
                <View className="flex-row">
                    {faqs && faqs.map((faq: FaqCardProps) => (
                        <FaqCard key={faq.id} question={faq.question} answer={faq.answer} />
                    ))}
                </View>
            </ScrollView>
            <Button variant="outline" onPress={() => router.push("/(app)/ask-a-question")}>Ask a Question</Button>
        </View>
    );
}

export { FAQs };