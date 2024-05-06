import {
    BasicTextStyleButton,
    BlockTypeSelect,
    CreateLinkButton,
    FormattingToolbar as Toolbar,
    FormattingToolbarController,
    ImageCaptionButton,
    ReplaceImageButton,
    blockTypeSelectItems,
  } from "@blocknote/react";
import { LuHeading } from "react-icons/lu";

export default function FormattingToolbar() {
    const suggestionItems = () => {
        const itemsToRemove = [
            'Image',
            'Heading 1',
            'Heading 2',
            'Table'
        ];
        const filteredSlashMenuItems = blockTypeSelectItems.filter((item) => {
            if (item.name === 'Heading 3') {
                item.name = 'Heading';
                item.icon = LuHeading;
            }

            return !itemsToRemove.includes(item.name);
        });

        return filteredSlashMenuItems;
    };

    return (
        <FormattingToolbarController
            formattingToolbar={() => (
                <Toolbar>
                    <BlockTypeSelect key={"blockTypeSelect"} items={suggestionItems()} />

                    <ImageCaptionButton key={"imageCaptionButton"} />
                    <ReplaceImageButton key={"replaceImageButton"} />

                    <BasicTextStyleButton
                        basicTextStyle={"bold"}
                        key={"boldStyleButton"}
                    />
                    <BasicTextStyleButton
                        basicTextStyle={"italic"}
                        key={"italicStyleButton"}
                    />
                    <BasicTextStyleButton
                        basicTextStyle={"underline"}
                        key={"underlineStyleButton"}
                    />
                    <BasicTextStyleButton
                        basicTextStyle={"strike"}
                        key={"strikeStyleButton"}
                    />
                    {/* Extra button to toggle code styles */}
                    <BasicTextStyleButton
                        key={"codeStyleButton"}
                        basicTextStyle={"code"}
                    />

                    <CreateLinkButton key={"createLinkButton"} />
                </Toolbar>
            )}
        />
    );
}
