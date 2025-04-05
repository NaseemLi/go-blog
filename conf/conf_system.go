package conf

import (
	"fmt"
	"strconv"
)

type System struct {
	IP      string `yaml:"ip"`
	Port    int    `yaml:"port"`
	GinMode string `yaml:"gin_mode"`
	Env     string `yaml:"env"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%s", s.IP, strconv.Itoa(s.Port))
}
