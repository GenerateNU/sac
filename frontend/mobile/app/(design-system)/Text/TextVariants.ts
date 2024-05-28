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
        fontWeight: '500',
        lineHeight: 'normal'
    },
    'body-1': {
        fontFamily: 'DMSans-Regular',
        fontSize: 16,
        fontStyle: 'normal',
        fontWeight: '400',
        lineHeight: 'normal'
    },
    'caption-1': {
        fontFamily: 'DMSans-Regular',
        fontSize: 12,
        fontStyle: 'normal',
        fontWeight: '400',
        lineHeight: 'normal'
    },
    'caption-2': {
        fontFamily: 'DMSans-Regular',
        fontSize: 10,
        fontStyle: 'normal',
        fontWeight: '400',
        lineHeight: 'normal'
    }
} as const;

export const TextVariants = {
    defaults: Texts['body-1'],
    ...Texts
};
