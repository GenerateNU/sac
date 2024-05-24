"use client"
import { useEditor, EditorContent } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'
import { Toolbar } from './toolbar'
import Underline from '@tiptap/extension-underline';
import { Heading } from '@tiptap/extension-heading';
import { mergeAttributes } from '@tiptap/core';
import { Link } from '@tiptap/extension-link';
import {BubbleMenu} from '@tiptap/extension-bubble-menu';
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
      Heading.extend({
        levels: [1, 2],
        renderHTML({ node, HTMLAttributes }) {
          const level = this.options.levels.includes(node.attrs.level)
            ? node.attrs.level
            : this.options.levels[0];
          const classes: { [index: number]: string } = {
            1: 'text-2xl',
            2: 'text-xl',
          };
          return [
            `h${level}`,
            mergeAttributes(this.options.HTMLAttributes, HTMLAttributes, {
              class: `${classes[level]}`,
            }),
            0,
          ];
        },
      }).configure({ levels: [1, 2] }),
    ],
    content: {"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Hello World! 🌎️"}]}]},
  })  
  
  return (
    <>
      <Toolbar editor={editor}/>
      {editor && <EditorBubbleMenu editor={editor}/>}
      <EditorContent className='bg-slate-200 h-15 pl-10 p-5 rounded-md' editor={editor} />
    </>
  )
}

export default Tiptap;