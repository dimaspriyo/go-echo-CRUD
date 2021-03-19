package repository

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func ESConnection() (es *elasticsearch.Client, err error) {
	es, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	return es, err
}
