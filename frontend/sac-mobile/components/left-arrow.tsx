import { router } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';

import { Button } from '@/components/button';

const LeftArrow = ({ onPress }: { onPress?: () => void }) => {
    return (
        <Button
            onPress={onPress || (() => router.back())}
            size={'icon'}
            variant={'icon'}
            className="flex items-center justify-center bg-gray-300 rounded-full"
        >
            <MaterialCommunityIcons
                name="chevron-left"
                size={28}
                color="rgb(75 85 99)"
            />
        </Button>
    );
};

export { LeftArrow };
