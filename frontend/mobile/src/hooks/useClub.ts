import { useEffect, useState } from 'react';

import { clubApi } from '@generatesac/lib';

import { setClub, setClubEvents, setClubTags } from '../store/slices/clubSlice';
import { useAppDispatch, useAppSelector } from '../store/store';

const useClub = (id?: string) => {
    const [retriggerFetch, setRetriggerFetch] = useState(false);

    const [getClub, { isLoading: clubLoading, error: clubError }] =
        clubApi.useLazyClubQuery();
    const [getClubTags, { isLoading: tagsLoading, error: tagsError }] =
        clubApi.useLazyClubTagsQuery();
    const [getEvents, { isLoading: eventsLoading, error: eventsError }] =
        clubApi.useLazyClubEventsQuery();

    const dispatch = useAppDispatch();
    const { id: clubId } = useAppSelector((state) => state.club);

    const apiLoading = clubLoading || tagsLoading || eventsLoading;
    const apiError = clubError || tagsError || eventsError;

    const idToFetch = id ?? clubId;

    useEffect(() => {
        if (idToFetch !== '') {
            getClub(idToFetch).then(({ data: clubData }) => {
                if (clubData) {
                    dispatch(setClub(clubData));
                }
            });

            getClubTags(clubId).then(({ data: tagData }) => {
                if (tagData) {
                    dispatch(setClubTags(tagData));
                }
            });

            getEvents({ id: idToFetch }).then(({ data: eventData }) => {
                if (eventData) {
                    const sortedEvents = [...eventData];
                    sortedEvents.sort((a, b) => {
                        return (
                            new Date(a.start_time).getTime() -
                            new Date(b.start_time).getTime()
                        );
                    });
                    const events = sortedEvents.filter(
                        (event) =>
                            new Date(event.end_time).getTime() >
                            new Date().getTime()
                    );
                    dispatch(setClubEvents(events));
                }
            });
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [retriggerFetch, idToFetch]);

    return {
        setRetriggerFetch,
        apiLoading,
        apiError
    };
};

export default useClub;
