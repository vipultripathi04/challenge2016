# Real Image Challenge 2016 - Distributors Permissions API

A RESTful API to determine global distribution rights for movies by checking territorial restrictions for distributors.

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

```bash
git clone https://github.com/vipultripathi04/challenge2016.git
cd challenge2016
go mod download
```

### Quick Start

```bash
make run
```

The API will start on `http://localhost:8080`

## API Endpoint

### Check Distributor Permissions

```http
POST /distributors/permissions
```

**Request Body Example:**

```json
[
  {
    "distributor": "DISTRIBUTOR1",
    "include": [
      { "country": "INDIA" },
      { "country": "UNITED STATES" }
    ],
    "exclude": [
      { "province": "KARNATAKA", "country": "INDIA" }
    ],
    "locations": [
      { "city": "CHICAGO", "province": "ILLINOIS", "country": "UNITED STATES" },
      { "city": "CHENNAI", "province": "TAMIL NADU", "country": "INDIA" }
    ]
  }
]
```

**Request Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `distributor` | `string` | Yes | Distributor identifier |
| `include` | `[]Location` | Yes | Regions where distributor has rights |
| `exclude` | `[]Location` | No | Regions to exclude from included rights |
| `locations` | `[]Location` | Yes | Locations to check permissions for |

**Response (200 OK):**

```json
[
  {
    "distributor": "DISTRIBUTOR1",
    "permissions": "DISTRIBUTOR1 can distribute in CHICAGO-ILLINOIS-UNITED STATES"
  },
  {
    "distributor": "DISTRIBUTOR1",
    "permissions": "DISTRIBUTOR1 cannot distribute in CHENNAI-TAMIL NADU-INDIA"
  }
]
```

## Project Structure

```
.
├── cmd/api/
│   └── main.go           # API entry point
├── internal/
│   ├── config/           # Configuration
│   ├── handler/          # HTTP handlers
│   ├── model/            # Request/response models
│   ├── repository/       # Data access layer
│   └── service/          # Business logic
├── data/
│   └── cities.csv        # City database
├── go.mod                # Module file
└── README.md
```

## Commands

```bash
make run      # Run the API
make build    # Build executable
make clean    # Remove build artifacts
make test     # Run tests
make help     # Display help
```
