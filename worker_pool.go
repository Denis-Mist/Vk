package main

import (
	"sync"
)

type WorkerPool struct {
	tasks   chan string
	workers []*Worker
	mu      sync.Mutex
	wg      sync.WaitGroup
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	pool := &WorkerPool{
		tasks:   make(chan string),
		workers: make([]*Worker, 0, numWorkers),
	}

	for i := 0; i < numWorkers; i++ {
		worker := &Worker{id: i + 1, quit: make(chan struct{})}
		pool.workers = append(pool.workers, worker)
		pool.wg.Add(1)
		go worker.Start(pool.tasks, &pool.wg)
	}

	return pool
}

func (p *WorkerPool) AddTask(task string) {
	p.tasks <- task
}

func (p *WorkerPool) Stop() {
	close(p.tasks) // Закрываем канал задач
	p.wg.Wait()    // Ждем завершения всех воркеров
}

func (p *WorkerPool) AddWorker() {
	p.mu.Lock()
	defer p.mu.Unlock()

	worker := &Worker{id: len(p.workers) + 1, quit: make(chan struct{})}
	p.workers = append(p.workers, worker)
	p.wg.Add(1)
	go worker.Start(p.tasks, &p.wg)
}

func (p *WorkerPool) RemoveWorker() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.workers) > 0 {
		worker := p.workers[len(p.workers)-1]
		close(worker.quit) // Отправляем сигнал завершения
		p.workers = p.workers[:len(p.workers)-1]
	}
}
