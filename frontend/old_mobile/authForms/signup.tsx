import {
    Keyboard,
    KeyboardAvoidingView,
    Platform,
    Pressable
} from 'react-native';

import { router } from 'expo-router';

import { Box, Text } from '../(design-system)';
import { SignupForm } from './components/signup-form';
import { SmallText } from './components/small-text';

const Signup = () => {
    return (
        <KeyboardAvoidingView
            style={{ flex: 1, backgroundColor: 'white' }}
            behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        >
            <Pressable onPress={() => Keyboard.dismiss()} style={{ flex: 1 }}>
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
                                Sign Up
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
                        style={{
                            borderTopLeftRadius: 30,
                            borderTopRightRadius: 30
                        }}
                        flex={2}
                        backgroundColor="white"
                        alignItems="center"
                    >
                        <Box
                            alignItems="center"
                            flex={1}
                            flexDirection="column"
                            justifyContent="space-between"
                        >
                            <SignupForm />
                            <SmallText
                                first="Already have an account?"
                                second="Login"
                                onPress={() => router.push('/(auth)/login')}
                            />
                        </Box>
                    </Box>
                </Box>
            </Pressable>
        </KeyboardAvoidingView>
    );
};

export default Signup;
