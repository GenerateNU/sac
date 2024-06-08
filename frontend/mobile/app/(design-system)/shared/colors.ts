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
    white: '#FFFFFF',
    black: '#000000',
    gray: '#C3C9D0',
    darkGray: '#7A7A7A',
    transparent: ''
};

export type SACColors = keyof typeof Colors;

export const defaultColor: keyof typeof Colors = 'black';

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
    gray: 'white',
    darkGray: 'white',
    transparent: 'black'
} as const;
