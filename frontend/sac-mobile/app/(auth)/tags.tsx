import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { Alert, SafeAreaView, ScrollView, Text, View } from 'react-native';

import { router } from 'expo-router';

import { ZodError } from 'zod';

import { Button } from '@/components/button';
import Error from '@/components/error';
import { tags } from '@/lib/const';
import Wordmark from '@/components/wordmark';

type TagsData = {
    tags: String[];
};

const Tags = () => {
    const [selectedTags, setSelectedTags] = useState<String[]>([]);
    const [buttonClicked, setButtonClicked] = useState<boolean>(false);

    const handleTagClick = (tag: string) => {
        let updatedTags;
        // deselect tag: if list of tags already have the tag, filter that tag out from the list
        if (selectedTags.includes(tag)) {
            updatedTags = selectedTags.filter(
                (selectedTag) => selectedTag !== tag
            );
        }
        // select tag: add it into the end of list
        else {
            updatedTags = [...selectedTags, tag];
        }
        setSelectedTags(updatedTags);
    };

    const { handleSubmit } = useForm<TagsData>();

    const onSubmit = (data: TagsData) => {
        setButtonClicked(true);
        if (selectedTags.length === 0) {
            return;
        }
        data.tags = selectedTags;
        try {
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

    const emptyTag = selectedTags.length === 0 && buttonClicked ? 'h-[53%]' : 'h-[54%]';

    return (
        <SafeAreaView>
            <View className="px-[8%] pt-[4%]">
                <View className="flex flex-row">
                    <Wordmark />
                </View>
                <Text className="text-5xl pt-[6%] font-bold">
                    What are you interested in?
                </Text>
                <Text className="text-xl pt-[6%] pb-[7%]">
                    Select one or more
                </Text>
                <ScrollView className={emptyTag}>
                    <View className="flex-row flex-wrap">
                        {tags.map((text, index) => (
                            <Button
                                key={index}
                                variant={
                                    selectedTags.includes(text)
                                        ? 'default'
                                        : 'outline'
                                }
                                size="tags"
                                onPress={() => handleTagClick(text)}
                            >
                                {text}
                            </Button>
                        ))}
                    </View>
                </ScrollView>
                {selectedTags.length === 0 && buttonClicked && (
                    <View className="pt-3">
                        <Error message="Please choose at least one interest" />
                    </View>
                )}
                <View className="flex-row justify-end mt-[7%]">
                    <Button size="lg" onPress={handleSubmit(onSubmit)}>
                        Finish
                    </Button>
                </View>
            </View>
        </SafeAreaView>
    );
};

export default Tags;
