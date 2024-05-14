import { Input } from "./input";
import { Label } from "./label";

interface InputProps {
    type: string;
    id: string;
    label: string;
    placeholder: string;
    value: string;
    onChange: () => void;
}

const InputWithLabel = (props: InputProps) => {
    return (
        <div className="grid w-full items-center gap-1.5">
            <Label htmlFor={props.id}>{props.label}</Label>
            <Input value={props.value} onChange={props.onChange} type={props.type} id={props.id} placeholder={props.placeholder} />
        </div>
    )
}

export {InputWithLabel}
