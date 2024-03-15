import React from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert, ScrollView, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';

import { Button } from '@/components/button';
import { DropdownComponent } from '@/components/dropdown';
import Error from '@/components/error';
import Input from '@/components/input';
import Wordmark from '@/components/wordmark';
import { Item, graduationYear } from '@/lib/utils';

// register form data
type RegisterFormData = {
    firstName: string;
    lastName: string;
    email: string;
    nuid: string;
    graduationYear: Item;
    password: string;
    passwordConfirm: string;
};

const registerSchema = z.object({
    firstName: z
        .string()
        .min(2, { message: 'First name must be at least 2 characters long' }),
    lastName: z
        .string()
        .min(2, { message: 'Last name must be at least 2 characters long' }),
    email: z.string().email({ message: 'Invalid email' }),
    nuid: z.string().length(9, { message: 'NUID must have 9 digits' }),
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
            const { graduationYear, passwordConfirm, ...rest } = data;
            const updatedData = {
                ...rest,
                graduationYear: graduationYear.value
            };
            registerSchema.parse(updatedData);
            Alert.alert('Form Submitted', JSON.stringify(updatedData));
            router.push('/(auth)/majorAndCollege');
        } catch (error) {
            if (error instanceof ZodError) {
                Alert.alert('Validation Error', error.errors[0].message);
            } else {
                console.error('An unexpected error occurred:', error);
            }
        }
    };

    return (
        <SafeAreaView className="bg-neutral-500" edges={['top']}>
            <ScrollView>
                <View className="px-[8%] pb-[9%]">
                    <Wordmark
                        textColor="text-white"
                        func={() => router.push('/(auth)/login')}
                        title="Login"
                    />
                    <View className="pt-[9%] pb-[7.5%]">
                        <Text className="text-white font-bold text-5xl">
                            Sign up
                        </Text>
                    </View>
                    <Text className="text-lg leading-6 text-white">
                        Discover, follow, and join all the clubs & events
                        Northeastern has to offer
                    </Text>
                </View>
                <View className="bg-white px-[8%] pt-[13%] rounded-t-3xl">
                    <View className="w-full mb-[8.5%]">
                        <Controller
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <Input
                                    title="First Name"
                                    autoCorrect={false}
                                    placeholder="Garrett"
                                    onChangeText={(text) => onChange(text)}
                                    value={value}
                                    onSubmitEditing={handleSubmit(onSubmit)}
                                    error={!!errors.firstName}
                                />
                            )}
                            name="firstName"
                            rules={{
                                required: 'First name is required',
                                validate: (value) => {
                                    const isValid = /^[a-zA-Z]+$/.test(value);
                                    if (!isValid) {
                                        return 'Please enter proper first name';
                                    } else if (value.length < 2) {
                                        return 'First name must be at least 2 characters long';
                                    }
                                    return true;
                                }
                            }}
                        />
                        {errors.firstName && (
                            <Error message={errors.firstName.message} />
                        )}
                    </View>

                    <View className="w-full mb-[8.5%]">
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
                                    error={!!errors.lastName}
                                />
                            )}
                            name="lastName"
                            rules={{
                                required: 'Last name is required',
                                validate: (value) => {
                                    const isValid = /^[a-zA-Z]+$/.test(value);
                                    if (!isValid) {
                                        return 'Please enter proper last name';
                                    } else if (value.length < 2) {
                                        return 'Last name must be at least 2 characters long';
                                    }
                                    return true;
                                }
                            }}
                        />
                        {errors.lastName && (
                            <Error message={errors.lastName.message} />
                        )}
                    </View>

                    <View className="w-full mb-[8.5%]">
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
                                    error={!!errors.email}
                                />
                            )}
                            name="email"
                            rules={{
                                required: 'Email is required',
                                validate: (value) => {
                                    if (!value.endsWith('@northeastern.edu')) {
                                        return 'Please enter your Northeastern email';
                                    }
                                    return true;
                                }
                            }}
                        />
                        {errors.email && (
                            <Error message={errors.email.message} />
                        )}
                    </View>

                    <View className="w-full mb-[8.5%]">
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
                                    error={!!errors.nuid}
                                />
                            )}
                            name="nuid"
                            rules={{
                                required: 'NUID is required',
                                validate: (value) => {
                                    if (
                                        !/^00\d+/.test(value) ||
                                        /[^\d]/.test(value)
                                    ) {
                                        return 'Please enter a proper NUID number';
                                    }
                                    if (value.length != 9) {
                                        return 'Please enter 9 digit number';
                                    }
                                    return true;
                                }
                            }}
                        />
                        {errors.nuid && <Error message={errors.nuid.message} />}
                    </View>
                    <View className="mb-[7%]">
                        <Controller
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <DropdownComponent
                                    title="Intended Graduation Year"
                                    item={graduationYear()}
                                    placeholder="Select Year"
                                    onChangeText={onChange}
                                    value={value}
                                    onSubmitEditing={handleSubmit(onSubmit)}
                                    error={!!errors.graduationYear}
                                />
                            )}
                            name="graduationYear"
                            rules={{ required: 'Graduation year is required' }}
                        />
                        {errors.graduationYear && (
                            <Error message={errors.graduationYear.message} />
                        )}
                    </View>

                    <View className="w-full mb-[8.5%]">
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
                                    error={!!errors.password}
                                />
                            )}
                            name="password"
                            rules={{
                                required: 'Password is required',
                                validate: (value) => {
                                    let specialChars =
                                        /[`!@#$%^&*()_\-+=\[\]{};':"\\|,.<>\/?~ ]/;
                                    if (value.length < 8) {
                                        return 'Password must have at least 8 characters';
                                    } else if (!specialChars.test(value)) {
                                        return 'Please contain at least one special character';
                                    } else if (!/^.*[0-9]+.*$/.test(value)) {
                                        return 'Please contain at least one number';
                                    }
                                    return true;
                                }
                            }}
                        />
                        {errors.password && (
                            <Error message={errors.password.message} />
                        )}
                    </View>

                    <View className="w-full mb-[8.5%]">
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
                                    error={!!errors.password}
                                />
                            )}
                            name="passwordConfirm"
                            rules={{
                                required: 'Please confirm your password',
                                validate: (value) => {
                                    const password =
                                        control._getWatch('password');
                                    if (value !== password) {
                                        return 'Passwords do not match';
                                    }
                                    return true;
                                }
                            }}
                        />
                        {errors.passwordConfirm && (
                            <Error message={errors.passwordConfirm.message} />
                        )}
                    </View>
                    <View className="pt-[2%] pb-[15%]">
                        <Button
                            onPress={handleSubmit(onSubmit)}
                            size="screenwide"
                        >
                            Submit
                        </Button>
                    </View>
                </View>
            </ScrollView>
        </SafeAreaView>
    );
};

export default Register;
