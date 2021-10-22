package config

import (
	"strconv"
)

type ConfigStruct struct {
	Region     string
	Endpoint   string
	DisableSSL bool
	Output     string
	Loglevel   string
}

func (c ConfigStruct) String() string {
	return "Region=" + c.Region + " ENDPOINT=" + c.Endpoint +
		" DISABLE_SSL=" + strconv.FormatBool(c.DisableSSL) +
		" OUTPUT=" + c.Output + " LOGLEVEL=" + c.Loglevel
}
