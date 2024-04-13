import React, { forwardRef, useCallback, useRef } from 'react'
import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet'
import { Text, View } from 'react-native'
import { Button } from '@/components/button'

type Ref = BottomSheet

const RSVPBottomSheet = forwardRef<Ref>((_, ref) => {
    const snapPoints = ['35%']

    if (!ref) {
        return null
    }

    // @ts-ignore
    const handleClosePress = () => ref.current?.close();
    const renderBackdrop = useCallback(
        (props: any) => <BottomSheetBackdrop appearsOnIndex={0} disappearsOnIndex={-1} {...props} />,
        []
    );

    return (
        <BottomSheet
            ref={ref}
            index={-1}
            snapPoints={snapPoints}
            enablePanDownToClose={true}
            backgroundStyle={{ backgroundColor: 'white' }}
            backdropComponent={renderBackdrop}
        >
            <View className="items-center justify-center p-4">
                <Text className="text-2xl font-bold">Save Event</Text>
                <Text className="mt-2 text-lg">
                    By saving this event, you are automatically signed up for notifications and this event will be added to your calendar.
                </Text>
                <View className="w-full mt-4">
                    <Button className='rounded-xl' onPress={handleClosePress} size={'screenwide'}>Save Event</Button>
                </View>
            </View>
        </BottomSheet>
    )
})

export default RSVPBottomSheet