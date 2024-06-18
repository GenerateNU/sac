import { IconDefinition } from '@fortawesome/fontawesome-svg-core';

import { firstLetterUppercase } from '@/src/utils/string';

import { SACColors } from '../../../shared/colors';
import { createStyles } from '../../../theme';
import { Box } from '../../Box/Box';
import { Icon } from '../../Icon/Icon';
import { Text } from '../../Text/Text';

interface RecruitmentItemProps {
    icon: IconDefinition;
    title: string;
    text: string;
    color: SACColors;
}

export const RecruitmentItem = ({
    icon,
    title,
    text,
    color
}: RecruitmentItemProps) => {
    return (
        <Box borderColor={color} {...styles.recruitmentItem}>
            <Box {...styles.recruitmentItemContent}>
                <Icon size={'md'} color={color} icon={icon} />
                <Text variant="caption-1" {...styles.recruitmentItemTitle}>
                    {firstLetterUppercase(title)}
                </Text>
                <Text variant="caption-1" style={{ textAlign: 'center' }}>
                    {firstLetterUppercase(text)}
                </Text>
            </Box>
        </Box>
    );
};

const styles = createStyles({
    recruitmentItem: {
        alignItems: 'center',
        borderRadius: 'sm',
        width: '33%',
        padding: 's'
    },
    recruitmentItemContent: {
        paddingVertical: 'l',
        paddingHorizontal: 'xxs',
        gap: 'xs',
        flexDirection: 'column',
        alignItems: 'center'
    },
    recruitmentItemTitle: {
        textAlign: 'center',
        fontWeight: '600'
    }
});
