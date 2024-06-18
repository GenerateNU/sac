import { IconDefinition } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Box, Text } from '../../design-system';

export interface ScreenProps {
    title: string;
    description: string;
    icon: IconDefinition;
}

export const Screen: React.FC<ScreenProps> = ({ title, description, icon }) => {
    return (
        <Box alignItems="center" justifyContent="center" gap="m">
            <FontAwesomeIcon size={200} icon={icon} />
            <Box
                marginTop="xl"
                marginBottom="xxl"
                alignItems="center"
                justifyContent="center"
                gap="l"
            >
                <Text
                    paddingHorizontal="xl"
                    textAlign="center"
                    variant="header-1"
                    fontSize={40}
                >
                    {title}
                </Text>
                <Text
                    paddingHorizontal="xs"
                    textAlign="center"
                    fontWeight="300"
                    variant="subheader-2"
                >
                    {description}
                </Text>
            </Box>
        </Box>
    );
};
