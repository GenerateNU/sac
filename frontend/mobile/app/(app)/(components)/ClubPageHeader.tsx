import React from 'react';
import Animated from 'react-native-reanimated';

import { Box } from '@/app/(design-system)';

interface AnimatedImageBoxProps {
    uri: string;
    animatedStyle: any;
}

export const AnimatedImageBox: React.FC<AnimatedImageBoxProps> = ({
    uri,
    animatedStyle
}) => {
    return (
        <Box>
            <Animated.Image
                source={{ uri }}
                style={[animatedStyle, { width: '100%', aspectRatio: 1 }]}
            />
        </Box>
    );
};
