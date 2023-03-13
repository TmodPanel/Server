package tmd

import "errors"

const (
	STOP    = 0
	RUNNING = 1
)

type ProcessorPool struct {
	Capacity int
	Running  int
	Workers  []*TModProc
}

func NewProcessorPool(capacity int) (*ProcessorPool, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity can not be zero")
	}
	return &ProcessorPool{
		Capacity: capacity,
		Running:  0,
		Workers:  make([]*TModProc, 0),
	}, nil
}

func (p *ProcessorPool) AddWorker(worker *TModProc) error {
	if len(p.Workers) >= p.Capacity {
		return errors.New("pool is full")
	}
	p.Workers = append(p.Workers, worker)
	return nil
}

func (p *ProcessorPool) RemoveWorker(worker *TModProc) error {
	for i, w := range p.Workers {
		if w == worker {
			p.Workers = append(p.Workers[:i], p.Workers[i+1:]...)
			return nil
		}
	}
	return errors.New("worker not found")
}

func (p *ProcessorPool) GetWorker(id int) (*TModProc, error) {
	if len(p.Workers) == 0 {
		return nil, errors.New("pool is empty")
	}
	for _, w := range p.Workers {
		if w.ID == id {
			return w, nil
		}
	}
	return nil, errors.New("worker not found")
}

func (p *ProcessorPool) Run(id int) error {
	worker, err := p.GetWorker(id)
	if err != nil {
		return err
	}
	if worker.IsOpen {
		return errors.New("worker is already running")
	}
	go worker.Start()
	p.Running++
	return nil
}

func (p *ProcessorPool) Stop(id int) error {
	worker, err := p.GetWorker(id)
	if err != nil {
		return err
	}
	if !worker.IsOpen {
		return errors.New("worker is already stopped")
	}
	p.Running--
	return worker.Stop()
}
