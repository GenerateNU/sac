import { Controller, useForm } from 'react-hook-form';
import { Alert, Text, View } from 'react-native';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';

import { Button } from '@/components/button';
import Error from '@/components/error';
import Input from '@/components/input';
import { useAuthStore } from '@/hooks/use-auth';
import { loginByEmail } from '@/services/auth';
import { useSignIn } from '@clerk/clerk-expo';
import { useState } from 'react';
import Spinner from 'react-native-loading-spinner-overlay';

type LoginFormData = {
    email: string;
    password: string;
};

const loginSchema = z.object({
    email: z.string().email({ message: 'Invalid email' }),
    password: z
        .string()
        .min(8, { message: 'Password must be at least 8 characters long' })
});

const LoginForm = () => {
    const {
        control,
        handleSubmit,
        formState: { errors }
    } = useForm<LoginFormData>();

    const { signIn, setActive, isLoaded } = useSignIn();
    const [loading, setLoading] = useState(false);

    const onSignInPress = async (loginData: LoginFormData) => {
        if (!isLoaded) {
            return;
        }

        setLoading(true);

        try {
            const validData = loginSchema.parse(loginData);

            const completeSignIn = await signIn.create({
                identifier: validData.email,
                password: validData.password
            });

            await setActive({ session: completeSignIn.createdSessionId });
        } catch (err: any) {
            if (err instanceof ZodError) {
                Alert.alert(err.errors[0].message);
            } else {
                Alert.alert('An error occurred', err.message);
            }
        } finally {
            setLoading(false);
        }
    };

    return (
        <>
            {loading && <Spinner visible={loading} />}
            <View>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <Input
                            title="Email"
                            autoCorrect={false}
                            placeholder="ladley.g@northeastern.edu"
                            onChangeText={onChange}
                            value={value}
                            onSubmitEditing={handleSubmit(onSignInPress)}
                            error={!!errors.email}
                        />
                    )}
                    name="email"
                    rules={{ required: 'Email is required' }}
                />
                {errors.email && <Error message={errors.email.message} />}
            </View>
            <View className="w-full mt-[8%] mb-[3%]">
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <Input
                            title="Password"
                            autoCorrect={false}
                            placeholder="Password"
                            onChangeText={onChange}
                            value={value}
                            secureTextEntry={true}
                            onSubmitEditing={handleSubmit(onSignInPress)}
                            error={!!errors.password}
                        />
                    )}
                    name="password"
                    rules={{ required: 'Password is required' }}
                />
                {errors.password && <Error message={errors.password.message} />}
            </View>

            <View className="pb-[8%] flex-row justify-end">
                <Text>Forgot password?</Text>
            </View>

            <View className="flex-row justify-between">
                <Button
                    size="lg"
                    variant="outline"
                    onPress={() => router.push('/(auth)/register')}
                >
                    Sign up
                </Button>
                <Button
                    size="lg"
                    variant="default"
                    onPress={handleSubmit(onSignInPress)}
                >
                    Log in
                </Button>
            </View>
        </>
    );
};

export default LoginForm;
