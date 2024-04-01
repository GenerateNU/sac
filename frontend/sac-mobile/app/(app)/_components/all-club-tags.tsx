import { ClubProfileTag, ClubProfileTagProps } from '@/components/club-profile-tag';
import { ScrollView, Text, View } from 'react-native';

const AllClubTags = () => {
    const { data: tags, isLoading, error } = useTags();

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View>
            <ScrollView horizontal={true} className="pb-[3%]">
                <View className="flex-row flex-wrap">
                    {tags.map((tag: ClubProfileTagProps) => (
                        <ClubProfileTag key={tag.id} name={tag.name} />
                    ))}
                </View>
            </ScrollView>
        </View>
    );
}

export { AllClubTags };

function useTags(): { data: ClubProfileTagProps[], isLoading: boolean, error: Error | null } {
    return {
        data: [
            { name: 'Science' },
            { name: 'Technology' },
            { name: 'Engineering' },
            { name: 'Mathematics' }
        ],
        isLoading: false,
        error: null
    };
}