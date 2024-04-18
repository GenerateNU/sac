import React, { forwardRef } from 'react';
import { Animated, View } from 'react-native';

import BottomSheet from '@gorhom/bottom-sheet/lib/typescript/components/bottomSheet/BottomSheet';

import EventDemoDay from '@/assets/images/svg/event-demo-day.svg';
import { Button } from '@/components/button';

const HEADER_HEIGHT = 175;
const LOGO_HEIGHT = 100;
const HALF_LOGO_HEIGHT = LOGO_HEIGHT / 2;
const HEADER_TRANSLATE = HEADER_HEIGHT - HALF_LOGO_HEIGHT;

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
            }
        };

        return (
            <View
                style={{
                    paddingBottom: type === 'club' ? HALF_LOGO_HEIGHT : 20
                }}
            >
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
                    className={`absolute px-[5%]`}
                    style={{
                        transform: `translateY(${HEADER_TRANSLATE}px)`
                    }}
                >
                    {type === 'club' && (
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
                                <Button
                                    size={'default'}
                                    variant={'default'}
                                    className="w-18"
                                    onPress={buttonOnPress[type].follow}
                                >
                                    Follow
                                </Button>
                            </View>
                        </View>
                    )}
                </View>
            </View>
        );
    }
);

export { HeaderImage };
