package dispatcher

import (
	"fedomn/converter/processor/galaxy"
	"reflect"
)

const (
	GalaxyJob JobType = iota
)

var DefaultDispatcher Dispatcher

func init() {
	DefaultDispatcher = Dispatcher{
		workerMap: make(map[JobType][]Worker),
	}
	DefaultDispatcher.workerMap[GalaxyJob] = []Worker{
		newWorker(GalaxyJob, galaxy.DefaultGuider),
		newWorker(GalaxyJob, galaxy.DefaultGuider),
		newWorker(GalaxyJob, galaxy.DefaultGuider),
	}
}

type Dispatcher struct {
	workerMap map[JobType][]Worker
}

func (d Dispatcher) Start() {
	for _, workers := range d.workerMap {
		for _, worker := range workers {
			worker.Start()
		}
	}
}

func (d Dispatcher) AddJob(j JobInput) {
	cases := make([]reflect.SelectCase, 0)
	for _, worker := range d.workerMap[j.Type] {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(worker.jobInputQueue),
			Send: reflect.ValueOf(j),
		})
	}
	reflect.Select(cases)
}

func (d Dispatcher) AcquireOutput(t JobType) JobOutput {
	cases := make([]reflect.SelectCase, 0)
	for _, worker := range d.workerMap[t] {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(worker.jobOutputQueue),
		})
	}
	_, receive, _ := reflect.Select(cases)
	return receive.Interface().(JobOutput)
}

func (d Dispatcher) Stop() {
	for _, workers := range d.workerMap {
		for _, worker := range workers {
			worker.Stop()
		}
	}
}
