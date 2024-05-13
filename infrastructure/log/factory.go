package log

import (
	"errors"

	"attendance-record/adapter/logger"
)

const (
	InstanceZapLogger int = iota
)

var (
	errInvalidLoggerInstance = errors.New("invalid logger instance")
)

func NewLoggerFactory(instance int) (logger.Logger, error) {
	switch instance {
	case InstanceZapLogger:
		return NewZapLogger()
	default:
		return nil, errInvalidLoggerInstance
	}
}
