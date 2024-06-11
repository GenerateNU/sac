import { TouchableOpacity } from 'react-native-gesture-handler';

import { Box, Text } from '../../(design-system)';

interface SmallTextProps {
    first: string;
    second: string;
    onPress?: () => void;
    disabled?: boolean;
}

export const SmallText: React.FC<SmallTextProps> = ({
    first,
    second,
    onPress,
    disabled = false
}) => {
    return (
        <Box flexDirection="row" paddingTop="m" gap="xxs">
            <Text color={disabled ? 'gray' : 'black'} variant="body-2">
                {first}
            </Text>
            <TouchableOpacity disabled={disabled} onPress={onPress}>
                <Text
                    fontWeight={500}
                    textDecorationLine="underline"
                    variant="body-2"
                    color={disabled ? 'gray' : 'black'}
                >
                    {second}
                </Text>
            </TouchableOpacity>
        </Box>
    );
};
