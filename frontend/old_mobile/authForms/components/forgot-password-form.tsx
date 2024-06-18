import { Controller, useForm } from 'react-hook-form';
import { Alert } from 'react-native';

import { zodResolver } from '@hookform/resolvers/zod';
import { ZodError, z } from 'zod';

import { Box } from '@/src/app/(design-system)';
import { Button } from '@/src/app/(design-system)/components/Button/Button';
import { TextBox } from '@/src/app/(design-system)/components/Textbox/Textbox';

type ForgotPasswordFormData = {
    email: string;
};

const forgotPasswordSchema = z.object({
    email: z.string().email({ message: 'Invalid email' })
});

export const ForgotPasswordForm = () => {
    const {
        control,
        handleSubmit,
        trigger,
        formState: { isValid, errors }
    } = useForm<ForgotPasswordFormData>({
        resolver: zodResolver(forgotPasswordSchema),
        mode: 'onTouched'
    });

    const resetPassword = async (loginData: ForgotPasswordFormData) => {
        try {
            const validData = forgotPasswordSchema.parse(loginData);
            console.log({ validData });
        } catch (err: any) {
            if (err instanceof ZodError) {
                Alert.alert(err.errors[0].message);
            } else {
                console.log(err);
            }
        }
    };

    return (
        <>
            <Box paddingVertical="m">
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
            </Box>
            <Button
                onPress={handleSubmit(resetPassword)}
                disabled={!isValid}
                variant="standardButton"
            >
                Reset Password
            </Button>
        </>
    );
};
