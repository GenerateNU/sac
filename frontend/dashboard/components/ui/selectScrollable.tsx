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
}

const SelectScrollable = (props : ScrollableProps) => {
    return (
        <div className="grid w-full items-center gap-1.5">
            <Label htmlFor="timezone">{props.label}</Label>
            <Select value={props.value} onValueChange={props.onChange}>
                <SelectTrigger>
                    <SelectValue placeholder={props.placeholder} />
                </SelectTrigger>
                <SelectContent>
                    <SelectGroup>
                        {props.data.map((item: Option) => (
                            <SelectItem key={item.value} value={item.value}>{item.label}</SelectItem>
                        ))}
                    </SelectGroup>
                </SelectContent>
            </Select>
        </div>
    )
}

export { SelectScrollable }