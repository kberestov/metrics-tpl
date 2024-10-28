package domain

import (
	"errors"
	"fmt"
)

func unknownKindErr(kind MetricKind) error {
	return fmt.Errorf(`unknown metric kind %s`, kind)
}

func invalidNameErr() error {
	return errors.New("metric value is blank")
}

func invalidValueErr(kind MetricKind, internalErr error) error {
	return fmt.Errorf(`invalid value for "%v" kind of metric: %w`, kind, internalErr)
}
