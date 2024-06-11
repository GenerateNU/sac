"use client"
import { useEditor, EditorContent } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'
import { Toolbar } from './toolbar'
import Underline from '@tiptap/extension-underline';
import { Link } from '@tiptap/extension-link';
import { BubbleMenu } from '@tiptap/extension-bubble-menu';
import { EditorBubbleMenu } from './bubbleMenu';
import React from 'react';
import './styles.css';

const Tiptap = () => {
  const editor = useEditor({
    autofocus: true,
    editable: false,
    extensions: [
      StarterKit, Underline,
      BubbleMenu,
      Link.configure({
        openOnClick: true,
        autolink: true,
      }),
    ],
    content: {"type":"doc","content":[{"type":"heading","attrs":{"level":1},"content":[{"type":"text","text":"Hello World"}]},{"type":"heading","attrs":{"level":2},"content":[{"type":"text","marks":[{"type":"bold"},{"type":"italic"},{"type":"strike"},{"type":"underline"}],"text":"This is generate"}]},{"type":"bulletList","content":[{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"Task 1"}]}]},{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"Task 2"}]}]}]},{"type":"orderedList","attrs":{"start":1},"content":[{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"Task 3"}]}]},{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"Task 4"}]}]}]},{"type":"paragraph","content":[{"type":"text","marks":[{"type":"link","attrs":{"href":"https://generatenu.com","target":"_blank","rel":"noopener noreferrer nofollow","class":null}}],"text":"link"}]}]},
  })  
  
  return (
    <>
      <Toolbar editor={editor}/>
      {editor && <EditorBubbleMenu editor={editor}/>}
      <EditorContent className="h-15 pl-10 p-5 rounded-md bg-slate-200" editor={editor} />
    </>
  )
}

export default Tiptap;