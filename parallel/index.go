package parallel

import (
	"golang.org/x/sync/errgroup"
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
