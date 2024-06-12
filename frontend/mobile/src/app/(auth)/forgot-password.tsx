import {
    Keyboard,
    KeyboardAvoidingView,
    Platform,
    Pressable
} from 'react-native';

import { router } from 'expo-router';

import { faUserLock } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { ForgotPasswordForm } from './components/forgot-password-form';
import { SmallText } from './components/small-text';
import { Box, Text } from '../(design-system)';
export const ForgotPassword = () => {
    return (
        <Box
            flex={1}
            backgroundColor="white"
            flexDirection="column"
            alignItems="center"
            paddingBottom="xl"
            justifyContent="space-between"
        >
            <KeyboardAvoidingView
                style={{ flex: 1, backgroundColor: 'white' }}
                behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
            >
                <Pressable
                    style={{ flex: 1 }}
                    onPress={() => Keyboard.dismiss()}
                >
                    <Box
                        paddingHorizontal="xl"
                        flex={1}
                        alignItems="center"
                        justifyContent="center"
                        flexDirection="column"
                    >
                        <FontAwesomeIcon size={200} icon={faUserLock} />
                        <Box paddingHorizontal="m">
                            <Text paddingTop="m" textAlign="center">
                                Please enter the email address associated with
                                your account
                            </Text>
                        </Box>
                        <ForgotPasswordForm />
                    </Box>
                </Pressable>
            </KeyboardAvoidingView>
            <SmallText
                first="Remember password?"
                second="Login"
                onPress={() => router.push('/(auth)/login')}
            />
        </Box>
    );
};

export default ForgotPassword;
