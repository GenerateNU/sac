import { useLocalSearchParams } from 'expo-router';
import { StatusBar } from 'expo-status-bar';
import React, { useRef } from 'react';
import { Animated, LogBox, RefreshControl, ScrollView, Text, View } from 'react-native';

import { TagList } from '@/components/all-tags';
import { Description } from '@/components/description';
import { HeaderImage } from '@/components/entity-header';
import { Title } from '@/components/title';
import { useAuthStore } from '@/hooks/use-auth';
import { useClubTags } from '@/hooks/use-club';
import { useClubQuery, useClubTagsQuery } from '@sac/lib';

import { Events } from './_components/all-events';
import { FAQs } from './_components/all-faqs';
import ApplyBottomSheet from './_components/application-form-bottom-sheet';
import { ClubAdmins } from './_components/club-admin-list';
import { ClubInsights } from './_components/club-insights';

LogBox.ignoreLogs(['Warning: ...']); // Ignore log notification by message
LogBox.ignoreAllLogs(); //Ignore all log notifications

const randomEventColorData = [
    {
        image: 'https://i.imgur.com/TM29405.png',
        color: '#edadc6'
    },
    {
        image: 'https://i.imgur.com/IfjHhrs.png',
        color: '#b0e0eb'
    },
    {
        image: 'https://i.imgur.com/TQQ7AEE.png',
        color: '#adc9ed'
    },
    {
        image: 'https://i.imgur.com/Gp0Wwa8.png',
        color: '#F0C584'
    }
];

const Club = () => {
    const { id } = useLocalSearchParams<{ id: string }>();
    const { user } = useAuthStore();

    const ref = useRef<BottomSheet>(null);
    const pan = useRef(new Animated.ValueXY()).current;

    const { data: club, isLoading, error, refetch } = useClubQuery(id);

    const randomEventColor =
        randomEventColorData[
        Math.round(Math.random() * (randomEventColorData.length - 1))
        ];

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error</Text>;
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
                        image={randomEventColor.image}
                        logo={randomEventColor.image}
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
                            padding: '5%',
                            paddingTop: '2%',
                            gap: 30
                        }}
                    >
                        <View
                            style={{
                                display: 'flex',
                                flexDirection: 'column',
                                gap: 20
                            }}
                        >
                            <Title title={club.name} />
                            <TagList
                                id={id}
                                tagQuery={useClubTags}
                                color={randomEventColor.color}
                            />
                        </View>
                        <ClubInsights />
                        <View>
                            <Description
                                title={'About the Club'}
                                description={club.description}
                            />
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
