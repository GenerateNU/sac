import React from 'react';
import { useState } from 'react';
import { Text, View } from 'react-native';
import {
    CodeField,
    Cursor,
    useBlurOnFulfill,
    useClearByFocusCell
} from 'react-native-confirmation-code-field';

import { Button } from '@/components/button';

const CELL_COUNT = 6;

type CodeProps = {
    onCompleted: (code: string) => void;
};

const Code = ({ onCompleted }: CodeProps) => {
    const [value, setValue] = useState('');
    const ref = useBlurOnFulfill({ value, cellCount: CELL_COUNT });
    const [props, getCellOnLayoutHandler] = useClearByFocusCell({
        value,
        setValue
    });

    return (
        <View className="flex-col gap-y-5">
            <CodeField
                ref={ref}
                {...props}
                value={value}
                onChangeText={setValue}
                cellCount={CELL_COUNT}
                rootStyle={{ marginTop: 20 }}
                keyboardType="number-pad"
                textContentType="oneTimeCode"
                renderCell={({ index, symbol, isFocused }) => (
                    <View
                        onLayout={getCellOnLayoutHandler(index)}
                        key={index}
                        style={{
                            width: 40,
                            height: 40,
                            justifyContent: 'center',
                            alignItems: 'center',
                            borderBottomWidth: 2,
                            borderBottomColor: isFocused ? 'gray' : 'black'
                        }}
                    >
                        <Text style={{ fontSize: 24, color: 'gray' }}>
                            {symbol || (isFocused ? <Cursor /> : null)}
                        </Text>
                    </View>
                )}
            />
            <View className="pt-[2%]">
                <Button
                    className="bg-primary-500"
                    onPress={() => onCompleted(value)}
                >
                Submit
                </Button>
            </View>
        </View>
    );
};

export default Code;
