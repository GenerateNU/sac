import React, { useRef } from 'react';
import { Animated, RefreshControl, ScrollView, Text, View } from 'react-native';

import { useLocalSearchParams } from 'expo-router';
import { StatusBar } from 'expo-status-bar';

import BottomSheet from '@gorhom/bottom-sheet/lib/typescript/components/bottomSheet/BottomSheet';

import { TagList } from '@/components/all-tags';
import { HeaderImage } from '@/components/entity-header';
import ExpandableText from '@/components/expandable-text';
import { Title } from '@/components/title';
import { useAuthStore } from '@/hooks/use-auth';
import { useClub, useClubTags } from '@/hooks/use-club';

import { Events } from './_components/all-events';
import { FAQs } from './_components/all-faqs';
import ApplyBottomSheet from './_components/application-form-bottom-sheet';
import { ClubAdmins } from './_components/club-admin-list';
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
                        {/* <SocialMediaIcons /> REMOVED Temporary needs new designs*/}
                        <TagList id={id} tagQueryResult={useClubTags} />
                        {/*
                         <View>
                            <Button variant='gray' size='screenwide' className='flex-col justify-between mb-4' onPress={handleOpenPress}>
                                <Text className='font-bold text-left text-black'>Application Form</Text>
                            </Button>
                        </View>
                        */}

                        <ClubInsights />
                        <View>
                            <View className="pb-[2%]">
                                <Text className="text-lg font-semibold text-black ">
                                    Who we are
                                </Text>
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
