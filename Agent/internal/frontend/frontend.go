package frontend

import "agent/internal/core"

type FrontEnd interface {
	StartWith(c *core.Core) error
}

type zeroFrontEnd struct{}

func (f zeroFrontEnd) StartWith(c *core.Core) error {
	return nil
}
