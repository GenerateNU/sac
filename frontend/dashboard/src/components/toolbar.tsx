import { Editor } from '@tiptap/react'
import {Bold, Italic, Strikethrough, 
    Heading1, Heading2, ListOrdered, 
     Undo, Redo, Unlink,
     ListIcon, UnderlineIcon} from 'lucide-react';
import { useState } from 'react';
import { Divider } from './divider';
import { HyperlinkButton } from './hyperlink';

interface ToolbarProps {
    editor: Editor | null;
}

export const Toolbar: React.FC<ToolbarProps> = ({ editor }) => {
    if (!editor) {
        return null
    }

    const [editable, setEditable] = useState(editor.isEditable)
    const [json, setJSON] = useState("")

    return (
        <div className="flex flex-row bg-gray-300 p-2 rounded-md items-center mb-2 justify-between flex-wrap">
            <div className="ml-1 flex flex-row items-center flex-wrap space-x-1">
                <button onClick={() => editor.chain().focus().toggleBold().run()}
                        className={editor.isActive('bold') ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}>
                    <Bold size={20} />
                </button>
                <button onClick={() => editor.chain().focus().toggleItalic().run()}
                        className={editor.isActive('italic') ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}>
                    <Italic size={20}/>
                </button>
                <button onClick={() => editor.chain().focus().toggleStrike().run()}
                        className={editor.isActive('strike') ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}>
                    <Strikethrough size={20} />
                </button>
                <button onClick={() => editor.chain().focus().toggleUnderline().run()}
                        className={editor.isActive('underline') ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}>
                    <UnderlineIcon size={20}/>
                </button>
                <Divider />
                <button onClick={() => editor.chain().focus().toggleHeading({ level: 1 }).run()}
                        className={editor.isActive('heading', { level: 1 }) ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}>
                    <Heading1 />
                </button>
                <button onClick={() => editor.chain().focus().toggleHeading({ level: 2 }).run()}
                        className={editor.isActive('heading', { level: 2 }) ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}>
                    <Heading2 />
                </button>
                <Divider />
                <button onClick={() => editor.chain().focus().toggleBulletList().run()}
                        className={editor.isActive('bulletList') ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}>
                    <ListIcon />
                </button>
                <button onClick={() => editor.chain().focus().toggleOrderedList().run()}
                        className={editor.isActive('orderedList') ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}>
                    <ListOrdered />
                </button>
                <HyperlinkButton editor={editor}/>
                <button
                    onClick={() => editor.chain().focus().unsetLink().run()}
                    disabled={!editor.isActive('link')}
                    className="p-1.5"
                >
                    <Unlink size={20}/>
                </button>
                <Divider />
                <button className="p-1.5" onClick={() => editor.chain().focus().undo().run()}>
                    <Undo />
                </button>
                <button className="p-1.5" onClick={() => editor.chain().focus().redo().run()}>
                    <Redo />
                </button>
            </div>

            <button onClick={() => {
                setEditable(!editable)
                setJSON(JSON.stringify(editor.getJSON()));
                editor.setEditable(!editable);
            }}
                className="px-4 py-1 bg-black rounded-md text-white text-sm mr-2">{editable ? "Save" : "Edit"}</button>
        </div>
    )
}