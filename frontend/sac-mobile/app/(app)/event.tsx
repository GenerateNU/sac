import React from 'react';
import {
    Linking,
    ScrollView,
    Text,
    TouchableOpacity,
    View
} from 'react-native';

import { Button } from '@/components/button';
import HostedBy from '@/components/hosted-by';
import { AllHosts } from '@/app/(app)/_components/all-hosts';
import AllTags from '@/components/all-tags';
import { useEvent } from '@/hooks/use-event';
import { Title } from './_components/title';
import { EventTime } from './_components/event-time';
import { EventLocation } from './_components/event-location';

// TODO: handle link OR location
const Event = () => {
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
            <View className='flex-col justify-between flex-1 w-full'>
                <View className='w-full bg-gray-600 h-[100]' />
                {/* <EventHeader /> */}
                <Title title={event.name} />
                <AllHosts eventID={event.id} />
                <View className='flex gap-y-2'>
                    <EventTime startTime={event.start_time} endTime={event.end_time} />
                    <EventLocation location={event.location} meetingLink={event.meeting_link} />
                    <AllTags eventID={event.id} />
                </View>
                {/* <HostedBy eventID={event.id} /> */}
            </View>
        </ScrollView>
    );
};

export default Event;

{/* <ScrollView showsVerticalScrollIndicator={false}>
                <View className="flex-col items-center justify-between flex-1 w-full">
                    <View className="w-full bg-gray-600 h-[100]" />
                    <View className="flex w-full px-6">
                        <EventHeader />
                        <Title title={event.name} />
                        <AllHosts eventID={event.id} />
                        <View className='flex gap-y-2'>
                            <EventTime startTime={event.start_time} endTime={event.end_time} />
                            <EventLocation location={event.location} meetingLink={event.meeting_link} />
                            <AllTags eventID={event.id} />
                        </View>
                        <View className="w-full py-2">
                            <Text className="text-2xl font-bold">
                                About Event
                            </Text>
                            <Text className="mt-2 text-lg">
                                {event.content}
                            </Text>
                        </View>
                        <View className="w-full px-6 py-2">
                            {/* <View className="flex-row items-center justify-between">
                                <Text className="text-2xl font-bold">
                                    Location
                                </Text>
                                <Button
                                    // TODO: Navigate to the native map app
                                    onPress={() =>
                                        console.log('Navigate to the map')
                                    }
                                    size={'default'}
                                    variant={'default'}
                                >
                                    <Text>Open in Maps</Text>
                                </Button>
                            </View> */}
{/* <Text className="py-2 text-lg">
                                {event.location}
                            </Text> */}
{/* {event.meeting_link && (
                                <TouchableOpacity
                                    onPress={() =>
                                        Linking.openURL(
                                            event.meeting_link!
                                        ).catch((err) =>
                                            console.error(
                                                'An error occurred',
                                                err
                                            )
                                        )
                                    }
                                >
                                    <Text className="text-blue-500 visited:text-purple-500">
                                        {event.meeting_link}
                                    </Text>
                                </TouchableOpacity>
                            )} */}
{/* <View className="w-full my-3 bg-gray-400 h-36 rounded-xl" /> */ }
//             </View>
//             {/* <View className="w-full px-6 py-2">
//                 <HostedBy eventID={event.id} />
//             </View> */}
//         </View >
//     </View >
// </ScrollView > * /}

const EventHeader = () => {
    // <View className="flex-row items-center justify-between w-full">
    //     <View className="w-24 h-24 bg-gray-300 rounded-xl" />
    //     <View className="flex-row space-x-3">
    //         <Button
    //             onPress={() => console.log('X Icon')}
    //             size={'icon'}
    //             variant={'secondary'}
    //         >
    //             <Text>X</Text>
    //         </Button>
    //         <Button
    //             onPress={() => console.log('Y Icon')}
    //             size={'icon'}
    //             variant={'secondary'}
    //         >
    //             <Text>Y</Text>
    //         </Button>
    //         <Button
    //             onPress={() => console.log('RSVP')}
    //             size={'default'}
    //             variant={'default'}
    //             className="w-18"
    //         >
    //             <Text>RSVP</Text>
    //         </Button>
    //     </View>
    // </View>
}

