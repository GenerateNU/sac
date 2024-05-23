// Add actual text variants here
const Texts = {
    'text-1': {
        fontSize: 20,
        fontWeight: 'bold'
    },
    'text-2': {
        fontSize: 16,
        fontWeight: 'normal'
    },
    'text-3': {
        fontSize: 14,
        fontWeight: 'normal'
    }
} as const;

export const TextVariants = {
    defaults: Texts['text-1'],
    ...Texts
};
