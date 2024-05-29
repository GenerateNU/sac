import React from 'react';
import { Text, View } from 'react-native';

import { Button } from '../(design-system)/Button/Button';

const Welcome = () => {
    return (
        <View>
            <Text>Welcome</Text>
            <Button variant="standardButton" color="blue">
                Hello
            </Button>
        </View>
    );
};

export default Welcome;
