package host_info

import (
	"context"
	"errors"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/caarlos0/env"
)

var (
	ErrReturnedCached = errors.New("returned cached value")
)

type Effector func(context.Context) (*HostInfo, error)

type HostInfo struct {
	OS       string
	Host     string `env:"HOSTNAME"`
	Home     string `env:"HOME"`
	Username string `env:"USERNAME"`
	Shell    string `env:"SHELL"`
	Term     string `env:"TERM"`
}

func GetHostInfo(ctx context.Context) (*HostInfo, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		inf := HostInfo{}
		if err := env.Parse(&inf); err != nil {
			return nil, err
		}

		if inf.Host == "" {
			host, err := os.Hostname()
			if err != nil {
				return nil, err
			}
			inf.Host = host
		}

		inf.OS = runtime.GOOS

		return &inf, nil
	}
}

// host_info.GetHostInfo, 1, 1, 5*time.Second, cached - TODO - descr
func Throttle(e Effector, max uint, refill uint, d time.Duration, cached *HostInfo) Effector {
	var tokens = max
	var once sync.Once

	return func(ctx context.Context) (*HostInfo, error) {
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
