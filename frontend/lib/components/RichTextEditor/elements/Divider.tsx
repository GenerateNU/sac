import { createReactBlockSpec } from "@blocknote/react";
import { BlockNoteSchemaCustom } from "../config/config";
import { insertOrUpdateBlock } from "@blocknote/core";
import { FaDivide } from "react-icons/fa6";

export const Divider = createReactBlockSpec(
  {
    type: "divider",
    content: "none",
    propSchema: {},
  },
  {
    render: (props) => {
      return (
        <hr 
          className={"divider"} 
          data-alert-type={props.block.type}
          style={{ margin: "10px 0 10px 0"}}
        />
      );
    },
  }
);

export const insertDivider = (editor: BlockNoteSchemaCustom) => ({
  title: "Divider",
  onItemClick: () => {
    insertOrUpdateBlock(editor, {
      type: "divider",
    });
  },
  aliases: [
    "divider",
    "line",
  ],
  group: 'Other',
  icon: <FaDivide />,
});

