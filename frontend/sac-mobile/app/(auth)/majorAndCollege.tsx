import { Controller, useForm } from 'react-hook-form';
import { Alert, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';

import { Button } from '@/components/button';
import { DropdownComponent } from '@/components/dropdown';
import Error from '@/components/error';
import Input from '@/components/input';
import MultiSelectComponent from '@/components/multiselect';
import Wordmark from '@/components/wordmark';
import { college } from '@/lib/const';
import { major } from '@/lib/utils';
import { graduationYear } from '@/lib/utils';
import { Item } from '@/types/item';

type MajorAndCollegeForm = {
    major: Item[];
    college: Item;
    nuid: string;
    graduationYear: Item;
};

const MajorAndCollege = () => {
    const {
        control,
        handleSubmit,
        formState: { errors }
    } = useForm<MajorAndCollegeForm>();

    const majorAndCollegeSchema = z.object({
        college: z.string(),
        major: z.string().array(),
        graduationYear: z.string(),
        nuid: z.string().length(9, { message: 'NUID must have 9 digits' })
    });

    const onSubmit = ({
        major,
        college,
        nuid,
        graduationYear
    }: MajorAndCollegeForm) => {
        try {
            const updatedData = {
                major,
                nuid,
                graduationYear: graduationYear.value,
                college: college.value
            };
            majorAndCollegeSchema.parse(updatedData);
            Alert.alert('Form Submitted', JSON.stringify(updatedData));
            router.push('/(auth)/tags');
        } catch (error) {
            if (error instanceof ZodError) {
                Alert.alert('Validation Error', error.errors[0].message);
            } else {
                console.error('An unexpected error occurred:', error);
            }
        }
    };

    return (
        <SafeAreaView>
            <View className="px-[8%] pb-[9%]">
                <Wordmark />
                <Text className="font-bold text-5xl pt-[5%] pb-[5%]">
                    Let's learn more about you
                </Text>
                <View className="w-full mb-[6%]">
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
                    {errors.college && (
                        <Error message={errors.college.message} />
                    )}
                </View>
                <View className="w-full">
                    <Controller
                        control={control}
                        render={({ field: { onChange } }) => (
                            <MultiSelectComponent
                                title="Major and Minor"
                                item={major()}
                                placeholder="Select up to 4 major or minor"
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
                <View className="flex-row justify-end pt-[5%]">
                    <Button
                        size="lg"
                        variant="default"
                        onPress={handleSubmit(onSubmit)}
                    >
                        Continue
                    </Button>
                </View>
            </View>
        </SafeAreaView>
    );
};

export default MajorAndCollege;
