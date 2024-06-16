import { Dimensions, Linking } from 'react-native';
import Animated, {
    interpolate,
    useAnimatedRef,
    useAnimatedStyle,
    useScrollViewOffset
} from 'react-native-reanimated';

import { Arrow, Box, PageTags, RecruitmentInfo, Tag, Text } from '@/src/app/(design-system)';
import { SACColors } from '@/src/app/(design-system)';
import { Button } from '@/src/app/(design-system)/components/Button/Button';
import AnimatedImageHeader from '../../(design-system)/components/AnimatedImageHeader/AnimatedImageHeader';
import { ClubIcon } from '../../(design-system)/components/ClubIcon/ClubIcon';
import useClub from '@/src/hooks/useClub';
import { Stack, useLocalSearchParams } from 'expo-router';
import { useAppSelector } from '@/src/store/store';
import { AboutSection } from '../../(design-system)/components/AboutSection/AboutSection';
import { Description } from '../event/components/description';
import { useRef } from 'react';
import BottomSheet from '@gorhom/bottom-sheet';
import { faExternalLink } from '@fortawesome/free-solid-svg-icons';

const color: SACColors = 'darkRed';

const ClubPage = () => {
    const { id } = useLocalSearchParams<{ id: string }>();
    const { width } = Dimensions.get('window');
    const IMG_HEIGHT = width;

    const scrollRef = useAnimatedRef<Animated.ScrollView>();
    const scrollOffset = useScrollViewOffset(scrollRef);

    const bottomSheet = useRef<BottomSheet>(null);

    const club = useAppSelector((state) => state.club);
    const { setRetriggerFetch, apiLoading, apiError } = useClub(id as string);

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
        <Box height='100%' backgroundColor='white'>
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
                    ),
                }}
            />
            <Animated.ScrollView
                showsVerticalScrollIndicator={false}
                showsHorizontalScrollIndicator={false}
                scrollEventThrottle={16}
                ref={scrollRef}
            >                
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
                        paddingVertical='s'
                        paddingHorizontal='l'
                    >
                        <ClubIcon imageUrl={club.logo as string} />
                        <Box width={116}>
                            <Button color={color} size='sm' variant="standardButton">
                                Follow
                            </Button>
                        </Box>
                    </Box>
                    <Box marginHorizontal='l' marginVertical='m' gap='m'>
                        <Text variant="header-1">{club.name}</Text>
                        <PageTags tags={club.tags} color={color} />
                        <AboutSection 
                            description={club.description}
                            onPress={() => bottomSheet.current?.snapToIndex(0)}
                            type='club'
                        />
                        {club.recruitment?.is_recruiting &&
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
                                <Box paddingBottom="xxxs" paddingTop="xxs" gap='xxs'>
                                    <Text color="white" variant="body-1">
                                        Application Form
                                    </Text>
                                    <Text color="white" variant="caption-1">
                                        Google Form
                                    </Text>
                                </Box>
                            </Button>
                        }
                        <Box>
                            <Text variant="subheader-1">Recruiting</Text>
                            <RecruitmentInfo
                                recruitingType={'application'}
                                isRecruiting={false}
                                recruitmentCycle={'fallSpring'}
                                color={color}
                            />
                        </Box>

                        <Text variant="subheader-1">Upcoming Events</Text>
                        {/* <EventCard
                            event={events[0].name}
                            variant="big"
                            tags={tags}
                            club={club.name}
                            startTime={events[0].start_time}
                            endTime={events[0].end_time}
                            logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                            image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                        /> */}
                        <Box height={1} backgroundColor="gray" marginVertical="m" />
                        {/* <EventCardList events={events} /> */}

                        <Text variant="subheader-1">Leadership</Text>
                        {/* <PointOfContactList
                            contacts={pointOfContacts}
                            color={color}
                        /> */}
                    </Box>
                </Box>
            </Animated.ScrollView>
            <Description
                description={club.description as string}
                ref={bottomSheet}
                type='club'
            />
        </Box>
    );
};

export default ClubPage;
