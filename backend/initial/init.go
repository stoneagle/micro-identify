package main

import (
	"fmt"
	"identify/backend/common"
	cm "identify/backend/models/card"
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	yamlFile, err := ioutil.ReadFile("../config/.config.yaml")
	if err != nil {
		panic(err)
	}
	config := &common.Conf{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic(err)
	}
	cardConfig := config.Card.Database
	initCard(cardConfig)
}

func initCard(dbConfig common.DBConf, mode string) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Target)

	engine, err := xorm.NewEngine(dbConfig.Type, source)
	if err != nil {
		panic(err)
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	engine.TZLocation = location
	engine.StoreEngine("InnoDB")
	engine.Charset("utf8")
	if mode == "debug" {
		err = engine.DropTables(new(cm.Card), new(cm.Album), new(cm.Config), new(cm.Message), new(cm.Tag), new(cm.CardTagMap), new(cm.Release))
		if err != nil {
			panic(err)
		}
	}
	err = engine.Sync2(new(cm.Card), new(cm.Album), new(cm.Config), new(cm.Message), new(cm.Tag), new(cm.CardTagMap), new(cm.Release))
	if err != nil {
		panic(err)
	}
}
