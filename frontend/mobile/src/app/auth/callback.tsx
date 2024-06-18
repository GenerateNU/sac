import { useEffect } from 'react';
import { Image } from 'react-native';

import { router, useLocalSearchParams } from 'expo-router';

import { OAuthCallbackRequestQueryParams, authApi } from '@generatesac/lib';

import Loading from '@/src/assets/gif/loading.gif';
import { setLoggedInState, setUser } from '@/src/store/slices/userSlice';
import { useAppDispatch, useAppSelector } from '@/src/store/store';

import { Box, Text } from '../design-system';

const OAuthCallback = () => {
    const params = useLocalSearchParams();
    const [oAuthCallback] = authApi.useLazyCallbackQuery();
    const dispatch = useAppDispatch();
    const { accessToken } = useAppSelector((state) => state.user);

    useEffect(() => {
        const { code, session_state, state } =
            params as OAuthCallbackRequestQueryParams;
        if (code || session_state || state) {
            oAuthCallback({ code, session_state, state }).then(
                ({ data, error }) => {
                    if (data) {
                        dispatch(setUser(data));
                        dispatch(setLoggedInState(true));
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
