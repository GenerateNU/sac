import { useState } from 'react';
import { Alert, Text, TouchableOpacity, View } from 'react-native';
import Spinner from 'react-native-loading-spinner-overlay';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';

import { useAuthStore } from '@/hooks/use-auth';
import { requestVerification } from '@/services/auth';

import Code from './code';

const verificationSchema = z.object({
    code: z.string().min(6, {
        message: 'Verification code must be 6 characters long'
    })
});

const VerificationForm = () => {
    const { user, onVerify } = useAuthStore();
    const [loading, setLoading] = useState(false);

    const onPressVerify = async (code: string) => {
        setLoading(true);

        try {
            verificationSchema.parse({ code });
            await onVerify(user?.email as string, code);
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
            <Code onCompleted={onPressVerify} />
            <View className="mt-[5%] flex-row justify-center">
                <Text className="text-gray-500 text-md">
                    Didn't receive the code?
                </Text>
                <TouchableOpacity
                    onPress={requestCode}
                    className="font-bold underline"
                >
                    <Text className="font-bold text-gray-500 underline text-md">
                        {' '}
                        Resend
                    </Text>
                </TouchableOpacity>
            </View>
        </>
    );
};

export { VerificationForm };
