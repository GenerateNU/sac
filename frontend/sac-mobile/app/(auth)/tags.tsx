import { Alert, Text, View, SafeAreaView, ScrollView } from 'react-native';
import { Button } from '@/components/button';
import Wordmark from '@/components/wordmark';
import React, { useState } from 'react';
import { ZodError, z } from 'zod';
import { Controller, useForm } from 'react-hook-form';
import Error from '@/components/error';
import { router } from 'expo-router';

type TagsData = {
    tags: String[];
}

const listOfTags = ["Pre-med","Pre-law","Judaism","Christianity","Hinduism","Islam","Latin America","African American","Asian American","LGBTQ","Performing Arts","Visual Arts","Creative Writing","Music","Soccer","Hiking","Climbing","Lacrosse","Mathematics","Physics","Biology","Chemistry","Environmental Science","Geology","Neuroscience","Psychology","Software Engineering","Artificial Intelligence","Data Science","Mechanical Engineering","Electrical Engineering","Industrial Engineering","Volunteerism","Environmental Advocacy","Human Rights","Community Outreach","Journalism","Broadcasting","Film","Public Relations","Other"]

const Tags = () => {
    const [selectedTags, setSelectedTags] = useState<String[]>([]);
    const [buttonClicked, setButtonClicked] = useState<boolean>(false);

    const handleTagClick = (tag: string) => {
        let updatedTags;
        // deselect tag: if list of tags already have the tag, filter that tag out from the list
        if (selectedTags.includes(tag)) {
            updatedTags = selectedTags.filter((selectedTag) => selectedTag !== tag);
        }
        // select tag: add it into the end of list
        else {
            updatedTags = [...selectedTags, tag];
        }
        setSelectedTags(updatedTags);
    };

    const {
        control,
        handleSubmit,
    } = useForm<TagsData>();

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

    return (
        <SafeAreaView>
            <View className="px-[8%] pt-[4%]">
                <View className="pb-[3%]"><Wordmark/></View>
                <Text className="text-5xl font-bold">What are you interested in?</Text>
                <Text className="text-xl pt-[3%] pb-[4%]">Select one or more</Text>
                <ScrollView className="h-[62%] pt-[3%]">
                    <View className="flex-row flex-wrap mb-[3%]">
                    {listOfTags.map((text, index) => (
                    <Controller
                        control={control}
                        render={() => (
                            <Button
                                key={index}
                                variant={selectedTags.includes(text) ? "default" : "outline"}
                                size="tags"
                                onPress={() => handleTagClick(text)}>
                            {text}
                            </Button>
                        )}
                        name="tags"
                    /> ))}
                    </View>
                </ScrollView>
                {selectedTags.length === 0 && buttonClicked && (
                    <Error message="Please choose at least one interest" />
                )}
                <View className="flex-row justify-end pt-[5%]">
                    <Button
                        size="lg"
                        onPress={handleSubmit(onSubmit)}
                    >Finish</Button>
                </View>
            </View>
        </SafeAreaView>
    );
}

export default Tags;

