import { Platform, TextInput, View, Text } from 'react-native';

import { MenuView } from '@react-native-menu/menu';

import { Button } from '@/components/button';
import BottomSheet, { BottomSheetView } from '@gorhom/bottom-sheet';
import { useRef, useCallback, forwardRef } from 'react';

type Ref = BottomSheet;

const EventHeader = forwardRef<Ref>((_, ref) => {

    if (!ref) {
        return null;
    }

    const handleOpenPress = () => ref.current?.snapToIndex(0);

    return (
        <View className="flex-row items-center justify-between -translate-y-10">
            <View className="w-24 h-24 bg-gray-300 rounded-xl" />
            <View className="flex-row space-x-3">
                <MenuView
                    onPressAction={({ nativeEvent }) => {
                        console.warn(JSON.stringify(nativeEvent));
                    }}
                    actions={[
                        {
                            id: 'yes',
                            title: 'Yes',
                            image: Platform.select({
                                ios: 'checkmark',
                                android: 'ic_menu_save'
                            })
                        },
                        {
                            id: 'maybe',
                            title: 'Maybe',
                            image: Platform.select({
                                ios: 'questionmark',
                                android: 'ic_menu_help'
                            })
                        },
                        {
                            id: 'no',
                            title: 'No',
                            image: Platform.select({
                                ios: 'xmark',
                                android: 'ic_menu_delete'
                            })
                        }
                    ]}
                    themeVariant="dark"
                    shouldOpenOnLongPress={false}
                    isAnchoredToRight={true}
                >
                    <Button
                        size={'default'}
                        variant={'default'}
                        className="w-18"
                    >
                        RSVP
                    </Button>
                </MenuView>

                <Button
                    size={'default'}
                    variant={'default'}
                    className="w-18"
                    onPress={handleOpenPress}
                >
                    Save
                </Button>
            </View>
        </View>
    );
});

export { EventHeader };
