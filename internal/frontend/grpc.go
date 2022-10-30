package frontend

import "github.com/s-vvardenfell/psyllium/internal/core"

type grpcFrontEnd struct {
	// c *core.Core
}

// TODO?
// NewFrontend()

func (g grpcFrontEnd) StartWith(c *core.Core) error {
	c.Agent.Run()
	return nil
}
