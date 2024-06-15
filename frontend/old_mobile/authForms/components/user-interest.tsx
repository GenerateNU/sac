import { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import { Alert, Dimensions, FlatList, TouchableOpacity } from 'react-native';

import { Category, Tag } from '@generatesac/lib';
import { ZodError } from 'zod';

import { useCategories } from '@/src/hooks/use-categories';
import { useTags } from '@/src/hooks/use-tags';

import { Box, Text } from '../../(design-system)';
import { Button } from '../../(design-system)/components/Button/Button';
import { Tag as TagComponent } from '../../(design-system)/components/Tag/Tag';

type UserInterestsData = {
    tags: Tag[];
};

interface UserInterestProps {
    onPress: () => void;
}

export const UserInterests: React.FC<UserInterestProps> = ({ onPress }) => {
    const { height } = Dimensions.get('window');
    const { handleSubmit } = useForm<UserInterestsData>();
    const [selectedTags, setSelectedTags] = useState<Tag[]>([]);

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

    const selected = (item: Tag) => {
        return selectedTags.includes(item);
    };

    const renderCategory = ({ item }: { item: Category }) => {
        return (
            <Box
                borderBottomColor="black"
                borderBottomWidth={item === selectedCategory ? 1 : 0}
                paddingHorizontal="m"
                paddingVertical="s"
            >
                <TouchableOpacity onPress={() => handleCategoryPress(item)}>
                    <Text fontWeight={300}>{item.name}</Text>
                </TouchableOpacity>
            </Box>
        );
    };

    const handleTagPress = (tag: Tag) => {
        if (selectedTags.includes(tag)) {
            setSelectedTags(selectedTags.filter((t) => t !== tag));
        } else {
            setSelectedTags([...selectedTags, tag]);
        }
    };

    const handleCategoryPress = (category: Category) => {
        setSelectedCategory(category);
    };

    const onSubmit = (data: UserInterestsData) => {
        setButtonClicked(true);
        onPress();
        if (selectedTags.length === 0) {
            return;
        }
        try {
            console.log({ ...data, tags: selectedTags });
        } catch (error) {
            if (error instanceof ZodError) {
                Alert.alert('Validation Error', error.errors[0].message);
            } else {
                console.error('An unexpected error occurred:', error);
            }
        }
    };

    return (
        <Box
            paddingHorizontal="xl"
            flexDirection="column"
            alignItems="flex-start"
            justifyContent="center"
            gap="l"
        >
            <Text fontSize={40} variant="header-1">
                What are you interested in?
            </Text>
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
            <Box height={height / 2.5} flexDirection="row" flexWrap="wrap">
                {tags &&
                    tags
                        .filter(
                            (tag) => selectedCategory?.id === tag.category_id
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
                                        selected(item) ? 'black' : 'gray'
                                    }
                                    color={selected(item) ? 'black' : 'white'}
                                >
                                    <Text
                                        color={
                                            selected(item) ? 'white' : 'black'
                                        }
                                    >
                                        {item.name}
                                    </Text>
                                </TagComponent>
                            </Box>
                        ))}
            </Box>
            <Button
                onPress={handleSubmit(onSubmit)}
                variant="standardButton"
                disabled={selectedTags.length === 0}
            >
                Continue
            </Button>
        </Box>
    );
};
