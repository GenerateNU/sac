import React, { forwardRef, useCallback } from 'react';
import { Text, View } from 'react-native';

import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet';

import { Button } from '@/components/button';

type Ref = BottomSheet;

const RSVPBottomSheet = forwardRef<Ref>((_, ref) => {
    const snapPoints = ['35%'];

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
            <View className="justify-center p-5">
                <Text className="text-xl font-bold">Register for Event</Text>
                <Text className="mt-2 text-base text-gray-500">
                    By registering for this event, you are automatically signed up for
                    notifications and this event will be added to your calendar.
                </Text>
                <View className="w-full mt-10">
                    <Button
                        className="rounded-xl"
                        onPress={handleClosePress}
                        size={'screenwide'}
                    >
                        Register
                    </Button>
                </View>
            </View>
        </BottomSheet>
    );
});

export { RSVPBottomSheet };
