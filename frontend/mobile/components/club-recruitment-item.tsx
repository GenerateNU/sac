import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Box, ColorName, Colors, Text } from '@/app/(design-system)';
import { firstLetterUppercase } from '@/utils/string';

interface RecruitmentItemProps {
    icon: IconDefinition;
    title: string;
    text: string;
    color: ColorName;
}

export const RecruitmentItem = ({
    icon,
    title,
    text,
    color
}: RecruitmentItemProps) => {
    return (
        <Box
            alignItems="center"
            borderWidth={1}
            borderColor={color}
            borderRadius={8}
            width={'30%'}
        >
            <Box
                paddingTop="l"
                paddingBottom="l"
                gap="xs"
                flexDirection="column"
                alignItems="center"
            >
                <FontAwesomeIcon size={30} color={Colors[color]} icon={icon} />
                <Text
                    style={{
                        textAlign: 'center',
                        fontWeight: '600',
                        fontSize: 14
                    }}
                >
                    {firstLetterUppercase(title)}
                </Text>
                <Text variant="caption-1" style={{ textAlign: 'center' }}>
                    {firstLetterUppercase(text)}
                </Text>
            </Box>
        </Box>
    );
};
