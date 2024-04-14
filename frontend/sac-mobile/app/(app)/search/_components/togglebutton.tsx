import { useEffect, useRef, useState } from "react";
import { GestureResponderEvent, TextInput, Touchable, TouchableOpacity, TouchableWithoutFeedback, View, Text, ViewProps } from "react-native";

import SearchIcon from '@/assets/images/svg/search.svg';
import { Svg } from "react-native-svg";
import { MaterialCommunityIcons } from "@expo/vector-icons";
import Animated, { Easing, ReduceMotion, useAnimatedStyle, useSharedValue, withSpring, withTiming } from "react-native-reanimated";

export interface ToggleButtonProps extends ViewProps {
    left: string,
    right: string,

    onPressed?: (side: "left" | "right") => void
}

// @refresh reset
export const ToggleButton = (p: ToggleButtonProps) => {
    const leftOffset = useSharedValue(0);

    const animatedStyles = useAnimatedStyle(() => ({
        left: withTiming(`${leftOffset.value}%`, {
            duration: 250,
            easing: Easing.out(Easing.quad),
            reduceMotion: ReduceMotion.System,
          })
    }));

    return (
        <View {...p} className="bg-white rounded-full flex flex-row items-center">
            <TouchableOpacity className="w-1/2 z-50 flex flex-row justify-center"
                onPress={(evt) => { 
                    leftOffset.value = 0;

                    if (p.onPressed) {
                        p.onPressed("left"); 
                    }
                }}>
                <Text className="font-bold text-xl">{p.left}</Text>
            </TouchableOpacity>
            <TouchableOpacity className="w-1/2 z-50 flex flex-row justify-center"
                onPress={(evt) => { 
                    leftOffset.value = 50;

                    if (p.onPressed) {
                        p.onPressed("right"); 
                    }
                }}>
                <Text className="font-bold text-xl">{p.right}</Text>
            </TouchableOpacity>
            <Animated.View style={animatedStyles} className="absolute w-1/2 rounded-full bg-[#C8102E] h-full">

            </Animated.View>
        </View>
    );
}