import { EBoardCard, EBoardCardProps } from '@/components/eboardCard';
import { useClubAdmin } from '@/hooks/use-club';
import { ScrollView, Text, View } from 'react-native';

const EBoardMembers = () => {
    const { data: eboardmembers, isLoading, error } = useClubAdmin("sac");

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View>
            <Text className="text-lg font-semibold text-black">E-Board</Text>
            <ScrollView horizontal={true} className="pb-[10%]" showsHorizontalScrollIndicator={false}>
                <View className="flex-row">
                    {eboardmembers && eboardmembers.map((member: EBoardCardProps) => (
                        <EBoardCard key={member.name} photo={member.photo} name={member.name} title={member.title}/>
                    ))}
                </View>
            </ScrollView>
        </View>
    );
}

export { EBoardMembers };