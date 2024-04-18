import React, { useRef, useState } from 'react';
import { Alert, Animated, Platform, RefreshControl, ScrollView, SectionList, Share, Text, Touchable, TouchableOpacity, View } from 'react-native';
import { Button } from '@/components/button';

import { useAuthStore } from '@/hooks/use-auth';
import { SafeAreaView } from 'react-native-safe-area-context';
import ExpandableText from '@/components/expandable-text';
import { EBoardMembers } from './_components/all-eboard-members';
import { FAQs } from './_components/all-faqs';
import { SocialMediaIcons } from './_components/all-social-media';
import { Events } from './_components/all-events';
import { Link, router, Stack, useLocalSearchParams } from 'expo-router';
import { MenuView } from '@react-native-menu/menu';
import BottomSheet from '@gorhom/bottom-sheet/lib/typescript/components/bottomSheet/BottomSheet';
import ApplyBottomSheet from './_components/application-form-bottom-sheet';
import { useClub, useClubTags } from '@/hooks/use-club';

import SandboxLogo from '@/assets/images/svg/sandbox-logo.svg';
import SandboxHeader from '@/assets/images/svg/sandbox-header.svg';
import { TagList } from '@/components/all-tags';

import CircleClock from '@/assets/images/svg/circle-clock.svg';
import MemberCount from '@/assets/images/svg/members.svg';
import PeaceHand from '@/assets/images/svg/peace-hand.svg';
import { pluralize } from '@/lib/utils';
import { useCreateClubFollowing, useDeleteClubFollowing, useUserFollowing } from '@/hooks/use-user';
import { HeaderImage } from '@/components/entity-header';

const Club = () => {
    const { id } = useLocalSearchParams<{ id: string }>();
    const { user } = useAuthStore();

    const ref = useRef<BottomSheet>(null);
    const pan = useRef(new Animated.ValueXY()).current;

    const handleOpenPress = () => { ref.current?.snapToIndex(0); };

    const { data: club, isLoading, error, refetch } = useClub(id);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    if (!club) {
        return <Text>Club not found</Text>;
    }

    const membersRange = club.num_members > 100 ? '100+' : club.num_members > 50 ? '50 - 100' : club.num_members > 25 ? '25 - 50' : '1 - 25';

    return (
        // TODO: this should be a section view
        <>
            <ScrollView
                refreshControl={
                    <RefreshControl
                        refreshing={isLoading}
                        onRefresh={refetch}
                    />
                }
                scrollEventThrottle={1}
                showsVerticalScrollIndicator={false}
                onScroll={Animated.event(
                    [{ nativeEvent: { contentOffset: { y: pan.y } } }],
                    {
                        useNativeDriver: false
                    }
                )}
            >
                <View className="flex-col justify-between flex-1 w-full">
                    <HeaderImage
                        image="https://blenderartists.org/uploads/default/optimized/4X/7/e/2/7e2d7bea4ac21388c4a96e1371f375c4ce00094b_2_1024x576.jpg"
                        pan={pan}
                        type="club"
                        ref={ref}
                    />
                    <View className="w-full px-3 pb-5">
                        {/* <ClubHeader userID={user?.id || ''} clubID={id} /> */}
                        <View className="flex-row">
                            <Text className="text-3xl font-bold text-black">{club.name}</Text>
                        </View>
                        {/* <SocialMediaIcons /> */}
                        <TagList id={id} tagQueryResult={useClubTags} />
                        {/* <View>
                            <Button variant='gray' size='screenwide' className='flex-col justify-between mb-4' onPress={handleOpenPress}>
                                <Text className='font-bold text-left text-black'>Application Form</Text>
                            </Button>
                        </View> */}

                        <View className='flex-row pb-5 gap-x-2'>
                            <View className='pt-3 pb-2 bg-white rounded-lg shadow-sm'>
                                <View className='flex-col items-center justify-center gap-y-2' style={{ width: 118 }}>
                                    <CircleClock width={24} height={24} />
                                    <Text className='font-semibold text-center text-black'>5 {pluralize(5, 'Hour', 'Hours')}</Text>
                                    <Text className='text-xs font-normal text-center text-neutral-400'>Weekly Time Commitment</Text>
                                </View>
                            </View>
                            <View className='pt-3 pb-2 bg-white rounded-lg shadow-sm'>
                                <View className='flex-col items-center justify-center gap-y-2' style={{ width: 118 }}>
                                    <MemberCount width={24} height={24} />
                                    <Text className='font-semibold text-center text-black'>{membersRange}</Text>
                                    <View>
                                        <Text className='text-xs font-normal text-center text-neutral-400 '>Member</Text>
                                        <Text className='text-xs font-normal text-center text-neutral-400 '>Count</Text>
                                    </View>
                                </View>
                            </View>
                            <View className='pt-3 pb-2 bg-white rounded-lg shadow-sm'>
                                <View className='flex-col items-center justify-center gap-y-2' style={{ width: 118 }}>
                                    <PeaceHand width={24} height={24} />
                                    <Text className='font-semibold text-center text-black'>Builders</Text>
                                    <Text className='text-xs font-normal text-center text-neutral-400 '>One Word To Describe Us</Text>
                                </View>
                            </View>
                        </View>
                        <View className="pb-[8%]">
                            <View className='pb-[2%]'>
                                <Text className="text-lg font-semibold text-black ">Who we are</Text>
                            </View>
                            <ExpandableText maxLines={5}>
                                {club.description}
                            </ExpandableText>
                        </View>


                        <EBoardMembers />
                        <Events />
                        <FAQs />
                    </View>
                </View>
            </ScrollView>
            <ApplyBottomSheet ref={ref} />
        </>
    );
};

export const ClubHeader = ({ userID, clubID }: { userID: string; clubID: string }) => {
    const followClub = useCreateClubFollowing();
    const unfollowClub = useDeleteClubFollowing();
    const userFollowing = useUserFollowing(userID);

    const isFollowing = userFollowing.data?.map((club) => club.id).includes(clubID);
    const [buttonClicked, setButtonClicked] = useState(isFollowing ? 'Following' : 'Follow');

    const handlePressFollow = async () => {
        if (buttonClicked === 'Following') {
            await unfollowClub.mutateAsync({
                userID: userID,
                clubID: clubID
            });
            setButtonClicked('Follow');
        } else {
            await followClub.mutateAsync({
                userID: userID,
                clubID: clubID
            });
            setButtonClicked('Following');
        }
    };

    return (
        <View className="flex-row items-center justify-between pb-2">
            <SandboxLogo width={100} height={100} />
            <View className="flex-row space-x-3">
                <Button
                    size={'default'}
                    variant={'default'}
                    className="w-18"
                    onPress={handlePressFollow}
                >
                    {buttonClicked === 'Following' ? 'Following' : 'Follow'}
                </Button>
            </View>
        </View>
    );
};

export default Club;