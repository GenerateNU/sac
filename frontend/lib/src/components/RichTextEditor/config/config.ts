import { BlockNoteSchema, defaultBlockSpecs } from "@blocknote/core";
import { Divider } from "../elements/Divider";
import { Theme } from "@blocknote/react";

export type EditorFonts = 'monospace' | 'timesNewRoman' | 'comicSans';

// Configuration of custom and default blocks:
export const schema = BlockNoteSchema.create({
  blockSpecs: {
    ...defaultBlockSpecs,

    // Custom Blocks:
    divider: Divider,
  },
});

export type BlockNoteSchemaCustom = typeof schema.BlockNoteEditor;

export const editorFonts: Record<EditorFonts, Theme> = {
    monospace: {
        fontFamily: "monospace",
    } satisfies Theme,
    timesNewRoman: {
        fontFamily: "Times New Roman",
    } satisfies Theme,
    comicSans: {
        fontFamily: "Comic Sans MS",
    } satisfies Theme,
}
