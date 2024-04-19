import React, { forwardRef, useState } from 'react';
import { Animated, View } from 'react-native';

import BottomSheet from '@gorhom/bottom-sheet/lib/typescript/components/bottomSheet/BottomSheet';

import EventDemoDay from '@/assets/images/svg/event-demo-day.svg';
import { Button } from '@/components/button';
import { useCreateClubFollowing, useDeleteClubFollowing, useUserFollowing } from '@/hooks/use-user';

const HEADER_HEIGHT = 175;
const LOGO_HEIGHT = 100;
const HALF_LOGO_HEIGHT = LOGO_HEIGHT / 2;
const HEADER_TRANSLATE = HEADER_HEIGHT - HALF_LOGO_HEIGHT;

type HeaderImageProps = {
    image: string;
    pan: Animated.ValueXY;
    type: 'club' | 'event';
    userID: string;
    clubID?: string;
};

type Ref = BottomSheet;

const HeaderImage = forwardRef<Ref, HeaderImageProps>(
    ({ image, pan, type, userID, clubID }, ref) => {

        if (!ref) {
            return null;
        }

        const followClub = useCreateClubFollowing();
        const unfollowClub = useDeleteClubFollowing();
        const userFollowing = useUserFollowing(userID);

        const isFollowing = userFollowing.data?.map((club) => club.id).includes(clubID || '');
        const [buttonClicked, setButtonClicked] = useState(isFollowing ? 'Following' : 'Follow');

        const handlePressFollow = async () => {
            if (buttonClicked === 'Following') {
                await unfollowClub.mutateAsync({
                    userID: userID,
                    clubID: clubID || ''
                });

                // force a refetch of the user following data
                userFollowing.refetch();
                setButtonClicked('Follow');
            } else {
                await followClub.mutateAsync({
                    userID: userID,
                    clubID: clubID || ''
                });

                // force a refetch of the user following data
                userFollowing.refetch();
                setButtonClicked('Following');
            }
        };
        const buttonOnPress = {
            club: {
                follow: handlePressFollow
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
                        <View className="flex-row items-center justify-between w-full">
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
                                    {buttonClicked === 'Following' ? 'Following' : 'Follow'}
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