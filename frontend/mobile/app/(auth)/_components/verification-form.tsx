import { useEffect, useState } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert, Text, TouchableOpacity, View } from 'react-native';
import Spinner from 'react-native-loading-spinner-overlay';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';

import { Button } from '@/components/button';
import { Error } from '@/components/error';
import { Input } from '@/components/input';
import { useAuthStore } from '@/hooks/use-auth';
import { requestVerification } from '@/services/auth';

type VerificationFormData = {
    code: string;
};

const verificationSchema = z.object({
    code: z.string().min(6, {
        message: 'Verification code must be 6 characters long'
    })
});

const VerificationForm = () => {
    const {
        control,
        handleSubmit,
        formState: { errors }
    } = useForm<VerificationFormData>();

    const { user, completeVerification } = useAuthStore();
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        console.log('[verification-form]', { user });
    }, [user]);

    const onVerify = async ({ code }: VerificationFormData) => {
        setLoading(true);

        try {
            verificationSchema.parse({ code });
            await completeVerification(user?.email as string, code);
            router.replace('/(auth)/user-details');
        } catch (error: any) {
            if (error instanceof ZodError) {
                Alert.alert('Error', error.errors[0].message);
            } else {
                Alert.alert('Error', error.message);
            }
        } finally {
            setLoading(false);
        }
    };

    const requestCode = async () => {
        try {
            await requestVerification(user?.email as string);
        } catch (error: any) {
            Alert.alert('Error', error.message);
        }
    };

    return (
        <>
            {loading && <Spinner visible={loading} />}
            <View className="w-full mb-[8.5%]">
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <Input
                            title="Verification code"
                            placeholder="Verification code"
                            onChangeText={onChange}
                            value={value}
                            keyboardType="numeric"
                            maxLength={6}
                        />
                    )}
                    name="code"
                    rules={{
                        required: 'Verification code is required'
                    }}
                />
                {errors.code && <Error message={errors.code.message} />}
            </View>
            <View className="pt-[2%]">
                <Button onPress={handleSubmit(onVerify)} size="screenwide">
                    Submit
                </Button>
            </View>
            <View className="mt-[5%] flex-row justify-center">
                <Text>Did not receive a code?</Text>
                <TouchableOpacity
                    onPress={requestCode}
                    className="font-bold underline"
                >
                    <Text> Resend</Text>
                </TouchableOpacity>
            </View>
        </>
    );
};

export { VerificationForm };
