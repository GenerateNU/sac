import { Text } from '@/app/(design-system)';

interface WordmarkProps {
    color?: 'black' | 'white';
}

export const Wordmark: React.FC<WordmarkProps> = ({ color = 'black' }) => {
    return (
        <Text color={color} fontWeight="500" variant="subheader-1">
            Wordmark
        </Text>
    );
};
