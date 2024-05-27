import { createTheme } from '@shopify/restyle';

import { Colors } from './Colors';
import { Spacing } from './Spacing';
import { TagVariants } from './Tag/TagVariants';
import { TextVariants } from './Text/TextVariants';
import { ButtonVariants } from './Button/ButtonVariants';

export const theme = createTheme({
    colors: Colors,
    spacing: Spacing,
    textVariants: TextVariants,
    tagVariants: TagVariants,
    buttonVariants: ButtonVariants,
});

export type Theme = typeof theme;
