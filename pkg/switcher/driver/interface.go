package metathings_switcher_driver

import (
	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type SwitcherState int32

const (
	STATE_UNKNOWN SwitcherState = iota
	STATE_ON
	STATE_OFF
	STATE_OVERFLOW
)

type Switcher struct {
	State SwitcherState
}

type SwitcherDriver interface {
	driver_helper.Driver
	Show() (Switcher, error)
	Turn(SwitcherState) (Switcher, error)
}

type NewDriverMethod func(opt_helper.Option) (SwitcherDriver, error)
