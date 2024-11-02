package main

import (
	"fmt"
	"time"
)

func main() {
	pool := NewWorkerPool(3)

	// Добавляем задачи
	for i := 0; i < 10; i++ {
		pool.AddTask(fmt.Sprintf("Task %d", i+1))
	}

	// Динамически добавляем воркера
	pool.AddWorker()

	// Добавляем новые задачи после добавления воркера
	for i := 10; i < 15; i++ {
		pool.AddTask(fmt.Sprintf("Task %d", i+1))
	}

	// Динамически удаляем воркера
	pool.RemoveWorker()
	pool.RemoveWorker()
	pool.RemoveWorker()

	// ожидание завершения воркеров
	time.Sleep(time.Second * 10)

	// Добавляем новые задачи после удаления воркера
	for i := 15; i < 20; i++ {
		pool.AddTask(fmt.Sprintf("Task %d", i+1))
	}

	// Ждем завершения обработки
	pool.Stop()
	fmt.Println("All tasks processed.")
}
