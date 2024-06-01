export interface Option {
    value: string;
    label: string;
    disable?: boolean;
    fixed?: boolean;
    [key: string]: string | boolean | undefined;
}
