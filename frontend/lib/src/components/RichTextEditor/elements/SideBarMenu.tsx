import {
    DragHandleButton,
    SideMenu,
    SideMenuController
} from '@blocknote/react';

import { RemoveBlockButton } from './RemoveBlockButton';

export default function SideBarMenu() {
    return (
        <SideMenuController
            sideMenu={(props) => (
                <SideMenu {...props}>
                    <RemoveBlockButton {...props} />
                    <DragHandleButton {...props} />
                </SideMenu>
            )}
        />
    );
}
