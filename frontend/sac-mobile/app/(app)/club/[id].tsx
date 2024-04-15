import React, { useState } from 'react';
import { Alert, Platform, ScrollView, Share, Text, Touchable, TouchableOpacity, View } from 'react-native';
import { Button } from '@/components/button';

import { useAuthStore } from '@/hooks/use-auth';
import { SafeAreaView } from 'react-native-safe-area-context';
import { AskAQuestionModal } from './_components/application-modal';
import ExpandableText from '@/components/expandable-text';
import { AllClubTags } from './_components/all-club-tags';
import { EBoardMembers } from './_components/all-eboard-members';
import { FAQs } from './_components/all-faqs';
import { SocialMediaIcons } from './_components/all-social-media';
import { Events } from './_components/all-events';
import { Link, router, Stack } from 'expo-router';
import { MenuView } from '@react-native-menu/menu';

const Club = () => {
    const { signOut } = useAuthStore();

    const [isModalVisible, setIsModalVisible] = useState(false);
    const openModal = () => {
        setIsModalVisible(true);
    }

    return (
        <ScrollView showsVerticalScrollIndicator={false}>
            <View className="flex-col justify-between flex-1 w-full">
                <View className="w-full bg-gray-600 h-[100]" />
                <View className="w-full pb-5 px-6">
                    <ClubHeader />
                    <View className="flex-row">
                        <Text className="text-black font-bold text-3xl">Club Name</Text>
                    </View>
                    <SocialMediaIcons/>
                    <AllClubTags />
                    <View>
                        <Button variant='gray' size='screenwide' className='mb-4 flex-col justify-between' onPress={() => router.push('/(app)/club/apply/apply-here')}>
                            <Text className='text-black font-bold text-left'>Application Form</Text>
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
                    <Events />
                    <EBoardMembers />
                    <FAQs />
                    {isModalVisible && <AskAQuestionModal onClose={() => setIsModalVisible(false)} />}
                </View>
            </View>
        </ScrollView>
    );
};

export const ClubHeader = () => {
    return (
        <View className="flex-row items-center justify-between -translate-y-10">
            <View className="w-24 h-24 bg-gray-300 rounded-xl" />
            <View className="flex-row space-x-3">

                <MenuView
                    onPressAction={({ nativeEvent }) => {
                        console.warn(JSON.stringify(nativeEvent));
                    }}
                    actions={[
                        {
                            id: 'yes',
                            title: 'Yes',
                            image: Platform.select({
                                ios: 'checkmark',
                                android: 'ic_menu_save'
                            })
                        },
                        {
                            id: 'maybe',
                            title: 'Maybe',
                            image: Platform.select({
                                ios: 'questionmark',
                                android: 'ic_menu_help'
                            })
                        },
                        {
                            id: 'no',
                            title: 'No',
                            image: Platform.select({
                                ios: 'xmark',
                                android: 'ic_menu_delete'
                            })
                        }
                    ]}
                    themeVariant="dark"
                    shouldOpenOnLongPress={false}
                    isAnchoredToRight={true}
                >
                    <Button
                        size={'default'}
                        variant={'default'}
                        className="w-18"
                    >
                        Follow
                    </Button>
                </MenuView>
            </View>
        </View>
    );
};

export default Club;