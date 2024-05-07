import { MdDelete } from 'react-icons/md';

import {
    SideMenuButton,
    SideMenuProps,
    useBlockNoteEditor
} from '@blocknote/react';

// Custom Side Menu button to remove the hovered block.
export default function RemoveBlockButton(props: SideMenuProps) {
    const editor = useBlockNoteEditor();

    return (
        <SideMenuButton>
            <MdDelete
                size={24}
                onClick={() => {
                    editor.removeBlocks([props.block]);
                }}
            />
        </SideMenuButton>
    );
}
