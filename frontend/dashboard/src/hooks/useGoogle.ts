import { calendarAuthRedirect } from "@generatesac/lib";

export default function useGoogle() {
    function googleAuthorize() {
        window.location.href = calendarAuthRedirect('google');
    }

    return { googleAuthorize }
}
