package frontend

import "github.com/s-vvardenfell/psyllium/internal/core"

type restFrontEnd struct {
	// c *core.Core
}

func (r restFrontEnd) StartWith(c *core.Core) error {
	c.Agent.Run()
	return nil
}
