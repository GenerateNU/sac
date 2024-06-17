import { TouchableOpacity } from 'react-native-gesture-handler';

import { Box } from '../Box/Box';
import { Text } from '../Text/Text';

type AboutSectionProps = {
    type: 'club' | 'event';
    onPress: () => void;
    description: string;
};

export const AboutSection = ({
    onPress,
    description,
    type
}: AboutSectionProps) => {
    return (
        <Box>
            <Text marginBottom="m" variant="subheader-1">
                About {type === 'club' ? 'Us' : 'Event'}
            </Text>
            <TouchableOpacity onPress={onPress}>
                <Box>
                    <Text lineHeight={22} numberOfLines={5}>
                        {description}
                    </Text>
                    <Text paddingTop="xs" variant="caption-1">
                        Read More...
                    </Text>
                </Box>
            </TouchableOpacity>
        </Box>
    );
};
