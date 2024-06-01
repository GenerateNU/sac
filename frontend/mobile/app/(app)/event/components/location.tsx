import MapView, { Marker } from 'react-native-maps';
import { createOpenLink } from 'react-native-open-maps';

import { Text } from '@/app/(design-system)';
import { Box } from '@/app/(design-system)';

interface LocationProps {
    location: string;
}

export const Location: React.FC<LocationProps> = ({ location }) => {
    const coordinates = { latitude: 42.3393326, longitude: -71.0869942 };
    const openMap = createOpenLink({ ...coordinates });

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
                <Marker
                    coordinate={{
                        ...coordinates
                    }}
                />
            </MapView>
        </Box>
    );
};
