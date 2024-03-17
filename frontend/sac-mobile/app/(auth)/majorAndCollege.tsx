import { Controller, useForm } from 'react-hook-form';
import { Alert, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { router } from 'expo-router';

import { ZodError, z } from 'zod';
import { useState } from 'react';
import { Button } from '@/components/button';
import { DropdownComponent } from '@/components/dropdown';
import MultiSelectComponent from '@/components/multiselect';

import Error from '@/components/error';
import Wordmark from '@/components/wordmark';
import { college } from '@/lib/const';
import { major } from '@/lib/utils';
import { Item } from '@/types/item';

type MajorAndCollegeForm = {
    major: Item[];
    college: Item;
};

const MajorAndCollege = () => {
    const {
        control,
        handleSubmit,
        formState: {errors}
    } = useForm<MajorAndCollegeForm>();

    const majorAndCollegeSchema = z.object({
        college: z.string(),
        major: z.string().array()
    });

    const onSubmit = ({major, college}: MajorAndCollegeForm) => {
        try {
            const updatedData = {
                major,
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
                <Text className="font-bold text-5xl pt-[9%] pb-[10%]">
                    Let's learn more about you
                </Text>
                <View className="w-full mb-[8.5%]">
                    <Controller
                        control={control}
                        render={({ field: { onChange, value } }) => (
                        <MultiSelectComponent
                            title="Major and Minor"
                            item={major()}
                            placeholder="Select up to 4 major or minor"
                            search={true}
                            onSubmitEditing={handleSubmit(onSubmit)}
                            error={!!errors.major}
                            onChange={onChange}
                            value={value}
                        />
                        )}
                        name="major"
                        rules={{
                            required: 'Major is required',
                        }}
                    />
                    {errors.major && <Error message={errors.major.message} />}
                </View>
                <View className="mb-[7%]">
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
                <View className="flex-row justify-end pt-[5%]">
                    <Button
                        size="lg"
                        variant="default"
                        onPress={handleSubmit(onSubmit)}
                    >Continue
                    </Button>
                </View>
            </View>
        </SafeAreaView>
    );
};

export default MajorAndCollege;
