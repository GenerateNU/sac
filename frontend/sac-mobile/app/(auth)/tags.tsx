import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { Alert, SafeAreaView, ScrollView, Text, View } from 'react-native';

import { router } from 'expo-router';

import { ZodError } from 'zod';

import { Button } from '@/components/button';
import Error from '@/components/error';
import Wordmark from '@/components/wordmark';
import { categories } from '@/lib/const';
import { Category } from '@/types/categories';

import TagForm from './_components/tag-form';

const Tags = () => {
    return (
        <SafeAreaView>
            <View className="px-[8%] pt-[4%]">
                <View className="flex flex-row">
                    <Wordmark />
                </View>
                <Text className="text-5xl pt-[6%] pb-[5%] font-bold">
                    What are you interested in?
                </Text>
                <TagForm />
            </View>
        </SafeAreaView>
    );
};

export default Tags;
