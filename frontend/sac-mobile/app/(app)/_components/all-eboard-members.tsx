import { EBoardCard, EBoardCardProps } from '@/components/eboardCard';
import { useAdmin } from '@/hooks/use-admin';
import { ScrollView, Text, View } from 'react-native';

const EBoardMembers = () => {
    const { data: eboardmembers, isLoading, error } = useAdmin("sac");

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View>
            <Text className="text-black font-bold">E-Board</Text>
            <ScrollView horizontal={true} className="pb-[10%]">
                <View className="flex-row">
                    {eboardmembers && eboardmembers.map((member: EBoardCardProps) => (
                        <EBoardCard key={member.id} photo={member.photo} name={member.name} title={member.title} />
                    ))}
                </View>
            </ScrollView>
        </View>
    );
}

export { EBoardMembers };