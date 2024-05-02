import { LuHeading } from 'react-icons/lu';

import { filterSuggestionItems } from '@blocknote/core';
import {
    SuggestionMenuController,
    getDefaultReactSlashMenuItems
} from '@blocknote/react';

import { BlockNoteSchemaCustom } from '../config/config';
import { insertDivider } from './Divider';

type SuggestionMenuProps = {
    editor: BlockNoteSchemaCustom;
};

export default function SuggestionMenu({ editor }: SuggestionMenuProps) {
    /* Used to remove the image option from the suggestion editor */
    const suggestionItems = (editor: BlockNoteSchemaCustom) => {
        const itemsToRemove = [
            'Image',
            'Heading 1',
            'Heading 2',
            'Paragraph',
            'Table'
        ];
        const filteredSlashMenuItems = getDefaultReactSlashMenuItems(
            editor
        ).filter((item) => {
            // Unset item fields:
            item.badge = '';
            item.subtext = '';

            // Update the header item:
            if (item.title === 'Heading 3') {
                item.title = 'Heading';
                item.icon = <LuHeading size={18} />;
            }

            return !itemsToRemove.includes(item.title);
        });

        return [...filteredSlashMenuItems, insertDivider(editor)];
    };

    return (
        <SuggestionMenuController
            triggerCharacter={'/'}
            getItems={async (query) =>
                filterSuggestionItems(suggestionItems(editor), query)
            }
        />
    );
}
