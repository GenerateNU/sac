import { useState } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert } from 'react-native';

import { router } from 'expo-router';

import { zodResolver } from '@hookform/resolvers/zod';
import { ZodError, z } from 'zod';

import { Box } from '@/app/(design-system)';
import { Button } from '@/app/(design-system)/components/Button/Button';
import { TextBox } from '@/app/(design-system)/components/Textbox/Textbox';

type ResetPasswordFormData = {
    password: string;
    passwordConfirm: string;
};

const passwordSchema = z
    .object({
        password: z
            .string()
            .regex(/[0-9]/, {
                message: 'Password must contain at least one number'
            })
            .regex(/[^A-Za-z0-9]/, {
                message: 'Password must contain at least one special character'
            })
            .min(8, { message: 'Password must be at least 8 characters long' }),
        passwordConfirm: z.string().min(1, { message: 'Required' })
    })
    .refine((data) => data.password === data.passwordConfirm, {
        path: ['passwordConfirm']
    });

export const ResetPasswordForm = () => {
    const {
        control,
        handleSubmit,
        trigger,
        reset,
        formState: { isValid, errors }
    } = useForm<ResetPasswordFormData>({
        resolver: zodResolver(passwordSchema),
        mode: 'onTouched'
    });

    const onLogInPress = async (resetPasswordData: ResetPasswordFormData) => {
        try {
            const validData = passwordSchema.parse(resetPasswordData);
            console.log({ validData });
            router.push('/(auth)/login');
        } catch (err: any) {
            if (err instanceof ZodError) {
                Alert.alert(err.errors[0].message);
            } else {
                reset();
            }
        }
    };

    const [isPasswordConfirmationTouched, setIsPasswordConfirmationTouched] =
        useState(false);

    return (
        <Box flexDirection="column">
            <Box paddingBottom="xl" gap="m">
                <Controller
                    name="password"
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <>
                            <TextBox
                                onChangeText={(text: string) => {
                                    onChange(text);
                                    trigger('password');
                                    setIsPasswordConfirmationTouched(false);
                                }}
                                helperText="8+ characters with numbers & special characters"
                                value={value}
                                secureTextEntry={true}
                                title="New Password"
                                placeholder="Enter password"
                                error={
                                    errors.password
                                        ? errors.password.message
                                        : undefined
                                }
                            />
                        </>
                    )}
                />
                <Controller
                    name="passwordConfirm"
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <>
                            <TextBox
                                title="Password Confirmation"
                                placeholder="Confirm password"
                                onChangeText={(text: string) => {
                                    onChange(text);
                                    trigger('passwordConfirm');
                                    setIsPasswordConfirmationTouched(true);
                                }}
                                value={value}
                                secureTextEntry={true}
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
                <Box
                    style={{ marginTop: -30 }}
                    width="100%"
                    flexDirection="row"
                    justifyContent="flex-end"
                />
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
                    Reset Password
                </Button>
            </Box>
        </Box>
    );
};
