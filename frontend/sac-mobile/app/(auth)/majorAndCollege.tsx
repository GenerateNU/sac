import { Alert, Text, View} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Button } from '@/components/button';
import Wordmark from '@/components/wordmark';
import Error from '@/components/error';
import { DropdownComponent } from '@/components/dropdown';
import { college } from '@/lib/const';
import { major } from '@/lib/utils';
import { Controller, useForm } from 'react-hook-form';
import { router } from 'expo-router';
import { ZodError, z } from 'zod';
import { Item } from '@/types/item';

type MajorAndCollegeForm = {
    major: Item;
    college: Item; 
}

const MajorAndCollege = () => {
    const {
        control,
        handleSubmit,
        formState: { errors }
    } = useForm<MajorAndCollegeForm>();

    const majorAndCollegeSchema = z.object({
        major: z
            .string()
            .min(2, { message: 'First name must be at least 2 characters long' }),
    });

    const onSubmit = (data: MajorAndCollegeForm) => {
        try {
            const { major, college } = data;
            const updatedData = {
                major: major.value,
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
                <Wordmark additionalClasses={"justify-end"}/> 
                <Text className="font-bold text-5xl pt-[9%] pb-[10%]">
                     Let's learn more about you
                </Text>
            <View className="w-full mb-[8.5%]">
            <Controller
                            control={control}
                            render={({ field: { onChange, value } }) => (
                                <DropdownComponent
                                    title="Major"
                                    item={major()}
                                    placeholder="Select your major"
                                    onChangeText={onChange}
                                    value={value}
                                    search={true}
                                    onSubmitEditing={handleSubmit(onSubmit)}
                                    error={!!errors.college}
                                    color={'#F6F6F6'}
                                />
                            )}
                            name="major"
                            rules={{ required: 'Major is required' }}
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
                                    color={'#F6F6F6'}
                                />
                            )}
                            name="college"
                            rules={{ required: 'College is required' }}
                        />
                        {errors.college && (
                            <Error message={errors.college.message} />
                        )}
                    </View>
                    <View className="flex-row justify-end"><Button
                        size="lg"
                        variant="default"
                        onPress={handleSubmit(onSubmit)}
                    >Continue</Button></View>
            </View>
        </SafeAreaView>
    )
}

export default MajorAndCollege;