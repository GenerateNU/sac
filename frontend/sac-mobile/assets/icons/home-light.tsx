import React from 'react';
import {View} from 'react-native';
import Svg, {Path} from 'react-native-svg';

const HomeLight = () => {
    return (
        <View className="w-[2%] h-[2%]">
            <Svg width="40" height="40" viewBox="0 0 40 40" fill="none" xmlns="http://www.w3.org/2000/svg">
                <Path d="M8.33398 21.266C8.33398 19.0031 8.33398 17.8716 8.79142 16.8771C9.24886 15.8825 10.1079 15.1462 11.8261 13.6735L13.4927 12.2449C16.5983 9.58302 18.151 8.25208 20.0007 8.25208C21.8503 8.25208 23.403 9.58302 26.5086 12.2449L28.1752 13.6735C29.8934 15.1462 30.7524 15.8825 31.2099 16.8771C31.6673 17.8716 31.6673 19.0031 31.6673 21.266V28.3334C31.6673 31.4761 31.6673 33.0474 30.691 34.0237C29.7147 35 28.1433 35 25.0007 35H15.0007C11.858 35 10.2866 35 9.3103 34.0237C8.33398 33.0474 8.33398 31.4761 8.33398 28.3334V21.266Z" stroke="#333333" stroke-width="2"/>
                <Path d="M24.1673 35V26C24.1673 25.4477 23.7196 25 23.1673 25H16.834C16.2817 25 15.834 25.4477 15.834 26V35" stroke="#333333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </Svg>
        </View>
    )
}

export default HomeLight;