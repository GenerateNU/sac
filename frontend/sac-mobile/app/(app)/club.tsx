import React from 'react';
import { ScrollView, Text, View } from 'react-native';
import { Button } from '@/components/button';

import { useAuthStore } from '@/hooks/use-auth';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Card } from '@/components/card';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { EBoardCard } from '@/components/eboardCard';
import { FaqCard } from '@/components/faqCard';

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
                <View className="pb-[10%]">
                    <View className="pt-[1%]">
                    </View>
                    <View className="pt-[20%] pb-[6%]">
                    </View>
                </View>

                <ScrollView className="bg-white pt-[20%] pb-[20%] flex-1 px-[8%]">
                    <View className="flex-row">
                        <Text className="text-black font-bold text-3xl">Club Name</Text>
                    </View>
                    {/* Rewrite with maybe the button template? */}
                    <View className='flex-row pt-[5%] pb-[5%]'>
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
                        <View className='pb-[2%]'>
                            <Text className=" text-black font-bold">Description</Text>
                        </View>
                        <Text className="text-black">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</Text>
                    </View>
                    <View>
                        <Text className="text-black font-bold">Events</Text>
                        <ScrollView horizontal={true} className="pt-[2%] pb-[10%]">
                            <View className="flex-row">
                                {/* To be replaced by event components */}
                                <Card variant="default" size="default" className="mr-2">
                                    <Text className="text-black">Event 1</Text>
                                </Card>
                            </View>
                        </ScrollView>
                    </View>
                    <View>
                        <Text className="text-black font-bold">E-Board</Text>
                        <ScrollView horizontal={true} className="pb-[10%]">
                            <View className="flex-row">
                                <EBoardCard name="Garrett Ladley" title="Tech Lead" className="mr-2">
                                    <Text className="text-black">Member1</Text>
                                </EBoardCard>
                                <EBoardCard name="David Oduneye" title="Tech Lead" className="mr-2">
                                    <Text className="text-black">Member1</Text>
                                </EBoardCard>
                            </View>
                        </ScrollView>
                    </View>
                    <View className = 'pb-[30%]'>
                        <Text className="text-black font-bold">FAQs</Text>
                        <ScrollView horizontal={true} className = 'pb-[10%]'>
                            <View className="flex-row">
                                <FaqCard question="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididun?"
                                    answer="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad "
                                    variant="default" size="default" className="mr-2"> </FaqCard>
                            </View>
                        </ScrollView>
                        <Button variant="outline">Ask a Question</Button>
                    </View>

                </ScrollView>
            </View>
        </SafeAreaView>
    );
};

export default Club;
