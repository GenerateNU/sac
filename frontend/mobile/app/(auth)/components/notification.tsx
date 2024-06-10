import { router } from 'expo-router';

import { faBell } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Box, Text } from '../../(design-system)';
import { Button } from '../../(design-system)/components/Button/Button';

export const Notification = () => {
    return (
        <Box paddingTop="m" alignItems="center" justifyContent="center">
            <FontAwesomeIcon icon={faBell} size={200} />
            <Box paddingVertical="xl" gap="s" paddingHorizontal="l">
                <Text textAlign="center" fontSize={40} variant="header-1">
                    Turn on notifications
                </Text>
                <Text textAlign="center">
                    Allow SAC to notify you of club announcements and your
                    upcoming events!
                </Text>
            </Box>
            <Box gap="m">
                <Button
                    onPress={() => router.push('/(app)/')}
                    variant="standardButton"
                >
                    Allow
                </Button>
                <Button
                    backgroundColor="white"
                    borderColor="black"
                    borderWidth={1}
                    variant="standardButton"
                    onPress={() => router.push('/(app)/')}
                >
                    <Text color="black">Another Time</Text>
                </Button>
            </Box>
        </Box>
    );
};
