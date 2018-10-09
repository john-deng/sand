
package handler
//go:build

import (
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/hidevopsio/cqrses/pkg/sand/isa"
	"github.com/hidevopsio/cqrses/pkg/aggregate"
	"github.com/hidevopsio/cqrses/pkg/eventstore"
	"github.com/hidevopsio/cqrses/pkg/command"
)

type pipelineCommandHandler struct {
	isa.CommandHandler
	repository eventstore.Repository
}

// eventstore.Repository will be injected into repository
func newPipelineCommandHandler(repository eventstore.Repository) *pipelineCommandHandler {
	return &pipelineCommandHandler{
		repository: repository,
	}
}

func init() {
	app.Component(newPipelineCommandHandler)
}

func (h *pipelineCommandHandler) StartPipeline(cmd *command.StartPipeline)  {
	h.repository.NewInstance(func() interface{} {
		return aggregate.NewPipelineAggregate(cmd.PipelineId, cmd.Spec)
	})
}

func (h *pipelineCommandHandler) CancelPipeline(cmd *command.CancelPipeline)  {
	a := h.repository.Load(cmd.PipelineId)
	a.Execute(func(ar isa.Aggregate) {
		aggregateRoot := ar.(aggregate.PipelineAggregate)
		aggregateRoot.Cancel(cmd.PipelineId)
	})
}