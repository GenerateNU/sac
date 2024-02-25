import React from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert, Button, StyleSheet, Text, TextInput, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';

import Wordmark from '@/components/Wordmark';
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
    // className="items-center justify-center flex-1 p-4"
    return (
        <SafeAreaView style={styles.container}>
            <View style={styles.innerContainer}>
                <View className="w-full p-5">
                    <Wordmark textColor="white" />
                    <Text style={styles.header}>Let's Go</Text>
                    <Text style={styles.description}>
                        Discover, follow, and join all the clubs & events
                        Northeastern has to offer
                    </Text>
                </View>
                <View>
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

                <View style={styles.buttonContainer}>
                    <Button title="Login" onPress={handleSubmit(onSubmit)} />
                    <Button
                        title="Register"
                        onPress={() => router.push('/(auth)/register')}
                    />
                </View>
                <Text>Not a student? Continue as a Guest</Text>
            </View>
        </SafeAreaView>
    );
};

const styles = StyleSheet.create({
    header: {
        fontSize: 40,
        fontWeight: 'bold',
        color: 'white'
    },
    description: {
        fontSize: 20,
        color: 'white'
    },
    container: {
        backgroundColor: 'grey',
        height: '100%'
    },
    innerContainer: {
        marginLeft: '5%',
        marginRight: '5%'
    },
    buttonContainer: {
        flexDirection: 'row',
        justifyContent: 'space-around',
    }
});

export default Login;
