import { TouchableOpacity } from 'react-native';

import { faEllipsis } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Box } from '../Box/Box';

interface KebabProps {
    onPress?: () => void;
    color?: 'black' | 'white';
}

export const KebabMenu = ({ onPress, color = 'white' }: KebabProps) => {
    return (
        <Box>
            <TouchableOpacity onPress={onPress}>
                <FontAwesomeIcon icon={faEllipsis} color={color} size={25} />
            </TouchableOpacity>
        </Box>
    );
};
