import { Controller, useForm } from 'react-hook-form';
import { Alert, View } from 'react-native';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';

import { Button } from '@/components/button';
import { DropdownComponent } from '@/components/dropdown';
import Error from '@/components/error';
import Input from '@/components/input';
import { graduationYear } from '@/lib/utils';
import { Item } from '@/types/item';

type RegisterFormData = {
    firstName: string;
    lastName: string;
    email: string;
    nuid: string;
    graduationYear: Item;
    password: string;
    passwordConfirm: string;
};

const registerSchema = z
    .object({
        firstName: z.string().min(2, {
            message: 'First name must be at least 2 characters long'
        }),
        lastName: z.string().min(2, {
            message: 'Last name must be at least 2 characters long'
        }),
        email: z.string().email({ message: 'Invalid email' }),
        nuid: z.string().length(9, { message: 'NUID must have 9 digits' }),
        password: z
            .string()
            .min(8, { message: 'Password must be at least 8 characters long' }),
        passwordConfirm: z.string(),
        graduationYear: z.object({
            label: z.string(),
            value: z.string()
        })
    })
    .refine((data) => data.password === data.passwordConfirm, {
        message: 'Passwords do not match',
        path: ['passwordConfirm']
    });

const RegistrationForm = () => {
    const {
        control,
        handleSubmit,
        formState: { errors }
    } = useForm<RegisterFormData>();

    const onSubmit = ({
        graduationYear: gradYear,
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        passwordConfirm,
        ...rest
    }: RegisterFormData) => {
        try {
            const updatedData = {
                ...rest,
                graduationYear: gradYear.value
            };
            registerSchema.parse(updatedData);
            Alert.alert('Form Submitted', JSON.stringify(updatedData));
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
        <>
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
                {errors.lastName && <Error message={errors.lastName.message} />}
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
                {errors.email && <Error message={errors.email.message} />}
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
                            if (!/^00\d+/.test(value)) {
                                return 'Please enter a proper NUID number';
                            }
                            if (value.length !== 9) {
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
                            const password = control._getWatch('password');
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
                <Button onPress={handleSubmit(onSubmit)} size="screenwide">
                    Submit
                </Button>
            </View>
        </>
    );
};

export default RegistrationForm;
