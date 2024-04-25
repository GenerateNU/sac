import React, { forwardRef, useEffect, useState } from 'react';
import { Animated, Image, View } from 'react-native';

import { Button } from '@/components/button';
import {
    useCreateClubFollowing,
    useDeleteClubFollowing,
    useUserFollowing
} from '@/hooks/use-user';
import { useCreateUserClubFollowingMutation } from '@sac/lib';

const HEADER_HEIGHT = 175;
const LOGO_HEIGHT = 100;
const HALF_LOGO_HEIGHT = LOGO_HEIGHT / 2;
const HEADER_TRANSLATE = HEADER_HEIGHT - HALF_LOGO_HEIGHT;

type HeaderImageProps = {
    image: string;
    logo: string;
    pan: Animated.ValueXY;
    type: 'club' | 'event';
    userID: string;
    clubID?: string;
};

type Ref = BottomSheet;

const HeaderImage = forwardRef<Ref, HeaderImageProps>(
    ({ image, logo, pan, type, userID, clubID }, ref) => {
        // const followClub = useCreateClubFollowing();
        const unfollowClub = useDeleteClubFollowing();
        const [createClubUserFollowing, { data, isLoading, error }] = useCreateUserClubFollowingMutation();
        const userFollowing = useUserFollowing(userID);
        const [buttonClicked, setButtonClicked] = useState(false);

        // MARK to resolve unused variables on L29
        console.log(data, isLoading, error);

        // Effects
        useEffect(() => {
            if (userFollowing.data) {
                const isFollowing = userFollowing.data
                    .map((club) => club.id)
                    .includes(clubID || '');
                setButtonClicked(isFollowing);
            }
        }, [userFollowing.data, clubID]);

        // Event Handlers
        const handlePressFollow = async () => {
            if (buttonClicked) {
                await unfollowClub.mutateAsync({
                    userID: userID,
                    clubID: clubID || ''
                });
            } else {
                await createClubUserFollowing({
                    userID: userID,
                    clubID: clubID || ''
                });
            }
            // Force a refetch of the user following data
            userFollowing.refetch();
            setButtonClicked(!buttonClicked);
        };

        const buttonOnPress = {
            club: {
                follow: handlePressFollow
            }
        };

        if (!ref) {
            return null;
        }

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
                                <Image
                                    source={{ uri: logo }}
                                    className="w-full h-full rounded-xl"
                                />
                            </View>
                            <View className="flex-row space-x-3">
                                <Button
                                    size={'default'}
                                    variant={'default'}
                                    onPress={buttonOnPress[type].follow}
                                >
                                    {buttonClicked ? 'Following' : 'Follow'}
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
