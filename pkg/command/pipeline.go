package command

import "github.com/hidevopsio/cqrses/pkg/entity"

type CancelPipeline struct {
	PipelineId string `identifier:"target"`
}

type StartPipeline struct {
	PipelineId string
	Spec       []entity.PipelineSpec
}