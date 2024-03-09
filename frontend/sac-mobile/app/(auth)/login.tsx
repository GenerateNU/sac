import React from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert, Text, TextInput, View } from 'react-native';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';

import { useAuthStore } from '@/hooks/use-auth';
import { loginByEmail } from '@/services/auth';
import Button from '@/components/button';

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

const Login = () => {
    const {
        control,
        handleSubmit,
        formState: { errors }
    } = useForm<LoginFormData>();
    const { login } = useAuthStore();

    const onSubmit = async (data: LoginFormData) => {
        try {
            loginSchema.parse(data);
            const { user, tokens } = await loginByEmail(
                data.email.toLowerCase(),
                data.password
            );
            login(tokens, user);
            router.push('/(app)/');
        } catch (e: unknown) {
            if (e instanceof ZodError) {
                Alert.alert('Validation Error', e.errors[0].message); // use a better way to display errors
            } else {
                console.error('An unexpected error occurred:', e);
            }
        }
    };

    return (
        <View className="items-center justify-center flex-1 p-4">
            <View className="w-full mb-4">
                <Text>Email</Text>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TextInput
                            autoCapitalize="none"
                            autoCorrect={false}
                            className="p-2 border border-gray-300"
                            placeholder="ladley.g@northeastern.edu"
                            onChangeText={onChange}
                            value={value}
                            onSubmitEditing={handleSubmit(onSubmit)}
                        />
                    )}
                    name="email"
                    rules={{ required: 'Email is required' }}
                />
                {errors.email && <Text>{errors.email.message}</Text>}
            </View>

            <View className="w-full mb-4">
                <Text>Password</Text>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TextInput
                            autoCapitalize="none"
                            autoCorrect={false}
                            className="p-2 border border-gray-300"
                            placeholder="Password"
                            onChangeText={onChange}
                            value={value}
                            secureTextEntry={true}
                            onSubmitEditing={handleSubmit(onSubmit)}
                        />
                    )}
                    name="password"
                    rules={{ required: 'Password is required' }}
                />
                {errors.password && <Text>{errors.password.message}</Text>}
            </View>

            <Button
                onPress={handleSubmit(onSubmit)}
                variant="secondary"
                size={"sm"}
            >
                Login
            </Button>

            <View className="mt-4">
                <Button
                    onPress={() => router.push('/(auth)/register')}
                    variant="outline"
                    size={"sm"}
                >
                    Register
                </Button>
            </View>
        </View>
    );
};

export default Login;
