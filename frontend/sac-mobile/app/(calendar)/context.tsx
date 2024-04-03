import { createContext } from "react";

export type DateContextState = {
    topViewDate: Date,
    selectedDate: Date,
}

export const DateContext = createContext<DateContextState>({
    topViewDate: new Date(),
    selectedDate: new Date()
});