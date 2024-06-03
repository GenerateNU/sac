import { createBox } from "@shopify/restyle";
import { createStyles } from "../../theme";
import { View } from "react-native";

type GlobalLayoutProps = {
    children: React.ReactNode;
}

const ViewBase = createBox(View);

export const GlobalLayout = ({ children }: GlobalLayoutProps) => {
    return (
        <ViewBase flex={1} paddingHorizontal={'m'}>
            {children}
        </ViewBase>
    )
}
