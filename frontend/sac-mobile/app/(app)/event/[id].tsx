import React from 'react';
import { Alert, Platform, ScrollView, Share, Text, View } from 'react-native';

import { MenuView } from '@react-native-menu/menu';

import { HostList } from '@/app/(app)/event/_components/host-list';
import ShareIcon from '@/assets/images/svg/share.svg';
import { TagList } from '@/components/all-tags';
import { Button } from '@/components/button';
import { useEvent } from '@/hooks/use-event';

import { Description } from '../_components/description';
import { Title } from '../_components/title';
import { EventLocation } from './_components/event-location';
import { EventTime } from './_components/event-time';
import { HostNames } from './_components/host-names';
import { LocationView } from './_components/location-view';

// TODO: handle link OR location
const Event = () => {
    // const { id } = useLocalSearchParams();

    const {
        data: event,
        isLoading: isEventLoading,
        error: eventError
    } = useEvent('tester');

    if (eventError) {
        console.error(eventError);
        return <Text>Error fetching event</Text>;
    }

    if (isEventLoading) {
        return <Text>Loading...</Text>;
    }

    if (!event) {
        return <Text>Event not found</Text>;
    }

    return (
        <ScrollView showsVerticalScrollIndicator={false}>
            <View className="flex-col justify-between flex-1 w-full">
                <View className="w-full bg-gray-600 h-[100]" />
                <View className="w-full px-6 pb-5">
                    <EventHeader />
                    <View className="flex gap-y-2">
                        <Title title={event.name} />
                        <HostNames eventID={event.id} />
                        <View className="flex py-1 m-0 gap-y-2">
                            <EventTime
                                startTime={event.start_time}
                                endTime={event.end_time}
                            />
                            <EventLocation
                                location={event.location}
                                meetingLink={event.meeting_link}
                            />
                        </View>
                        <TagList id={event.id} />
                        <Description
                            title={'About Event'}
                            description={event.content}
                        />
                        <LocationView
                            location={event.location}
                            meetingLink={event.meeting_link}
                        />
                        <HostList eventID={event.id} />
                        <Button
                            onPress={() => console.log('RSVP')}
                            size={'default'}
                            variant={'default'}
                            className="w-20 rounded-xl"
                        >
                            <Text>Report</Text>
                        </Button>
                    </View>
                </View>
            </View>
        </ScrollView>
    );
};

export default Event;

const EventHeader = () => {
    const onShare = async () => {
        try {
            const result = await Share.share({
                message:
                    'React Native | A framework for building native apps using React'
            });
            if (result.action === Share.sharedAction) {
                if (result.activityType) {
                    // shared with activity type of result.activityType
                } else {
                    // shared
                }
            } else if (result.action === Share.dismissedAction) {
                // dismissed
            }
        } catch (error: any) {
            Alert.alert(error.message);
        }
    };

    return (
        <View className="flex-row items-center justify-between -translate-y-10">
            <View className="w-24 h-24 bg-gray-300 rounded-xl" />
            <View className="flex-row space-x-3">
                <Button
                    onPress={() => onShare()}
                    size={'icon'}
                    variant={'icon'}
                    className="w-18"
                >
                    <ShareIcon />
                </Button>

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
                        RSVP
                    </Button>
                </MenuView>
            </View>
        </View>
    );
};
