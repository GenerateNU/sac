import {BubbleMenu} from '@tiptap/react';
import { Editor } from '@tiptap/react'
import { Bold, Heading1, Heading2, Italic, ListIcon, ListOrdered, UnderlineIcon } from 'lucide-react';

interface BubbleMenuProps {
    editor: Editor | null;
}

export const EditorBubbleMenu: React.FC<BubbleMenuProps> = ({ editor }) => {
    if (!editor) {
        return null
    }

    return (
        <BubbleMenu className="flex items-center flex-row bg-slate-300 rounded-md space-x-3 p-2" editor={editor} tippyOptions={{ duration: 100 }}>
            <div className={editor.isActive('bold') ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}
                onClick={() => editor.chain().focus().toggleBold().run()}>
                <Bold size={20} />
            </div>
            <div className={editor.isActive('italic') ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}
                onClick={() => editor.chain().focus().toggleItalic().run()}>
                <Italic size={20}/>
            </div>
            <div className={editor.isActive('heading', { level: 1 }) ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}
                onClick={() => editor.chain().focus().toggleHeading({ level: 1 }).run()}>
                <Heading1 />
            </div>
            <div className={editor.isActive('heading', { level: 2 }) ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}
                onClick={() => editor.chain().focus().toggleHeading({ level: 2 }).run()}>
                <Heading2 />
            </div>
            <div className={editor.isActive('bulletList') ? 'bg-slate-400 rounded-md p-1.5' : 'p-1.5'}
                onClick={() => editor.chain().focus().toggleBulletList().run()}>
                <ListIcon />
            </div>
        </BubbleMenu>
    )
}