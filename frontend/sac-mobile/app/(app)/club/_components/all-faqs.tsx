import { ScrollView, Text, TouchableOpacity, View } from 'react-native';

import { router } from 'expo-router';

import { FaqCard } from '@/components/faqCard';
import { useClubFAQs } from '@/hooks/use-club';

const FAQs = ({}: { clubID?: string }) => {
    // const { data: faqs, isLoading, error } = useClubFAQs(clubID);
    const { data: faqs, isLoading, error } = useClubFAQs('sac');

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    if (!faqs) {
        return <Text>FAQs not found</Text>;
    }

    return (
        <View className="flex-col w-full gap-y-1">
            <View className="flex-row items-center justify-between">
                <Text className="text-xl font-light leading-relaxed text-zinc-900">
                    FAQ
                </Text>
                <TouchableOpacity
                    onPress={() => router.push('/(app)/club/faq/view-all-faqs')}
                >
                    <Text className="text-neutral-400 text-[15px] font-normal">
                        View All
                    </Text>
                </TouchableOpacity>
            </View>
            <ScrollView className="pt-[2%]" horizontal={true}>
                {faqs &&
                    faqs.map((faq, index) => (
                        <FaqCard
                            key={index}
                            question={faq.question}
                            answer={faq.answer}
                        />
                    ))}
            </ScrollView>
        </View>
    );
};

export { FAQs };
