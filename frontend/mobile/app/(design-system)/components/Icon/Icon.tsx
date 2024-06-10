import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Colors, SACColors, defaultColor } from '../../shared/colors';
import { ComponentSizes } from '../../shared/types';
import { createStyles } from '../../theme';

type IconProps = {
    icon: IconDefinition;
    size?: ComponentSizes;
    color?: SACColors;
};

export const Icon: React.FC<IconProps> = ({
    icon,
    size = 'md',
    color = defaultColor
}) => {
    return (
        <FontAwesomeIcon
            icon={icon}
            size={styles[size].fontSize}
            color={Colors[color]}
        />
    );
};

const styles = createStyles({
    xs: {
        fontSize: 12
    },
    sm: {
        fontSize: 16
    },
    md: {
        fontSize: 24
    },
    lg: {
        fontSize: 75
    },
    full: {
        fontSize: 24
    }
});
