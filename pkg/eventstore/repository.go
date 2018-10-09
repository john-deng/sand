package eventstore

import "github.com/hidevopsio/cqrses/pkg/sand/isa"

type Repository interface {
	NewInstance(cb func() interface{})
	Load(id interface{}) isa.Aggregate
}
