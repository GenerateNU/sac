import { Tag as TagProps } from '@generatesac/lib';

import { Box } from '../Box/Box';
import { ClubIcon } from '../ClubIcon/ClubIcon';
import { Tag } from '../Tag/Tag';
import { Text } from '../Text/Text';

interface ClubCardProps {
    club: string;
    imageUrl: string;
    tags?: TagProps[];
}

export const ClubCard: React.FC<ClubCardProps> = ({ club, imageUrl, tags }) => {
    return (
        <Box
            flexDirection="row"
            alignItems="center"
            padding="m"
            borderRadius="sm"
            backgroundColor="white"
            shadowOffset={{ width: 0, height: 2 }}
            elevation={2}
            shadowOpacity={0.1}
        >
            <ClubIcon size={44} imageUrl={imageUrl} />
            <Box marginLeft="m">
                <Text variant="body-1">{club}</Text>
                {tags && (
                    <Box flexDirection="row" flexWrap="wrap">
                        {tags.map((tag) => (
                            <Tag key={tag.id} variant="eventCard">
                                <Text variant="caption-1">{tag.name}</Text>
                            </Tag>
                        ))}
                    </Box>
                )}
            </Box>
        </Box>
    );
};
