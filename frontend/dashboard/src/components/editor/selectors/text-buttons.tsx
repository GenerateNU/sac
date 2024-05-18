import { cn } from "@/lib/utils";
import { EditorBubbleItem, useEditor } from "novel";
import {
  BoldIcon,
  ItalicIcon,
  UnderlineIcon,
  StrikethroughIcon,
  CodeIcon,
  AlignLeft,
  AlignRight,
  AlignCenter,
  AlignJustify,
} from "lucide-react";
import type { SelectorItem } from "./node-selector";
import { Button } from "@/components/ui/button";

export const TextButtons = () => {
  const { editor } = useEditor();
  if (!editor) return null;

  const items: SelectorItem[] = [
    {
      name: "bold",
      isActive: (editor) => editor.isActive("bold"),
      command: (editor) => editor.chain().focus().toggleBold().run(),
      icon: BoldIcon,
    },
    {
      name: "italic",
      isActive: (editor) => editor.isActive("italic"),
      command: (editor) => editor.chain().focus().toggleItalic().run(),
      icon: ItalicIcon,
    },
    {
      name: "underline",
      isActive: (editor) => editor.isActive("underline"),
      command: (editor) => editor.chain().focus().toggleUnderline().run(),
      icon: UnderlineIcon,
    },
    {
      name: "strike",
      isActive: (editor) => editor.isActive("strike"),
      command: (editor) => editor.chain().focus().toggleStrike().run(),
      icon: StrikethroughIcon,
    },
    {
      name: "code",
      isActive: (editor) => editor.isActive("code"),
      command: (editor) => editor.chain().focus().toggleCode().run(),
      icon: CodeIcon,
    },
    {
      name: "align left",
      isActive: (editor) => editor.isActive("left"),
      command: (editor) => editor.chain().focus().setTextAlign('left').run(),
      icon: AlignLeft,
    },
    {
      name: "align right",
      isActive: (editor) => editor.isActive("right"),
      command: (editor) => editor.chain().focus().setTextAlign('right').run(),
      icon: AlignRight,
    },
    {
      name: "align center",
      isActive: (editor) => editor.isActive("center"),
      command: (editor) => editor.chain().focus().setTextAlign('center').run(),
      icon: AlignCenter,
    },
    {
      name: "align justify",
      isActive: (editor) => editor.isActive("justify"),
      command: (editor) => editor.chain().focus().setTextAlign('justify').run(),
      icon: AlignJustify,
    },
  ];

  return (
    <div className="flex">
      {items.map((item, index) => (
        <EditorBubbleItem
          key={index}
          onSelect={(editor) => {
            item.command(editor);
          }}
        >
          <Button size="sm" className="rounded-none" variant="ghost">
            <item.icon
              className={cn("h-4 w-4", {
                "text-blue-500": item.isActive(editor),
              })}
            />
          </Button>
        </EditorBubbleItem>
      ))}
    </div>
  );
};
