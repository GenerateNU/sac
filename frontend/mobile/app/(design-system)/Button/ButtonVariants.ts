import { BackgroundColorVariants, defaultColor } from "../Colors";

export const ButtonVariants = {
    defaults: {
        flex: '1',
        alignItems: 'center',
        justifyContent: 'center',
        paddingHorizontal: 'l',
        paddingVertical: 'm',
        backgroundColor: defaultColor,
        borderRadius: 8,
        minWidth: '100%'
    },
    ...BackgroundColorVariants
};