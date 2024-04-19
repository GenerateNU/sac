import { ScrollView, Text, View } from 'react-native';

import { ClubAdminCard } from '@/app/(app)/club/_components/admin-card';
import { useClubAdmin } from '@/hooks/use-club';

const ClubAdmins = ({}: { clubID?: string }) => {
    // const { data: admins, isLoading, error } = useClubAdmin(clubID);
    const { data: admins, isLoading, error } = useClubAdmin('sac');

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View>
            <Text className="text-lg font-semibold text-black">E-Board</Text>
            <ScrollView
                horizontal={true}
                className="pb-[10%]"
                showsHorizontalScrollIndicator={false}
            >
                <View className="flex-row">
                    {admins &&
                        admins.map((member) => (
                            <ClubAdminCard
                                key={member.name}
                                name={member.name}
                                title={member.title}
                                photo={'https://i.pravatar.cc/300'}
                            />
                        ))}
                </View>
            </ScrollView>
        </View>
    );
};

export { ClubAdmins };
