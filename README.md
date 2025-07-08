# Meilisearch Product Catalog Indexer

A Go application that indexes product catalog data into Meilisearch for fast, typo-tolerant search functionality.

## 🚀 Features

- **Bulk Data Import**: Efficiently imports large JSON datasets into Meilisearch
- **Data Cleaning**: Automatically converts "NULL" strings to proper null values
- **Batch Processing**: Uploads documents in configurable batches for optimal performance
- **Progress Tracking**: Real-time progress updates during indexing
- **Search Testing**: Built-in search functionality to verify indexing
- **Error Handling**: Comprehensive error handling and logging

## 📋 Prerequisites

- Go 1.24.4 or higher
- Docker (for running Meilisearch)
- Git

## 🛠️ Installation & Setup

### 1. Clone the Repository

```bash
git clone <your-repo-url>
cd Meilisearch
```

### 2. Initialize Go Module

```bash
go mod init meilisearch-project
go mod tidy
```

### 3. Install Dependencies

```bash
go get github.com/meilisearch/meilisearch-go
```

### 4. Start Meilisearch Server

Using Docker (recommended):

```bash
# Run Meilisearch with data persistence
docker run -p 7700:7700 -v $(pwd)/data.ms:/data.ms getmeili/meilisearch:latest

# Or run without analytics for development
docker run -p 7700:7700 -v $(pwd)/data.ms:/data.ms getmeili/meilisearch:latest --no-analytics
```

### 5. Set Environment Variables (Optional)

For production, set the Meilisearch master key:

```bash
export MASTER_KEY="your-master-key-here"
```

For development, you can run without authentication.

## 📁 Project Structure

```
Meilisearch/
├── cmd/
│   └── main.go          # Main application file
├── dto/                 # Data Transfer Objects
├── handler/             # HTTP handlers
├── service/             # Business logic services
├── sku.json             # Product catalog data
├── query_result.json    # Alternative data source
├── query_result.csv     # CSV data source
├── go.mod               # Go module file
├── go.sum               # Go dependencies checksum
└── README.md           # This file
```

## 🚀 Usage

### Basic Usage

```bash
# Run the application
go run cmd/main.go
```

### What the Application Does

1. **Connects to Meilisearch** at `http://localhost:7700`
2. **Loads data** from `sku.json` (789 product records)
3. **Cleans the data** by converting "NULL" strings to null values
4. **Uploads documents** in batches of 1000
5. **Waits for indexing** to complete
6. **Runs search tests** to verify functionality
7. **Displays statistics** about the indexed data

### Expected Output

```
✅ Connected to Meilisearch successfully
📊 Loaded 789 documents from sku.json
🚀 Starting upload of 789 documents in 1 batches...
📦 Uploading batch 1/1 (789 documents)...
✅ Batch 1/1 uploaded successfully
🎉 All documents uploaded successfully!
⏳ Waiting for indexing task 12345 to complete...
✅ Indexing complete!
📈 Index stats: 789 documents indexed

🔍 Testing search functionality...
Test 1: Searching for 'FEVICOL'...
✅ Found 8 results for 'FEVICOL'
   First result: FEVICOL SH  1 KG
...
```

## 🔍 Testing

### Built-in Tests

The application includes three search tests:

1. **Product Name Search**: Searches for "FEVICOL" products
2. **Category Search**: Searches for "Adhesives" category
3. **SKU Search**: Searches for specific SKU "ADH1"

### Manual Testing with curl

```bash
# Check if index exists
curl "http://localhost:7700/indexes/sku"

# Search for products
curl "http://localhost:7700/indexes/sku/search" \
  -H "Content-Type: application/json" \
  -d '{"q": "FEVICOL", "limit": 5}'

# Get index statistics
curl "http://localhost:7700/indexes/sku/stats"
```

### Meilisearch Dashboard

If you have the dashboard enabled, visit:
- `http://localhost:7700` in your browser
- Use the web interface to test searches

## 📊 Data Structure

The application expects JSON data with the following structure:

```json
[
  {
    "id": 1,
    "sku": "ADH1",
    "name": "FEVICOL SH  1 KG",
    "category_id": 1,
    "description": "This is an adhesive SKU",
    "image_urls": ["url1", "url2"],
    "mrp": null,
    "status": "Active",
    "created_by": 24385,
    "updated_by": 24385,
    "created_at": "2025-06-10 03:39:32",
    "updated_at": "2025-06-12 09:54:02",
    "per_unit_mrp_price": null,
    "unit_type": null,
    "per_unit_selling_price": null,
    "unit_value": null,
    "selling_price": null,
    "category_brand_index_id": null,
    "is_active": 1,
    "discount": null,
    "category_name": "Adhesives"
  }
]
```

## ⚙️ Configuration

### Batch Size

Modify the batch size in `cmd/main.go`:

```go
batchSize := 1000  // Change this value as needed
```

### Meilisearch Connection

Update the connection settings in `cmd/main.go`:

```go
client := meilisearch.New("http://localhost:7700", meilisearch.WithAPIKey(os.Getenv("MASTER_KEY")))
```

### Data Source

Change the data source file in `cmd/main.go`:

```go
jsonFile, err := os.Open("sku.json")  // Change to your data file
```

## 🐛 Troubleshooting

### Common Issues

1. **Connection Failed**
   - Ensure Meilisearch is running on port 7700
   - Check if Docker container is active: `docker ps`

2. **Primary Key Error**
   - The application explicitly sets "id" as the primary key
   - Ensure your data has unique "id" values

3. **File Not Found**
   - Ensure `sku.json` exists in the project root
   - Check file permissions

4. **Indexing Failed**
   - Check Meilisearch logs: `docker logs <container-id>`
   - Verify data format is correct

### Logs

View Meilisearch logs:

```bash
docker logs $(docker ps -q --filter ancestor=getmeili/meilisearch:latest)
```

## 🔧 Development

### Adding New Features

1. **New Search Tests**: Add to the testing section in `main()`
2. **Data Processing**: Modify the `cleanData()` function
3. **Error Handling**: Add specific error cases as needed

### Code Structure

- `main()`: Application entry point and orchestration
- `cleanData()`: Data cleaning and normalization
- Search tests: Built-in verification functionality

## 📝 License

[Add your license information here]

## 🤝 Contributing

[Add contribution guidelines here]

## 📞 Support

[Add support contact information here] 