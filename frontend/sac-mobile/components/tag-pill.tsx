import { Text } from 'react-native';

import { Tag } from '@/types/tag';

import { Button } from './button';

const TagPill = ({ tag }: { tag: Tag }) => {
    return (
        <Button onPress={() => console.log('Pressed:', tag)}>
            <Text>{tag.name}</Text>
        </Button>
    );
};

export default TagPill;
