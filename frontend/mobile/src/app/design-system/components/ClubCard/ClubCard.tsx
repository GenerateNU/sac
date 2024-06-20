import { Tag } from '@generatesac/lib';

import ClubCardStandard from './Variants/ClubCardStandard';

export interface ClubCardProps {
    id: string;
    logo: string;
    name: string;
    tags?: Tag[];
    variant?: 'standard';
}

export const Clubcard = ({
    id,
    logo,
    name,
    variant = 'standard'
}: ClubCardProps) => {
    switch (variant) {
        case 'standard':
            return <ClubCardStandard id={id} logo={logo} name={name} />;
        default:
            return <ClubCardStandard id={id} logo={logo} name={name} />;
    }
};
