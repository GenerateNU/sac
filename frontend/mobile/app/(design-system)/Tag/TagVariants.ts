import { BackgroundColorVariants, defaultColor } from "../Colors";

export const TagVariants = {
    defaults: {
        flex: '1',
        alignItems: 'center',
        justifyContent: 'center',
        paddingHorizontal: 'l',
        paddingVertical: 'xs',
        backgroundColor: defaultColor,
        borderRadius: 105
    },
    ...BackgroundColorVariants
};
