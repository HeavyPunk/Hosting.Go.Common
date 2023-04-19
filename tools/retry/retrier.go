package tools_retry

import (
	"errors"
	"time"
)

func Retry[TRes any](
	action func() (TRes, error),
	isSuccess func(TRes, error) bool,
	maxAttempt uint,
	delay time.Duration,
	onError func(TRes, error),
) (TRes, error) {
	res, err := action()
	for i := 0; i < int(maxAttempt); i++ {
		if isSuccess(res, err) {
			return res, err
		}
		if onError != nil && err != nil {
			onError(res, err)
		}
		time.Sleep(delay)
		res, err = action()
	}
	return res, errors.New("action spent all attempts")
}
