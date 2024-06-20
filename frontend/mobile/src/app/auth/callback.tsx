import { useEffect } from 'react';
import { Image } from 'react-native';

import { router, useLocalSearchParams } from 'expo-router';

import { Club, OAuthCallbackRequestQueryParams, authApi, userApi } from '@generatesac/lib';

import Loading from '@/src/assets/gif/loading.gif';
import { setLoggedIn } from '@/src/store/slices/globalSlice';
import { setUser, setUserFollowing } from '@/src/store/slices/userSlice';
import { useAppDispatch, useAppSelector } from '@/src/store/store';

import { Box, Text } from '../design-system';

const OAuthCallback = () => {
    const params = useLocalSearchParams();
    const [oAuthCallback] = authApi.useLazyCallbackQuery();
    const [getUserFollowing] = userApi.useLazyGetUserFollowingQuery();
    const dispatch = useAppDispatch();
    const { accessToken } = useAppSelector((state) => state.global);

    const parseUserFollowing = (clubs: Club[]) => {
        return clubs.filter((club) => club.name !== 'SAC').map((club) => club.id);
    }

    useEffect(() => {
        const { code, session_state, state } =
            params as OAuthCallbackRequestQueryParams;
        if (code || session_state || state) {
            oAuthCallback({ code, session_state, state }).then(
                async ({ data, error }) => {
                    if (data) {
                        // Set the user:
                        dispatch(setUser(data));

                        // Retrieve their following:
                        await getUserFollowing(data.id)
                            .then(({ data: followingData }) => {
                                if (followingData) {
                                    const following = parseUserFollowing(followingData);
                                    dispatch(setUserFollowing(following));
                                }
                            });
                                                
                        // Set the user as logged in and redirect to the app:
                        dispatch(setLoggedIn(true));
                        router.push('/app/');
                    }
                    if (error) {
                        router.push('/auth/');
                    }
                }
            );
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [accessToken]);

    return (
        <Box
            backgroundColor="white"
            height="100%"
            width="100%"
            alignItems="center"
            justifyContent="center"
            gap="m"
        >
            <Image
                source={Loading}
                style={{
                    width: 30,
                    height: 30
                }}
            />
            <Text>Logging you in...</Text>
        </Box>
    );
};

export default OAuthCallback;
