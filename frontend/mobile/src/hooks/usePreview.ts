import { useEffect, useRef } from 'react';

import BottomSheet from '@gorhom/bottom-sheet';

import { setClubShouldPreview } from '../store/slices/clubSlice';
import { setEventShouldPreview } from '../store/slices/eventSlice';
import { useAppDispatch, useAppSelector } from '../store/store';
import useClub from './useClub';
import useEvent from './useEvent';

const usePreview = () => {
    const dispatch = useAppDispatch();
    const eventPreviewRef = useRef<BottomSheet>(null);
    const clubPreviewRef = useRef<BottomSheet>(null);
    const { shouldPreview: eventShouldPreview, id: eventId } = useAppSelector(
        (state) => state.event
    );
    const { shouldPreview: clubShouldPreview, id: clubId } = useAppSelector(
        (state) => state.club
    );
    const { apiLoading: eventApiLoading, apiError: eventApiError } =
        useEvent(eventId);
    const { apiLoading: clubApiLoading, apiError: clubApiError } =
        useClub(clubId);

    useEffect(() => {
        dispatch(setEventShouldPreview(false));
        dispatch(setClubShouldPreview(false));
    }, [dispatch]);

    useEffect(() => {
        if (eventShouldPreview) {
            eventPreviewRef.current?.snapToIndex(0);
        } else {
            eventPreviewRef.current?.close();
        }
        if (clubShouldPreview) {
            clubPreviewRef.current?.snapToIndex(0);
        } else {
            clubPreviewRef.current?.close();
        }
    }, [eventShouldPreview, clubShouldPreview]);

    return {
        eventPreviewRef,
        clubPreviewRef,
        eventId,
        clubId,
        eventApiLoading,
        eventApiError,
        clubApiLoading,
        clubApiError
    };
};

export default usePreview;
