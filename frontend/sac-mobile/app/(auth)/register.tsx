import React from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert, Text, TextInput, View } from 'react-native';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';

import { Button } from '@/components/button';

type RegisterFormData = {
    firstName: string;
    lastName: string;
    email: string;
    password: string;
};

const registerSchema = z.object({
    firstName: z
        .string()
        .min(2, { message: 'First name must be at least 2 characters long' }),
    lastName: z
        .string()
        .min(2, { message: 'Last name must be at least 2 characters long' }),
    email: z.string().email({ message: 'Invalid email' }),
    password: z
        .string()
        .min(8, { message: 'Password must be at least 8 characters long' })
});

const Register = () => {
    const {
        control,
        handleSubmit,
        formState: { errors }
    } = useForm<RegisterFormData>();

    const onSubmit = (data: RegisterFormData) => {
        try {
            registerSchema.parse(data);
            Alert.alert('Form Submitted', JSON.stringify(data));
            router.push('/(app)/');
        } catch (error) {
            if (error instanceof ZodError) {
                Alert.alert('Validation Error', error.errors[0].message);
            } else {
                console.error('An unexpected error occurred:', error);
            }
        }
    };

    return (
        <View className="items-center justify-center flex-1 p-4">
            <View className="w-full mb-4">
                <Text>First Name</Text>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TextInput
                            className="p-2 border border-gray-300"
                            placeholder="Ladley"
                            onChangeText={onChange}
                            value={value}
                            onSubmitEditing={handleSubmit(onSubmit)}
                        />
                    )}
                    name="firstName"
                    rules={{ required: 'First name is required' }}
                />
                {errors.firstName && <Text>{errors.firstName.message}</Text>}
            </View>

            <View className="w-full mb-4">
                <Text>Last Name</Text>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TextInput
                            className="p-2 border border-gray-300"
                            placeholder="G"
                            onChangeText={onChange}
                            value={value}
                            onSubmitEditing={handleSubmit(onSubmit)}
                        />
                    )}
                    name="lastName"
                    rules={{ required: 'Last name is required' }}
                />
                {errors.lastName && <Text>{errors.lastName.message}</Text>}
            </View>

            <View className="w-full mb-4">
                <Text>Email</Text>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TextInput
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
                size={'sm'}
            >
                Register
            </Button>
            <View className="mt-4">
                <Button
                    variant="outline"
                    size={'sm'}
                    onPress={() => router.push('/(auth)/login')}
                >
                    Login
                </Button>
            </View>
        </View>
    );
};

export default Register;
