import { useState } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert, Keyboard, TouchableWithoutFeedback, View } from 'react-native';
import Spinner from 'react-native-loading-spinner-overlay';

import { ZodError, z } from 'zod';

import { Button } from '@/components/button';
import Error from '@/components/error';
import Input from '@/components/input';
import { useAuthStore } from '@/hooks/use-auth';

type RegisterFormData = {
    first_name: string;
    last_name: string;
    email: string;
    password: string;
    password_confirmation: string;
};

const registerSchema = z.object({
    first_name: z.string().min(2, {
        message: 'First name must be at least 2 characters long'
    }),
    last_name: z.string().min(2, {
        message: 'Last name must be at least 2 characters long'
    }),
    email: z.string().email({ message: 'Invalid email' }),
    password: z
        .string()
        .min(8, { message: 'Password must be at least 8 characters long' }),
    password_confirmation: z.string()
});

const RegistrationForm = () => {
    const {
        control,
        handleSubmit,
        formState: { errors }
    } = useForm<RegisterFormData>();

    const { signUp } = useAuthStore();
    const [loading, setLoading] = useState(false);

    const onSignUp = async ({
        password_confirmation,
        ...rest
    }: RegisterFormData) => {
        setLoading(true);

        try {
            registerSchema.parse({ password_confirmation, ...rest });
            const updatedData = { ...rest };
            console.log(updatedData);
            await signUp(updatedData);
        } catch (error: any) {
            if (error instanceof ZodError) {
                Alert.alert(error.errors[0].message);
            } else {
                Alert.alert('An error occurred', error.message);
            }
        } finally {
            setLoading(false);
        }
    };

    return (
        <>
            {loading && <Spinner visible={loading} />}
            <View className="w-full mb-[8.5%]">
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
                            <Input
                                title="First Name"
                                autoCorrect={false}
                                placeholder="Garrett"
                                onChangeText={(text) => onChange(text)}
                                value={value}
                                onSubmitEditing={handleSubmit(onSignUp)}
                                error={!!errors.first_name}
                            />
                        </TouchableWithoutFeedback>
                    )}
                    name="first_name"
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
                {errors.first_name && (
                    <Error message={errors.first_name.message} />
                )}
            </View>

            <View className="w-full mb-[8.5%]">
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
                            <Input
                                title="Last Name"
                                autoCorrect={false}
                                placeholder="Ladley"
                                onChangeText={onChange}
                                value={value}
                                onSubmitEditing={handleSubmit(onSignUp)}
                                error={!!errors.last_name}
                            />
                        </TouchableWithoutFeedback>
                    )}
                    name="last_name"
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
                {errors.last_name && (
                    <Error message={errors.last_name.message} />
                )}
            </View>

            <View className="w-full mb-[8.5%]">
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
                            <Input
                                title="Email"
                                autoCorrect={false}
                                placeholder="doe.j@northeastern.edu"
                                onChangeText={onChange}
                                value={value}
                                onSubmitEditing={handleSubmit(onSignUp)}
                                error={!!errors.email}
                            />
                        </TouchableWithoutFeedback>
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
                {errors.email && <Error message={errors.email.message} />}
            </View>

            <View className="w-full mb-[8.5%]">
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
                            <Input
                                title="Password"
                                autoCorrect={false}
                                placeholder="Password"
                                onChangeText={onChange}
                                value={value}
                                onSubmitEditing={handleSubmit(onSignUp)}
                                secureTextEntry={true}
                                error={!!errors.password}
                            />
                        </TouchableWithoutFeedback>
                    )}
                    name="password"
                    rules={{
                        required: 'Password is required',
                        validate: (value) => {
                            const validationErrors = [];

                            if (value.length < 8) {
                                validationErrors.push(
                                    'Password must be at least 8 characters long.'
                                );
                            }

                            const specialCharRegex =
                                /[!@#$%^&*()_+\-=;:'",<>.]/;
                            if (!specialCharRegex.test(value)) {
                                validationErrors.push(
                                    'Password must contain at least one special character.'
                                );
                            }

                            if (!/\d/.test(value)) {
                                validationErrors.push(
                                    'Password must contain at least one number (0-9).'
                                );
                            }

                            return validationErrors.length > 0
                                ? validationErrors.join(' ')
                                : true;
                        }
                    }}
                />
                {errors.password && <Error message={errors.password.message} />}
            </View>

            <View className="w-full mb-[8.5%]">
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
                            <Input
                                title="Password Confirmation"
                                autoCorrect={false}
                                placeholder="Confirm your password"
                                onChangeText={onChange}
                                value={value}
                                onSubmitEditing={handleSubmit(onSignUp)}
                                secureTextEntry={true}
                                error={!!errors.password}
                            />
                        </TouchableWithoutFeedback>
                    )}
                    name="password_confirmation"
                    rules={{
                        required: 'Please confirm your password',
                        validate: (value) => {
                            const password = control._getWatch('password');
                            if (value !== password) {
                                return 'Passwords do not match';
                            }
                            return true;
                        }
                    }}
                />
                {errors.password_confirmation && (
                    <Error message={errors.password_confirmation.message} />
                )}
            </View>
            <View className="pt-[2%] pb-[15%]">
                <Button onPress={handleSubmit(onSignUp)} size="screenwide">
                    Submit
                </Button>
            </View>
        </>
    );
};

export default RegistrationForm;
