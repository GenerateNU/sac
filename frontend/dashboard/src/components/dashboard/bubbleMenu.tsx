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
        <BubbleMenu className="flex items-center flex-row bg-gray-300 rounded-md space-x-1 p-2" editor={editor} tippyOptions={{ duration: 100 }}>
            <div className={editor.isActive('bold') ? 'bg-gray-400 rounded-md p-1.5' : 'p-1.5'}
                onClick={() => editor.chain().focus().toggleBold().run()}>
                <Bold size={20} />
            </div>
            <div className={editor.isActive('italic') ? 'bg-gray-400 rounded-md p-1.5' : 'p-1.5'}
                onClick={() => editor.chain().focus().toggleItalic().run()}>
                <Italic size={20}/>
            </div>
            <div className={editor.isActive('heading', { level: 1 }) ? 'bg-gray-400 rounded-md p-1.5' : 'p-1.5'}
                onClick={() => editor.chain().focus().toggleHeading({ level: 1 }).run()}>
                <Heading1 size={20} />
            </div>
            <div className={editor.isActive('heading', { level: 2 }) ? 'bg-gray-400 rounded-md p-1.5' : 'p-1.5'}
                onClick={() => editor.chain().focus().toggleHeading({ level: 2 }).run()}>
                <Heading2 size={20}/>
            </div>
            <div className={editor.isActive('bulletList') ? 'bg-gray-400 rounded-md p-1.5' : 'p-1.5'}
                onClick={() => editor.chain().focus().toggleBulletList().run()}>
                <ListIcon size={20}/>
            </div>
        </BubbleMenu>
    )
}