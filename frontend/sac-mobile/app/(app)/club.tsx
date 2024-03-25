import React, { useState } from 'react';
import { GestureResponderEvent, ScrollView, Text, View } from 'react-native';
import { Button } from '@/components/button';

import { useAuthStore } from '@/hooks/use-auth';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Card } from '@/components/card';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { EBoardCard } from '@/components/eboardCard';
import { FaqCard } from '@/components/faqCard';
import { AskAQuestionModal } from './_components/ask-a-question';
import ExpandableText from '@/components/expandable-text';
import { ClubProfileTag } from '@/components/club-profile-tag';

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

    const [isModalVisible, setIsModalVisible] = useState(false);
    const openModal = () => {
        setIsModalVisible(true);
    }

    return (
        <SafeAreaView className="bg-neutral-500 h-[100%]" edges={['top']}>
            <ScrollView className="flex-1">
                {/* Div for the grey */}
                <View className="pt-[40%] z-10">
                    <View className="aspect-square rounded-lg w-20 bg-card-bg absolute top-20 translate-y-14 left-10" />
                </View>
                {/* Div for the club */}
                <View className="bg-white pt-[18%] px-[8%]">
                    <View className="flex-row">
                        <Text className="text-black font-bold text-3xl">Club Name</Text>
                    </View>
                    <View>
                        <ScrollView horizontal={true} className="pt-[2%]">
                            <View className="flex-row">
                                <ClubProfileTag name="Tag1" />
                                <ClubProfileTag name="Tag2" />
                                <ClubProfileTag name="Tag3" />
                            </View>
                        </ScrollView>
                    </View>
                    {/* Rewrite with maybe the button template? */}
                    <View className='flex-row pt-[5%] pb-[5%]'>
                        <Button variant="secondary" onPress={() => handleSocialMediaClick('email')}>
                            <EmailIcon color="black" />
                        </Button>

                        <Button variant="secondary" onPress={() => handleSocialMediaClick('instagram')}>
                            <InstagramIcon color="black" />
                        </Button>

                        <Button variant="secondary" onPress={() => handleSocialMediaClick('slack')}>
                            <SlackIcon color="black" />
                        </Button>
                    </View>
                    <View>
                    <Button variant='gray' size='screenwide' className='mb-4'>
                            Application Form
                        </Button>
                    </View>

                    <View className="pb-[8%]">
                        <View className='pb-[2%]'>
                            <Text className=" text-black font-bold">Description</Text>
                        </View>
                        <ExpandableText maxLines={5}>
                            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                        </ExpandableText>
                    </View>
                    <View>
                        <Text className="text-black font-bold">Events</Text>
                        <ScrollView horizontal={true} className="pt-[2%] pb-[10%]">
                            <View className="flex-row">
                                {/* To be replaced by event components */}
                                <Card title="Orientation" variant="default" size="default" className="mr-2">
                                    <Text className='text-gray-600'>{"08:00 - 09:30 PM"}</Text>
                                    <Text className='text-gray-600'>{"April 23, 2024"}</Text>
                                    <Text className='text-gray-600'>{"Ryder Hall, Room 294"}</Text>
                                </Card>
                                <Card title="Orientation" variant="default" size="default" className="mr-2">
                                    <Text className='text-gray-600'>{"08:00 - 09:30 PM"}</Text>
                                    <Text className='text-gray-600'>{"April 23, 2024"}</Text>
                                    <Text className='text-gray-600'>{"Ryder Hall, Room 294"}</Text>
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
                    <View className='pb-[30%]'>
                        <Text className="text-black font-bold pb-4">FAQs</Text>
                        <ScrollView horizontal={true} className='pb-[10%]'>
                            <View className="flex-row">
                                <FaqCard question="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididun?"
                                    answer="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad "
                                    className="mr-2"> </FaqCard>
                            </View>
                        </ScrollView>
                        <Button variant="outline" onPress={openModal}>Ask a Question</Button>
                    </View>
                </View>
            </ScrollView>
            {isModalVisible && <AskAQuestionModal onClose={() => setIsModalVisible(false)} />}
        </SafeAreaView>
    );
};

export default Club;