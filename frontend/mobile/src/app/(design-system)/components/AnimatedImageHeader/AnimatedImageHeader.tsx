import React from 'react';
import Animated, {
    SharedValue,
    interpolate,
    useAnimatedStyle
} from 'react-native-reanimated';

import { Box } from '../Box/Box';

interface AnimatedImageBoxProps {
    uri: string;
    scrollOffset: SharedValue<number>;
    imageHeight: number;
    aspectRatio: number;
}

const AnimatedImageHeader: React.FC<AnimatedImageBoxProps> = ({
    uri,
    scrollOffset,
    imageHeight,
    aspectRatio
}) => {
    const imageAnimatedStyle = useAnimatedStyle(() => {
        return {
            transform: [
                {
                    translateY: interpolate(
                        scrollOffset.value,
                        [-imageHeight, 0, imageHeight],
                        [
                            -imageHeight / (2 * aspectRatio),
                            0,
                            imageHeight * 0.75
                        ]
                    )
                },
                {
                    scale: interpolate(
                        scrollOffset.value,
                        [-imageHeight, 0, imageHeight],
                        [2 * aspectRatio, 1, 1]
                    )
                }
            ]
        };
    });

    return (
        <Box>
            <Animated.Image
                source={{ uri }}
                style={[
                    imageAnimatedStyle,
                    { width: '100%', aspectRatio: aspectRatio }
                ]}
            />
        </Box>
    );
};

export default AnimatedImageHeader;
