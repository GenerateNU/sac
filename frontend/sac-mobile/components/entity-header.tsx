import React, { forwardRef } from 'react';
import { Animated, Platform, View } from 'react-native';

import BottomSheet from '@gorhom/bottom-sheet/lib/typescript/components/bottomSheet/BottomSheet';
import { MenuView } from '@react-native-menu/menu';

import EventDemoDay from '@/assets/images/svg/event-demo-day.svg';
import { Button } from '@/components/button';

const HEADER_HEIGHT = 175;
const LOGO_HEIGHT = 100;

type HeaderImageProps = {
    image: string;
    pan: Animated.ValueXY;
    type: 'club' | 'event';
};

type Ref = BottomSheet;

const HeaderImage = forwardRef<Ref, HeaderImageProps>(
    ({ image, pan, type }: HeaderImageProps, ref) => {
        if (!ref) {
            return null;
        }

        const buttonOnPress = {
            club: {
                follow: () => {}
            },
            event: {
                // @ts-ignore
                save: () => ref.current?.snapToIndex(0)
            }
        };

        return (
            <View className={`pb-[${LOGO_HEIGHT / 2}px]`}>
                <Animated.Image
                    source={{ uri: image }}
                    resizeMode="cover"
                    style={{
                        width: 'auto',
                        height: HEADER_HEIGHT,
                        transform: [
                            {
                                translateY: pan.y.interpolate({
                                    inputRange: [-1000, 0],
                                    outputRange: [-500, 0],
                                    extrapolate: 'clamp'
                                })
                            },
                            {
                                scale: pan.y.interpolate({
                                    inputRange: [-3000, 0],
                                    outputRange: [20, 1],
                                    extrapolate: 'clamp'
                                })
                            }
                        ]
                    }}
                />
                <View
                    className={`absolute translate-y-[${HEADER_HEIGHT - LOGO_HEIGHT / 2}px] px-[5%]`}
                >
                    <View className="flex-row w-full items-center justify-between">
                        <View
                            className="p-2 bg-[#f2f2f2] rounded-xl"
                            style={{
                                width: LOGO_HEIGHT,
                                height: LOGO_HEIGHT
                            }}
                        >
                            <EventDemoDay
                                className="rounded-xl"
                                width="100%"
                                height="100%"
                            />
                        </View>
                        <View className="flex-row space-x-3">
                            {type === 'event' ? (
                                <>
                                    <MenuView
                                        onPressAction={({ nativeEvent }) => {
                                            console.warn(
                                                JSON.stringify(nativeEvent)
                                            );
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
                                        onPress={buttonOnPress[type].save}
                                    >
                                        Save
                                    </Button>
                                </>
                            ) : (
                                <Button
                                    size={'default'}
                                    variant={'default'}
                                    className="w-18"
                                    onPress={buttonOnPress[type].follow}
                                >
                                    Follow
                                </Button>
                            )}
                        </View>
                    </View>
                </View>
            </View>
        );
    }
);

export { HeaderImage };
