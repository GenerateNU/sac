import React from 'react';
import { Text, View, TextInput, Button, Alert } from 'react-native';
import { useForm, Controller } from 'react-hook-form';
import { ZodError, z } from 'zod'; // Import Zod
import { useAuthStore } from '@/hooks/use-auth';
import { router } from 'expo-router';
import { loginByEmail } from '@/services/auth';

type LoginFormData = {
    email: string;
    password: string;
};

// Define Zod schema
const loginSchema = z.object({
    email: z.string().email({ message: 'Invalid email' }), // Email validation
    password: z.string().min(8, { message: 'Password must be at least 8 characters long' }), // Password validation
});

const Login = () => {
    const {
        control,
        handleSubmit,
        formState: { errors },
    } = useForm<LoginFormData>();
    const { login } = useAuthStore();

    const onSubmit = async (data: LoginFormData) => {
        try {
            loginSchema.parse(data);
            const { user, tokens } = await loginByEmail(data.email, data.password)
            login(tokens, user)
            console.log(`Logged in, ${user}, ${tokens}`)
            router.push('/(app)/');
        } catch (e: unknown) {
            if (e instanceof ZodError) {
                Alert.alert('Error 2', e.errors[0].message);
            } else {
                Alert.alert('Error 1')
            }
        }
    }

    return (
        <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center', padding: 20 }}>
            <View style={{ width: '100%', marginBottom: 20 }}>
                <Text>Email</Text>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TextInput
                            style={{ padding: 10, borderColor: 'gray', borderWidth: 1 }}
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

            <View style={{ width: '100%', marginBottom: 20 }}>
                <Text>Password</Text>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TextInput
                            style={{ padding: 10, borderColor: 'gray', borderWidth: 1 }}
                            placeholder="Password"
                            onChangeText={onChange}
                            value={value}
                            // secureTextEntry={true}
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

            <View style={{ marginTop: 20 }}>
                <Button title="Register" onPress={() => router.push('/(auth)/register')} />
            </View>
        </View>
    );
};

export default Login;
