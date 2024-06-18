import {
    Keyboard,
    KeyboardAvoidingView,
    Platform,
    Pressable
} from 'react-native';

import { router } from 'expo-router';

import { Box, Text } from '../(design-system)';
import { ResetPasswordForm } from './components/reset-password-form';
import { SmallText } from './components/small-text';

const ResetPassword = () => {
    return (
        <KeyboardAvoidingView
            style={{ flex: 1, backgroundColor: 'white' }}
            behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        >
            <Pressable style={{ flex: 1 }} onPress={() => Keyboard.dismiss()}>
                <Box backgroundColor="darkRed" flex={1}>
                    <Box
                        paddingHorizontal="xl"
                        paddingBottom="xxl"
                        justifyContent="flex-end"
                        flex={0.9}
                        flexDirection="column"
                    >
                        <Box flexDirection="column" gap="l">
                            <Text
                                color="white"
                                fontSize={40}
                                variant="header-1"
                            >
                                Reset Password
                            </Text>
                            <Text
                                color="white"
                                variant="subheader-2"
                                fontWeight={400}
                            >
                                Discover, follow, and join all the clubs &
                                events Northeastern has to offer
                            </Text>
                        </Box>
                    </Box>
                    <Box
                        padding="xl"
                        justifyContent="space-between"
                        style={{
                            borderTopLeftRadius: 30,
                            borderTopRightRadius: 30
                        }}
                        flex={2}
                        backgroundColor="white"
                        alignItems="center"
                    >
                        <ResetPasswordForm />
                        <SmallText
                            first="Not a student?"
                            second="Continue as Guest"
                            onPress={() => router.push('/(app)/')}
                        />
                    </Box>
                </Box>
            </Pressable>
        </KeyboardAvoidingView>
    );
};

export default ResetPassword;
