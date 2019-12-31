package main

import (
	"encoding/json"

	"github.com/micro/go-plugins/config/source/url"

	//"outback/micro-go/configration/url"
	"time"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
)

func main() {
	urlSource := url.NewSource(
		url.WithURL("http://localhost:10010/config/mysql"),
	)
	// Create new config
	conf := config.NewConfig(
		config.WithSource(urlSource),
	)
	var defautTicker time.Duration = time.Second * 6
	// Load url source
	for {
		err := conf.Load(urlSource)
		if err != nil {
			log.Info("load error ", err.Error())
			<-time.Tick(defautTicker)
			continue
		}
		changeSet, err := urlSource.Read()
		if err != nil {
			log.Info("Read error ", err.Error())
			<-time.Tick(defautTicker)
			continue
		}
		log.Info("cheange set  Data    is ", changeSet.Data)
		log.Info("cheange set checksum is ", changeSet.Checksum)
		log.Info("cheange set Format   is ", changeSet.Format)
		log.Info("cheange set Source   is ", changeSet.Source)
		log.Info("cheange Timestamp    is ", changeSet.Timestamp)
		log.Info("cheange set sum      is ", changeSet.Sum())
		d := make(map[string]interface{}, 0)
		json.Unmarshal(changeSet.Data, &d)
		log.Info("d is ", d)
		<-time.Tick(defautTicker)
	}

}
