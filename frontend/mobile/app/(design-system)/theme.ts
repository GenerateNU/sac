import { createTheme } from '@shopify/restyle';

import { Colors } from './Colors';
import { Spacing } from './Spacing';
import { TagVariants } from './Tag/TagVariants';
import { TextVariants } from './Text/TextVariants';
import { TextBoxVariants } from './Textbox/TextboxVariants';

export const theme = createTheme({
    colors: Colors,
    spacing: Spacing,
    textVariants: TextVariants,
    tagVariants: TagVariants,
    textBoxVariants: TextBoxVariants
});

export type Theme = typeof theme;
