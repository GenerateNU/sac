import { ScrollView, Text, View } from 'react-native';

import { useEventTags } from '@/hooks/use-event-tags';
import { uuid } from '@/types/uuid';

import TagPill from './tag-pill';

const TagPillLazySusan = ({ eventID }: { eventID: uuid }) => {
    const { data: tags, error, isLoading } = useEventTags(eventID);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    if (!tags || tags.length === 0) {
        return <Text>No tags found</Text>;
    }

    return (
        <>
            <ScrollView
                className="w-full"
                horizontal={true}
                showsHorizontalScrollIndicator={false}
            >
                <View className="flex-row justify-between space-x-5">
                    {tags.map((tag) => (
                        <View key={tag.id}>
                            <TagPill tag={tag} />
                        </View>
                    ))}
                </View>
            </ScrollView>
        </>
    );
};

export default TagPillLazySusan;
