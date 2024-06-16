import { clubApi, eventApi } from "@generatesac/lib";
import { useEffect, useState } from "react";
import { useAppDispatch } from "../store/store";
import { setClub, setClubEvents, setClubTags } from "../store/slices/clubSlice";

const useClub = (clubId: string) => {
    const [retriggerFetch, setRetriggerFetch] = useState(false);

    const [getClub, { isLoading: clubLoading, error: clubError }] = clubApi.useLazyClubQuery();
    const [getClubTags, { isLoading: tagsLoading, error: tagsError }] = clubApi.useLazyClubTagsQuery();
    const [getEvents, { isLoading: eventsLoading, error: eventsError }] = clubApi.useLazyClubEventsQuery();

    const dispatch = useAppDispatch();

    useEffect(() => {
        if (clubId !== "") {
            getClub(clubId)
                .then(({ data: clubData }) => {
                if (clubData) {
                    dispatch(setClub(clubData));
                }
                });

            getClubTags(clubId)
                .then(({ data: tagData }) => {
                    if (tagData) {
                        dispatch(setClubTags(tagData));
                    }
                })

            getEvents({ queryParams: { page: 1, limit: 5 }, id: clubId })
                .then(({ data: eventData }) => {
                    if (eventData) {
                        dispatch(setClubEvents(eventData));
                    }
                });
        }
    }, [retriggerFetch, clubId, getClub, getClubTags]);

    const apiLoading = clubLoading || tagsLoading || eventsLoading;
    const apiError = clubError || tagsError || eventsError;

    return {
        setRetriggerFetch,
        apiLoading,
        apiError,
    }
}

export default useClub;