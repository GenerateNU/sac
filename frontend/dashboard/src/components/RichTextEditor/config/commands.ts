import { Editor, Extension, Range } from "@tiptap/react";
import Suggestion from "@tiptap/suggestion";
import { CommandItem } from "./commandItems";

export type CommandsProps = {
  editor: Editor,
  range: Range,
  props: CommandItem,
}

export type CommandProps = {
  editor: Editor,
  range: Range,
}

export const Commands = Extension.create({
  name: 'commands',

  addOptions() {
    return {
      suggestion: {
        char: '/',
        command: ({ editor, range, props }: CommandsProps) => {
          props.command({ editor, range });
        },
      },
    }
  },

  addProseMirrorPlugins() {
    return [
      Suggestion({
        editor: this.editor,
        ...this.options.suggestion,
      })
    ]
  },
});
