import { Text } from 'react-native';

import { useHosts } from '@/hooks/use-hosts';
import { uuid } from '@/types/uuid';

const HostedByText = ({ eventID }: { eventID: uuid }) => {
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

    if (hosts.length === 1) {
        return (
            <Text className="text-lg">
                Hosted by
                <Text
                    onPress={() => console.log(`${hosts[0]} pressed`)}
                    className="font-bold"
                >
                    {hosts[0].name}
                </Text>
            </Text>
        );
    } else {
        return (
            <Text className="text-lg">
                Hosted by{' '}
                {hosts.slice(0, -1).map((host, index) => (
                    <Text
                        key={index}
                        onPress={() => console.log(`${host} pressed`)}
                        className="font-bold"
                    >
                        {host.name}
                        {index === hosts.length - 2 ? '' : ', '}
                    </Text>
                ))}{' '}
                &{' '}
                <Text
                    onPress={() =>
                        console.log(`${hosts[hosts.length - 1]} pressed`)
                    }
                    className="font-bold"
                >
                    {hosts[hosts.length - 1].name}
                </Text>
            </Text>
        );
    }
};

export default HostedByText;
