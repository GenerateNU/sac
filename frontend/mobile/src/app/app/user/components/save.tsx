import { TouchableOpacity } from 'react-native-gesture-handler';

import { Text } from '@/src/app/design-system';

interface SaveProps {
    onPress: () => void;
    isValid: boolean;
}

export const Save: React.FC<SaveProps> = ({ onPress, isValid }) => {
    return (
        <TouchableOpacity disabled={!isValid} onPress={onPress}>
            <Text color={isValid ? 'black' : 'darkGray'}>Save</Text>
        </TouchableOpacity>
    );
};
