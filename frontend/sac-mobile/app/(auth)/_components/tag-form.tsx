import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { Alert, SafeAreaView, ScrollView, Text, View } from 'react-native';

import { router } from 'expo-router';

import { ZodError } from 'zod';

import { Button } from '@/components/button';
import Error from '@/components/error';
import { categories } from '@/lib/const';
import { Category } from '@/types/categories';

type TagsData = {
    tags: String[];
};

// combine all tags of different categories into one for All tab
let allTags: string[] = [];
for (let i = 0; i < categories.length; i++) {
    allTags = allTags.concat(categories[i].tags);
}
const categoriesMenu = [{ name: 'All', tags: allTags }, ...categories];

const TagForm = () => {
    const { handleSubmit } = useForm<TagsData>();
    const [selectedTags, setSelectedTags] = useState<String[]>([]);
    const [buttonClicked, setButtonClicked] = useState<boolean>(false);
    const [selectedCategory, setSelectedCategory] = useState('All');

    // when a category is selected, set selected category to the category's name
    const handleCategoryPress = (category: Category) => {
        setSelectedCategory(category.name);
    };

    const handleTagPress = (tag: string) => {
        // if a tag is selected and tag has been selected before, filter it out
        if (selectedTags.includes(tag)) {
            setSelectedTags(selectedTags.filter((t) => t !== tag));
        }
        // appended into selectedTags list
        else {
            setSelectedTags([...selectedTags, tag]);
        }
    };

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

    const Tag = (props: { tag: string;}) => {
        return (
            <Button
                onPress={() => handleTagPress(props.tag)}
                variant={
                    selectedTags.includes(props.tag) ? 'default' : 'outline'
                }
                size="tags"
            >
                {props.tag}
            </Button>
        );
    };

    return (
        <>
            <ScrollView horizontal showsHorizontalScrollIndicator={false}>
                {categoriesMenu.map((category, key) => (
                    <View>
                        <Button
                            onPress={() => handleCategoryPress(category)}
                            variant={
                                category.name === selectedCategory
                                    ? 'underline'
                                    : 'menu'
                            }
                            key={key}
                            size="menu"
                        >
                            {category.name}
                        </Button>
                    </View>
                ))}
            </ScrollView>

            <Text className="text-lg pb-[6%] pt-[5%]">Select one or more</Text>
            <ScrollView className={heightAdjust}>
                <View className="flex-row flex-wrap">
                    {selectedCategory === 'All'
                        ? allTags.map((tag, key) => <Tag tag={tag} key={key} />)
                        : categories
                              .find((c) => c.name === selectedCategory)
                              ?.tags.map((tag, key) => (
                                  <Tag tag={tag} key={key} />
                              ))}
                </View>
            </ScrollView>
            {selectedTags.length === 0 && buttonClicked && (
                <View className="pt-[2%]">
                    <Error message="Please choose at least one interest" />
                </View>
            )}
            <View className="flex-row justify-end mt-[8%]">
                <Button size="lg" onPress={handleSubmit(onSubmit)}>
                    Finish
                </Button>
            </View>
        </>
    );
};

export default TagForm;
