import { Pressable, Text, View } from 'react-native';

import { useHosts } from '@/hooks/use-hosts';
import { uuid } from '@/types/uuid';

const AllHosts = ({ eventID }: { eventID: uuid }) => {
    const { data: hosts, error, isLoading } = useHosts(eventID);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    if (!hosts || hosts.length === 0) {
        return <Text>No hosts found</Text>;
    }


    return (
        <View className='flex-row flex-wrap w-full text-lg'>
            <Text className="text-lg">Hosted by{' '}</Text>
            {hosts.map((host, index) => (
                <>
                    <Pressable
                        key={index}
                        onPress={() => console.log(`${host.name} pressed`)}
                        className={'font-bold'}
                    >
                        <Text className="text-lg font-bold">
                            {host.name}
                        </Text>
                    </Pressable>
                    <Text className='text-lg'>{index < hosts.length - 2 ? ', ' : (index === hosts.length - 2 ? ' & ' : '')}</Text>
                </>
            ))
            }
        </View>
    );
};

export { AllHosts };
