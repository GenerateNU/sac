import React from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert, Pressable, Text, View } from 'react-native';

import { ZodError, z } from 'zod';

import { Error } from '@/components/error';
import { Input } from '@/components/input';
import { FAQ } from '@/types/faq';

type FAQData = {
    question: string;
};

const FAQSchema = z.object({
    question: z.string()
});

const FAQHomePageCard = ({ faq }: { faq: FAQ }) => {
    const length = () => {
        if (faq.club_name.length > 11) {
            return faq.club_name.substring(0, 11) + '...';
        } else {
            return faq.club_name;
        }
    };

    const {
        control,
        handleSubmit,
        formState: { errors },
        reset
    } = useForm<FAQData>();

    const onSubmit = ({ question }: FAQData) => {
        try {
            FAQSchema.parse({ question });
            Alert.alert('Form Submitted', JSON.stringify(question));
            reset();
        } catch (error) {
            if (error instanceof ZodError) {
                Alert.alert('Validation Error', error.errors[0].message);
            } else {
                console.error('An unexpected error occurred:', error);
            }
        }
    };

    return (
        <Pressable className="bg-gray-200 rounded-2xl my-[2%] p-[5%]">
            <View className="flex-row">
                <View className="w-16 h-16 bg-gray-300 rounded-xl" />
                <View className="ml-[5%]">
                    <Text className="text-base font-bold leading-6 ">
                        {faq.club_name}
                    </Text>
                    <Text className="text-sm font-medium leading-5 text-gray-500">
                        Frequently Asked
                    </Text>
                    <Text className="text-sm font-medium leading-5 text-gray-500">
                        Questions
                    </Text>
                </View>
            </View>
            <View className="mt-[7%]">
                <Text className="text-base font-bold">Question:</Text>
                <Text>{faq.question}</Text>
                <Text className="text-base font-bold mt-[4%]">Answer:</Text>
                <Text numberOfLines={2} ellipsizeMode="tail">
                    {faq.answer}
                </Text>
                <View className="mt-[6%] mb-[1.5%]">
                    <Controller
                        control={control}
                        render={({ field: { onChange, value } }) => (
                            <Input
                                variant="faq"
                                placeholder={
                                    'Submit a question for ' + length()
                                }
                                onSubmitEditing={handleSubmit(onSubmit)}
                                autoCorrect={false}
                                onChangeText={onChange}
                                value={value}
                            />
                        )}
                        name="question"
                        rules={{
                            required: 'Question is required'
                        }}
                    />
                    {errors.question && (
                        <View className="mt-[2%]">
                            <Error message={errors.question.message} />
                        </View>
                    )}
                </View>
            </View>
        </Pressable>
    );
};

export { FAQHomePageCard };
