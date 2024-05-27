import React, { PropsWithChildren } from 'react';
import { ViewProps } from 'react-native';

import { BoxProps, createBox } from '@shopify/restyle';

import { Theme } from './theme';

type Props = ViewProps & PropsWithChildren & BoxProps<Theme>;

const BoxDefault = createBox<Theme>();

export const Box: React.FC<Props> = ({ style, ...rest }) => {
    return <BoxDefault style={[style]} {...rest} />;
};
