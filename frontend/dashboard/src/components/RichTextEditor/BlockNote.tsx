"use client"; // this registers <Editor> as a Client Component
import { BlockNoteView, useCreateBlockNote } from "@blocknote/react";
import "@blocknote/core/fonts/inter.css";
import "@blocknote/react/style.css";
import "@/src/components/RichTextEditor/styles/BlockNote.styles.css"
import { PartialBlock } from "@blocknote/core";
import { Dispatch, SetStateAction } from "react";
import SuggestionMenu from "./elements/SuggestionMenu";
import SideBarMenu from "./elements/SideBarMenu";
import { EditorFonts, editorFonts, schema } from "./config/config";

type EditorProps = {
  description: PartialBlock[];
  setDescription: Dispatch<SetStateAction<Record<string, string>>>;
  editable: boolean;
  fontFamily: EditorFonts;
};
 
export default function Editor({ description, setDescription, editable, fontFamily }: EditorProps) {
  // Configuration for initial content:
  const editorConfig = description.length !== 0 ? {
    initialContent: description,
    schema: schema,
  }: {
    schema: schema,
  };

  // Instantiate the editor:
  const editor = useCreateBlockNote(editorConfig);

  // Save the content to state on any change:
  function onChange() {
    console.log(editor.document)
    const content: any = editor.document;
    setDescription(content as Record<string, string>);
  }
 
  return (
    <div>
      <BlockNoteView 
        editor={editor} 
        onChange={onChange}
        editable={editable}
        style={{
          border: '1px solid #eaeaea',
          minHeight: '500px',
          borderRadius: '30px',
          padding: '30px 0 0 0',
        }}
        theme={editorFonts[fontFamily]}
        slashMenu={false}
        sideMenu={false}
      >
        <SuggestionMenu editor={editor}/>
        <SideBarMenu />
      </BlockNoteView>
    </div>
  );
}