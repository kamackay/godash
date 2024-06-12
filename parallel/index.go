package parallel

import "sync"

func ForEach[T any](list []T, threads int) {
	var wg sync.WaitGroup
	wg.Add(threads)
}
