import { Button } from '@/components/button';
import { FaqCard, FaqCardProps } from '@/components/faqCard';
import { router } from 'expo-router';
import { ScrollView, Text, View } from 'react-native';

const FAQs = () => {
    const { data: faqs, isLoading, error } = useFAQs();

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
                    {faqs.map((faq: FaqCardProps) => (
                        <FaqCard key={faq.id} question={faq.question} answer={faq.answer} />
                    ))}
                </View>
            </ScrollView>
            <Button variant="outline" onPress={() => router.push("/(app)/ask-a-question")}>Ask a Question</Button>
        </View>
    );
}

export { FAQs };

function useFAQs(): { data: FaqCardProps[], isLoading: boolean, error: Error | null } {
    return {
        data: [
            { question: 'What is the purpose of the club?', answer: 'The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field. The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field. The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field.' },
            { question: 'How can I join the club?', answer: 'To join the club, you can attend our regular meetings and express your interest in becoming a member. We welcome anyone who shares our passion and is willing to contribute.' },
            { question: 'What activities does the club organize?', answer: 'The club organizes various activities such as workshops, guest lectures, competitions, and networking events to enhance the learning and networking opportunities for its members.' },
        ],
        isLoading: false,
        error: null
    };
}