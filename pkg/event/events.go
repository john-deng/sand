package event


type PipelineTemplate struct {

}


type Pipeline struct {

}

type PipelineStarted struct {
	Pipeline
	Id string
}


type PipelineRestarted struct {
	Pipeline
	Id string
}


type PipelineCanceled struct {
	Pipeline
}


