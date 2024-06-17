import { useState } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert } from 'react-native';

import { router } from 'expo-router';

import { zodResolver } from '@hookform/resolvers/zod';
import { ZodError, z } from 'zod';

import { Box } from '@/src/app/(design-system)';
import { Button } from '@/src/app/(design-system)/components/Button/Button';
import { TextBox } from '@/src/app/(design-system)/components/Textbox/Textbox';

type SignupFormData = {
    fullName: string;
    email: string;
    password: string;
    passwordConfirmation: string;
};

const signupSchema = z
    .object({
        fullName: z.string().min(2, {
            message: 'First name must be at least 2 characters long'
        }),
        email: z
            .string()
            .email({ message: 'Must be Northeastern email' })
            .refine((email) => email.endsWith('@northeastern.edu'), {
                message: 'Must be Northeastern email'
            }),
        password: z
            .string()
            .regex(/[0-9]/, {
                message: 'Password must contain at least one number'
            })
            .regex(/[^A-Za-z0-9]/, {
                message: 'Password must contain at least one special character'
            })
            .min(8, { message: 'Password must be at least 8 characters long' }),
        passwordConfirmation: z.string()
    })
    .refine((data) => data.password === data.passwordConfirmation, {
        path: ['passwordConfirm']
    });

export const SignupForm = () => {
    const {
        control,
        handleSubmit,
        trigger,
        formState: { isValid, errors }
    } = useForm<SignupFormData>({
        resolver: zodResolver(signupSchema),
        mode: 'onTouched'
    });

    const onSignUpPress = async (signupData: SignupFormData) => {
        try {
            const validData = signupSchema.parse(signupData);
            console.log({ validData });
            router.push('/(auth)/email-verification');
        } catch (err: any) {
            if (err instanceof ZodError) {
                Alert.alert(err.errors[0].message);
            } else {
                Alert.alert(err.errors[0].message);
            }
        }
    };

    const [isPasswordConfirmationTouched, setIsPasswordConfirmationTouched] =
        useState(false);

    return (
        <Box flexDirection="column" gap="xs">
            <Controller
                name="fullName"
                control={control}
                render={({ field: { onChange, value } }) => (
                    <>
                        <TextBox
                            onChangeText={(text: string) => {
                                onChange(text);
                                trigger('fullName');
                            }}
                            value={value}
                            title="Full Name"
                            placeholder="Emma Pon"
                            error={errors.fullName && errors.fullName.message}
                        />
                    </>
                )}
            />
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
                            helperText="Must be Northeastern email"
                            placeholder="email@northeastern.edu"
                            error={errors.email && errors.email.message}
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
                            onChangeText={(text: string) => {
                                onChange(text);
                                trigger('password');
                                setIsPasswordConfirmationTouched(false); // Reset password confirmation touched state
                            }}
                            helperText="8+ characters with numbers & special characters"
                            value={value}
                            title="Password"
                            secureTextEntry={true}
                            placeholder="Enter password"
                            error={errors.password && errors.password.message}
                        />
                    </>
                )}
            />
            <Controller
                name="passwordConfirmation"
                control={control}
                render={({ field: { onChange, value } }) => (
                    <>
                        <TextBox
                            onChangeText={(text: string) => {
                                onChange(text);
                                setIsPasswordConfirmationTouched(true);
                            }}
                            value={value}
                            title="Password Confirmation"
                            secureTextEntry={true}
                            placeholder="Re-enter password"
                            error={
                                isPasswordConfirmationTouched &&
                                value !== control._getWatch('password')
                                    ? 'Passwords do not match'
                                    : undefined
                            }
                        />
                    </>
                )}
            />
            <Box paddingTop="s">
                <Button
                    disabled={!isValid}
                    onPress={handleSubmit(onSignUpPress)}
                    variant="standardButton"
                >
                    Sign Up
                </Button>
            </Box>
        </Box>
    );
};
