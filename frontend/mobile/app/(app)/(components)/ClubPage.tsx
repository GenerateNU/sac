import { Dimensions } from 'react-native';
import { Linking } from 'react-native';
import Animated, {
    interpolate,
    useAnimatedRef,
    useAnimatedStyle,
    useScrollViewOffset
} from 'react-native-reanimated';

import { faExternalLink } from '@fortawesome/free-solid-svg-icons';
import {
    Club,
    Contact,
    Event,
    PointOfContact as POCProps,
    Tag as TagProps
} from '@generatesac/lib';

import {
    Box,
    Button,
    RecruitmentInfo,
    SACColors,
    Tag,
    Text
} from '@/app/(design-system)';
import { ClubIcon } from '@/app/(design-system)/components/ClubIcon/ClubIcon';
import { EventCard } from '@/app/(design-system)/components/EventCard';
import { PointOfContactList } from '@/app/(design-system)/components/PointOfContactCard/PointofContactsList';

import { AnimatedImageBox } from './ClubPageHeader';
import { EventCardList } from './EventCardList';

interface ClubPageProps {
    club: Club;
    tags: TagProps[];
    events: Event[];
    contacts?: Contact[];
    color: SACColors;
    pointOfContacts: POCProps[];
}

export const ClubPage: React.FC<ClubPageProps> = ({
    club,
    tags,
    events,
    color,
    pointOfContacts
}) => {
    const { width } = Dimensions.get('window');
    const IMG_HEIGHT = width;

    const scrollRef = useAnimatedRef<Animated.ScrollView>();
    const scrollOffset = useScrollViewOffset(scrollRef);

    const imageAnimatedStyle = useAnimatedStyle(() => {
        return {
            transform: [
                {
                    translateY: interpolate(
                        scrollOffset.value,
                        [-IMG_HEIGHT, 0, IMG_HEIGHT],
                        [-IMG_HEIGHT / 2, 0, IMG_HEIGHT * 0.75]
                    )
                },
                {
                    scale: interpolate(
                        scrollOffset.value,
                        [-IMG_HEIGHT, 0, IMG_HEIGHT],
                        [2, 1, 1]
                    )
                }
            ]
        };
    });

    return (
        <Box>
            <Animated.ScrollView
                showsVerticalScrollIndicator={false}
                showsHorizontalScrollIndicator={false}
                scrollEventThrottle={16}
                ref={scrollRef}
            >
                <AnimatedImageBox
                    uri="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSLF3ord7lnV_5Je-pC2AUgUiesHNPcZlpI7A&s"
                    animatedStyle={imageAnimatedStyle}
                />
                <Box rowGap="m" backgroundColor="white" padding="s">
                    <Box
                        flexDirection="row"
                        justifyContent="space-between"
                        marginTop="negativeXl"
                        alignItems="center"
                    >
                        <ClubIcon imageUrl="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSLF3ord7lnV_5Je-pC2AUgUiesHNPcZlpI7A&s" />
                        <Box width={116}>
                            <Button color={color} variant="standardButton">
                                Follow
                            </Button>
                        </Box>
                    </Box>
                    <Text variant="header-1">{club.name}</Text>
                    <Box
                        flexDirection="row"
                        flexWrap="wrap"
                        rowGap="xs"
                        columnGap="s"
                    >
                        {tags.map((tag) => (
                            <Tag color={color} key={tag.id}>
                                {tag.name}
                            </Tag>
                        ))}
                    </Box>
                    <Box gap="xxs" marginBottom="s">
                        <Text variant="subheader-1">About Us</Text>
                        <Text variant="body-1">{club.description}</Text>
                    </Box>

                    <Button
                        variant="iconButton"
                        icon={faExternalLink}
                        color={color}
                        iconPosition="right"
                        size="sm"
                        justifyContent="space-between"
                        flexDirection="row"
                        onPress={() => Linking.openURL(club.application_link)}
                    >
                        <Box paddingBottom="xxxs" paddingTop="xxs">
                            <Text color="white" variant="body-1">
                                Application Form
                            </Text>
                            <Text color="white" variant="caption-1">
                                Google Form
                            </Text>
                        </Box>
                    </Button>

                    <Box>
                        <Text variant="subheader-1">Recruiting</Text>
                        <RecruitmentInfo
                            recruitmentText="Fill out our application"
                            recruitingType={club.recruitment_type}
                            isRecruiting={club.is_recruiting}
                            recruitmentCycle={club.recruitment_cycle}
                            color={color}
                        />
                    </Box>

                    <Text variant="subheader-1">Upcoming Events</Text>
                    <EventCard
                        event={events[0].name}
                        variant="big"
                        tags={tags}
                        club={club.name}
                        startTime={events[0].start_time}
                        endTime={events[0].end_time}
                        logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                        image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                    />
                    <Box height={1} backgroundColor="gray" marginVertical="m" />
                    <EventCardList events={events} />

                    <Text variant="subheader-1">Leadership</Text>
                    <PointOfContactList
                        contacts={pointOfContacts}
                        color={color}
                    />
                </Box>
            </Animated.ScrollView>
        </Box>
    );
};
