import { View } from 'react-native';

import { createBox } from '@shopify/restyle';

type GlobalLayoutProps = {
    children: React.ReactNode;
};

const ViewBase = createBox(View);

export const GlobalLayout = ({ children }: GlobalLayoutProps) => {
    return (
        <ViewBase marginHorizontal={'l'}>
            {children}
        </ViewBase>
    );
};
