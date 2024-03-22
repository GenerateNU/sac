import React from 'react';
import {View} from 'react-native';
import Svg, {Path, Circle, Line} from 'react-native-svg';

interface SearchIconProp {
    focused: boolean
}

const SearchAltLight = ({focused} : SearchIconProp) => {
    return (
        <View>
            <Svg width="40" height="40" viewBox="0 0 40 40">
      <Circle cx="18.334" cy="18.3334" r="10" stroke="#333333" strokeWidth={2} fill="none" />
      <Path
        d="M18.334 13.3334C17.6774 13.3334 17.0272 13.4627 16.4206 13.714C15.8139 13.9652 15.2627 14.3335 14.7985 14.7978C14.3342 15.2621 13.9659 15.8133 13.7146 16.42C13.4633 17.0266 13.334 17.6768 13.334 18.3334"
        stroke="#333333"
        strokeWidth={2}
        strokeLinecap="round"
        fill="none"
      />
      <Line x1="33.334" y1="33.3334" x2="28.334" y2="28.3334" stroke="#333333" strokeWidth={2} strokeLinecap="round" />
    </Svg>
        </View>
    )
}

export default SearchAltLight;