import React from 'react';
import { Text, View } from 'react-native';
import { TextBox } from '../(design-system)/Textbox/Textbox';

const Welcome = () => {
    return (
        <View style={{ paddingTop: 100, paddingLeft: 20, paddingRight: 20 }}>
            {/* <Text>Welcome</Text> */}
            <TextBox
                placeholder='000000000'
                helperText='Please input your NUID'
                title='NUID'
            />
        </View>
    );
};

export default Welcome;
