package errors

import "errors"

var (
	RequestBodyIsMalformed         = errors.New("error.request_body_is_malformed")
	ItemCountIsEqualOrLessThanZero = errors.New("error.item_count_is_equal_or_less_than_zero")
	InternalError                  = errors.New("error.internal_error")
)

func New(text string) error {
	return errors.New(text)
}
