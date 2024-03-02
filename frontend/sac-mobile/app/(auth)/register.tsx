import React from 'react';
import { Controller, useForm } from 'react-hook-form';
import {
    Alert,
    ScrollView,
    Text,
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
import Error from '@/components/error';

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

// register form data
type RegisterFormData = {
    firstName: string;
    lastName: string;
    email: string;
    id: string;
    password: string;
    passwordConfirm: string;
    graduationYear: string;
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
                                onChangeText={onChange}
                                value={value}
                                onSubmitEditing={handleSubmit(onSubmit)}
                                error={!!errors.firstName}
                            />
                        )}
                        name="firstName"
                        rules={{ required: 'First name is required' }}
                    />
                    {errors.firstName && (
                        <Error message={errors.firstName.message}/>
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
                        rules={{ required: 'Last name is required' }}
                    />
                    {errors.lastName && <Error message={errors.lastName.message}/>}
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
                        rules={{ required: 'Email is required' }}
                    />
                    {errors.email && <Error message={errors.email.message}/>}
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
                                error={!!errors.id}
                            />
                        )}
                        name="id"
                        rules={{ required: 'NUID is required' }}
                    />
                    {errors.id && <Error message={errors.id.message}/>}
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
                                height="8.5%"
                                error={!!errors.graduationYear}
                            />
                        )}
                        name="graduationYear"
                        rules={{ required: 'Graduation year is required' }}
                    />
                    {errors.graduationYear && <View className="mt-[-3.5%]"><Error message={errors.graduationYear.message}/></View>}
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
                        rules={{ required: 'Password is required' }}
                    />
                    {errors.password && <Error message={errors.password.message}/>}
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
                                error={!!errors.passwordConfirm}
                            />
                        )}
                        name="passwordConfirm"
                        rules={{ required: 'Please confirm your password' }}
                    />
                    {errors.passwordConfirm && <Error message={errors.passwordConfirm.message}/>}
                </View>
                <View className="pt-[2%] pb-[25%]">
                    <Button
                        title="Submit"
                        fullWidth={true}
                        color="white"
                        padding="5%"
                        onPress={handleSubmit(onSubmit)}
                    />
                </View>
            </View>
        </SafeAreaView>
        </ScrollView>
    );
};

export default Register;
