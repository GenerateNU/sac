import { FlatList, Text } from 'react-native';

import { useEventTagsByEventIDQuery } from '@sac/lib/api/event';
import { uuid } from '@sac/lib/types/uuid';

import { Tag } from '@/components/tag';

// TODO: generify this component

type TagListProps = {
    id: uuid;
};

const TagList = ({ id }: TagListProps) => {
    const { data: tags, isLoading, error } = useEventTagsByEventIDQuery(id);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        console.error(error);
        return <Text>Error</Text>;
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
