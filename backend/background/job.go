package background

import "log/slog"

type JobFunc func()

func Go(fn JobFunc) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("panic in background job", r)
			}
		}()

		fn()
	}()
}
