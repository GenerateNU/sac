import { CommandProps, CommandsProps } from "./commands";

export type CommandItem = {
    title: string;
    command: (props: CommandProps) => void;
};

// The list of commands made available to users when interacting with the rich text editor:
export const getCommandItems = (query: string): CommandItem[] => {
    return [
      {
        title: "H1",
        command: ({ editor, range }: CommandProps) => {
          editor
            .chain()
            .focus()
            .deleteRange(range)
            .setNode("heading", { level: 1 })
            .run();
        }
      },
      {
        title: "H2",
        command: ({ editor, range }: CommandProps) => {
          editor
            .chain()
            .focus()
            .deleteRange(range)
            .setNode("heading", { level: 2 })
            .run();
        }
      },
      {
        title: "H3",
        command: ({ editor, range }: CommandProps) => {
          editor
            .chain()
            .focus()
            .deleteRange(range)
            .setNode("heading", { level: 3 })
            .run();
        }
      },
      {
        title: "bold",
        command: ({ editor, range }: CommandProps) => {
          editor.chain().focus().deleteRange(range).setMark("bold").run();
        }
      },
      {
        title: "italic",
        command: ({ editor, range }: CommandProps) => {
          editor.chain().focus().deleteRange(range).setMark("italic").run();
        }
      },
    ]
      .filter((item) => item.title.toLowerCase().startsWith(query.toLowerCase()))
      .slice(0, 10);
  };
  