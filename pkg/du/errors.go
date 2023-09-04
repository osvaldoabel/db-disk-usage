package du

import "errors"

var (
	ErrNotImplemented              = errors.New("Browser not implemented")
	ErrNotFoundDiskUsageExecutable = errors.New("Disk Usage executable not found")
)

type error interface {
	Error() string
}
