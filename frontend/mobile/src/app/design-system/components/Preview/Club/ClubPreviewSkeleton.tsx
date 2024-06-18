import { Skeleton } from '@rneui/base';

import { Colors } from '../../../shared/colors';
import { Box } from '../../Box/Box';

const ClubPreviewSkeleton = () => {
    return (
        <Box marginHorizontal="xxl" gap="m">
            <Box height={100} width="75%">
                <Skeleton
                    height="auto"
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
            </Box>
            <Box height={30}>
                <Skeleton
                    height="auto"
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
            </Box>
            <Box height={20}>
                <Skeleton
                    width={'30%'}
                    height="auto"
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
            </Box>
            <Box height={125}>
                <Skeleton
                    height="auto"
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
            </Box>
        </Box>
    );
};

export default ClubPreviewSkeleton;
