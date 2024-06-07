import { StyleSheet } from 'react-native';
import { CodeField, Cursor } from 'react-native-confirmation-code-field';

import { Colors, Spacing, Text } from '@/app/(design-system)';
import { Border } from '@/app/(design-system)/shared/border';

interface EmailFormProps {
    onComplete: (text: string) => void;
    error: string;
    value: string;
}

export const EmailVerificationForm: React.FC<EmailFormProps> = ({
    onComplete,
    value,
    error
}) => {
    return (
        <>
            {error !== '' ? <Text color="red">{error}</Text> : null}
            <CodeField
                value={value}
                onChangeText={onComplete}
                cellCount={6}
                autoFocus
                rootStyle={styles.codeFieldRoot}
                keyboardType="number-pad"
                renderCell={({ index, symbol, isFocused }) => (
                    <Text
                        key={index}
                        style={[
                            styles.cell,
                            { borderColor: error !== '' ? 'red' : 'gray' },
                            isFocused && styles.focusCell
                        ]}
                    >
                        {symbol || (isFocused ? <Cursor /> : null)}
                    </Text>
                )}
            />
        </>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1
    },
    codeFieldRoot: {
        width: '80%',
        marginLeft: 'auto',
        marginRight: 'auto',
        marginTop: Spacing.l,
        alignItems: 'center',
        justifyContent: 'center',
        gap: 5
    },
    cell: {
        width: 50,
        height: 50,
        lineHeight: 46,
        fontSize: 24,
        borderWidth: 1,
        textAlign: 'center',
        borderRadius: Border.base
    },
    focusCell: {
        borderColor: Colors.black
    }
});

export default EmailVerificationForm;
