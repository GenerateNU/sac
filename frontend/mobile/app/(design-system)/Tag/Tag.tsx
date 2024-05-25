import React from 'react';

import {
    BoxProps,
    VariantProps,
    createBox,
    createRestyleComponent,
    createVariant
} from '@shopify/restyle';

import { Text } from '../Text/Text';
import { Theme } from '../theme';

const Box = createBox<Theme>();

const tagVariant = createVariant<Theme, 'tagVariants'>({
    themeKey: 'tagVariants'
});

type TagProps = VariantProps<Theme, 'tagVariants'> &
    BoxProps<Theme> & {
        children?: React.ReactNode;
    };

const TagBase = createRestyleComponent<TagProps, Theme>([tagVariant], Box);

export const Tag: React.FC<TagProps> = ({ variant, children, ...rest }) => {
    return (
        <TagBase variant={variant} {...rest}>
            {typeof children === 'string' ? (
                <Text variant="body-1">{children}</Text>
            ) : (
                children
            )}
        </TagBase>
    );
};
