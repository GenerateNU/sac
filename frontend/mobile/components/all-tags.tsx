import { FlatList, Text } from 'react-native';

import { Tag } from '@/components/tag';
import { useEventTags } from '@/hooks/use-event';
import { uuid } from '@/types/uuid';

// TODO: generify this component

type TagListProps = {
    id: uuid;
};

const TagList = ({ id }: TagListProps) => {
    const { data: tags, isLoading, error } = useEventTags(id);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    if (tags?.length === 0) {
        return null;
    }

    return (
        <FlatList
            horizontal
            showsHorizontalScrollIndicator={false}
            alwaysBounceHorizontal
            data={tags}
            renderItem={({ item }) => <Tag key={item.id} tag={item} />}
            keyExtractor={(item) => item.id}
        />
    );
};

export { TagList };
