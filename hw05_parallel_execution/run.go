package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var errCounter int32

	tasksCh := make(chan Task, len(tasks))
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(errCounter *int32) {
			for task := range tasksCh {
				if atomic.LoadInt32(errCounter) >= int32(m) {
					break
				}
				if err := task(); err != nil {
					atomic.AddInt32(errCounter, 1)
				}
			}
			wg.Done()
		}(&errCounter)
	}

	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)
	wg.Wait()
	if m != 0 && errCounter >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
