import React from 'react';
import {View} from 'react-native';
import Svg, {Path, Circle} from 'react-native-svg';

interface UserIconProp {
    focused: boolean
}

const UserLight = ({focused} : UserIconProp) => {
    return (
        <View className="w-[2%] h-[2%]">
            <Svg width="40" height="40" viewBox="0 0 40 40" fill="none" xmlns="http://www.w3.org/2000/svg">
                <Path d="M32.8784 34.0784C32.1187 31.9521 30.4448 30.0732 28.1162 28.7331C25.7876 27.393 22.9345 26.6666 19.9994 26.6666C17.0642 26.6666 14.2111 27.393 11.8825 28.7331C9.55394 30.0732 7.88001 31.9521 7.12034 34.0784" stroke="#333333" strokeWidth={2} stroke-linecap="round"/>
                <Circle cx="20.0007" cy="13.3333" r="6.66667" stroke="#333333" strokeWidth={2} stroke-linecap="round"/>
            </Svg>
        </View>
    )
}

export default UserLight;