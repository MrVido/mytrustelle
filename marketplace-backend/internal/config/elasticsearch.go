package config

import (
    "context"
    "log"
    "os"
    "net/http"
    "crypto/tls"
    "time"

    "github.com/elastic/go-elasticsearch/v8"
)

// NewElasticsearchClient creates and configures an instance of the Elasticsearch client.
func NewElasticsearchClient() (*elasticsearch.Client, error) {
    addresses := os.Getenv("ELASTICSEARCH_ADDRESSES")
    username := os.Getenv("ELASTICSEARCH_USERNAME")
    password := os.Getenv("ELASTICSEARCH_PASSWORD")

    // Setup a custom HTTP client with expanded timeouts and disabled SSL verification (if needed)
    httpClient := &http.Client{
        Timeout: time.Second * 30, // Set request timeout
        Transport: &http.Transport{
            MaxIdleConnsPerHost:   10, // Set max idle connections per host
            ResponseHeaderTimeout: time.Second * 10, // Set response header timeout
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true, // Set to true to disable SSL verification (not recommended for production)
            },
        },
    }

    cfg := elasticsearch.Config{
        Addresses: []string{
            addresses, // Use environment variable to configure Elasticsearch addresses
        },
        Username: username, // Basic Authentication Username (optional)
        Password: password, // Basic Authentication Password (optional)
        Transport: httpClient, // Use custom HTTP client
    }

    // Create the Elasticsearch client
    client, err := elasticsearch.NewClient(cfg)
    if err != nil {
        log.Fatalf("Error creating the Elasticsearch client: %s", err)
    }

    // Optionally, you can add a ping to the Elasticsearch server to verify the connection
    res, err := client.Ping(client.Ping.WithContext(context.Background()))
    if err != nil {
        log.Fatalf("Error pinging Elasticsearch: %s", err)
    }
    defer res.Body.Close() // Don't forget to close the response body

    if res.IsError() {
        log.Fatalf("Error response from Elasticsearch: %s", res.String())
    } else {
        log.Println("Successfully connected to Elasticsearch")
    }

    return client, nil
}
