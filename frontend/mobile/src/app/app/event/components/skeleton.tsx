import { Skeleton } from '@rneui/base';

import { Box, Colors } from '@/src/app/design-system';

const EventPageSkeleton = () => {
    return (
        <Box flex={1} height="100%">
            <Box flex={1} aspectRatio={1}>
                <Skeleton
                    width="auto"
                    height="auto"
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
            </Box>
            <Box flex={1} margin="xl" gap="m">
                <Skeleton
                    width={'100%'}
                    height={80}
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
                <Skeleton
                    width={'50%'}
                    height={20}
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
                <Skeleton
                    width={'65%'}
                    height={70}
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
                <Skeleton
                    width={'100%'}
                    height={120}
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
                <Skeleton
                    width={'35%'}
                    height={15}
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
            </Box>
        </Box>
    );
};

export default EventPageSkeleton;
