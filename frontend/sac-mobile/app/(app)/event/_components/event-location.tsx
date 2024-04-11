import React from 'react';
import { Linking, Text, TouchableOpacity, View } from 'react-native';

import Location from '@/assets/images/svg/location.svg';

type EventLocationProps = {
    location: string;
    meetingLink?: string;
};

const EventLocation = ({ location, meetingLink }: EventLocationProps) => {
    const [city, rest] = location.split(/,(.+)/);

    return (
        <View className="flex-row items-center py-1">
            <Location />
            <View className="flex-col pb-1">
                <View className="flex-row pl-2">
                    <Text className="text-black">{city}</Text>
                    <Text className="text-black">, {rest}</Text>
                </View>
                {meetingLink && (
                    <View className="flex-row items-center">
                        <TouchableOpacity
                            className="pl-2"
                            onPress={() =>
                                Linking.openURL(meetingLink!).catch((err) =>
                                    console.error('An error occurred', err)
                                )
                            }
                        >
                            <Text className="text-black underline">
                                {meetingLink}
                            </Text>
                        </TouchableOpacity>
                    </View>
                )}
            </View>
        </View>
    );
};

export { EventLocation };