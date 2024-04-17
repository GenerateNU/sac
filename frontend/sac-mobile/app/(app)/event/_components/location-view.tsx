import React from 'react';
import { Linking, Text, TouchableOpacity, View } from 'react-native';
import MapView from 'react-native-maps';
import { createOpenLink } from 'react-native-open-maps';

import { Button } from '@/components/button';

type LocationViewProps = {
    location: string;
    meetingLink?: string;
};

const LocationView = ({ location, meetingLink }: LocationViewProps) => {
    const coordinates = { latitude: 42.3393326, longitude: -71.0869942 };
    const openMap = createOpenLink({ ...coordinates });

    return (
        <>
            <View className="inline-flex flex-row items-center justify-between">
                <Text className="text-2xl font-normal leading-relaxed text-zinc-900">
                    Location
                </Text>
                <Button onPress={openMap} size={'default'} variant={'default'}>
                    <Text>Open in Maps</Text>
                </Button>
            </View>

            <Text className="text-base font-light tracking-tight text-zinc-900">
                {location}
            </Text>

            {meetingLink && (
                <TouchableOpacity
                    onPress={() =>
                        Linking.openURL(meetingLink!).catch((err) =>
                            console.error('An error occurred', err)
                        )
                    }
                >
                    <Text className="text-blue-500 visited:text-purple-500">
                        {meetingLink}
                    </Text>
                </TouchableOpacity>
            )}

            <View className="w-full mt-4 overflow-hidden border-2 rounded-xl border-zinc-300">
                <MapView
                    onPress={openMap}
                    style={{ width: '100%', height: 200 }}
                    initialRegion={{
                        ...coordinates,
                        latitudeDelta: 0.0922,
                        longitudeDelta: 0.0421
                    }}
                />
            </View>
        </>
    );
};

export { LocationView };
