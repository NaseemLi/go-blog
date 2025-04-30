package server

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/sirupsen/logrus"
)

func GetCpuPercent() float64 {
	// 获取CPU使用率
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		logrus.Errorf("Error getting CPU percent: %v", err)
		return 0
	}
	return cpuPercent[0]
}

func GetMemPercent() float64 {
	// 获取内存占用情况
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		logrus.Errorf("Error getting memory info: %v", err)
		return 0
	}
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	// 获取所有挂载点的磁盘使用率信息
	partitions, err := disk.Partitions(false)
	if err != nil {
		logrus.Errorf("Error getting memory info: %v", err)
		return 0
	}

	var total uint64
	var used uint64

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			log.Printf("Error getting usage for %s: %v", partition.Mountpoint, err)
			continue
		}

		total += usage.Total
		used += usage.Used
	}

	// 计算总使用率
	usagePercent := float64(used) / float64(total) * 100

	// 打印总磁盘使用率信息
	return usagePercent
}
