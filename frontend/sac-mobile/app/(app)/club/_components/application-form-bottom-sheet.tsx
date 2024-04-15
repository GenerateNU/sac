import React, { forwardRef, useCallback, useState } from 'react';
import { Text, TouchableOpacity, View } from 'react-native';

import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet';

import { Button } from '@/components/button';

type Ref = BottomSheet;

const ApplyBottomSheet = forwardRef<Ref>((_, ref) => {
    const snapPoints = ['40%'];

    const renderBackdrop = useCallback(
        (props: any) => (
            <BottomSheetBackdrop
                appearsOnIndex={0}
                disappearsOnIndex={-1}
                {...props}
            />
        ),
        []
    );

    const [selectedOption, setSelectedOption] = useState('');

    const handleOptionChange = (optionValue: string) => {
        setSelectedOption(optionValue);
    };

    const handleSubmit = () => {
        if (!selectedOption) {
            console.log('Please select an option!');
            return;
        }
        if (selectedOption === 'browser') {
            console.log('Opening browser');
            handleClosePress();
        }
        if (selectedOption === 'email') {
            console.log('Sending email');
            handleClosePress();
        }
        if (selectedOption === 'both') {
            console.log('Opening browser and sending email');
            handleClosePress();
        }
    };

    // @ts-ignore
    const handleClosePress = () => ref.current?.close();

    return (
        <BottomSheet
            ref={ref}
            index={-1}
            snapPoints={snapPoints}
            enablePanDownToClose={true}
            backgroundStyle={{ backgroundColor: 'white' }}
            backdropComponent={renderBackdrop}
        >
            <View className="flex flex-col p-4">
                <Text className="text-xl font-bold mb-2 text-center">Application Form</Text>
                <View className="flex flex-col space-y-2">
                    <TouchableOpacity
                        className={`flex-row items-center p-2 rounded-md
                       ${selectedOption === 'browser' ? 'bg-gray-400' : 'bg-gray-100'}`}
                        onPress={() => handleOptionChange('browser')}
                    >
                        <Text className="text-lg">Open in browser</Text>
                    </TouchableOpacity>
                    <TouchableOpacity
                        className={`flex-row items-center p-2 rounded-md
                       ${selectedOption === 'email' ? 'bg-gray-400' : 'bg-gray-100'}`}
                        onPress={() => handleOptionChange('email')}
                    >
                        <Text className="text-lg">Send to email</Text>
                    </TouchableOpacity>
                    <TouchableOpacity
                        className={`flex-row items-center p-2 rounded-md
                       ${selectedOption === 'both' ? 'bg-gray-400' : 'bg-gray-100'}`}
                        onPress={() => handleOptionChange('both')}
                    >
                        <Text className="text-lg">Open in browser and send to email</Text>
                    </TouchableOpacity>
                </View>
                <Button className="p-3 mt-4 rounded-md" onPress={handleSubmit}>
                    <Text className="text-white font-medium text-center">Send</Text>
                </Button>
            </View>
        </BottomSheet>
    );
});

export default ApplyBottomSheet;
