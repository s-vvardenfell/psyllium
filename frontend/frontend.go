package frontend

import "github.com/s-vvardenfell/psyllium/internal/core"

type FrontEnd interface {
	StartWith(c *core.Core) error
}

type zeroFrontEnd struct{}

func (f zeroFrontEnd) StartWith(c *core.Core) error {
	return nil
}
