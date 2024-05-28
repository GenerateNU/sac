import { Input } from "./input";
import { Label } from "./label";

interface InputProps {
    type: string;
    id: string;
    label: string;
    placeholder: string;
    value: string;
    onChange: (event: React.FocusEvent<HTMLInputElement>) => void;
    onBlur: (event: React.FocusEvent<HTMLInputElement>) => void;
}

const InputWithLabel = ({ type, id, label, placeholder, value, onChange, onBlur }: InputProps) => {
    return (
        <div className="grid w-full items-center gap-1.5">
            <Label htmlFor={id}>{label}</Label>
            <Input onBlur={onBlur} value={value} onChange={onChange} type={type} id={id} placeholder={placeholder} />
        </div>
    )
}

export {InputWithLabel}
