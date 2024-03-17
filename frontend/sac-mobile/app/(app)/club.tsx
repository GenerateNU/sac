import React from 'react';
import { ScrollView, Text, View } from 'react-native';
import { Button } from '@/components/button';

import { useAuthStore } from '@/hooks/use-auth';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Card } from '@/components/card';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { EBoardCard } from '@/components/eboardCard';
import { FaqCard, faqCard } from '@/components/faqCard';

// import SlackIcon from '@/components/icons/SlackIcon';

const SlackIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="slack" size={24} color={color} />
);

const EmailIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="email" size={24} color={color} />
);

const InstagramIcon = ({ color }: { color: string }) => (
    <MaterialCommunityIcons name="instagram" size={24} color={color} />
);

const Club = () => {
    const { logout } = useAuthStore();

    // TODO: Implement social media click handling
    function handleSocialMediaClick(arg0: string): void {
        throw new Error('Function not implemented.');
    }

    return (
        <SafeAreaView className="bg-neutral-500 h-[100%]" edges={['top']}>
            <View className="flex-1">
                <View className="px-[8%] pb-[10%]">
                    <View className="pt-[1%]">
                    </View>
                    <View className="pt-[20%] pb-[6%]">
                    </View>
                </View>

                <ScrollView className="bg-white pt-[13%] pb-[2%] flex-1 rounded-tl-3xl rounded-tr-3xl px-[8%]">
                    <View className="pb-[8%] flex-row justify-between">
                        <Text className="text-black font-bold text-4xl">Club Name</Text>
                    </View>
                    {/* Rewrite with maybe the button template? */}
                    <View style={{ flexDirection: 'row' }}>
                        <Button variant="outline" onPress={() => handleSocialMediaClick('email')}>
                            <EmailIcon color="black" />
                        </Button>

                        <Button variant="outline" onPress={() => handleSocialMediaClick('instagram')}>
                            <InstagramIcon color="black" />
                        </Button>

                        <Button variant="outline" onPress={() => handleSocialMediaClick('slack')}>
                            <SlackIcon color="black" />
                        </Button>
                    </View>

                    <View className="pb-[8%]">
                        <Text className="text-black font-bold">Club Description</Text>
                        <Text className="text-black">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</Text>
                    </View>
                    <ScrollView horizontal={true} className="pb-[10%]">
                        <View className="flex-row">
                            {/* To be replaced by event components */}
                            <Card variant="default" size="default" className="mr-2">
                                <Text className="text-black">Event 1</Text>
                            </Card>

                        </View>
                    </ScrollView>

                    <ScrollView horizontal={true} className="pb-[10%]">
                        <View className="flex-row">
                            {/* To be replaced by e-board components */}
                            <EBoardCard name = "Garrett Ladley" title = "Tech Lead" variant="default" size="default" className="mr-2">
                                <Text className="text-black">Member1</Text>
                            </EBoardCard>
                            <EBoardCard name = "David Oduneye" title = "Tech Lead" variant="default" size="default" className="mr-2">
                                <Text className="text-black">Member1</Text>
                            </EBoardCard>
                        </View>
                    </ScrollView>

                    <ScrollView horizontal={true}>
                        <View className="flex-row">
                            <FaqCard question="Question 1" answer="Answer 1" variant="default" size="default" className="mr-2"> </FaqCard>
                        </View>
                    </ScrollView>

                    <Button variant="outline">Ask a Question</Button>

                </ScrollView>
            </View>
        </SafeAreaView>
    );
};

export default Club;
