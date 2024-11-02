package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	id   int
	quit chan struct{}
}

func (w *Worker) Start(tasks <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				return // Канал закрыт, завершить работу
			}
			fmt.Printf("Worker %d: processing task: %s\n", w.id, task)
			time.Sleep(time.Second) // Имитация обработки задачи
		case <-w.quit:
			return // Завершить работу по сигналу
		}
	}
}
