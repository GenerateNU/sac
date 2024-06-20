import { useEffect, useRef } from 'react';
import { Dimensions, Linking } from 'react-native';
import Animated, {
    interpolate,
    useAnimatedRef,
    useAnimatedStyle,
    useScrollViewOffset
} from 'react-native-reanimated';

import { Stack, useLocalSearchParams } from 'expo-router';

import { faExternalLink } from '@fortawesome/free-solid-svg-icons';
import BottomSheet from '@gorhom/bottom-sheet';

import {
    Arrow,
    Box,
    PageTags,
    RecruitmentInfo,
    Text
} from '@/src/app/design-system';
import { SACColors } from '@/src/app/design-system';
import { Button } from '@/src/app/design-system/components/Button/Button';
import useClub from '@/src/hooks/useClub';
import { setClubId } from '@/src/store/slices/clubSlice';
import { useAppDispatch, useAppSelector } from '@/src/store/store';

import { AboutSection } from '../../design-system/components/AboutSection/AboutSection';
import AnimatedImageHeader from '../../design-system/components/AnimatedImageHeader/AnimatedImageHeader';
import { ClubIcon } from '../../design-system/components/ClubIcon/ClubIcon';
import { EventCard } from '../../design-system/components/EventCard';
import { EventCardList } from '../../design-system/components/EventCard/EventCardList';
import PageError from '../../design-system/components/PageError/PageError';
import { Description } from '../event/components/description';
import ClubPageSkeleton from './components/skeleton';
import { clubApi } from '@generatesac/lib';
import { setUserFollowing } from '@/src/store/slices/userSlice';

const color: SACColors = 'darkRed';

const ClubPage = () => {
    const { id } = useLocalSearchParams<{ id: string }>();
    const { width } = Dimensions.get('window');
    const IMG_HEIGHT = width;

    const scrollRef = useAnimatedRef<Animated.ScrollView>();
    const scrollOffset = useScrollViewOffset(scrollRef);

    const bottomSheet = useRef<BottomSheet>(null);

    const club = useAppSelector((state) => state.club);
    const { id: user_id, following } = useAppSelector((state) => state.user)
    const dispatch = useAppDispatch();
    const { setRetriggerFetch, apiLoading, apiError } = useClub(id);
    const [followClub] = clubApi.useCreateClubFollowerMutation();
    const [unFollowClub] = clubApi.useDeleteClubFollowerMutation();
    
    const clubFollowed = following.includes(id as string);

    const handleUserFollow = (follow: boolean) => {
        if (id) {
            if (follow) {
                followClub({ club_id: id, user_id })
                    .then(({ error }) => {
                        if (!error) {
                            dispatch(setUserFollowing([...following, id]));
                        }
                    })
            } else {
                unFollowClub({ club_id: id, user_id })
                    .then(({ error }) => {
                        if (!error) {
                            dispatch(setUserFollowing(following.filter(clubId => clubId !== id)));
                        }
                    })
            }
        }
    };

    const headerAnimatedStyle = useAnimatedStyle(() => {
        return {
            transform: [
                {
                    translateY: interpolate(
                        scrollOffset.value,
                        [-IMG_HEIGHT, 0, IMG_HEIGHT],
                        [0, 0, -IMG_HEIGHT * 0.75]
                    )
                }
            ]
        };
    });

    return (
        <Box height="100%" backgroundColor="white">
            <Stack.Screen
                options={{
                    headerTitle: '',
                    headerTransparent: true,
                    headerShown: true,
                    headerTitleStyle: {
                        color: 'white'
                    },
                    headerLeft: () => (
                        <Animated.View style={headerAnimatedStyle}>
                            <Arrow color={apiError ? 'black' : 'white'} />
                        </Animated.View>
                    )
                }}
            />
            <Animated.ScrollView
                showsVerticalScrollIndicator={false}
                showsHorizontalScrollIndicator={false}
                scrollEventThrottle={16}
                ref={scrollRef}
            >
                {apiLoading || club.id !== id ? (
                    <ClubPageSkeleton />
                ) : apiError ? (
                    <PageError refetch={setRetriggerFetch} />
                ) : (
                    <>
                        <AnimatedImageHeader
                            uri="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSLF3ord7lnV_5Je-pC2AUgUiesHNPcZlpI7A&s"
                            scrollOffset={scrollOffset}
                            imageHeight={IMG_HEIGHT}
                            aspectRatio={2.5}
                        />
                        <Box backgroundColor="white">
                            <Box
                                flexDirection="row"
                                justifyContent="space-between"
                                marginTop="negativeXl"
                                alignItems="center"
                                paddingVertical="s"
                                paddingHorizontal="l"
                            >
                                <ClubIcon imageUrl={club.logo as string} />
                                <Box width={116}>
                                    <Button
                                        color={color}
                                        size="sm"
                                        variant="standardButton"
                                        onPress={() => handleUserFollow(!clubFollowed)}
                                    >
                                        {clubFollowed ? 'Unfollow' : 'Follow'}
                                    </Button>
                                </Box>
                            </Box>
                            <Box
                                marginHorizontal="l"
                                marginVertical="m"
                                gap="m"
                            >
                                <Text variant="header-1">{club.name}</Text>
                                <PageTags tags={club.tags} color={color} />
                                <AboutSection
                                    description={club.description}
                                    onPress={() =>
                                        bottomSheet.current?.snapToIndex(0)
                                    }
                                    type="club"
                                />
                                {club.recruitment?.is_recruiting && (
                                    <Button
                                        variant="iconButton"
                                        icon={faExternalLink}
                                        color={color}
                                        iconPosition="right"
                                        size="sm"
                                        justifyContent="space-between"
                                        flexDirection="row"
                                        onPress={() => Linking.openURL('')}
                                    >
                                        <Box
                                            paddingBottom="xxxs"
                                            paddingTop="xxs"
                                            gap="xxs"
                                        >
                                            <Text
                                                color="white"
                                                variant="body-1"
                                            >
                                                Application Form
                                            </Text>
                                            <Text
                                                color="white"
                                                variant="caption-1"
                                            >
                                                Google Form
                                            </Text>
                                        </Box>
                                    </Button>
                                )}
                                <Box>
                                    <Text variant="subheader-1">
                                        Recruiting
                                    </Text>
                                    <RecruitmentInfo
                                        recruitingType={'application'}
                                        isRecruiting={false}
                                        recruitmentCycle={'fallSpring'}
                                        color={color}
                                    />
                                </Box>
                                <Box gap="l">
                                    <Text variant="subheader-1">
                                        Upcoming Events
                                    </Text>
                                    {club.events.length > 0 ? (
                                        <>
                                            <EventCard
                                                eventId={club.events[0].id}
                                                event={club.events[0].name}
                                                variant="big"
                                                tags={[]}
                                                club={club.name}
                                                startTime={
                                                    club.events[0].start_time
                                                }
                                                endTime={
                                                    club.events[0].end_time
                                                }
                                                logo={club.logo as string}
                                                image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                                            />
                                            <EventCardList
                                                events={club.events.slice(1)}
                                                club={club.name}
                                            />
                                        </>
                                    ) : (
                                        <></>
                                    )}
                                </Box>
                            </Box>
                        </Box>
                    </>
                )}
            </Animated.ScrollView>
            <Description
                description={club.description as string}
                ref={bottomSheet}
                type="club"
            />
        </Box>
    );
};

export default ClubPage;
