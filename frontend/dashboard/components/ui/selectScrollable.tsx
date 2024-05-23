import { Option } from "@/types/option";
import { Label } from "./label";

import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectTrigger,
    SelectValue,
  } from "@/components/ui/select";
  
interface ScrollableProps {
    data: Option[];
    label: string;
    placeholder: string;
    value: string;
    onChange: () => void;
    onBlur: (open: boolean) => void;
}

const SelectScrollable = ({label, value, placeholder, data, onBlur, onChange}: ScrollableProps) => {
    return (
        <div className="grid w-full items-center gap-1.5">
            <Label htmlFor="timezone">{label}</Label>
            <Select value={value} onOpenChange={onBlur} onValueChange={onChange}>
                <SelectTrigger>
                    <SelectValue placeholder={placeholder} />
                </SelectTrigger>
                <SelectContent>
                    <SelectGroup>
                        {data.map((item: Option) => (
                            <SelectItem key={item.value} value={item.value}>{item.label}</SelectItem>
                        ))}
                    </SelectGroup>
                </SelectContent>
            </Select>
        </div>
    )
}

export { SelectScrollable }