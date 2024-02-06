import React from 'react';
import { Text, View, TextInput, Button, Alert } from 'react-native';
import { useForm, Controller } from 'react-hook-form';
import { ZodError, z } from 'zod'; // Import Zod
import { useAuthStore } from '@/hooks/use-auth';
import { router } from 'expo-router';

type RegisterFormData = {
    firstName: string;
    lastName: string;
    email: string;
    password: string;
};

// Define Zod schema
const registerSchema = z.object({
    firstName: z.string().min(2, { message: 'First name must be at least 2 characters long' }),
    lastName: z.string().min(2, { message: 'Last name must be at least 2 characters long' }),
    email: z.string().email({ message: 'Invalid email' }), // Email validation
    password: z.string().min(8, { message: 'Password must be at least 8 characters long' }), // Password validation
});

const Register = () => {
    const {
        control,
        handleSubmit,
        formState: { errors },
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
        <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center', padding: 20 }}>
            <View style={{ width: '100%', marginBottom: 20 }}>
                <Text>First Name</Text>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TextInput
                            style={{ padding: 10, borderColor: 'gray', borderWidth: 1 }}
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

            <View style={{ width: '100%', marginBottom: 20 }}>
                <Text>Last Name</Text>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TextInput
                            style={{ padding: 10, borderColor: 'gray', borderWidth: 1 }}
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
            <View style={{ marginTop: 20 }}>
                <Button title="Login" onPress={() => router.push('/(auth)/login')} />
            </View>
        </View>
    );
}

export default Register;