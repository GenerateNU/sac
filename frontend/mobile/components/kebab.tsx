import { MaterialCommunityIcons } from '@expo/vector-icons';
import {
    MenuAction,
    MenuView,
    NativeActionEvent
} from '@react-native-menu/menu';

type KebabProps = {
    items: MenuAction[];
    onPressAction: (event: NativeActionEvent) => void;
};

const Kebab = ({ items, onPressAction }: KebabProps) => {
    return (
        <MenuView onPressAction={onPressAction} actions={items}>
            <MaterialCommunityIcons
                name="dots-vertical"
                size={24}
                color="white"
            />
        </MenuView>
    );
};

export { Kebab };
