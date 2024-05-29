import { AllProps, createTheme } from '@shopify/restyle';

import { TextVariants } from './components/Text/TextVariants';
import { Border } from './shared/border';
import { Colors } from './shared/colors';
import { Spacing } from './shared/spacing';

export const theme = createTheme({
    colors: Colors,
    spacing: Spacing,
    borderRadii: Border,
    textVariants: TextVariants
});

export const createStyles = <Keys extends string>(
    styles: Record<Keys, AllProps<Theme>>
): Record<Keys, AllProps<Theme>> => {
    return styles;
};

export type Theme = typeof theme;
