import { Text } from 'react-native';

import { router } from 'expo-router';

import { Button } from '@/components/button';
import { Tag as T } from '@/types/tag';

type TagProps = {
    tag: T;
    color: string;
};

const Tag = ({ tag, color }: TagProps) => {
    return (
        <Button
            onPress={() => router.push(`/(app)/search/results/${tag.name}`)}
            size={'pill'}
            variant={'pill'}
            color={color}
            className="my-1 mr-1"
        >
            <Text>{tag.name}</Text>
        </Button>
    );
};

export { Tag };
