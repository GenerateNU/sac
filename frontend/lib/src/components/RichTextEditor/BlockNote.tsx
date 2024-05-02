'use client';

import { Dispatch, SetStateAction } from 'react';

import { PartialBlock } from '@blocknote/core';
import { BlockNoteView, useCreateBlockNote } from '@blocknote/react';

import { EditorFonts, editorFonts, schema } from './config/config';
import SideBarMenu from './elements/SideBarMenu';
import SuggestionMenu from './elements/SuggestionMenu';
import './styles/BlockNote.styles.css';

type EditorProps = {
    description: PartialBlock[];
    setDescription: Dispatch<SetStateAction<Record<string, string>>>;
    editable: boolean;
    fontFamily: EditorFonts;
};

export default function Editor({
    description,
    setDescription,
    editable,
    fontFamily
}: EditorProps) {
    // Configuration for initial content:
    const editorConfig =
        description.length !== 0
            ? {
                  initialContent: description,
                  schema: schema
              }
            : {
                  schema: schema
              };

    // Instantiate the editor:
    const editor = useCreateBlockNote(editorConfig);

    // Save the content to state on any change:
    function onChange() {
        console.log(editor.document);
        const content: any = editor.document;
        setDescription(content as Record<string, string>);
    }

    return (
        <div>
            <BlockNoteView
                editor={editor}
                onChange={onChange}
                editable={editable}
                style={{
                    border: '1px solid #eaeaea',
                    minHeight: '500px',
                    borderRadius: '30px',
                    padding: '30px 0 0 0'
                }}
                theme={editorFonts[fontFamily]}
                slashMenu={false}
                sideMenu={false}
            >
                <SuggestionMenu editor={editor} />
                <SideBarMenu />
            </BlockNoteView>
        </div>
    );
}
