import React from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { ZodError, z } from 'zod';
import Wordmark from '@/components/Wordmark';
import Button from '@/components/button';
import Header from '@/components/header';
import Input from '@/components/input';
import Error from '@/components/error';
import { useAuthStore } from '@/hooks/use-auth';
import { loginByEmail } from '@/services/auth';

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
        <SafeAreaView className="bg-neutral-500 h-[100%]" edges={['top']}>
            <View className="flex-1">
                <View className="px-[8%] pb-[10%]">
                    <View className="pt-[1%]"><Wordmark textColor="white" /></View>
                    <View className="pt-[9.5%] pb-[6%]">
                        <Header
                            text="Let's go"
                            fontSize="45"
                            color="white"
                        ></Header>
                    </View>
                    <Text className="text-white pt-[0.9%] text-lg leading-6">
                        Discover, follow, and join all the clubs & events
                        Northeastern has to offer
                    </Text>
                </View>
                <View className="bg-white pt-[13%] pb-[2%] flex-1 rounded-tl-3xl rounded-tr-3xl px-[8%]">
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
                                    onSubmitEditing={handleSubmit(onSubmit)}
                                    error={!!errors.email}
                                />
                            )}
                            name="email"
                            rules={{required: 'Email is required'}}
                        />
                        {errors.email && <Error message={errors.email.message}/>}
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
                                    onSubmitEditing={handleSubmit(onSubmit)}
                                    error={!!errors.password}
                                />
                            )}
                            name="password"
                            rules={{required: 'Password is required'}}
                        />
                        {errors.password && <Error message={errors.password.message}/>}
                    </View>

                    <View className="pb-[8%] flex-row justify-end">
                        <Text>Forgot password?</Text>
                    </View>

                    <View className="flex-row justify-between">
                        <Button
                            backgroundColor="white"
                            title="Sign up"
                            borderColor="gray"
                            fullWidth={false}
                            onPress={() => router.push('/(auth)/register')}
                        />
                        <Button
                            title="Login"
                            color="white"
                            borderColor="gray"
                            backgroundColor="gray"
                            fullWidth={false}
                            onPress={handleSubmit(onSubmit)}
                        />
                    </View>
                    <View className="mt-[9%] flex-row justify-center">
                        <Text className="font-bold">Not a student?</Text>
                        <Text> Continue as a <Text className="underline">guest</Text></Text>
                    </View>
                </View>
            </View>
        </SafeAreaView>
    );
};

export default Login;
