import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { SACColors } from '../Colors';
import { ComponentSizes } from '../shared/types';
import { createStyles } from '../theme';

type IconProps = {
    icon: IconDefinition;
    size?: ComponentSizes;
    color?: SACColors;
};

export const Icon: React.FC<IconProps> = ({ icon, size = 'medium', color }) => {
    return (
        <FontAwesomeIcon
            icon={icon}
            size={styles[size].fontSize}
            color={color}
        />
    );
};

const styles = createStyles({
    small: {
        fontSize: 16
    },
    medium: {
        fontSize: 24
    },
    full: {
        fontSize: 24
    }
});
