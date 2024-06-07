import React, { useEffect, useState } from 'react';
import {
    Keyboard,
    KeyboardAvoidingView,
    Platform,
    Pressable,
    StyleSheet
} from 'react-native';

import { router } from 'expo-router';

import {
    faClockRotateLeft,
    faEnvelopeCircleCheck
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Box, Colors, Spacing, Text } from '../(design-system)';
import { Border } from '../(design-system)/shared/border';
import EmailVerificationForm from './components/email-verification';
import { SmallText } from './components/small-text';

const EmailVerification: React.FC = () => {
    const email = 'email@northeastern.edu';
    const [code, setCode] = useState<string>('');
    const [sendCodeDisabled, setSendCodeDisabled] = useState(false);
    const [timer, setTimer] = useState<number>(0);
    const [error, setError] = useState('');
    const MOCK_TEST = '123456';

    const handleResendCode = () => {
        setSendCodeDisabled(true);
        setTimer(90);
    };

    const handleCodeChange = (text: string) => {
        setCode(text);
        if (text.length === 6) {
            if (text === MOCK_TEST) {
                // try catch api hooks here
                setError('');
                router.push('/(auth)/user-details');
            } else {
                setError('Wrong code. Please try again.');
            }
        } else {
            setError('');
        }
    };

    useEffect(() => {
        let interval: NodeJS.Timeout | null = null;
        if (sendCodeDisabled) {
            interval = setInterval(() => {
                setTimer((prevTimer) => {
                    if (prevTimer <= 1) {
                        clearInterval(interval!);
                        setSendCodeDisabled(false);
                        return 0;
                    }
                    return prevTimer - 1;
                });
            }, 1000);
        }
        return () => {
            if (interval) {
                clearInterval(interval);
            }
        };
    }, [sendCodeDisabled]);

    const formatTime = (seconds: number) => {
        const mins = Math.floor(seconds / 60);
        const secs = seconds % 60;
        return `${mins}:${secs < 10 ? '0' : ''}${secs}`;
    };

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
                        <FontAwesomeIcon
                            size={200}
                            icon={faEnvelopeCircleCheck}
                        />
                        <Box flexDirection="column" paddingHorizontal="m">
                            <Text
                                textAlign="center"
                                fontSize={40}
                                variant="header-1"
                            >
                                Confirm your email address
                            </Text>
                            <Text paddingTop="xl" textAlign="center">
                                Enter the 6-digit confirmation code sent to:
                            </Text>
                            <Text fontWeight={500} textAlign="center">
                                {email}
                            </Text>
                        </Box>
                        <Box style={styles.codeFieldRoot}>
                            <EmailVerificationForm
                                error={error}
                                onComplete={handleCodeChange}
                                value={code}
                            />
                        </Box>
                        {timer !== 0 && (
                            <Box
                                gap="s"
                                paddingTop="xxl"
                                flexDirection="row"
                                alignItems="center"
                            >
                                <FontAwesomeIcon
                                    color={Colors.gray}
                                    icon={faClockRotateLeft}
                                />
                                <Text color="darkGray">
                                    {formatTime(timer)}
                                </Text>
                            </Box>
                        )}
                    </Box>
                </Pressable>
            </KeyboardAvoidingView>
            <SmallText
                disabled={sendCodeDisabled}
                onPress={handleResendCode}
                first="Didn't receive an email?"
                second="Resend code"
            />
        </Box>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1
    },
    codeFieldRoot: {
        width: '80%',
        marginLeft: 'auto',
        marginRight: 'auto',
        marginTop: Spacing.m,
        alignItems: 'center',
        justifyContent: 'center',
        gap: 5
    },
    cell: {
        width: 50,
        height: 50,
        lineHeight: 46,
        fontSize: 24,
        borderWidth: 2,
        borderColor: '#00000030',
        textAlign: 'center',
        borderRadius: Border.base
    },
    focusCell: {
        borderColor: Colors.black
    }
});

export default EmailVerification;
