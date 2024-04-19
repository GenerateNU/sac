import { FlatList, Text } from 'react-native';

import { Tag as TagComponent } from '@/components/tag';
import { uuid } from '@/types/uuid';
import { UseQueryResult } from '@tanstack/react-query';
import { Tag } from '@/types/tag';

type TagListProps = {
    id: uuid
    tagQueryResult: (id: uuid) => UseQueryResult<Tag[], Error>
};

const TagList = ({ tagQueryResult, id }: TagListProps) => {
    const { data: tags, isLoading, error } = tagQueryResult(id);

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
            data={tags}
            renderItem={({ item }) => <TagComponent key={item.id} tag={item} />}
            keyExtractor={(item) => item.id}
            contentContainerStyle={{ flexDirection: 'row', flexWrap: 'wrap' }}
        />
    );
};

export { TagList };
