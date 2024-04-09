import { Controller, useForm } from 'react-hook-form';
import { Alert, View } from 'react-native';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';

import { Button } from '@/components/button';
import { DropdownComponent } from '@/components/dropdown';
import Error from '@/components/error';
import MultiSelectComponent from '@/components/multiselect';
import { college } from '@/lib/const';
import { graduationCycle } from '@/lib/const';
import { major } from '@/lib/utils';
import { graduationYear } from '@/lib/utils';
import { Item } from '@/types/item';
import { useUpdateUser } from '@/hooks/use-update-users';

type UserDetailsData = {
    major: Item[];
    college: Item;
    graduationYear: Item;
    graduationCycle: Item;
};

const UserDetailsForm = () => {
    const {
        control,
        handleSubmit,
        formState: { errors }
    } = useForm<UserDetailsData>();

    const UserDetailsSchema = z.object({
        college: z.string(),
        major: z.string().array(),
        graduationYear: z.string(),
        graduationCycle: z.string()
    });

    const updateUserInfo = useUpdateUser();

    const onSubmit = ({ major, college, graduationYear, graduationCycle }: UserDetailsData) => {
        try {
            const updatedData = {
                major,
                graduationYear: graduationYear.value,
                college: college.value,
                graduationCycle: graduationCycle.value, 
            };
            UserDetailsSchema.parse(updatedData);
            Alert.alert('Form Submitted', JSON.stringify(updatedData));
            router.push('/(auth)/user-interests');
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
            <View className="mb-[6%]">
                <Controller
                    control={control}
                    render={({ field: { onChange } }) => (
                        <MultiSelectComponent
                            title="Major and Minor"
                            item={major()}
                            placeholder="Select up to 4 major & minor"
                            search={true}
                            onSubmitEditing={handleSubmit(onSubmit)}
                            error={!!errors.major}
                            maxSelect={4}
                            onChange={(selectedItems) => {
                                onChange(selectedItems);
                            }}
                        />
                    )}
                    name="major"
                    rules={{
                        required: 'Major is required'
                    }}
                />
                {errors.major && <Error message={errors.major.message} />}
            </View>
            <View className="mb-[6%]">
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <DropdownComponent
                            title="College"
                            item={college}
                            placeholder="Select your college"
                            onChangeText={onChange}
                            value={value}
                            onSubmitEditing={handleSubmit(onSubmit)}
                            error={!!errors.college}
                        />
                    )}
                    name="college"
                    rules={{ required: 'College is required' }}
                />
                {errors.college && <Error message={errors.college.message} />}
            </View>
            <View className="mb-[6%]">
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
            <View className="mb-[6%]">
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <DropdownComponent
                            title="Intended Graduation Cycle"
                            item={graduationCycle}
                            placeholder="Select Year"
                            onChangeText={onChange}
                            value={value}
                            onSubmitEditing={handleSubmit(onSubmit)}
                            error={!!errors.graduationCycle}
                        />
                    )}
                    name="graduationCycle"
                    rules={{ required: 'Graduation cycle is required' }}
                />
                {errors.graduationCycle && (
                    <Error message={errors.graduationCycle.message} />
                )}
            </View>
            <View className="flex-row justify-end pt-[5%]">
                <Button
                    onPress={handleSubmit(onSubmit)}
                    size="lg"
                    variant="default"
                >
                    Continue
                </Button>
            </View>
        </>
    );
};

export default UserDetailsForm;
