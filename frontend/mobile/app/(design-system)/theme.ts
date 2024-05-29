import { AllProps, createTheme } from '@shopify/restyle';

import { TextVariants } from './components/Text/TextVariants';
import { Border } from './shared/border';
import { Colors } from './shared/colors';
import { Spacing } from './shared/spacing';
import { TextBoxVariants } from './Textbox/TextboxVariants';

export const theme = createTheme({
    colors: Colors,
    spacing: Spacing,
    borderRadii: Border,
    textVariants: TextVariants,
    textBoxVariants: TextBoxVariants
});

export type Theme = typeof theme;
