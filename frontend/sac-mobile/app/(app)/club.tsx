import React, { useState } from 'react';
import {ScrollView, Text, View } from 'react-native';
import { Button } from '@/components/button';

import { useAuthStore } from '@/hooks/use-auth';
import { SafeAreaView } from 'react-native-safe-area-context';
import { EventCard } from '@/components/eventCard';
import { AskAQuestionModal } from './_components/application-modal';
import ExpandableText from '@/components/expandable-text';
import { AllClubTags } from './_components/all-club-tags';
import { EBoardMembers } from './_components/all-eboard-members';
import { FAQs } from './_components/all-faqs';
import { SocialMediaIcons } from './_components/all-social-media';

const Club = () => {
    const { logout } = useAuthStore();

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
                <View className="bg-white pt-[18%] px-[6%]">
                    <Button className="aspect-square rounded-lg w-[22%] absolute translate-y-9 right-9">Follow</Button>
                    <View className="flex-row">
                        <Text className="text-black font-bold text-3xl">Club Name</Text>
                    </View>
                    <SocialMediaIcons></SocialMediaIcons>
                    <AllClubTags></AllClubTags>
                    <View>
                        <Button variant='gray' size='screenwide' className='mb-4'>
                            Application Form
                        </Button>
                    </View>

                    <View className="pb-[8%]">
                        <View className='pb-[2%]'>
                            <Text className=" text-black font-medium">Description</Text>
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
                                <EventCard title={'Orientation'} time={'8:00am-9:30am'} date={'February 29'} location={'Ryder Hall, Room 294'}></EventCard>
                                <EventCard title={'Orientation'} time={'8:00am-9:30am'} date={'February 29'} location={'Ryder Hall, Room 294'}></EventCard>
                                <EventCard title={'Orientation'} time={'8:00am-9:30am'} date={'February 29'} location={'Ryder Hall, Room 294'}></EventCard>
                            </View>
                        </ScrollView>
                    </View>
                    <EBoardMembers></EBoardMembers>
                    <FAQs></FAQs>
                </View>
            </ScrollView>
            {isModalVisible && <AskAQuestionModal onClose={() => setIsModalVisible(false)} />}
        </SafeAreaView>
    );
};

export default Club;