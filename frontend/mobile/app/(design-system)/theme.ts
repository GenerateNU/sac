import { createTheme } from '@shopify/restyle';

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

export type Theme = typeof theme;
