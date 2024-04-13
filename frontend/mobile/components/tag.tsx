import { Text } from 'react-native';

import { Tag as T } from '@sac/lib/api/tag';

import { Button } from '@/components/button';

type TagProps = {
    tag: T;
};

const Tag = ({ tag }: TagProps) => {
    return (
        <Button
            onPress={() => console.log('Pressed:', tag)}
            size={'pill'}
            variant={'pill'}
            className="my-1 mr-1"
        >
            <Text>{tag.name}</Text>
        </Button>
    );
};

export { Tag };
