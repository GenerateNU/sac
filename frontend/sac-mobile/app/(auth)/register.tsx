import React from 'react';
import { Controller, useForm } from 'react-hook-form';
import {
    Alert,
    ScrollView,
    ScrollViewBase,
    StyleSheet,
    Text,
    TextInput,
    View
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';

import Wordmark from '@/components/Wordmark';
import Button from '@/components/button';
import { DropdownComponent } from '@/components/dropdown';
import Header from '@/components/header';
import Input from '@/components/input';

// list of items for dropdown menu
type Item = {
    label: string,
    value: string,
}

// list of graduation year
const graduationYear = () => {
    var year = new Date().getFullYear();
    const graduationYear: Item[] = [];
    for (let i = 0; i < 5; i++) {
        graduationYear.push({
            label: String(year + i),
            value: String(year + i)
        });
    }
    return graduationYear;
}

type RegisterFormData = {
    firstName: string;
    lastName: string;
    email: string;
    password: string;
    passwordConfirm: string;
    id: string;
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
        <ScrollView className="scroll-smooth">
        <SafeAreaView
            edges={['top']}
            className="bg-neutral-500"
        >
            <View className="px-[8%] pb-[9%]">
                <Wordmark
                    textColor="white"
                    button={true}
                    func={() => router.push('/(auth)/login')}
                    title="Login"
                />
                <View className="pt-[9%] pb-[6%]">
                    <Header text="Sign up" fontSize="45" color="white"></Header>
                </View>
                <Text className="text-lg text-white">
                    Discover, follow, and join all the clubs & events
                    Northeastern has to offer
                </Text>
            </View>
            <View className="bg-white px-[8%] pt-[13%] rounded-t-3xl">
                <View className="w-full mb-7">
                    <Controller
                        control={control}
                        render={({ field: { onChange, value } }) => (
                            <Input
                                title="First Name"
                                autoCorrect={false}
                                placeholder="Garrett"
                                onChangeText={onChange}
                                value={value}
                                onSubmitEditing={handleSubmit(onSubmit)}
                            />
                        )}
                        name="firstName"
                        rules={{ required: 'First name is required' }}
                    />
                    {errors.firstName && (
                        <Text>{errors.firstName.message}</Text>
                    )}
                </View>

                <View className="w-full mb-7">
                    <Controller
                        control={control}
                        render={({ field: { onChange, value } }) => (
                            <Input
                                title="Last Name"
                                autoCorrect={false}
                                placeholder="Ladley"
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

                <View className="w-full mb-7">
                    <Controller
                        control={control}
                        render={({ field: { onChange, value } }) => (
                            <Input
                                title="Email"
                                autoCorrect={false}
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

                <View className="w-full mb-7">
                    <Controller
                        control={control}
                        render={({ field: { onChange, value } }) => (
                            <Input
                                title="NUID"
                                autoCorrect={false}
                                placeholder="9 digit student ID number"
                                onChangeText={onChange}
                                value={value}
                                onSubmitEditing={handleSubmit(onSubmit)}
                            />
                        )}
                        name="id"
                        rules={{ required: 'Password is required' }}
                    />
                    {errors.password && <Text>{errors.password.message}</Text>}
                </View>

                <View className="mb-6">
                    <DropdownComponent
                        title="Intended Graduation Year"
                        item={graduationYear()}
                        placeholder="Select Year"
                    />
                </View>

                <View className="w-full mb-7">
                    <Controller
                        control={control}
                        render={({ field: { onChange, value } }) => (
                            <Input
                                title="Password"
                                autoCorrect={false}
                                placeholder="Password"
                                onChangeText={onChange}
                                value={value}
                                onSubmitEditing={handleSubmit(onSubmit)}
                                secureTextEntry={true}
                            />
                        )}
                        name="password"
                        rules={{ required: 'Password is required' }}
                    />
                    {errors.password && <Text>{errors.password.message}</Text>}
                </View>

                <View className="w-full mb-7">
                    <Controller
                        control={control}
                        render={({ field: { onChange, value } }) => (
                            <Input
                                title="Password Confirmation"
                                autoCorrect={false}
                                placeholder="Confirm your password"
                                onChangeText={onChange}
                                value={value}
                                onSubmitEditing={handleSubmit(onSubmit)}
                                secureTextEntry={true}
                            />
                        )}
                        name="passwordConfirm"
                        rules={{ required: 'Password is required' }}
                    />
                    {errors.password && <Text>{errors.password.message}</Text>}
                </View>
                <View className="pt-[2%] pb-[25%]">
                    <Button
                        title="Submit"
                        fullWidth={true}
                        color="white"
                        onPress={handleSubmit(onSubmit)}
                    />
                </View>
            </View>
        </SafeAreaView>
        </ScrollView>
    );
};

export default Register;
