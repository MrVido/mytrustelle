package util

import (
	"context"
	"encoding/json"
	"fmt"
	"marketplace-backend/internal/config"
	"marketplace-backend/internal/model"
)

// IndexListing indexes a listing document in Elasticsearch.
func IndexListing(listing *model.Listing) error {
	es, err := config.NewElasticsearchClient()
	if err != nil {
		return err
	}

	// Serialize listing
	body, err := json.Marshal(listing)
	if err != nil {
		return err
	}

	// Index the listing document
	res, err := es.Index(
		"listings", // Index name
		json.RawMessage(body),
		es.Index.WithDocumentID(fmt.Sprintf("%d", listing.ID)), // Use listing ID as document ID
		es.Index.WithRefresh("true"),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

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
