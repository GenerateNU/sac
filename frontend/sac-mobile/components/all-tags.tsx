import { ScrollView, Text, View } from 'react-native';

import { Tag } from '@/components/tag';
import { useEventTags } from '@/hooks/use-event';
import { uuid } from '@/types/uuid';

// TODO: generify this component

type TagListProps = {
    id: uuid;
};

const TagList = ({ id }: TagListProps) => {
    const { data: tags = [], isLoading, error } = useEventTags(id);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    if (tags.length === 0) {
        return <Text>No tags found</Text>;
    }

    return (
        <ScrollView
            className="w-full py-2"
            horizontal={true}
            showsHorizontalScrollIndicator={false}
        >
            <View className="flex-row justify-start items-start gap-1.5 inline-flex">
                {tags.map((tag) => (
                    <View key={tag.id}>
                        <Tag tag={tag} />
                    </View>
                ))}
            </View>
        </ScrollView>
    );
};

export { TagList };
