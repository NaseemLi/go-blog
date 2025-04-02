package core

import (
	"fmt"
	"strings"

	ipUtils "goblog/utils/ip"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
)

var searcher *xdb.Searcher

func InitIPDB() {
	var dbPath = "init/ip2region.xdb"
	_searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		logrus.Fatalf("ip地址数据库加载 %s", err)
		return
	}

	searcher = _searcher
}

func GetipADDR(ip string) (addr string) {
	if ipUtils.HasLocalIPAddr(ip) {
		return
	}
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		logrus.Warnf("错误的 ip 地址%s", err)
		return "错误的 ip 地址"
	}

	_addrList := strings.Split(region, "|")
	if len(_addrList) != 5 {
		logrus.Warnf("错误的 ip 地址%s", err)
		return "未知的 ip 地址"
	}
	//国家 省份 市 运营商
	country := _addrList[0]
	province := _addrList[2]
	city := _addrList[3]

	if province != "0" && city != "0" {
		return fmt.Sprintf("%s %s", province, country)
	}
	if country != "0" && province != "0" {
		return fmt.Sprintf("%s %s", country, province)
	}
	if country != "0" {
		return country
	}

	return region
}
