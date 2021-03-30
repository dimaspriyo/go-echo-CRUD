package elasticsearch

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func ESConnection() *elasticsearch.Client {
	ES, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return ES
}
