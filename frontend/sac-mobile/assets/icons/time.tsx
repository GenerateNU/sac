import React from 'react';
import {View} from 'react-native';
import Svg, {Path, Circle, Line} from 'react-native-svg';

interface TimeIconProp {
    focused?: boolean
}

const Time = ({focused} : TimeIconProp) => {
    return (
        <View className='w-[5%] h-[5%'>
            <Svg width={14} height={14} viewBox="0 0 10 11" fill="none" xmlns="http://www.w3.org/2000/svg">
                <Path fill-rule="evenodd" clip-rule="evenodd" d="M4.99999 0.862793C4.79236 0.862793 4.62405 1.0311 4.62405 1.23874V2.87209C4.62405 3.07971 4.79236 3.24802 4.99999 3.24802C5.20762 3.24802 5.37593 3.07971 5.37593 2.87209V1.63108C7.54602 1.82137 9.24812 3.6433 9.24812 5.86279C9.24812 8.20896 7.34617 10.1109 4.99999 10.1109C2.65383 10.1109 0.751881 8.20896 0.751881 5.86279C0.751881 4.81477 1.1309 3.85625 1.75978 3.11536C1.89414 2.95708 1.87474 2.71984 1.71645 2.58548C1.55817 2.45112 1.32093 2.47051 1.18656 2.6288C0.446671 3.50046 0 4.62994 0 5.86279C0 8.62423 2.23857 10.8628 4.99999 10.8628C7.76143 10.8628 10 8.62423 10 5.86279C10 3.10136 7.76143 0.862793 4.99999 0.862793ZM4.42895 6.29943L2.53658 3.66158C2.48295 3.58683 2.49134 3.48425 2.55639 3.4192C2.62143 3.35414 2.72402 3.34576 2.79876 3.39939L5.43662 5.29176C5.79871 5.55152 5.84141 6.07401 5.52631 6.38912C5.2112 6.70422 4.68871 6.66151 4.42895 6.29943Z" fill="black"/>
            </Svg>
        </View>
    )
}

export default Time;