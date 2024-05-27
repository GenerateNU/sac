import { TextVariants } from "../Text/TextVariants";

export const TextBoxColorVariants = {
    error: {
        borderColor: 'darkRed',
    },
}

export const TextBoxVariants = {
    defaults: {
        borderWidth: 1,
        borderRadius: 8,
        borderColor: 'gray',
        backgroundColor: 'white',
        padding: 'm',
        marginBottom: 'xs',
        marginTop: 'xs',
        ...TextVariants["body-1"],
        minWidth: '100%',
    },
    ...TextBoxColorVariants
}