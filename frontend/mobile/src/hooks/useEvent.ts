import { useEffect, useState } from 'react';

import { clubApi, eventApi } from '@generatesac/lib';

import { setClub } from '../store/slices/clubSlice';
import { setEvent, setEventTags } from '../store/slices/eventSlice';
import { useAppDispatch, useAppSelector } from '../store/store';

const useEvent = (id: string) => {
    const [retriggerFetch, setRetriggerFetch] = useState(false);

    const [getEvent, { isLoading: eventLoading, error: eventError }] =
        eventApi.useLazyEventQuery();
    const [getClub, { isLoading: clubLoading, error: clubError }] =
        clubApi.useLazyClubQuery();
    const [getEventTags, { isLoading: tagsLoading, error: tagsError }] =
        eventApi.useLazyEventTagsQuery();

    const dispatch = useAppDispatch();
    const { shouldPreview, id: eventId } = useAppSelector(
        (state) => state.event
    );

    const apiLoading = eventLoading || clubLoading || tagsLoading;
    const apiError = eventError || clubError || tagsError;

    useEffect(() => {
        if ((id !== '' && (id !== eventId || shouldPreview)) || apiError) {
            // Fetch events
            getEvent(id).then(({ data: eventData }) => {
                dispatch(setEvent(eventData));
                if (eventData) {
                    // Fetch club
                    getClub(eventData.host as string).then(
                        ({ data: clubData }) => {
                            if (clubData) {
                                dispatch(setClub(clubData));
                            }
                        }
                    );
                }
            });

            getEventTags(id).then(({ data: tagData }) => {
                if (tagData) {
                    dispatch(setEventTags(tagData));
                }
            });
        }
    }, [
        retriggerFetch,
        id,
        getClub,
        getEvent,
        getEventTags,
        dispatch,
        eventId,
        shouldPreview,
        apiError
    ]);

    return {
        setRetriggerFetch,
        apiLoading,
        apiError
    };
};

export default useEvent;
