import { EBoardCard, EBoardCardProps } from '@/components/eboardCard';
import { ScrollView, Text, View } from 'react-native';

const EBoardMembers = () => {
    const { data: eboardmembers, isLoading, error } = useEBoardMembers();

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
                    {eboardmembers.map((member: EBoardCardProps) => (
                        <EBoardCard key={member.id} photo={member.photo} name={member.name} title={member.title} />
                    ))}
                </View>
            </ScrollView>
        </View>
    );
}

export { EBoardMembers };

function useEBoardMembers(): { data: EBoardCardProps[], isLoading: boolean, error: Error | null } {
    return {
        data: [
            { name: 'Alder Whiteford', title: 'Project Lead' },
            { name: 'Amanda Kerr', title: 'Design Lead' },
            { name: 'David Oduneye', title: 'Tech Lead' },
            { name: 'Garrett Ladley', title: 'Tech Lead' },
        ],
        isLoading: false,
        error: null
    };
}