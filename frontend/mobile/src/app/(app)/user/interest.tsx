import { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import { Alert, Dimensions, FlatList, TouchableOpacity } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { Stack, router } from 'expo-router';

import { Category, Tag } from '@generatesac/lib';
import { ZodError } from 'zod';

import { Arrow, Box, Spacing, Text } from '@/src/app/(design-system)';
import { Tag as TagComponent } from '@/src/app/(design-system)';
import { useCategories } from '@/src/hooks/use-categories';
import { useTags } from '@/src/hooks/use-tags';
import { formatCategoryOrTag } from '@/src/utils/string';

import { Save } from './components/save';

type UserInterestsData = {
    tags: Tag[];
};

const mockTags: Tag[] = [
    {
        category_id: '9b45050e-4a55-4cd1-bd12-b6322af78a0f',
        created_at: '2024-03-26T13:03:44.22905Z',
        id: '4628601d-65d9-405a-9cdc-c68fd1a17733',
        name: 'Prelaw',
        updated_at: '2024-03-26T13:03:44.22905Z'
    },
    {
        category_id: 'b7808076-f1fb-48e1-88bf-6eb181c1d7ee',
        created_at: '2024-03-26T13:03:44.22905Z',
        id: 'e690247a-1b17-4992-9e6f-d8ba4be084bc',
        name: 'Christianity',
        updated_at: '2024-03-26T13:03:44.22905Z'
    },
    {
        category_id: 'b7808076-f1fb-48e1-88bf-6eb181c1d7ee',
        created_at: '2024-03-26T13:03:44.22905Z',
        id: '7fabc19e-298d-446e-9200-e3019ebfa738',
        name: 'LatinAmerica',
        updated_at: '2024-03-26T13:03:44.22905Z'
    },
    {
        category_id: 'f1e79d31-895d-4c99-bb66-9ce09fb0e3d4',
        created_at: '2024-03-26T13:03:44.22905Z',
        id: '812b3594-8984-4718-b0be-2d2fe409955c',
        name: 'Music',
        updated_at: '2024-03-26T13:03:44.22905Z'
    }
];

const UserInterest = () => {
    const { height } = Dimensions.get('window');
    const { handleSubmit } = useForm<UserInterestsData>();

    const [_, setButtonClicked] = useState<boolean>(false);
    const {
        data: categoryData,
        isLoading: loadingCategories,
        error: errorCategory
    } = useCategories();
    const {
        data: tagsData,
        isLoading: loadingTags,
        error: errorTag
    } = useTags();

    const [selectedCategory, setSelectedCategory] = useState<Category>();
    const [selectedTags, setSelectedTags] = useState<Tag[]>(mockTags);
    const isButtonDisabled =
        JSON.stringify(selectedTags) === JSON.stringify(mockTags);

    useEffect(() => {
        if (categoryData && categoryData.length > 0) {
            setSelectedCategory(categoryData[0]);
        }
    }, [categoryData]);

    if (loadingTags || loadingCategories) {
        return <Text>Loading...</Text>;
    }

    if (errorTag || errorCategory) {
        return <Text>Error</Text>;
    }

    const tags = tagsData ? tagsData : [];
    const categories = categoryData ? categoryData : [];

    const isSelected = (item: Tag) => {
        return selectedTags.map((tag) => tag.id).includes(item.id);
    };

    const renderCategory = ({ item }: { item: Category }) => {
        return (
            <Box
                borderBottomColor={'black'}
                borderBottomWidth={item === selectedCategory ? 1 : 0}
                paddingHorizontal="m"
                paddingVertical="s"
            >
                <TouchableOpacity onPress={() => handleCategoryPress(item)}>
                    <Text
                        color={
                            item.id === selectedCategory?.id
                                ? 'black'
                                : 'darkGray'
                        }
                        fontWeight={
                            item.id === selectedCategory?.id ? 500 : 300
                        }
                    >
                        {formatCategoryOrTag(item.name)}
                    </Text>
                </TouchableOpacity>
            </Box>
        );
    };

    const handleTagPress = (tag: Tag) => {
        if (selectedTags.includes(tag)) {
            setSelectedTags(selectedTags.filter((t) => t.id !== tag.id));
        } else {
            setSelectedTags([...selectedTags, tag]);
        }
    };

    const handleCategoryPress = (category: Category) => {
        setSelectedCategory(category);
    };

    const onSubmit = (data: UserInterestsData) => {
        setButtonClicked(true);
        console.log(isButtonDisabled);
        if (selectedTags.length === 0) {
            return;
        }
        try {
            console.log({ ...data, tags: selectedTags });
            router.back();
        } catch (error) {
            if (error instanceof ZodError) {
                Alert.alert('Validation Error', error.errors[0].message);
            } else {
                console.error('An unexpected error occurred:', error);
            }
        }
    };

    return (
        <>
            <Stack.Screen
                options={{
                    headerTitle: () => (
                        <Text variant="subheader-2">Interests</Text>
                    ),
                    headerTransparent: true,
                    headerShown: true,
                    headerLeft: () => <Arrow color="black" />,
                    headerRight: () => (
                        <Save
                            onPress={handleSubmit(onSubmit)}
                            isValid={
                                !isButtonDisabled && selectedTags.length > 0
                            }
                        />
                    )
                }}
            />
            <SafeAreaView style={{ paddingTop: Spacing.xxl }}>
                <Box
                    paddingHorizontal="xl"
                    flexDirection="column"
                    alignItems="flex-start"
                    justifyContent="center"
                    paddingTop="xxl"
                    gap="l"
                >
                    <Text>
                        Get personalized recommendations. Select one or more to
                        continue.
                    </Text>
                    <FlatList
                        horizontal
                        showsHorizontalScrollIndicator={false}
                        data={categories}
                        renderItem={renderCategory}
                        keyExtractor={(item) => item.id.toString()}
                    />
                    <Box
                        height={height / 2.5}
                        flexDirection="row"
                        flexWrap="wrap"
                    >
                        {tags &&
                            tags
                                .filter(
                                    (tag) =>
                                        selectedCategory?.id === tag.category_id
                                )
                                .map((item: Tag) => (
                                    <Box
                                        key={item.id}
                                        marginRight="xxs"
                                        marginBottom="xs"
                                    >
                                        <TagComponent
                                            onPress={() => handleTagPress(item)}
                                            borderColor={
                                                isSelected(item)
                                                    ? 'black'
                                                    : 'gray'
                                            }
                                            color={
                                                isSelected(item)
                                                    ? 'black'
                                                    : 'white'
                                            }
                                        >
                                            <Text
                                                color={
                                                    isSelected(item)
                                                        ? 'white'
                                                        : 'black'
                                                }
                                            >
                                                {formatCategoryOrTag(item.name)}
                                            </Text>
                                        </TagComponent>
                                    </Box>
                                ))}
                    </Box>
                </Box>
            </SafeAreaView>
        </>
    );
};

export default UserInterest;
