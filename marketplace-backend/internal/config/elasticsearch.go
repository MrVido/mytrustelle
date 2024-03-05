package config

import (
	"github.com/elastic/go-elasticsearch/v8"
)

// NewElasticsearchClient creates and configures an instance of the Elasticsearch client.
func NewElasticsearchClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200", // Adjust this to your Elasticsearch server address
		},
		// Add Authentication details if needed
	}
	return elasticsearch.NewClient(cfg)
}
