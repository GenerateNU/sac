import { View, Text, TouchableOpacity, Linking } from 'react-native'
import React from 'react'
import { MaterialCommunityIcons } from '@expo/vector-icons';

type EventLocationProps = {
    location: string;
    meetingLink?: string;
}

const EventLocation = ({ location, meetingLink }: EventLocationProps) => {
    return (
        <View className="flex-row items-center py-2">
            <MaterialCommunityIcons name="map-marker" size={36} color="black" />
            <View className="flex-col px-2 py-1">
                <Text className="text-base font-medium ">{location}</Text>
                {meetingLink && (
                    <TouchableOpacity
                        onPress={() =>
                            Linking.openURL(meetingLink).catch((err) =>
                                console.error('An error occurred', err)
                            )
                        }
                    >
                        <Text className="text-blue-500 visited:text-purple-500">
                            {meetingLink}
                        </Text>
                    </TouchableOpacity>
                )}
            </View>
        </View>
    )
}

export { EventLocation };