package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"marketplace-backend/internal/config"
	"marketplace-backend/internal/model"
	"log"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// IndexListing indexes a listing document in Elasticsearch.
func IndexListing(listing *model.Listing) error {
	es, err := config.NewElasticsearchClient()
	if err != nil {
		return fmt.Errorf("error creating Elasticsearch client: %w", err)
	}

	body, err := json.Marshal(listing)
	if err != nil {
		return fmt.Errorf("error serializing listing: %w", err)
	}

	req := esapi.IndexRequest{
		Index:      "listings",
		DocumentID: fmt.Sprintf("%d", listing.ID),
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return fmt.Errorf("error indexing listing document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error indexing document ID=%d: %s", listing.ID, res.String())
		return fmt.Errorf("error from Elasticsearch: %s", res.String())
	}

	return nil
}
// SearchListings searches for listings in Elasticsearch based on a query.
func SearchListings(query string) ([]model.Listing, error) {
	es, err := config.NewElasticsearchClient()
	if err != nil {
		return nil, err
	}

	// Construct a search query
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"title^3", "description"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	// Perform the search query
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("listings"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Parse the response body to extract listings
	// Implementation depends on your Elasticsearch version and setup

	return listings, nil
}
