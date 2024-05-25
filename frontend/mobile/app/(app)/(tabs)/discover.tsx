import { Box, Tag, Text } from '@/app/(design-system)';

const DiscoverPage = () => {
    return (
        <Box flex={1} alignItems="center" justifyContent="center" padding="m">
            <Tag variant="darkRed">
                <Text variant="body-1">Discover</Text>
            </Tag>
        </Box>
    );
};

export default DiscoverPage;
