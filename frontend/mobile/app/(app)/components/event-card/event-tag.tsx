import { Tag } from '@generatesac/lib';

import { Box, Tag as TagComponent, Text } from '@/app/(design-system)';

interface EventTagProps {
    tags: Tag[];
    variant: 'limited' | 'plus';
}

export const EventTags = ({ tags, variant }: EventTagProps) => {
    const displayLimitedTags = variant === 'limited';

    const renderTags = () => {
        if (displayLimitedTags) {
            const shouldDisplayOnlyFirstTag =
                tags.length > 0 &&
                tags[0].name.length + (tags[1]?.name?.length || 0) > 30;
            return shouldDisplayOnlyFirstTag ? [tags[0]] : tags.slice(0, 2);
        } else {
            return tags.slice(0, 1);
        }
    };

    const renderPlusTag = () => {
        const remainingTagsCount = tags.length - 1;
        return (
            remainingTagsCount > 0 && (
                <TagComponent variant="eventCard">
                    <Text variant="caption-1">{`+${remainingTagsCount}`}</Text>
                </TagComponent>
            )
        );
    };

    return (
        <Box flexDirection="row" paddingTop="xxs" gap="xxs">
            {renderTags().map((tag, index) => (
                <TagComponent key={index} variant="eventCard">
                    <Text variant="caption-1">{tag.name}</Text>
                </TagComponent>
            ))}
            {variant === 'plus' && renderPlusTag()}
        </Box>
    );
};
