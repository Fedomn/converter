package dispatcher

import (
	"fedomn/converter/processor/galaxy"
)

const (
	GalaxyJob JobType = iota
)

var DefaultDispatcher Dispatcher

func init() {
	DefaultDispatcher = Dispatcher{
		workers: make(map[JobType]Worker),
	}
	DefaultDispatcher.workers[GalaxyJob] = newWorker(galaxy.DefaultGuider)
}

type Dispatcher struct {
	workers map[JobType]Worker
}

func (d Dispatcher) Start() {
	for _, worker := range d.workers {
		worker.Start()
	}
}

func (d Dispatcher) AddJob(j JobInput) {
	d.workers[j.Type].AddJob(j)
}

func (d Dispatcher) AcquireOutput(t JobType) JobOutputQueue {
	return d.workers[t].AcquireJobOutput()
}

func (d Dispatcher) Stop() {
	for _, worker := range d.workers {
		worker.Stop()
	}
}
