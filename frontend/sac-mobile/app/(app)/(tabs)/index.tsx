
import { Link } from 'expo-router';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';
import { useEvents } from '@/hooks/use-event';
import { View, Text } from 'react-native';

const Home = () => {
    const { signOut, user } = useAuthStore();
    const { data: events, isLoading, error } = useEvents();

    const handleSignOut = async () => {
        signOut();
    };

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View className="items-center justify-center flex-1 gap-5">
            <Text>Welcome {user?.first_name}</Text>
            <Button onPress={handleSignOut}>Sign Out</Button>
            {events?.map((event) => (
                <Link
                    key={event.id}
                    href={{
                        pathname: '/(app)/event/[id]',
                        params: { id: event.id }
                    }}
                >
                    <Text>{event.name}</Text>
                </Link>
            ))}

            <Link href={`/club/${1}`}>Club</Link>
            <Link href={`/club/modal/${1}`}>Modal</Link>

        </View>
    );
};

export default Home;
