package parallel

import (
	"context"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

func ForEach[T any](list []T, threads int, action func(t T) error) error {
	var eg errgroup.Group
	eg.SetLimit(threads)
	for _, item := range list {
		i := item
		eg.Go(func() error {
			return action(i)
		})
	}
	return eg.Wait()
}

func Map[T any, S any](list []T, threads int, action func(t T) S) ([]S, error) {
	mappedValues := make([]S, len(list))
	sem := semaphore.NewWeighted(1)
	var eg errgroup.Group
	eg.SetLimit(threads)
	for index, item := range list {
		i := item
		x := index
		eg.Go(func() error {
			val := action(i)
			defer sem.Release(1)
			_ = sem.Acquire(context.Background(), 1)
			mappedValues[x] = val
			return nil
		})
	}
	return mappedValues, eg.Wait()
}
