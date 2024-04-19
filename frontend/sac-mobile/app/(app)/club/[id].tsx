import React, { useRef, useState } from 'react';
import { Alert, Animated, Platform, RefreshControl, ScrollView, SectionList, Share, Text, Touchable, TouchableOpacity, View } from 'react-native';
import { Button } from '@/components/button';

import { useAuthStore } from '@/hooks/use-auth';
import { SafeAreaView } from 'react-native-safe-area-context';
import ExpandableText from '@/components/expandable-text';
import { ClubAdmins } from './_components/club-admin-list';
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
import { StatusBar } from 'expo-status-bar';
import { Title } from '@/components/title';
import { ClubInsights } from './_components/club-insights';

const Club = () => {
    const { id } = useLocalSearchParams<{ id: string }>();
    const { user } = useAuthStore();

    const ref = useRef<BottomSheet>(null);
    const pan = useRef(new Animated.ValueXY()).current;

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

    return (
        <>
            <StatusBar style="light" />
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
                        userID={user?.id || ''}
                        clubID={id}
                    />
                    <View
                        style={{
                            display: 'flex',
                            flexDirection: 'column',
                            padding: '3%',
                            paddingTop: '2%',
                            gap: 30
                        }}
                    >
                        <Title title={club.name} />
                        {/* <SocialMediaIcons /> */}
                        <TagList id={id} tagQueryResult={useClubTags} />
                        {/* <View>
                            <Button variant='gray' size='screenwide' className='flex-col justify-between mb-4' onPress={handleOpenPress}>
                                <Text className='font-bold text-left text-black'>Application Form</Text>
                        </Button>
                        </View> */}

                        <ClubInsights />
                        <View>
                            <View className='pb-[2%]'>
                                <Text className="text-lg font-semibold text-black ">Who we are</Text>
                            </View>
                            <ExpandableText maxLines={5}>
                                {club.description}
                            </ExpandableText>
                        </View>

                        <ClubAdmins clubID={id} />
                        <Events clubID={id} />
                        <FAQs clubID={id} />
                    </View>
                </View>
            </ScrollView>
            <ApplyBottomSheet ref={ref} />
        </>
    );
};


// export const ClubHeader = ({ userID, clubID }: { userID: string; clubID: string }) => {
//     const followClub = useCreateClubFollowing();
//     const unfollowClub = useDeleteClubFollowing();
//     const userFollowing = useUserFollowing(userID);

//     const isFollowing = userFollowing.data?.map((club) => club.id).includes(clubID);
//     const [buttonClicked, setButtonClicked] = useState(isFollowing ? 'Following' : 'Follow');

//     const handlePressFollow = async () => {
//         if (buttonClicked === 'Following') {
//             await unfollowClub.mutateAsync({
//                 userID: userID,
//                 clubID: clubID
//             });
//             setButtonClicked('Follow');
//         } else {
//             await followClub.mutateAsync({
//                 userID: userID,
//                 clubID: clubID
//             });
//             setButtonClicked('Following');
//         }
//     };

//     return (
//         <View className="flex-row items-center justify-between pb-2">
//             <SandboxLogo width={100} height={100} />
//             <View className="flex-row space-x-3">
//                 <Button
//                     size={'default'}
//                     variant={'default'}
//                     className="w-18"
//                     onPress={handlePressFollow}
//                 >
//                     {buttonClicked === 'Following' ? 'Following' : 'Follow'}
//                 </Button>
//             </View>
//         </View>
//     );
// };

export default Club;