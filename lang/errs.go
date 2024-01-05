package lang

import (
	"bytes"
	"fmt"
)

func DefaultErrorFn(err error, defErrFn func() error) error {
	if err != nil {
		return err
	}
	return defErrFn()
}

func DefaultErrorMsg(err error, format string, p ...any) error {
	return DefaultErrorFn(err, func() error {
		return fmt.Errorf(format, p...)
	})
}

type MultiError []error

func (errs MultiError) Error() string {
	if len(errs) == 0 {
		return ""
	}
	buf := &bytes.Buffer{}
	_, _ = fmt.Fprintf(buf, "%d error(s) occurred:", len(errs))
	for _, err := range errs {
		_, _ = fmt.Fprintf(buf, "\n* %s", err)
	}
	return buf.String()
}

func (errs *MultiError) Append(err error) {
	if err != nil {
		*errs = append(*errs, err)
	}
}

func (errs MultiError) MaybeUnwrap() error {
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errs
	}
}
