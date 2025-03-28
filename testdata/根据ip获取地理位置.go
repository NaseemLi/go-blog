// testdata/10.获取地理位置.go
package main

import (
	"fmt"
	"goblog/core"
)

func main() {
	core.InitIPDB()
	addr, err := core.GetipADDR("175.0.201.207")
	fmt.Println(addr, err)
}
