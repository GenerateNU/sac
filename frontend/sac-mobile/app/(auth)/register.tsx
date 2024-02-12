import React from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert, Button, Text, TextInput, View } from 'react-native';

import { router } from 'expo-router';
import { ZodError, z } from 'zod';

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
        // Validate form data using Zod schema
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
                    defaultValue=""
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
                    defaultValue=""
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
                    defaultValue=""
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
                    defaultValue=""
                />
                {errors.password && <Text>{errors.password.message}</Text>}
            </View>

            <Button title="Submit" onPress={handleSubmit(onSubmit)} />
            <View className="mt-4">
                <Button
                    title="Login"
                    onPress={() => router.push('/(auth)/login')}
                />
            </View>
        </View>
    );
};

export default Register;
