package output_port

import "errors"

var (
	ErrInvalidTransaction = errors.New("invalid transaction")
)

type GormTransaction interface {
	StartTransaction(func(tx interface{}) error) error
}