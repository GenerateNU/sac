import { AllProps, createTheme } from '@shopify/restyle';

import { Colors } from './Colors';
import { Spacing } from './Spacing';
import { TagVariants } from './Tag/TagVariants';
import { TextVariants } from './Text/TextVariants';

export const theme = createTheme({
    colors: Colors,
    spacing: Spacing,
    textVariants: TextVariants,
    tagVariants: TagVariants
});

export const createStyles = <Keys extends string>(
    styles: Record<Keys, AllProps<Theme>>
): Record<Keys, AllProps<Theme>> => {
    return styles;
};

export type Theme = typeof theme;
