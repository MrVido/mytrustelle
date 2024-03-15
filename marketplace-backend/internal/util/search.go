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

// IndexListing indexes a listing document in Elasticsearch with improved error handling.
func IndexListing(listing *model.Listing) error {
	es, err := config.NewElasticsearchClient()
	if err != nil {
		return fmt.Errorf("failed to create Elasticsearch client: %w", err)
	}

	body, err := json.Marshal(listing)
	if err != nil {
		return fmt.Errorf("failed to serialize listing: %w", err)
	}

	req := esapi.IndexRequest{
		Index:      "listings",
		DocumentID: fmt.Sprintf("%d", listing.ID),
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return fmt.Errorf("error indexing listing: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error from Elasticsearch: %s", res.String())
	}

	return nil
}

// SearchListings searches for listings in Elasticsearch with refined query and result parsing.
func SearchListings(query string) ([]model.Listing, error) {
	es, err := config.NewElasticsearchClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create Elasticsearch client: %w", err)
	}

	var buf bytes.Buffer
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"title^3", "description"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return nil, fmt.Errorf("failed to encode search query: %w", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("listings"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error performing search query: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error from Elasticsearch: %s", res.String())
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	// Parse search results
	var listings []model.Listing
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var listing model.Listing
		source := hit.(map[string]interface{})["_source"]
		data, err := json.Marshal(source)
		if err != nil {
			log.Printf("failed to marshal listing source: %v", err)
			continue
		}
		if err := json.Unmarshal(data, &listing); err != nil {
			log.Printf("failed to unmarshal listing data: %v", err)
			continue
		}
		listings = append(listings, listing)
	}

	return listings, nil
}
