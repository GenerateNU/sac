import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { Alert, FlatList, Text, View } from 'react-native';

import { router } from 'expo-router';

import { ZodError } from 'zod';

import { Button } from '@/components/button';
import { Error } from '@/components/error';
import { useAuthStore } from '@/hooks/use-auth';
import { useCategories } from '@/hooks/use-categories';
import { useTags } from '@/hooks/use-tags';
import { useCreateUserTags } from '@/hooks/use-user';
import { Category } from '@/types/category';
import { Tag } from '@/types/tag';

type UserInterestsData = {
    tags: Tag[];
};

const UserInterestsForm = () => {
    const { user } = useAuthStore();
    const userId: string = user?.id!;

    const {
        data: categoriesData,
        isLoading: loadingCategories,
        error: errorCategory
    } = useCategories();

    const {
        data: tagsData,
        isLoading: loadingTags,
        error: errorTag
    } = useTags();

    const { handleSubmit } = useForm<UserInterestsData>();

    const AllCategory: Category = {
        name: 'All',
        id: '',
        created_at: new Date(),
        updated_at: new Date(),
        tags: tagsData?.length ? tagsData : []
    };

    const CategoriesData = categoriesData
        ? [AllCategory, ...categoriesData]
        : [];

    const [selectedTags, setSelectedTags] = useState<Tag[]>([]);
    const [buttonClicked, setButtonClicked] = useState<boolean>(false);
    const [selectedCategory, setSelectedCategory] =
        useState<Category>(AllCategory);
    const createUserTags = useCreateUserTags();

    // when a category is selected, set selected category to this category
    const handleCategoryPress = (category: Category) => {
        setSelectedCategory(category);
    };

    const handleTagPress = (tag: Tag) => {
        if (selectedTags.includes(tag)) {
            setSelectedTags(selectedTags.filter((t) => t !== tag));
        } else {
            setSelectedTags([...selectedTags, tag]);
        }
    };

    const onSubmit = (data: UserInterestsData) => {
        setButtonClicked(true);
        if (selectedTags.length === 0) {
            return;
        }
        data.tags = selectedTags;
        const selectedTagsUUID = data.tags.map((tag) => tag.id);
        createUserTags.mutateAsync({
            userId: userId,
            tagIDs: selectedTagsUUID
        });
        try {
            Alert.alert('Form Submitted', JSON.stringify(selectedTagsUUID));
            router.push('/(app)/(tabs)/');
        } catch (error) {
            if (error instanceof ZodError) {
                Alert.alert('Validation Error', error.errors[0].message);
            } else {
                console.error('An unexpected error occurred:', error);
            }
        }
    };

    if (loadingCategories || loadingTags) {
        return <Text>Loading...</Text>;
    }

    if (errorCategory || errorTag) {
        return <Text>Error</Text>;
    }

    const heightAdjust =
        selectedTags.length === 0 && buttonClicked ? 'h-[48%]' : 'h-[50%]';

    const Tag = ({ tag }: { tag: Tag }) => {
        return (
            <Button
                onPress={() => handleTagPress(tag)}
                variant={selectedTags.includes(tag) ? 'default' : 'outline'}
                className="h-12 rounded-3xl px-[5%] mr-[4%] w-fit mb-[4%]"
            >
                {tag.name}
            </Button>
        );
    };

    const renderCategory = ({ item }: { item: Category }) => {
        return (
            <View>
                <Button
                    onPress={() => handleCategoryPress(item)}
                    variant={
                        item.name === selectedCategory.name
                            ? 'underline'
                            : 'menu'
                    }
                    key={item.id}
                    className={`h-12 px-5`}
                >
                    {item.name}
                </Button>
            </View>
        );
    };

    const renderTag = ({ item }: { item: Tag }) => {
        return (
            <View className="m-1">
                <Tag tag={item} key={item.id} />
            </View>
        );
    };

    return (
        <>
            <FlatList
                horizontal
                showsHorizontalScrollIndicator={false}
                data={CategoriesData}
                renderItem={renderCategory}
                keyExtractor={(item) => item.id.toString()}
            />
            <Text className="text-lg pb-[4%]">Select one or more</Text>
            <View className={heightAdjust}>
                {selectedCategory.name === AllCategory.name ? (
                    <FlatList
                        data={AllCategory.tags}
                        renderItem={renderTag}
                        keyExtractor={(item) => item.id.toString()}
                        contentContainerStyle={{
                            flexDirection: 'row',
                            flexWrap: 'wrap'
                        }}
                    />
                ) : (
                    <FlatList
                        data={tagsData?.filter(
                            (tag) => selectedCategory.id === tag.category_id
                        )}
                        renderItem={renderTag}
                        keyExtractor={(item) => item.id.toString()}
                        contentContainerStyle={{
                            flexDirection: 'row',
                            flexWrap: 'wrap'
                        }}
                    />
                )}
            </View>
            {selectedTags.length === 0 && buttonClicked && (
                <View className="pt-[2%]">
                    <Error message="Please choose at least one interest" />
                </View>
            )}
            <View className="flex-row justify-end mt-[8%] mb-[10%]">
                <Button size="lg" onPress={handleSubmit(onSubmit)}>
                    Finish
                </Button>
            </View>
        </>
    );
};

export default UserInterestsForm;
