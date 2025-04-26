package core

import (
	"goblog/global"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func EsConnect() *elastic.Client {
	es := global.Config.ES
	if !es.Enable || es.Addr == "" {
		logrus.Infof("es未启用或地址为空, es连接失败!")
		return nil
	}
	client, err := elastic.NewClient(
		elastic.SetURL(es.Url()),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(es.Username, es.Password),
	)
	if err != nil {
		logrus.Panicf("es 连接失败 %s!", err)
		return nil
	}
	logrus.Infof("es连接成功!")
	return client
}
