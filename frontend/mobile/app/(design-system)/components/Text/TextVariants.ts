const Texts = {
    'header-1': {
        fontFamily: 'DMSans-Bold',
        fontSize: 32,
        fontStyle: 'normal',
        fontWeight: '700',
        lineHeight: 'normal'
    },
    'subheader-1': {
        fontFamily: 'DMSans-Medium',
        fontSize: 24,
        fontStyle: 'normal',
        fontWeight: '500'
    },
    'body-1': {
        fontFamily: 'DMSans-Regular',
        fontSize: 16,
        fontStyle: 'normal',
        fontWeight: '400'
    },
    'caption-1': {
        fontFamily: 'DMSans-Regular',
        fontSize: 12,
        fontStyle: 'normal',
        fontWeight: '200'
    },
    'caption-2': {
        fontFamily: 'DMSans-Regular',
        fontSize: 10,
        fontStyle: 'normal',
        fontWeight: '200'
    }
} as const;

export const TextVariants = {
    defaults: Texts['body-1'],
    ...Texts
};

