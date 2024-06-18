import { Linking, Platform } from 'react-native';
import MapView, { Marker } from 'react-native-maps';

import { Text } from '@/src/app/design-system';
import { Box } from '@/src/app/design-system';

interface LocationProps {
    location: string;
}

const scheme = Platform.select({ ios: 'maps://0,0?q=', android: 'geo:0,0?q=' });

export const Location: React.FC<LocationProps> = ({ location }) => {
    const coordinates = { latitude: 42.3393326, longitude: -71.0869942 };
    const openMap = () => Linking.openURL(`${scheme}${location}`);

    return (
        <Box flexDirection="column" gap="s">
            <Text variant="subheader-1">Location</Text>
            <Text>{location}</Text>
            <MapView
                onPress={() => openMap()}
                style={{ width: '100%', aspectRatio: 1.8, borderRadius: 16 }}
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
                <Marker coordinate={coordinates} />
            </MapView>
        </Box>
    );
};
