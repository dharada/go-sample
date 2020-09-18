package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"io/ioutil"
	"log"
)

func main() {

	fmt.Println("version:", elasticsearch.Version)
	fmt.Printf("indexer: %+v\n", esutil.BulkIndexerConfig{})

	onCloud := true

	if onCloud {

		cfg := elasticsearch.Config{
			CloudID: "v7-experience:ZXVyb3BlLXdlc3QxLmdjcC5jbG91ZC5lcy5pbyQ5ZjIyODNkOGY2ODg0ZWMyOTdlM2E2YzliN2Y5NzUzMCQ4NDdjZjVjZDg2MTA0Y2IyODBjMDAzNTM4M2YxOTc3MA==",
			APIKey:  "YjA1S24zUUIyc2VHa0phcW54QnA6bzJUMWtiV3BRcDI4bC03empNRHN5dw==",
		}

		es, err := elasticsearch.NewClient(cfg)

		res, err := es.Info()
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		log.Println(res)

	} else {

		cert, _ := ioutil.ReadFile("/home/daisuke/d/7.9.1/elasticsearch/config/certs/ca.crt")

		var (
			clusterURLs = []string{"https://daisuke:9200"}
			username    = "elastic"
			password    = "changeme"
		)

		cfg := elasticsearch.Config{
			Addresses: clusterURLs,
			Username:  username,
			Password:  password,
			CACert: cert,
		}

		es, err := elasticsearch.NewClient(cfg)

		res, err := es.Info()
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}

		log.Println(res)
	}

}
