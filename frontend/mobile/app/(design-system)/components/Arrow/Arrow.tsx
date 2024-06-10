import { TouchableOpacity } from 'react-native-gesture-handler';

import { router } from 'expo-router';

import { faAngleRight } from '@fortawesome/free-solid-svg-icons';
import { faAngleLeft } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Box } from '../Box/Box';

interface ArrowProps {
    direction?: 'left' | 'right';
    onPress?: () => void;
    color?: 'black' | 'white';
}

export const Arrow = ({
    direction = 'left',
    onPress,
    color = 'white'
}: ArrowProps) => {
    return (
        <Box>
            <TouchableOpacity onPress={onPress || (() => router.back())}>
                <FontAwesomeIcon
                    icon={direction === 'left' ? faAngleLeft : faAngleRight}
                    color={color}
                    size={25}
                />
            </TouchableOpacity>
        </Box>
    );
};
