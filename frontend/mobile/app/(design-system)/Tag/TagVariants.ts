const TagColorVariants = {
    darkBlue: {
        backgroundColor: 'darkBlue'
    },
    darkRed: {
        backgroundColor: 'darkRed'
    },
    green: {
        backgroundColor: 'green'
    },
    blue: {
        backgroundColor: 'blue'
    },
    aqua: {
        backgroundColor: 'aqua'
    },
    purple: {
        backgroundColor: 'purple'
    },
    red: {
        backgroundColor: 'red'
    },
    orange: {
        backgroundColor: 'orange'
    },
    yellow: {
        backgroundColor: 'yellow'
    }
} as const;

export const TagVariants = {
    defaults: {
        flex: '1',
        alignItems: 'center',
        justifyContent: 'center',
        paddingHorizontal: 'l',
        paddingVertical: 'xs',
        backgroundColor: 'aqua',
        borderRadius: 105
    },
    eventCard: {
        flex: '1',
        alignItems: 'center',
        justifyContent: 'center',
        paddingHorizontal: 's',
        paddingVertical: 'xxs',
        borderColor: 'darkGray',
        borderWidth: 0.3,
        borderRadius: 10,
        backgroundColor: 'white'
    },
    ...TagColorVariants
};
