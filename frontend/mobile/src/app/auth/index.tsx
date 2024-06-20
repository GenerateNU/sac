import React from 'react';
import { Image, Linking, SafeAreaView } from 'react-native';

import { authApi } from '@generatesac/lib';

import Loading from '@/src/assets/gif/loading.gif';
import { useAppDispatch } from '@/src/store/store';

import { Box, Button, Text } from '../design-system';
import { GlobalLayout } from '../design-system/components/GlobalLayout/GlobalLayout';
import { setAccessToken } from '@/src/store/slices/globalSlice';

const WelcomePage = () => {
    const [login, { isLoading }] = authApi.useLazyLoginQuery();
    const dispatch = useAppDispatch();

    const handleLogin = () => {
        login().then(async ({ data }) => {
            if (data) {
                await Linking.openURL(data.redirect_uri);
                dispatch(setAccessToken(data.sac_session));
            }
        });
    };

    return (
        <SafeAreaView style={{ backgroundColor: 'white', height: '100%' }}>
            <GlobalLayout>
                <Box
                    flexDirection="column"
                    justifyContent="space-between"
                    height="100%"
                    paddingVertical="xxl"
                >
                    <Box backgroundColor="white">
                        <Text variant="subheader-1">Welcome to the</Text>
                        <Text variant="header-1" color="darkRed">
                            Student Activity Calendar
                        </Text>
                    </Box>
                    <Button
                        variant="standardButton"
                        color="darkRed"
                        onPress={() => handleLogin()}
                    >
                        {isLoading ? (
                            <Image
                                source={Loading}
                                style={{
                                    width: 15,
                                    height: 15
                                }}
                            />
                        ) : (
                            'Continue with Outlook'
                        )}
                    </Button>
                </Box>
            </GlobalLayout>
        </SafeAreaView>
    );
};

export default WelcomePage;
