package output_port

type GormTransaction interface {
	StartTransaction(func(tx interface{}) error) error
}