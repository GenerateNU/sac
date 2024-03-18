import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { Alert, SafeAreaView, ScrollView, Text, View } from 'react-native';

import { router } from 'expo-router';

import { ZodError } from 'zod';

import { Button } from '@/components/button';
import Error from '@/components/error';
import Wordmark from '@/components/wordmark';
import { categories } from '@/lib/const';
import { Categories } from '@/types/categories';

type TagsData = {
    tags: String[];
};

let allTags: string[] = [];
for (let i = 0; i < categories.length; i++) {
    allTags = allTags.concat(categories[i].tags);
}
const categoriesMenu = [{ name: 'All', tags: allTags }, ...categories];

const Tags = () => {
    const [selectedTags, setSelectedTags] = useState<String[]>([]);
    const [buttonClicked, setButtonClicked] = useState<boolean>(false);
    const [selectedCategory, setSelectedCategory] = useState('All');

    const handleCategoryPress = (category: Categories) => {
        setSelectedCategory(category.name);
    };

    const handleTagPress = (tag: string) => {
        if (selectedTags.includes(tag)) {
            setSelectedTags(selectedTags.filter((t) => t !== tag));
        } else {
            setSelectedTags([...selectedTags, tag]);
        }
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

    const heightAdjust =
        selectedTags.length === 0 && buttonClicked ? 'h-[46%]' : 'h-[47%]';

    return (
        <SafeAreaView>
            <View className="px-[8%] pt-[4%]">
                <View className="flex flex-row">
                    <Wordmark />
                </View>
                <Text className="text-5xl pt-[6%] pb-[5%] font-bold">
                    What are you interested in?
                </Text>
                <ScrollView horizontal showsHorizontalScrollIndicator={false}>
                    {categoriesMenu.map((category) => (
                        <View>
                            <Button
                                onPress={() => handleCategoryPress(category)}
                                variant={
                                    category.name === selectedCategory
                                        ? 'underline'
                                        : 'menu'
                                }
                                size="menu"
                            >
                                {category.name}
                            </Button>
                        </View>
                    ))}
                </ScrollView>

                <Text className="text-lg pb-[6%] pt-[5%]">
                    Select one or more
                </Text>
                <ScrollView className={heightAdjust}>
                    <View className="flex-row flex-wrap">
                        {selectedCategory === 'All'
                            ? allTags.map((tag) => (
                                  <Button
                                      onPress={() => handleTagPress(tag)}
                                      variant={
                                          selectedTags.includes(tag)
                                              ? 'default'
                                              : 'outline'
                                      }
                                      size="tags"
                                  >
                                      {tag}
                                  </Button>
                              ))
                            : categories
                                  .find((c) => c.name === selectedCategory)
                                  ?.tags.map((tag) => (
                                      <Button
                                          onPress={() => handleTagPress(tag)}
                                          variant={
                                              selectedTags.includes(tag)
                                                  ? 'default'
                                                  : 'outline'
                                          }
                                          size="tags"
                                      >
                                          {tag}
                                      </Button>
                                  ))}
                    </View>
                </ScrollView>
                {selectedTags.length === 0 && buttonClicked && (
                    <View className="pt-3">
                        <Error message="Please choose at least one interest" />
                    </View>
                )}
                <View className="flex-row justify-end mt-[8%]">
                    <Button size="lg" onPress={handleSubmit(onSubmit)}>
                        Finish
                    </Button>
                </View>
            </View>
        </SafeAreaView>
    );
};

export default Tags;
