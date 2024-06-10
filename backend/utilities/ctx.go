package utilities

import (
	"context"
)

func ExecuteWithTimeout(ctx context.Context, fn func() error) error {
	errChan := make(chan error)

	go func() {
		err := fn()
		errChan <- err
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

func ExecuteWithTimeoutResult[Result any](ctx context.Context, fn func() (*Result, error)) (*Result, error) {
	resultChan := make(chan *Result)
	errChan := make(chan error)

	go func() {
		result, err := fn()
		if err != nil {
			errChan <- err
		} else {
			resultChan <- result
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errChan:
		return nil, err
	case result := <-resultChan:
		return result, nil
	}
}
