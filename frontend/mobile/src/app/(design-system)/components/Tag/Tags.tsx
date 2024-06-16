import { Tag } from '@generatesac/lib';

import { SACColors } from '../../shared/colors';
import { Box } from '../Box/Box';
import { Text } from '../Text/Text';
import { Tag as TagComponent } from './Tag';

interface EventTagsProps {
    tags: Tag[];
    color: SACColors;
}

export const EventPageTags: React.FC<EventTagsProps> = ({ tags, color }) => {
    const renderTag = (item: Tag) => (
        <TagComponent key={item.id} variant="selected" color={color}>
            <Text variant="body-2">{item.name}</Text>
        </TagComponent>
    );

    return (
        <Box flexDirection="row" flexWrap="wrap" gap="xs">
            {tags.length >= 5
                ? tags.slice(0, 5).map((item) => renderTag(item))
                : tags.map((item) => renderTag(item))}
        </Box>
    );
};
