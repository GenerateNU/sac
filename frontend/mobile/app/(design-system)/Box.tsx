import React, { PropsWithChildren } from 'react';

import { BoxProps, createBox } from '@shopify/restyle';

import { Theme } from './theme';
import { ViewProps } from 'react-native';

type Props = ViewProps & PropsWithChildren & BoxProps<Theme>;

const BoxDefault = createBox<Theme>();

export const Box: React.FC<Props> = ({ style, ...rest }) => {
    return <BoxDefault style={[style]} {...rest} />;
};
