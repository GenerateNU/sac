import { backgroundColor } from "@shopify/restyle";

// TODO: We will have to change these
export const Colors = {
    darkBlue: '#0C3354',
    darkRed: '#C8102E',
    green: '#14AC3F',
    blue: '#2663FF',
    aqua: '#35B8E2',
    purple: '#9226FF',
    red: '#EC2215',
    orange: '#FF7426',
    yellow: '#FFB626',
    white: "#FFFFFF",
    black: "#000000",
    disabled: "#BDBDBD",
};

export const defaultColor: keyof typeof Colors = 'black';

export const BackgroundColorVariants = {
    darkBlue: {
        backgroundColor: 'darkBlue',
    },
    darkRed: {
        backgroundColor: 'darkRed',
    },
    green: {
        backgroundColor: 'green',
    },
    blue: {
        backgroundColor: 'blue',
    },
    aqua: {
        backgroundColor: 'aqua',
    },
    purple: {
        backgroundColor: 'purple',
    },
    red: {
        backgroundColor: 'red',
    },
    orange: {
        backgroundColor: 'orange',
    },
    yellow: {
        backgroundColor: 'yellow',
    },
    white: {
        backgroundColor: 'white',
    },
    black: {
        backgroundColor: 'black',
    },
    disabled: {
        backgroundColor: 'disabled',
    }
} as const;

export const textColorVariants = {
    darkBlue: 'white',
    darkRed: 'white',
    green: 'white',
    blue: 'white',
    aqua: 'white',
    purple: 'white',
    red: 'white',
    orange: 'white',
    yellow: 'white',
    white: 'black',
    black: 'white',
    disabled: 'white',
} as const;
