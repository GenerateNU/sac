'use client'

import { useEditor, EditorContent, Editor } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'
import { useEffect, useState } from 'react';
import { Commands } from './config/commands';
import { getCommandItems } from './config/commandItems';

export function Tiptap() {
  const [editor, setEditor] = useState<Editor | null>(null);

  useEffect(() => {
    const editor = new Editor({
      extensions: [
        StarterKit,
        Commands.configure({
          suggestion: {
            items: getCommandItems,
          }
        })
      ],
      content: '<p>Hello World! 🌎️</p>',
    });

    setEditor(editor);

    return () => {
      editor.destroy();
    }
  }, [])

  return (
    <EditorContent editor={editor} />
  )
};

export default Tiptap