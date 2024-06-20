import { useEffect, useState } from 'react';

import { eventApi } from '@generatesac/lib';

import { setClubId } from '../store/slices/clubSlice';
import { setEvent, setEventTags } from '../store/slices/eventSlice';
import { useAppDispatch, useAppSelector } from '../store/store';

const useEvent = () => {
    const [retriggerFetch, setRetriggerFetch] = useState(false);

    const [getEvent, { isLoading: eventLoading, error: eventError }] =
        eventApi.useLazyEventQuery();
    const [getEventTags, { isLoading: tagsLoading, error: tagsError }] =
        eventApi.useLazyEventTagsQuery();

    const dispatch = useAppDispatch();
    const { id: eventId } = useAppSelector((state) => state.event);

    const apiLoading = eventLoading || tagsLoading;
    const apiError = eventError || tagsError;

    useEffect(() => {
        if (eventId !== '') {
            // Fetch events
            getEvent(eventId).then(({ data: eventData }) => {
                if (eventData) {
                    console.log(eventData);
                    dispatch(setEvent(eventData));
                    dispatch(setClubId(eventData?.host));
                }
            });

            getEventTags(eventId).then(({ data: tagData }) => {
                if (tagData) {
                    dispatch(setEventTags(tagData));
                }
            });
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [retriggerFetch, eventId]);

    return {
        setRetriggerFetch,
        apiLoading,
        apiError
    };
};

export default useEvent;
