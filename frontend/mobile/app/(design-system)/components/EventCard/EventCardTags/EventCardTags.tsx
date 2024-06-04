import { Tag } from '@generatesac/lib';

import { Box, Tag as TagComponent, Text } from '@/app/(design-system)';

interface EventTagProps {
    tags: Tag[];
}

export const EventTags = ({ tags }: EventTagProps) => {
    const renderTags = () => {
        let currentLength = 0;
        return tags.filter((tag) => {
            currentLength += tag.name.length;
            return currentLength <= 25;
        });
    };

    const renderPlusTag = (remainingTagCount: number) => {
        return (
            remainingTagCount > 0 && (
                <TagComponent variant="eventCard">
                    <Text variant="caption-1">{`+${remainingTagCount}`}</Text>
                </TagComponent>
            )
        );
    };

    const renderedTags = renderTags();
    const remainingTagCount = tags.length - renderedTags.length;

    return (
        <Box flexDirection="row" paddingTop="xxs" gap="xxs">
            {renderedTags.map((tag, index) => (
                <TagComponent key={index} variant="eventCard">
                    <Text variant="caption-1">{tag.name}</Text>
                </TagComponent>
            ))}
            {renderPlusTag(remainingTagCount)}
        </Box>
    );
};
