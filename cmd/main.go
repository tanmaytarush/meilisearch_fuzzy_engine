// In the command line:
// go get -u github.com/meilisearch/meilisearch-go

// In your .go file:
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/meilisearch/meilisearch-go"
)

// cleanData converts "NULL" strings to nil and ensures proper data types
func cleanData(data []map[string]interface{}) []map[string]interface{} {
	for _, item := range data {
		for key, value := range item {
			// Convert "NULL" strings to actual null
			if str, ok := value.(string); ok && strings.ToUpper(str) == "NULL" {
				item[key] = nil
			}
			// Ensure numeric fields are properly typed
			if key == "id" || key == "category_id" || key == "created_by" || key == "updated_by" || key == "is_active" {
				if str, ok := value.(string); ok {
					if str == "NULL" {
						item[key] = nil
					}
				}
			}
		}
	}
	return data
}

func main() {
	// Initialize Meilisearch client for v0.32.0
	client := meilisearch.New("http://localhost:7700", meilisearch.WithAPIKey(os.Getenv("MASTER_KEY")))

	// Test connection to Meilisearch
	_, err := client.Health()
	if err != nil {
		log.Fatalf("Failed to connect to Meilisearch: %v", err)
	}
	fmt.Println("✅ Connected to Meilisearch successfully")

	// Create or get the index
	indexName := "sku"
	index := client.Index(indexName)

	// Read and parse JSON file
	jsonFile, err := os.Open("sku.json")
	if err != nil {
		log.Fatalf("Failed to open sku.json: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read sku.json: %v", err)
	}

	var sku []map[string]interface{}
	if err := json.Unmarshal(byteValue, &sku); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	fmt.Printf("📊 Loaded %d documents from sku.json\n", len(sku))

	// Clean the data before sending to Meilisearch
	sku = cleanData(sku)

	// Upload documents in batches
	batchSize := 1000
	totalBatches := (len(sku) + batchSize - 1) / batchSize

	fmt.Printf("🚀 Starting upload of %d documents in %d batches...\n", len(sku), totalBatches)

	var lastTaskUID int64

	for i := 0; i < len(sku); i += batchSize {
		end := i + batchSize
		if end > len(sku) {
			end = len(sku)
		}

		batchNum := (i / batchSize) + 1
		fmt.Printf("📦 Uploading batch %d/%d (%d documents)...\n", batchNum, totalBatches, end-i)

		task, err := index.AddDocuments(sku[i:end], "id")
		if err != nil {
			log.Fatalf("Failed to upload batch %d-%d: %v", i, end, err)
		}
		lastTaskUID = task.TaskUID

		fmt.Printf("✅ Batch %d/%d uploaded successfully\n", batchNum, totalBatches)

		// Small delay between batches to avoid overwhelming the server
		if i+batchSize < len(sku) {
			time.Sleep(100 * time.Millisecond)
		}
	}

	fmt.Println("🎉 All documents uploaded successfully!")

	// Wait for the last task to complete before getting stats
	if lastTaskUID != 0 {
		fmt.Printf("⏳ Waiting for indexing task %d to complete...\n", lastTaskUID)
		for {
			taskStatus, err := client.GetTask(lastTaskUID)
			if err != nil {
				fmt.Printf("⚠️  Could not get task status: %v\n", err)
				break
			}
			if taskStatus.Status == "succeeded" {
				fmt.Println("✅ Indexing complete!")
				break
			} else if taskStatus.Status == "failed" {
				fmt.Printf("❌ Indexing failed: %v\n", taskStatus.Error)
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
	}

	// Get index stats
	stats, err := index.GetStats()
	if err != nil {
		fmt.Printf("⚠️  Could not get index stats: %v\n", err)
	} else {
		fmt.Printf("📈 Index stats: %d documents indexed\n", stats.NumberOfDocuments)
	}

	// Test search functionality
	fmt.Println("\n🔍 Testing search functionality...")

	// Test 1: Search for "FEVICOL"
	fmt.Println("Test 1: Searching for 'FEVICOL'...")
	searchRes, err := index.Search("FEVICOL", &meilisearch.SearchRequest{
		Limit: 5,
	})
	if err != nil {
		fmt.Printf("❌ Search failed: %v\n", err)
	} else {
		fmt.Printf("✅ Found %d results for 'FEVICOL'\n", len(searchRes.Hits))
		if len(searchRes.Hits) > 0 {
			fmt.Printf("   First result: %s\n", searchRes.Hits[0].(map[string]interface{})["name"])
		}
	}

	// Test 2: Search for "Adhesives" category
	fmt.Println("\nTest 2: Searching for 'Adhesives' category...")
	searchRes2, err := index.Search("Adhesives", &meilisearch.SearchRequest{
		Limit: 3,
	})
	if err != nil {
		fmt.Printf("❌ Search failed: %v\n", err)
	} else {
		fmt.Printf("✅ Found %d results for 'Adhesives'\n", len(searchRes2.Hits))
		if len(searchRes2.Hits) > 0 {
			fmt.Printf("   First result: %s\n", searchRes2.Hits[0].(map[string]interface{})["name"])
		}
	}

	// Test 3: Search by SKU code
	fmt.Println("\nTest 3: Searching for SKU 'ADH1'...")
	searchRes3, err := index.Search("ADH1", &meilisearch.SearchRequest{
		Limit: 1,
	})
	if err != nil {
		fmt.Printf("❌ Search failed: %v\n", err)
	} else {
		fmt.Printf("✅ Found %d results for 'ADH1'\n", len(searchRes3.Hits))
		if len(searchRes3.Hits) > 0 {
			result := searchRes3.Hits[0].(map[string]interface{})
			fmt.Printf("   SKU: %s, Name: %s, Category: %s\n",
				result["sku"], result["name"], result["category_name"])
		}
	}

	fmt.Println("\n🎉 All tests completed!")
}
