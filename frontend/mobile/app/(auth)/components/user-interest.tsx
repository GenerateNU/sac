import { useState } from 'react';
import { useForm } from 'react-hook-form';
import {
    Alert,
    Dimensions,
    FlatList,
    SafeAreaView,
    TouchableOpacity
} from 'react-native';

import { Category, Tag } from '@generatesac/lib';
import { ZodError } from 'zod';

import { CATEGORIES } from '@/consts/user-interest';

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
    const [selectedCategory, setSelectedCategory] = useState<Category>(
        CATEGORIES[0]
    );
    const [_, setButtonClicked] = useState<boolean>(false);

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
        <SafeAreaView>
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
                    data={CATEGORIES}
                    renderItem={renderCategory}
                    keyExtractor={(item) => item.id.toString()}
                />
                <Box height={height / 2.5} flexDirection="row" flexWrap="wrap">
                    {selectedCategory.tags.map((item: Tag) => (
                        <Box key={item.id} marginRight="xxs" marginBottom="xs">
                            <TagComponent
                                onPress={() => handleTagPress(item)}
                                borderColor={selected(item) ? 'black' : 'gray'}
                                color={selected(item) ? 'black' : 'white'}
                            >
                                <Text
                                    color={selected(item) ? 'white' : 'black'}
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
        </SafeAreaView>
    );
};
