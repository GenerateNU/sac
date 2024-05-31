import { Tag } from '@generatesac/lib';

import { Box, Text } from '@/app/(design-system)';

interface EventTagProps {
    tags: Tag[];
    variant: 'limited' | 'plus';
}

export const EventTags = ({ tags, variant }: EventTagProps) => {
    const displayLimitedTags = variant === 'limited';

    const renderTags = () => {
        if (displayLimitedTags) {
            const shouldDisplayOnlyFirstTag =
                tags.length >= 2 &&
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
                <Box
                    borderWidth={1}
                    borderColor="gray"
                    borderRadius="medium"
                    paddingHorizontal="xs"
                    paddingVertical="xxs"
                >
                    <Text variant="caption-1">{`+${remainingTagsCount}`}</Text>
                </Box>
            )
        );
    };

    return (
        <Box flexDirection="row" paddingTop="xxs" gap="xxs">
            {renderTags().map((tag, index) => (
                <Box
                    borderWidth={1}
                    borderColor="gray"
                    borderRadius="medium"
                    key={index}
                    paddingHorizontal="xs"
                    paddingVertical="xxs"
                >
                    <Text variant="caption-1">{tag.name}</Text>
                </Box>
            ))}
            {variant === 'plus' && renderPlusTag()}
        </Box>
    );
};
