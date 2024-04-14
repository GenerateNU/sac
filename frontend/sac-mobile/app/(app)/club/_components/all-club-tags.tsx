import { ClubProfileTag, ClubProfileTagProps } from '@/components/club-profile-tag';
import { ScrollView, Text, View } from 'react-native';
import { useTags } from '@/hooks/use-tags';

const AllClubTags = () => {
    const { data: tags, isLoading, error } = useTags("sac");

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View>
            <ScrollView horizontal={true} showsHorizontalScrollIndicator={false} className="pb-[3%]">
                <View className="flex-row flex-wrap">
                    {tags && tags.map((tag: ClubProfileTagProps) => (
                        <ClubProfileTag key={tag.id} name={tag.name} />
                    ))}
                </View>
            </ScrollView>
        </View>
    );
}

export { AllClubTags };