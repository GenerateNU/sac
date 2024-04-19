import React from 'react';
import { Text, View } from 'react-native';
import MapView, { Marker } from 'react-native-maps';
import { createOpenLink } from 'react-native-open-maps';

type LocationViewProps = {
    location: string;
};

const LocationView = ({ location }: LocationViewProps) => {
    const [city, rest] = location.split(/,(.+)/);
    const coordinates = { latitude: 42.3393326, longitude: -71.0869942 };
    const openMap = createOpenLink({ ...coordinates });

    return (
        <View className="flex flex-col gap-2">
            <View className="inline-flex flex-col justify-between gap-y-2">
                <Text className="text-xl font-light leading-relaxed text-zinc-900">
                    How to get there
                </Text>
                <Text className="text-sm font-medium leading-relaxed text-zinc-900">
                    {city},<Text className="font-normal">{rest}</Text>
                </Text>
            </View>
            <View className="rounded-xl overflow-hidden border-[1px] border-gray-200">
                <MapView
                    onPress={() => openMap()}
                    style={{ width: '100%', height: 200 }}
                    cameraZoomRange={{
                        minCenterCoordinateDistance: 150,
                        maxCenterCoordinateDistance: 200
                    }}
                    initialRegion={{
                        ...coordinates,
                        latitudeDelta: 0.0922,
                        longitudeDelta: 0.0421
                    }}
                >
                    <Marker
                        coordinate={{
                            ...coordinates
                        }}
                    />
                </MapView>
            </View>
        </View>
    );
};

export { LocationView };
