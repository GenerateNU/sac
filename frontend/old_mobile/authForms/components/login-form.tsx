import { useState } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert } from 'react-native';
import { TouchableOpacity } from 'react-native-gesture-handler';

import { router } from 'expo-router';

import { zodResolver } from '@hookform/resolvers/zod';
import { ZodError, z } from 'zod';

import { Box, Text } from '@/src/app/(design-system)';
import { Button } from '@/src/app/(design-system)/components/Button/Button';
import { TextBox } from '@/src/app/(design-system)/components/Textbox/Textbox';

type LoginFormData = {
    email: string;
    password: string;
};

const loginSchema = z.object({
    email: z.string().email({ message: 'Invalid email' }),
    password: z.string().min(1, { message: 'Password required' })
});

export const LoginForm = () => {
    const {
        control,
        handleSubmit,
        trigger,
        reset,
        formState: { isValid, errors }
    } = useForm<LoginFormData>({
        resolver: zodResolver(loginSchema),
        mode: 'onTouched'
    });

    const [loginError, setLoginError] = useState(false);
    const failedToLogin =
        'The email or password you entered did not match our records. Please try again.';

    const onLogInPress = async (loginData: LoginFormData) => {
        try {
            const validData = loginSchema.parse(loginData);
            console.log({ validData });
            router.push('/(app)/');
        } catch (err: any) {
            if (err instanceof ZodError) {
                Alert.alert(err.errors[0].message);
            } else {
                setLoginError(true);
                reset();
            }
        }
    };

    return (
        <Box flexDirection="column">
            <Box paddingBottom="xl" gap="m">
                {loginError && <Text color="darkRed">{failedToLogin}</Text>}
                <Controller
                    name="email"
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <>
                            <TextBox
                                onChangeText={(text: string) => {
                                    onChange(text);
                                    trigger('email');
                                }}
                                value={value}
                                title="Email"
                                placeholder="email@northeastern.edu"
                                error={
                                    errors.email
                                        ? errors.email.message
                                        : undefined
                                }
                            />
                        </>
                    )}
                />
                <Controller
                    name="password"
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <>
                            <TextBox
                                title="Password"
                                placeholder="Enter password"
                                onChangeText={(text: string) => {
                                    onChange(text);
                                    trigger('password');
                                }}
                                value={value}
                                secureTextEntry={true}
                                error={
                                    errors.password
                                        ? errors.password.message
                                        : undefined
                                }
                            />
                        </>
                    )}
                />
                <Box
                    style={{ marginTop: -30 }}
                    width="100%"
                    flexDirection="row"
                    justifyContent="flex-end"
                >
                    <TouchableOpacity
                        onPress={() => router.push('/(auth)/forgot-password')}
                    >
                        <Text>Forgot password?</Text>
                    </TouchableOpacity>
                </Box>
            </Box>
            <Box
                flexDirection="column"
                alignItems="center"
                justifyContent="space-between"
                gap="m"
            >
                <Button
                    disabled={!isValid}
                    onPress={handleSubmit(onLogInPress)}
                    variant="standardButton"
                >
                    Login
                </Button>
                <Button
                    backgroundColor="white"
                    borderColor="black"
                    borderWidth={1}
                    onPress={() => router.push('/(auth)/signup')}
                    variant="standardButton"
                >
                    <Text color="black">Sign Up</Text>
                </Button>
            </Box>
        </Box>
    );
};
