import React from 'react';

import { Skeleton } from '@rneui/base';

import { Box, createStyles } from '@/app/(design-system)';

export const EventCardCalendarSkeleton = () => {
    return (
        <Box {...styles.cardContainer}>
            <Box {...styles.cardSubContainer}>
                <Box borderRadius='sm' aspectRatio={1}>
                    <Skeleton 
                        width={'auto'}
                        height={'auto'}
                        skeletonStyle={{ backgroundColor: '#FAFAFA' }} 
                    />
                </Box>
                <Box flex={1} gap={'xxs'}>
                    <Skeleton 
                        width={'50%'}
                        height={20}
                        skeletonStyle={{ backgroundColor: '#FAFAFA' }} 
                    />
                    <Skeleton 
                            width={'75%'} 
                            height={15} 
                            skeletonStyle={{ backgroundColor: '#FAFAFA' }} 
                        />
                    <Box flexDirection='row' gap='xxs'>
                        <Skeleton 
                            width={'20%'} 
                            height={20} 
                            skeletonStyle={{ backgroundColor: '#FAFAFA' }} 
                        />
                        <Skeleton 
                            width={'20%'} 
                            height={20} 
                            skeletonStyle={{ backgroundColor: '#FAFAFA' }} 
                        />
                        <Skeleton 
                            width={'20%'} 
                            height={20} 
                            skeletonStyle={{ backgroundColor: '#FAFAFA' }} 
                        />
                    </Box>
                </Box>
            </Box>
        </Box>
    );
};

const styles = createStyles({
    calendarImage: {
        flex: 1,
        borderRadius: 'sm',
    },
    isHappeningContainer: {
        zIndex: 1,
        position: 'absolute',
        marginRight: 'xl',
        top: '-15%',
        right: '-5%',
        flexDirection: 'row',
        justifyContent: 'flex-end'
    },
    isHappeningContent: {
        paddingHorizontal: 's',
        paddingVertical: 'xxs',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        borderRadius: 'sm',
        backgroundColor: 'darkRed',
        shadowColor: 'black',
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.3,
        shadowRadius: 2
    },
    cardContainer: {
        borderRadius: 'md',
        shadowColor:"black",
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.1,
        shadowRadius: 2,
        backgroundColor: 'white',
    },
    cardSubContainer: {
        width:"100%",
        borderRadius: 'md',
        padding: "m",
        backgroundColor: "white",
        flexDirection: "row",
        justifyContent: "space-between",
        gap: "s"
    },
    cardContentContainer: {
        flexDirection: "row",
        alignItems: "center",
        gap: "xxs",
        flexWrap: "wrap",
    },
});
