package event

import "github.com/hidevopsio/cqrses/pkg/entity"

type OrderPlacedEvent struct {
	OrderId string
	OrderLine []entity.PipelineSpec
}
