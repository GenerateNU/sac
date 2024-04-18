import { Text } from 'react-native';

import { router } from 'expo-router';

import { Button } from '@/components/button';
import { Tag as T } from '@/types/tag';

type TagProps = {
    tag: T;
};

const Tag = ({ tag }: TagProps) => {
    return (
        <Button
            onPress={() => router.push(`/tag/${tag.id}`)}
            size={'pill'}
            variant={'pill'}
            className="my-1 mr-1"
        >
            <Text>{tag.name}</Text>
        </Button>
    );
};

export { Tag };
