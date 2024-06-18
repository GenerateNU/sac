import React, { useEffect } from 'react';

import { Box, Button, Text } from '../design-system';
import { Image, Linking, SafeAreaView } from 'react-native';
import { GlobalLayout } from '../design-system/components/GlobalLayout/GlobalLayout';
import { authApi } from '@generatesac/lib';
import Loading from '@/src/assets/gif/loading.gif';
import { useLocalSearchParams } from 'expo-router';
import { useAppDispatch } from '@/src/store/store';
import { setSacLoginSession } from '@/src/store/slices/userSlice';

const WelcomePage = () => {
    const [login, { isLoading, error, data }] = authApi.useLazyLoginQuery();
    const dispatch = useAppDispatch();
    const search = useLocalSearchParams();
    console.log(search);

    const handleLogin = async () => {
        await login().then(async () => {
            if (data) {
                dispatch(setSacLoginSession(data.sac_session));
                await Linking.openURL(data.redirect_uri);
            }
        })
    };

    return (
        <SafeAreaView style={{ backgroundColor: 'white', height: '100%' }}>
            <GlobalLayout>
                <Box 
                    flexDirection='column'
                    justifyContent='space-between'
                    height='100%'
                    paddingVertical='xxl'
                >
                    <Box backgroundColor="white">
                        <Text variant='subheader-1'>
                            Welcome to the
                        </Text>
                        <Text variant='header-1' color='darkRed'>
                            Student Activity Calendar
                        </Text>
                    </Box>
                    <Button 
                        variant='standardButton'
                        color='darkRed'
                        onPress={() => handleLogin()}
                    >
                        {isLoading ? 
                            <Image 
                                source={Loading}
                                style={{
                                    width: 15,
                                    height: 15
                                }}
                            /> 
                            : 'Continue with Outlook'
                        }
                    </Button>
                </Box>
            </GlobalLayout>
        </SafeAreaView>
    );
};

export default WelcomePage;
