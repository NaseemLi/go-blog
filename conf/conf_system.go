package conf

import (
	"fmt"
	"strconv"
)

type System struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%s", s.IP, strconv.Itoa(s.Port))
}
