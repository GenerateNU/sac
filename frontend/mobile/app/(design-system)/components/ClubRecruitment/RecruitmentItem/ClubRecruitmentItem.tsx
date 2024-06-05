import { IconDefinition } from '@fortawesome/fontawesome-svg-core';

import { Box, SACColors, Text, createStyles } from '@/app/(design-system)';
import { firstLetterUppercase } from '@/utils/string';

import { Icon } from '../../Icon/Icon';

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
        <Box {...styles.recruitmentItem}>
            <Box {...styles.recruitmentItemContent}>
                <Icon size={'md'} color={color} icon={icon} />
                <Text variant="body-1" {...styles.recruitmentItemTitle}>
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
        borderWidth: 1,
        borderRadius: 'sm',
        borderColor: 'gray',
        width: '33%'
    },
    recruitmentItemContent: {
        paddingTop: 'm',
        paddingBottom: 'm',
        gap: 'xs',
        flexDirection: 'column',
        alignItems: 'center'
    },
    recruitmentItemTitle: {
        textAlign: 'center',
        fontWeight: '600'
    }
});
