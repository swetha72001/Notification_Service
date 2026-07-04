package internal

import (
	"Backend/dto"
	"context"
	"sync"
)

func Start(ctx context.Context, worker int, d *DispatcherService, value <-chan *dto.Events, wg *sync.WaitGroup) {

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():

					return
				case val, ok := <-value:
					if !ok {
						return
					}

					d.DispatcherProcess(ctx, val)

				}
			}

		}()

	}

}
