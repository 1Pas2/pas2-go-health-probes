package pas2

import (
	"context"
	"fmt"
	"sort"
	"sync"
)

type Probe func(context.Context) error
type Result struct {
	Name string
	Err  error
}

func Run(ctx context.Context, probes map[string]Probe, concurrency int) ([]Result, error) {
	if concurrency < 1 {
		return nil, fmt.Errorf("concurrency must be positive")
	}
	names := make([]string, 0, len(probes))
	for name, probe := range probes {
		if probe == nil {
			return nil, fmt.Errorf("nil probe: %s", name)
		}
		names = append(names, name)
	}
	sort.Strings(names)
	results := make([]Result, len(names))
	jobs := make(chan int)
	var wait sync.WaitGroup
	for n := 0; n < min(concurrency, len(names)); n++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			for i := range jobs {
				var err error
				if ctx.Err() != nil {
					err = ctx.Err()
				} else {
					err = probes[names[i]](ctx)
				}
				results[i] = Result{names[i], err}
			}
		}()
	}
	for i := range names {
		jobs <- i
	}
	close(jobs)
	wait.Wait()
	return results, nil
}
