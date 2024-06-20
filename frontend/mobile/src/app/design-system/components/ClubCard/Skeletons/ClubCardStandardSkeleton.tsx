import { Skeleton } from '@rneui/base';

import { Colors } from '../../../shared/colors';
import { createStyles } from '../../../theme';
import { Box } from '../../Box/Box';

const ClubCardStandardSkeleton = () => {
    return (
        <Box {...styles.cardContainer}>
            <Box {...styles.cardSubContainer}>
                <Skeleton
                    circle
                    height={55}
                    width={55}
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
                <Skeleton
                    width="50%"
                    height={30}
                    skeletonStyle={{ backgroundColor: Colors.skeleton }}
                />
            </Box>
        </Box>
    );
};

const styles = createStyles({
    cardContainer: {
        shadowColor: 'black',
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.1,
        shadowRadius: 2,
        backgroundColor: 'white',
        borderRadius: 'md'
    },
    cardSubContainer: {
        width: '100%',
        borderRadius: 'md',
        padding: 'm',
        backgroundColor: 'white',
        flexDirection: 'row',
        alignItems: 'center',
        gap: 's',
        overflow: 'hidden'
    }
});

export default ClubCardStandardSkeleton;
