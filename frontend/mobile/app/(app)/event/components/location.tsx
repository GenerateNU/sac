import MapView, { Marker } from 'react-native-maps';
import { createOpenLink } from 'react-native-open-maps';

import { Text } from '@/app/(design-system)';
import { Box } from '@/app/(design-system)';

// import Geocoder from 'react-native-geocoding';
// import { useEffect, useState } from 'react';

interface LocationProps {
    location: string;
}

export const Location: React.FC<LocationProps> = ({ location }) => {
    const coordinates = { latitude: 42.3393326, longitude: -71.0869942 };
    const openMap = createOpenLink({ ...coordinates });
    // Geocoder.init(process.env.GOOGLE_MAPS_API_KEY | '');
    // const [coordinates, setCoordinates] = useState<{ latitude: number; longitude: number }>({
    //     latitude: 0,
    //     longitude: 0
    // });
    // const openMap = createOpenLink({ ...coordinates });
    // useEffect(() => {
    //     async function fetchCoordinates() {
    //         try {
    //             const response = await Geocoder.from(location);
    //             const { lat, lng } = response.results[0].geometry.location;
    //             setCoordinates({ latitude: lat, longitude: lng });
    //         } catch (error) {
    //             console.error('Error fetching coordinates:', error);
    //         }
    //     }
    //     fetchCoordinates();
    // }, [location]);

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
