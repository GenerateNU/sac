
import { API_BASE_URL, baseApi } from "./base";
import { CalendarType } from "../types/calendar";

const CALENDAR_API_BASE_URL = "/calendar";

export function calendarAuthRedirect(calendarType: CalendarType) {
    return `${API_BASE_URL}${CALENDAR_API_BASE_URL}/authorize?type=${calendarType}`;
}

export const calendarApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    authorize: builder.query<undefined, CalendarType>({
      query: (calendarType) => ({
        url: `${CALENDAR_API_BASE_URL}/authorize?type=${calendarType}`,
        method: "GET",
      }),
    }),
  }),
});
