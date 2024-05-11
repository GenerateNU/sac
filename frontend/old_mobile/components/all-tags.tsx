import { FlatList, Text } from 'react-native';
import { ScrollView } from 'react-native-gesture-handler';

import { UseQueryResult } from '@tanstack/react-query';

import { Tag as TagComponent } from '@/components/tag';
import { Tag } from '@/types/tag';
import { uuid } from '@/types/uuid';

type TagListProps = {
    id: uuid;
    color: string;
    tagQuery: (id: uuid) => UseQueryResult<Tag[], Error>;
};
// TODO: Fix this type
// UseQuery<QueryDefinition<uuid, BaseQueryFn<string | FetchArgs, unknown, FetchBaseQueryError, {}, FetchBaseQueryMeta>, "Clubs", Tag[], "clubsApi">>

const TagList = ({ tagQuery, id, color }: TagListProps) => {
    const { data: tags, isLoading, error } = tagQuery(id);

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
        <ScrollView horizontal>
            <FlatList
                data={tags}
                renderItem={({ item }) => (
                    <TagComponent color={color} key={item.id} tag={item} />
                )}
                keyExtractor={(item) => item.id}
                contentContainerStyle={{
                    flexDirection: 'row',
                    gap: 10
                }}
            />
        </ScrollView>
    );
};

export { TagList };
