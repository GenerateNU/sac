import { router } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';

import { Button } from '@/components/button';

type Arrow = {
    direction?: 'left' | 'right';
    onPress?: () => void;
};

const Arrow = ({ direction, onPress }: Arrow) => {
    return (
        <Button
            onPress={onPress || (() => router.back())}
            size={'icon'}
            variant={'icon'}
            className="flex items-center justify-center bg-gray-300 rounded-full"
        >
            <MaterialCommunityIcons
                name={`chevron-${direction || 'left'}`}
                size={28}
                color="rgb(75 85 99)"
            />
        </Button>
    );
};

export { Arrow };
