package aggregate

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/hidevopsio/cqrses/pkg/sand/isa"
	"github.com/hidevopsio/cqrses/pkg/sand"
	"github.com/hidevopsio/cqrses/pkg/event"
)

type Job interface {}

type Stage struct {
	Jobs []Job
}

type PipelineSpec struct {
	Stages []Stage
}

type PipelineStatus struct {
}


type Pipeline struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec PipelineSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" aggregate:"member"`

	// Most recently observed status of the Deployment.
	// +optional
	Status PipelineStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PipelineAggregate struct {
	// Embedded isa.Aggregate to annotate that this struct is a aggregate
	isa.Aggregate `json:"omitempty"`
	PipelineId string `aggregate:"identifier"`
	Pipeline
}

func NewPipelineAggregate(id string, spec interface{}) *PipelineAggregate  {
	sand.Apply(&event.PipelineStarted{
		Id: id,
	})
	return &PipelineAggregate{}
}

func init()  {
	sand.Aggregate(NewPipelineAggregate)
}


func (a *PipelineAggregate) Cancel(id string)  {
}


func (a *PipelineAggregate) PipelineStarted(e *event.PipelineStarted)  {
	a.PipelineId = e.Id
}

func (a *PipelineAggregate) PipelineCanceled(e *event.PipelineCanceled)  {

}

func (a *PipelineAggregate) PipelineRestarted(e *event.PipelineRestarted)  {

}