import { Text } from 'react-native';
import { Tag as T } from '@/types/tag';
import { Button } from '@/components/button';

type TagProps = {
    tag: T;
};

const Tag = ({ tag }: TagProps) => {
    return (
        <Button onPress={() => console.log('Pressed:', tag)} size={'pill'} variant={'pill'}>
            <Text>{tag.name}</Text>
        </Button>
    );
};

export default Tag;
