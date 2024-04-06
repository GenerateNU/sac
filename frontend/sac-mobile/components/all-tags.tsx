import { ScrollView, Text, View } from 'react-native';

import { useEventTags } from '@/hooks/use-event-tags';
import { uuid } from '@/types/uuid';

import Tag from '@/components/tag';

const AllTags = ({ eventID }: { eventID: uuid }) => {
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
        <ScrollView
            className="w-full py-2"
            horizontal={true}
            showsHorizontalScrollIndicator={false}
        >
            <View className="flex-row flex-wrap w-full space-x-2">
                {tags.map((tag) => (
                    <View key={tag.id}>
                        <Tag tag={tag} />
                    </View>
                ))}
            </View>
        </ScrollView>
    );
};

export default AllTags;
