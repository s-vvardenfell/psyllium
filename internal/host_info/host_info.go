package host_info

import (
	"context"
	"errors"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/caarlos0/env"
	"github.com/s-vvardenfell/psyllium/internal/core"
)

var (
	ErrReturnedCached = errors.New("returned cached value")
)

func GetHostInfo(ctx context.Context) (core.HostInfo, error) {
	select {
	case <-ctx.Done():
		return core.HostInfo{}, ctx.Err()
	default:
		inf := core.HostInfo{}
		if err := env.Parse(&inf); err != nil {
			return core.HostInfo{}, err
		}

		if inf.Host == "" {
			host, err := os.Hostname()
			if err != nil {
				return core.HostInfo{}, err
			}

			inf.Host = host
		}

		inf.OS = runtime.GOOS

		return inf, nil
	}
}

// host_info.GetHostInfo, 1, 1, 5*time.Second, cached - TODO - descr
func Throttle(
	e core.Effector, max uint, refill uint, d time.Duration, cached *core.HostInfo) core.Effector {
	var (
		tokens = max
		once   sync.Once
	)

	return func(ctx context.Context) (*core.HostInfo, error) {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		once.Do(func() {
			ticker := time.NewTicker(d)
			go func() {
				defer ticker.Stop()
				for {
					select {
					case <-ctx.Done():
						return
					case <-ticker.C:
						t := tokens + refill
						if t > max {
							t = max
						}
						tokens = t
					}
				}
			}()
		})

		if tokens <= 0 {
			return cached, ErrReturnedCached
		}
		tokens--

		return e(ctx)
	}
}
