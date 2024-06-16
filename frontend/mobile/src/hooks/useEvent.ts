import { useEffect, useState } from "react";
import { useAppDispatch } from "../store/store";
import { clubApi, eventApi } from "@generatesac/lib";
import { setEvent, setEventTags } from "../store/slices/eventSlice";
import { setClub } from "../store/slices/clubSlice";

const useEvent = (id: string) => {
    const [retriggerFetch, setRetriggerFetch] = useState(false);
    
    const [getEvent, { isLoading: eventLoading, error: eventError }] = eventApi.useLazyEventQuery();
    const [getClub, { isLoading: clubLoading, error: clubError }] = clubApi.useLazyClubQuery();
    const [getEventTags, { isLoading: tagsLoading, error: tagsError }] = eventApi.useLazyEventTagsQuery();

    const dispatch = useAppDispatch();

    useEffect(() => {
        if (id !== "") {
            // Fetch events
            getEvent(id)
                .then(({ data: eventData }) => {
                    dispatch(setEvent(eventData));
                    if (eventData) {
                        // Fetch club
                        getClub(eventData.host as string)
                            .then(({ data: clubData }) => {
                                if (clubData) {
                                    dispatch(setClub(clubData));
                                }
                            })
                    }
                });
            
            getEventTags(id)
                .then(({ data: tagData }) => {
                    if (tagData) {
                        dispatch(setEventTags(tagData));
                    }
                })

        }
    }, [retriggerFetch, id, getClub, getEvent, getEventTags]);

    const apiLoading = eventLoading || clubLoading || tagsLoading;
    const apiError = eventError || clubError || tagsError;

    return {
        setRetriggerFetch,
        apiLoading,
        apiError,
    }
};

export default useEvent;