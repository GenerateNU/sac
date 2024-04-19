import { ClubAdminCardProps, ClubAdminCard } from '@/app/(app)/club/_components/admin-card';
import { useClubAdmin } from '@/hooks/use-club';
import { uuid } from '@/types/uuid';
import { ScrollView, Text, View } from 'react-native';

const ClubAdmins = ({ clubID }: { clubID: string }) => {
    // const { data: admins, isLoading, error } = useClubAdmin(clubID);
    const { data: admins, isLoading, error } = useClubAdmin("sac");

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
                    {admins && admins.map((member: ClubAdminCardProps) => (
                        <ClubAdminCard key={member.name} {...member} />
                    ))}
                </View>
            </ScrollView>
        </View>
    );
}

export { ClubAdmins };