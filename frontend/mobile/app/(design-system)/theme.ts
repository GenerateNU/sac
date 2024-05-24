import { createTheme } from '@shopify/restyle';

import { Colors } from './Colors';
import { Spacing } from './Spacing';
import { TextVariants } from './TextVariants';

export const theme = createTheme({
    colors: Colors,
    spacing: Spacing,
    textVariants: TextVariants
});

export type Theme = typeof theme;
