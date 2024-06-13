import { Controller, useForm } from 'react-hook-form';
import { Alert } from 'react-native';
import { ScrollView } from 'react-native-gesture-handler';
import { SafeAreaView } from 'react-native-safe-area-context';

import { Stack, router } from 'expo-router';

import { ZodError } from 'zod';

import { Box, Spacing } from '@/src/app/(design-system)';
import { MultiSelect, SelectOne } from '@/src/app/(design-system)';
import { Text } from '@/src/app/(design-system)';
import { Arrow } from '@/src/app/(design-system)';
import { TextBox } from '@/src/app/(design-system)/components/Textbox/Textbox';
import {
    COLLEGE,
    GRADUATION_CYCLE,
    MAJORS,
    graduationYear,
    startYear
} from '@/src/consts/user-info';
import { UserInfoData } from '@/src/types/userInfo';
import { convertUserData } from '@/src/utils/userInfo';

import { Save } from './components/save';

const mockUserData = {
    name: 'Jane Doe',
    major0: 'economics',
    major1: 'biology',
    major2: null,
    college: 'KCCS',
    startYear: 2023,
    graduationYear: null,
    graduationCycle: null,
    email: 'doe.j@northeastern.edu'
};

const UserDetails = () => {
    const defaultValues: Partial<UserInfoData> = convertUserData(mockUserData);

    const {
        control,
        handleSubmit,
        trigger,
        formState: { errors, isDirty }
    } = useForm<UserInfoData>({
        mode: 'onTouched',
        defaultValues
    });

    const onSubmit = (data: UserInfoData) => {
        try {
            console.log(data);
            router.back();
        } catch (error) {
            if (error instanceof ZodError) {
                Alert.alert('Validation Error', error.errors[0].message);
            } else {
                console.error('An unexpected error occurred:', error);
            }
        }
    };

    console.log(defaultValues);

    return (
        <>
            <Stack.Screen
                options={{
                    headerTitle: () => (
                        <Text variant="subheader-2">Profile</Text>
                    ),
                    headerTransparent: true,
                    headerShown: true,
                    headerLeft: () => <Arrow color="black" />,
                    headerRight: () => (
                        <Save
                            onPress={handleSubmit(onSubmit)}
                            isValid={isDirty}
                        />
                    )
                }}
            />
            <SafeAreaView
                edges={['top']}
                style={{ flex: 1, marginTop: Spacing.xxl }}
            >
                <ScrollView
                    showsVerticalScrollIndicator={false}
                    contentContainerStyle={{ paddingVertical: 40 }}
                >
                    <Box
                        paddingHorizontal="xl"
                        justifyContent="center"
                        alignItems="flex-start"
                        flexDirection="column"
                        gap="xs"
                    >
                        <Controller
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <TextBox
                                    readOnly={true}
                                    onChangeText={(text: string) => {
                                        onChange(text);
                                        trigger('name');
                                    }}
                                    value={value}
                                    error={
                                        errors.name ? errors.name.message : ''
                                    }
                                    placeholder="Enter your full name"
                                    title="Full Name"
                                />
                            )}
                            name="name"
                        />
                        <Controller
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <TextBox
                                    readOnly={true}
                                    onChangeText={(text: string) => {
                                        onChange(text);
                                        trigger('email');
                                    }}
                                    error={
                                        errors.email ? errors.email.message : ''
                                    }
                                    value={value}
                                    title="Email"
                                />
                            )}
                            name="email"
                        />
                        <Controller
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <MultiSelect
                                    onChange={onChange}
                                    title="Major"
                                    data={MAJORS}
                                    limit={3}
                                    value={value || []}
                                    placeholder="Select up to 3 majors"
                                />
                            )}
                            name="major"
                        />
                        <Controller
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <SelectOne
                                    value={value}
                                    onChange={onChange}
                                    placeholder="Select home college"
                                    title="College"
                                    data={COLLEGE}
                                />
                            )}
                            name="college"
                        />
                        <Controller
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <SelectOne
                                    value={value}
                                    onChange={onChange}
                                    placeholder="Select graduation cycle"
                                    title="Graduation Cycle"
                                    data={GRADUATION_CYCLE}
                                />
                            )}
                            name="graduationCycle"
                        />
                        <Controller
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <SelectOne
                                    value={value}
                                    onChange={onChange}
                                    placeholder="Select start year"
                                    title="Start Year"
                                    data={startYear()}
                                />
                            )}
                            name="startYear"
                        />
                        <Controller
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <SelectOne
                                    value={value}
                                    onChange={onChange}
                                    placeholder="Select intended graduation year"
                                    title="Intended Graduation Year"
                                    data={graduationYear()}
                                />
                            )}
                            name="graduationYear"
                        />
                    </Box>
                </ScrollView>
            </SafeAreaView>
        </>
    );
};

export default UserDetails;
