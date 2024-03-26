import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { Alert, ScrollView, Text, View } from 'react-native';

import { router } from 'expo-router';

import { ZodError } from 'zod';

import { Button } from '@/components/button';
import Error from '@/components/error';
import {useCategories} from '@/hooks/use-categories';
import {useTags} from '@/hooks/use-tags';
import {Category} from '@/types/category';
import { Tag } from '@/types/tag';

type UserInterestsData = {
    tags: Tag[];
};


const UserInterestsForm = () => {
    const { data: categoriesData, isLoading: loadingCategories, error: errorCategory } = useCategories() as { data: Category[], isLoading: boolean, error: Error };
    const { data: tagsData, isLoading: loadingTags, error: errorTag } = useTags() as { data: Tag[], isLoading: boolean, error: Error };
    const { handleSubmit } = useForm<UserInterestsData>();
    const AllCategory: Category = {
        name: 'All',
        id: '',
        createdAt: new Date(),
        updatedAt: new Date(),
        tags: tagsData?.length ? tagsData : [],
    }
    const [selectedTags, setSelectedTags] = useState<Tag[]>([]);
    const [buttonClicked, setButtonClicked] = useState<boolean>(false);
    const [selectedCategory, setSelectedCategory] = useState<Category>(AllCategory);

    // when a category is selected, set selected category to this category
    const handleCategoryPress = (category: Category) => {
        setSelectedCategory(category);
    };

    const handleTagPress = (tag: Tag) => {
        // if a tag is selected and tag has been selected before, filter it out
        if (selectedTags.includes(tag)) {
            setSelectedTags(selectedTags.filter((t) => t !== tag));
        }
        // appended into selectedTags list
        else {
            setSelectedTags([...selectedTags, tag]);
        }
    };

    const onSubmit = (data: UserInterestsData) => {
        setButtonClicked(true);
        if (selectedTags.length === 0) {
            return;
        }
        data.tags = selectedTags;
        const selectedTagsUUID = data.tags.map(tag => tag.id);
        try {
            Alert.alert('Form Submitted', JSON.stringify(selectedTagsUUID));
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

    const Tag = ({tag}: { tag: Tag }) => {
        return (
            <Button
                onPress={() => handleTagPress(tag)}
                variant={
                    selectedTags.includes(tag) ? 'default' : 'outline'
                }
                size="tags"
            >
                {tag.name}
            </Button>
        );
    };

    return (
        <>
            <ScrollView horizontal showsHorizontalScrollIndicator={false}>
                <Button
                    onPress={() => handleCategoryPress(AllCategory)}
                    variant={
                    'All' === selectedCategory.name
                        ? 'underline'
                        : 'menu'
                    }
                >All</Button>
                {categoriesData && categoriesData.map((category: Category, key) => (
                    <View>
                        <Button
                            onPress={() => handleCategoryPress(category)}
                            variant={
                                category.name === selectedCategory.name
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
                    {selectedCategory.name === AllCategory.name
                        ? AllCategory.tags.map((tag, key) => (
                            <Tag tag={tag} key={key} />
                        ))
                        : tagsData && tagsData.filter((tag) => selectedCategory.id === tag.category_id)
                        .map((tag, key) => <Tag tag={tag} key={key} />)
                    }
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

export default UserInterestsForm;
