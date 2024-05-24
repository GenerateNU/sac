import { Editor } from '@tiptap/react'
import {Bold, Italic, Strikethrough, 
    Heading1, Heading2, ListOrdered, 
     Undo, Redo, Unlink,
     ListIcon, UnderlineIcon} from 'lucide-react';
import { useState } from 'react';
import { Divider } from './divider';
import { HyperlinkButton } from './hyperlink';
import { Button } from "@/components/ui/button";

interface ToolbarProps {
    editor: Editor | null;
}

export const Toolbar: React.FC<ToolbarProps> = ({ editor }) => {
    const [json, setJSON] = useState("")
    
    if (!editor) {
        return null
    }

    return (
        <div className="flex flex-row bg-slate-200 p-2 rounded-md items-center mb-2 justify-between flex-wrap">
            <div className="ml-1 flex flex-row items-center flex-wrap space-x-1">
                <Button onClick={() => editor.chain().focus().toggleBold().run()}
                        disabled={!editor.isEditable}
                        variant="link"
                        size="link"
                        className={editor.isActive('bold') && editor.isEditable ? 'bg-slate-300 rounded-md p-1.5' : 'p-1.5'}>
                    <Bold />
                </Button>
                <Button disabled={!editor.isEditable}
                        variant="link"
                        size="link"
                        onClick={() => editor.chain().focus().toggleItalic().run()}
                        className={editor.isActive('italic') && editor.isEditable ? 'bg-slate-300 rounded-md p-1.5' : 'p-1.5'}>
                    <Italic />
                </Button>
                <Button variant="link"
                        size="link" disabled={!editor.isEditable} onClick={() => editor.chain().focus().toggleStrike().run()}
                        className={editor.isActive('strike') && editor.isEditable ? 'bg-slate-300 rounded-md p-1.5' : 'p-1.5'}>
                    <Strikethrough />
                </Button>
                <Button variant="link"
                        size="link" disabled={!editor.isEditable} onClick={() => editor.chain().focus().toggleUnderline().run()}
                        className={editor.isActive('underline') && editor.isEditable ? 'bg-slate-300 rounded-md p-1.5' : 'p-1.5'}>
                    <UnderlineIcon />
                </Button>
                <Divider />
                <Button variant="link"
                        size="link" disabled={!editor.isEditable} onClick={() => editor.chain().focus().toggleHeading({ level: 1 }).run()}
                        className={editor.isActive('heading', { level: 1 }) && editor.isEditable ? 'bg-slate-300 rounded-md p-1.5' : 'p-1.5'}>
                    <Heading1 />
                </Button>
                <Button variant="link"
                        size="link" disabled={!editor.isEditable} onClick={() => editor.chain().focus().toggleHeading({ level: 2 }).run()}
                        className={editor.isActive('heading', { level: 2 }) && editor.isEditable ? 'bg-slate-300 rounded-md p-1.5' : 'p-1.5'}>
                    <Heading2 />
                </Button>
                <Divider />
                <Button variant="link"
                        size="link" disabled={!editor.isEditable} onClick={() => editor.chain().focus().toggleBulletList().run()}
                        className={editor.isActive('bulletList') && editor.isEditable ? 'bg-slate-300 rounded-md p-1.5' : 'p-1.5'}>
                    <ListIcon />
                </Button>
                <Button variant="link"
                        size="link" disabled={!editor.isEditable} onClick={() => editor.chain().focus().toggleOrderedList().run()}
                        className={editor.isActive('orderedList') && editor.isEditable ? 'bg-slate-300 rounded-md p-1.5' : 'p-1.5'}>
                    <ListOrdered />
                </Button>
                <HyperlinkButton disabled={!editor.isEditable} editor={editor}/>
                <Button
                    variant="link"
                    size="link"
                    onClick={() => {
                        editor.chain().focus().unsetLink().run();
                    }}
                    disabled={!editor.isActive('link') || !editor.isEditable}
                    className="p-1.5"
                >
                    <Unlink />
                </Button>
                <Divider />
                <Button variant="link"
                        size="link" disabled={!editor.isEditable} className="p-1.5" onClick={() => editor.chain().focus().undo().run()}>
                    <Undo />
                </Button>
                <Button variant="link"
                        size="link" disabled={!editor.isEditable} className="p-1.5" onClick={() => editor.chain().focus().redo().run()}>
                    <Redo />
                </Button>
            </div>

            <Button onClick={() => {
                setJSON(JSON.stringify(editor.getJSON()));
                editor.setEditable(!editor.isEditable);
            }}
                className="px-4 py-1 bg-black rounded-md text-white text-sm">{editor.isEditable ? "Save" : "Edit"}</Button>
        </div>
    )
}