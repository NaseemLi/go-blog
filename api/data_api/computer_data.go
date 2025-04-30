package dataapi

import (
	"goblog/common/res"
	"goblog/utils/server"

	"github.com/gin-gonic/gin"
)

type ComputerDataResponse struct {
	CpuPercent  float64 `json:"cpuPercent"`  // CPU使用率
	MemPercent  float64 `json:"memPercent"`  // 内存使用率
	DiskPercent float64 `json:"diskPercent"` // 磁盘使用率
}

func (DataApi) ComputerDataView(c *gin.Context) {

	var data = ComputerDataResponse{
		CpuPercent:  server.GetCpuPercent(),
		MemPercent:  server.GetMemPercent(),
		DiskPercent: server.GetDiskPercent(),
	}

	res.OkWithData(data, c)
}
