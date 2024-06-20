import { useEffect, useState } from 'react';

import { eventApi } from '@generatesac/lib';

import { setClubId } from '../store/slices/clubSlice';
import {
    setEvent,
    setEventAttendees,
    setEventTags
} from '../store/slices/eventSlice';
import { useAppDispatch, useAppSelector } from '../store/store';

const useEvent = () => {
    const [retriggerFetch, setRetriggerFetch] = useState(false);

    const [getEvent, { isLoading: eventLoading, error: eventError }] =
        eventApi.useLazyEventQuery();
    const [getEventTags, { isLoading: tagsLoading, error: tagsError }] =
        eventApi.useLazyEventTagsQuery();
    const [
        getEventAttendees,
        { isLoading: attendeesLoading, error: attendeesError }
    ] = eventApi.useLazyEventRegistrationsQuery();

    const dispatch = useAppDispatch();
    const { id: eventId } = useAppSelector((state) => state.event);

    const apiLoading = eventLoading || tagsLoading || attendeesLoading;
    const apiError = eventError || tagsError || attendeesError;

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

            getEventAttendees(eventId).then(({ data: attendeeData }) => {
                if (attendeeData) {
                    const attendeeIds = attendeeData.map(
                        (attendeeData) => attendeeData.id
                    );
                    dispatch(setEventAttendees(attendeeIds));
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
